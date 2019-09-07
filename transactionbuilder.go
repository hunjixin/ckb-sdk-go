package ckb_sdk_go

import (
	"ckb-sdk-go/core"
	"ckb-sdk-go/crypto"
	"github.com/decred/dcrd/dcrec/secp256k1"
)

type TransactionBuilder struct {
	Version   uint32
	Deps      []core.OutPoint
	Inputs    []core.CellInput
	Outputs   []core.CellOutput
	Witnesses []core.Witness
}

func NewTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{
		Deps:      []core.OutPoint{},
		Inputs:    []core.CellInput{},
		Outputs:   []core.CellOutput{},
		Witnesses: []core.Witness{},
	}
}

func FromTransction(tx core.Transaction) *TransactionBuilder {
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

func (builder *TransactionBuilder) AppendOutPoint(dep core.OutPoint) *TransactionBuilder {
	builder.Deps = append(builder.Deps, dep)
	return builder
}

func (builder *TransactionBuilder) ClearDeps(dep core.OutPoint) *TransactionBuilder {
	builder.Deps = []core.OutPoint{}
	return builder
}

func (builder *TransactionBuilder) AppendInput(input core.CellInput) *TransactionBuilder {
	builder.Inputs = append(builder.Inputs, input)
	return builder
}

func (builder *TransactionBuilder) ClearInput(input core.CellInput) *TransactionBuilder {
	builder.Inputs = []core.CellInput{}
	return builder
}

func (builder *TransactionBuilder) AppendOutput(output core.CellOutput) *TransactionBuilder {
	builder.Outputs = append(builder.Outputs, output)
	return builder
}

func (builder *TransactionBuilder) ClearOutput() *TransactionBuilder {
	builder.Outputs = []core.CellOutput{}
	return builder
}

func (builder *TransactionBuilder) AppendWitness(witness core.Witness) *TransactionBuilder {
	builder.Witnesses = append(builder.Witnesses, witness)
	return builder
}

func (builder *TransactionBuilder) ClearWitness() *TransactionBuilder {
	builder.Witnesses = nil
	return builder
}

func (builder *TransactionBuilder) Build() core.Transaction {
	return core.Transaction{
		Version:   builder.Version,
		Deps:      builder.Deps,
		Inputs:    builder.Inputs,
		Outputs:   builder.Outputs,
		Witnesses: builder.Witnesses,
	}
}

func SignTx(tx *TransactionBuilder, priv *secp256k1.PrivateKey) []byte {
	raw := tx.Build()
	txHash := raw.TxHash()
	hash := crypto.Black256(txHash[:])
	sig := crypto.SignMesage(hash[:], priv)
	tx.AppendWitness([][]byte{sig})
	return sig
}
