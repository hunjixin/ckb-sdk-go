package core

import (
	"ckb-sdk-go/bincode"
	"ckb-sdk-go/crypto"
)

type RawTransaction struct {
	Version    uint32
	CellDeps   []CellDep
	HeadDeps   []H256
	Inputs     []CellInput
	Outputs    []CellOutput
	OutputData [][]byte
}

func (tx *Transaction) RawTransaction() RawTransaction {
	return RawTransaction{
		Version:    tx.Version,
		CellDeps:   tx.Cell_deps,
		HeadDeps:   tx.Header_deps,
		Inputs:     tx.Inputs,
		Outputs:    tx.Outputs,
		OutputData: tx.Outputs_data,
	}
}

func (tx *Transaction) TxHash() H256 {
	rawTx := &RawTransaction{
		Version:    tx.Version,
		CellDeps:   tx.Cell_deps,
		HeadDeps:   tx.Header_deps,
		Inputs:     tx.Inputs,
		Outputs:    tx.Outputs,
		OutputData: tx.Outputs_data,
	}
	rawTxBytes, _ := bincode.Marshal(rawTx)
	hBytes := crypto.Black256(rawTxBytes)
	h256 := H256{}
	h256.SetBytes(hBytes)
	return h256
}

func (tx *Transaction) WitnessHash() H256 {
	bytes, _ := bincode.Marshal(tx)
	hBytes := crypto.Black256(bytes)
	h256 := H256{}
	h256.SetBytes(hBytes)
	return h256
}
