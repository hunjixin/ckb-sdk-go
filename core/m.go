package core

import (
	"ckb-sdk-go/bincode"
	"ckb-sdk-go/crypto"
)

type RawTransaction struct {
	Version uint32
	Deps    []OutPoint
	Inputs  []CellInput
	Outputs []CellOutput
}

func (tx *Transaction) TxHash() H256 {
	rawTx := &RawTransaction{
		Version: tx.Version,
		Deps:    tx.Deps,
		Inputs:  tx.Inputs,
		Outputs: tx.Outputs,
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