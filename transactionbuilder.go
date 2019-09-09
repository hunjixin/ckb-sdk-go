package ckb_sdk_go

import (
	"ckb-sdk-go/core"
	"ckb-sdk-go/crypto"
	"github.com/decred/dcrd/dcrec/secp256k1"
)

type TransactionBuilder struct {
	Version   uint32
	CellDeps       []core.CellDep
	HeadDeps      []core.H256
	Inputs    []core.CellInput
	Outputs   []core.CellOutput
	Witnesses []core.Witness
	OutputsData [][]byte
}

func NewTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{
		CellDeps:      []core.CellDep{},
		HeadDeps:       []core.H256{},
		Inputs:    []core.CellInput{},
		Outputs:   []core.CellOutput{},
		Witnesses: []core.Witness{},
		OutputsData:       [][]byte{},
	}
}

func FromTransction(tx core.Transaction) *TransactionBuilder {
	return &TransactionBuilder{
		Version:   tx.Version,
		CellDeps:      tx.Cell_deps,
		HeadDeps:      tx.Header_deps,
		Inputs:    tx.Inputs,
		Outputs:   tx.Outputs,
		Witnesses: tx.Witnesses,
		OutputsData:tx.Outputs_data,
	}
}

func (builder *TransactionBuilder) SetVersion(version uint32) *TransactionBuilder {
	builder.Version = version
	return builder
}

func (builder *TransactionBuilder) AppendCellDeps(dep core.CellDep) *TransactionBuilder {
	builder.CellDeps = append(builder.CellDeps, dep)
	return builder
}

func (builder *TransactionBuilder) ClearDeps(dep core.OutPoint) *TransactionBuilder {
	builder.CellDeps = []core.CellDep{}
	return builder
}

func (builder *TransactionBuilder) AppendHeadDeps(dep core.H256) *TransactionBuilder {
	builder.HeadDeps = append(builder.HeadDeps, dep)
	return builder
}

func (builder *TransactionBuilder) ClearHeadDeps(dep core.H256) *TransactionBuilder {
	builder.HeadDeps = []core.H256{}
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

func (builder *TransactionBuilder) AppendOutputData(output []byte) *TransactionBuilder {
	builder.OutputsData = append(builder.OutputsData, output)
	return builder
}

func (builder *TransactionBuilder) ClearOutputsData() *TransactionBuilder {
	builder.OutputsData = [][]byte{}
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
		Cell_deps:      builder.CellDeps,
		Header_deps:      builder.HeadDeps,
		Inputs:    builder.Inputs,
		Outputs:   builder.Outputs,
		Witnesses: builder.Witnesses,
		Outputs_data:builder.OutputsData,
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
