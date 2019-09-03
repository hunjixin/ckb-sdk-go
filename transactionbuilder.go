package ckb_sdk_go

import (
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/dchest/blake2b"
	"reflect"
)

type Witness [][]byte
type H256 string
type U256 string
type Marshaler interface {
	Marshal(binCode *BinCodeSerizlize, val reflect.Value) error
}

type UnMarshaler interface {
	UnMarshal(binCode *BinCodeDeSerizlize) (reflect.Value,error)
}

func (_ H256) UnMarshal(binCode *BinCodeDeSerizlize) (reflect.Value,error) {
	strBytes, err := binCode.SliceBytes()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	h256Val := H256(string(strBytes.Bytes()))
	return reflect.ValueOf(h256Val), nil
}

type  TransactionBuilder struct {
	Version uint32
	Deps []OutPoint
	Inputs []CellInput
	Outputs []CellOutput
	Witnesses []Witness
}

func FromTransction(tx Transaction) *TransactionBuilder {
	return &TransactionBuilder{
		Version:tx.Version,
		Deps:tx.Deps,
		Inputs:tx.Inputs,
		Outputs:tx.Outputs,
		Witnesses:tx.Witnesses,
	}
}

func (builder *TransactionBuilder) SetVersion(version uint32) *TransactionBuilder{
	builder.Version= version
	return builder
}

func (builder *TransactionBuilder) AppendOutPoint(dep OutPoint) *TransactionBuilder{
	builder.Deps= append(builder.Deps, dep)
	return builder
}

func (builder *TransactionBuilder) ClearDeps(dep OutPoint) *TransactionBuilder{
	builder.Deps = []OutPoint{}
	return builder
}

func (builder *TransactionBuilder) AppendInput(input CellInput) *TransactionBuilder{
	builder.Inputs = append(builder.Inputs,input)
	return builder
}

func (builder *TransactionBuilder) ClearInput(input CellInput) *TransactionBuilder{
	builder.Inputs = []CellInput{}
	return builder
}

func (builder *TransactionBuilder) AppendOutput(output CellOutput) *TransactionBuilder{
	builder.Outputs = append(builder.Outputs,output)
	return builder
}

func (builder *TransactionBuilder) ClearOutput() *TransactionBuilder{
	builder.Outputs = []CellOutput{}
	return builder
}

func (builder *TransactionBuilder) AppendWitness(witness Witness) *TransactionBuilder{
	builder.Witnesses = append(builder.Witnesses, witness)
	return builder
}

func (builder *TransactionBuilder) ClearWitness() *TransactionBuilder{
	builder.Witnesses = nil
	return builder
}


type RawTransaction struct{
	Version uint32
	Deps  []OutPoint
	Inputs []CellInput
	Outputs []CellOutput
}


func  (builder *TransactionBuilder)Build() Transaction {
	return Transaction{
		Version:builder.Version,
		Deps:builder.Deps,
		Inputs:builder.Inputs,
		Outputs:builder.Outputs,
		Witnesses:builder.Witnesses,
	}
}

func (tx *Transaction)  TxHash() [32]byte {
	rawTx := &RawTransaction{
		Version:tx.Version,
		Deps: tx.Deps,
		Inputs:tx.Inputs,
		Outputs:tx.Outputs,
	}
	bytes, _ := Marshal(rawTx)
	str := "170000000000000000000000010000000000000001420000000000000030783030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303000000000007B0000000000000001000000000000000088526A740000000300000000000000010203000000000000000042000000000000003078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030300000000000"
	bytes222, _ := hex.DecodeString(str)
	t := reflect.TypeOf(RawTransaction{})
	dddd, _ := UnMarshal(bytes222, t)
	fmt.Println(dddd)
	hashBytes := blake2b.Sum256(bytes)
	return hashBytes
}
func (tx *Transaction) WitnessHash() [32]byte {
	bytes, _ := Marshal(tx)
	hashBytes := blake2b.Sum256(bytes)
	return hashBytes
}



func sign(tx TransactionBuilder, priv *secp256k1.PrivateKey) Transaction {

	raw := tx.Build()
	data, _ := Marshal(raw)
	hash := blake2b.Sum256(data)
	sig, _  := priv.Sign(hash[:])

	tx.AppendWitness([][]byte{sig.Serialize()})
	return tx.Build()
}

