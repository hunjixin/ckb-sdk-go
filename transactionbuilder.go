package ckb_sdk_go

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/ethereum/go-ethereum/crypto"
	"reflect"
)

type Witness [][]byte
type H256 [32]byte
type U256 string

var (
	ZeroH256 = H256{}
)

type Marshaler interface {
	Marshal(binCode *BinCodeSerizlize, val reflect.Value) error
}

type UnMarshaler interface {
	UnMarshal(binCode *BinCodeDeSerizlize) (reflect.Value, error)
}

func (h256 H256) Bytes() []byte {
	return h256[:]
}
func (h256 *H256) SetBytes(h256Bytes []byte) {
	copy(h256[:], h256Bytes)
}
func (_ H256) UnMarshal(binCode *BinCodeDeSerizlize) (reflect.Value, error) {
	strBytes, err := binCode.SliceBytes()
	if err != nil {
		return reflect.ValueOf(nil), err
	}

	hexBytes, err := hex.DecodeString(string(strBytes.Bytes()[2:]))
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	h256 := [32]byte{}
	copy(h256[:], hexBytes)
	return reflect.ValueOf(H256(h256)), nil
}

func (_ H256) Marshal(binCode *BinCodeSerizlize, val reflect.Value) error {
	hexH256 := "0x" + hex.EncodeToString(val.Interface().(H256).Bytes())
	return binCode.SliceBytes(reflect.ValueOf([]byte(hexH256)))
}

type TransactionBuilder struct {
	Version   uint32
	Deps      []OutPoint
	Inputs    []CellInput
	Outputs   []CellOutput
	Witnesses []Witness
}

func FromTransction(tx Transaction) *TransactionBuilder {
	return &TransactionBuilder{
		Version:   tx.Version,
		Deps:      tx.Deps,
		Inputs:    tx.Inputs,
		Outputs:   tx.Outputs,
		Witnesses: tx.Witnesses,
	}
}

func (builder *TransactionBuilder) SetVersion(version uint32) *TransactionBuilder {
	builder.Version = version
	return builder
}

func (builder *TransactionBuilder) AppendOutPoint(dep OutPoint) *TransactionBuilder {
	builder.Deps = append(builder.Deps, dep)
	return builder
}

func (builder *TransactionBuilder) ClearDeps(dep OutPoint) *TransactionBuilder {
	builder.Deps = []OutPoint{}
	return builder
}

func (builder *TransactionBuilder) AppendInput(input CellInput) *TransactionBuilder {
	builder.Inputs = append(builder.Inputs, input)
	return builder
}

func (builder *TransactionBuilder) ClearInput(input CellInput) *TransactionBuilder {
	builder.Inputs = []CellInput{}
	return builder
}

func (builder *TransactionBuilder) AppendOutput(output CellOutput) *TransactionBuilder {
	builder.Outputs = append(builder.Outputs, output)
	return builder
}

func (builder *TransactionBuilder) ClearOutput() *TransactionBuilder {
	builder.Outputs = []CellOutput{}
	return builder
}

func (builder *TransactionBuilder) AppendWitness(witness Witness) *TransactionBuilder {
	builder.Witnesses = append(builder.Witnesses, witness)
	return builder
}

func (builder *TransactionBuilder) ClearWitness() *TransactionBuilder {
	builder.Witnesses = nil
	return builder
}

type RawTransaction struct {
	Version uint32
	Deps    []OutPoint
	Inputs  []CellInput
	Outputs []CellOutput
}

func (builder *TransactionBuilder) Build() Transaction {
	return Transaction{
		Version:   builder.Version,
		Deps:      builder.Deps,
		Inputs:    builder.Inputs,
		Outputs:   builder.Outputs,
		Witnesses: builder.Witnesses,
	}
}

func (tx *Transaction) TxHash() H256 {
	rawTx := &RawTransaction{
		Version: tx.Version,
		Deps:    tx.Deps,
		Inputs:  tx.Inputs,
		Outputs: tx.Outputs,
	}
	rawTxBytes, _ := Marshal(rawTx)
	return Black256(rawTxBytes)
}

func (tx *Transaction) WitnessHash() H256 {
	bytes, _ := Marshal(tx)
	return Black256(bytes)
}

func SignTx(tx TransactionBuilder, priv *secp256k1.PrivateKey) Transaction {
	raw := tx.Build()
	txHash := raw.TxHash()
	hashBytes := make([][]byte, len(tx.Witnesses)+1)
	hashBytes[0] = txHash[:]
	for index, witness := range tx.Witnesses {
		for _, bytes := range witness {
			index++
			hashBytes[index] = bytes
		}
	}
	hash := Black256M(hashBytes...)
	sig, _ := crypto.Sign(hash[:], (*ecdsa.PrivateKey)(priv))

	tx.AppendWitness([][]byte{sig})
	return tx.Build()
}
