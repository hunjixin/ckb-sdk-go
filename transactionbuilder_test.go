package ckb_sdk_go

import (
	"encoding/hex"
	"testing"
)

func Test_TxHash(t *testing.T) {
	builder := new(TransactionBuilder).
		AppendOutput(CellOutput{
			Capacity:5000,
			Data:[]byte{1,2,3},
			Lock:Script{},
			Type_:nil,
	}).
	AppendInput(CellInput{
		Previous_output:OutPoint{
			Block_hash:nil,
			Cell: &CellOutPoint{
				Tx_hash: [32]byte{},
				Index:0,
			},

		},
		Since:0,
	}).
	AppendWitness([][]byte{[]byte{7, 8, 9}}).
	Build()
	expected := "6e9d9e6a6d5be5adafe7eac9f159b439cf4a4a400400cf98c231a341eb318bc2"
	hashStr := hex.EncodeToString(builder.TxHash()[:])
	if hashStr != expected {
		t.Errorf("expect %s but got %s", expected, hashStr)
	}
}
