package ckb_sdk_go

import (
	"bytes"
	"ckb-sdk-go/bincode"
	"ckb-sdk-go/ckbserize"
	"ckb-sdk-go/core"
	"ckb-sdk-go/crypto"
	"encoding/hex"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"reflect"
	"strings"
	"testing"
)

func Test_UnMarshal(t *testing.T) {
	rawTxStr := "170000000000000000000000010000000000000001420000000000000030783030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303000000000007B0000000000000001000000000000000088526A740000000300000000000000010203000000000000000042000000000000003078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030300000000000"
	rawTxBytes, _ := hex.DecodeString(rawTxStr)
	tRawTx := reflect.TypeOf(core.RawTransaction{})
	dddd, err := bincode.UnMarshal(rawTxBytes, tRawTx)
	if err != nil {
		t.Error(err)
	}
	rawTx := dddd.(core.RawTransaction)
	if rawTx.Inputs[0].Since != 123 {
		t.Errorf("expect since %d but got %d", 123, rawTx.Inputs[0].Since)
	}

	marshalBytes, err := bincode.Marshal(dddd)
	if err != nil {
		t.Error(err)
	}
	marshalStr := hex.EncodeToString(marshalBytes)
	if marshalStr != strings.ToLower(rawTxStr) {
		t.Errorf("expect equal but got %s", marshalStr)
	}
}
func Test_TxHash(t *testing.T) {
	arg, _ := hex.DecodeString("59a27ef3ba84f061517d13f42cf44ed020610061")
	builder := NewTransactionBuilder()
	builder.
		SetVersion(0).
		AppendCellDeps(core.CellDep{
			Out_point: core.OutPoint{
				*core.StringToHash("c12386705b5cbb312b693874f3edf45c43a274482e27b8df0fd80c8d3f5feb8b"),
				0,
			},
			Dep_type: core.DepGroup,
		}).
		AppendCellDeps(core.CellDep{
			Out_point: core.OutPoint{
				*core.StringToHash("0fb4945d52baf91e0dee2a686cdd9d84cad95b566a1d7409b970ee0a0f364f60"),
				2,
			},
			Dep_type: core.Code,
		}).
		AppendOutput(core.CellOutput{
			Capacity: 100000000000, //5000 * 00000000
			Lock: core.Script{
				Args:      [][]byte{arg},
				Code_hash: *core.StringToHash("68d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e88"),
				Hash_type: core.Type,
			},
			Type_: &core.Script{
				Args:      [][]byte{},
				Code_hash: *core.StringToHash("0xece45e0979030e2f8909f76258631c42333b1e906fd9701ec3600a464a90b8f6"),
				Hash_type: core.Data,
			},
		}).
		AppendOutput(core.CellOutput{
			Capacity: 98824000000000, //5000 * 00000000
			Lock: core.Script{
				Args:      [][]byte{arg},
				Code_hash: *core.StringToHash("68d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e88"),
				Hash_type: core.Type,
			},
			Type_: nil,
		}).
		AppendInput(core.CellInput{
			Previous_output: core.OutPoint{
				Tx_hash: *core.StringToHash("31f695263423a4b05045dd25ce6692bb55d7bba2965d8be16b036e138e72cc65"),
				Index:   1,
			},
			Since: 0,
		}).
		AppendOutputData([]byte{}).
		AppendOutputData([]byte{})
	expectSerize := "130100000c000000aa0000009e00000010000000180000006900000000e87648170000005100000010000000300000003100000068d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e880120000000080000001400000059a27ef3ba84f061517d13f42cf44ed02061006135000000100000003000000031000000ece45e0979030e2f8909f76258631c42333b1e906fd9701ec3600a464a90b8f600040000006900000010000000180000006900000000506a41e15900005100000010000000300000003100000068d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e880120000000080000001400000059a27ef3ba84f061517d13f42cf44ed020610061"
	serilizeByte, _ := hex.DecodeString(expectSerize)

	tx := builder.Build()
    rawTx := tx.RawTransaction()

	bytes2 := ckbserize.SerializeOutputs(rawTx.Outputs)
	fmt.Println(hex.EncodeToString(bytes2))
	if !bytes.Equal(bytes2, serilizeByte)  {
		fmt.Println(bytes2)
		fmt.Println(serilizeByte)
		t.Error("serize error")
	}

	/*expectedHash := "0x9d1bf801b235ce62812844f01381a070c0cc72876364861e00492eac1d8b54e7"
	hash := tx.TxHash()
	hashStr := hex.EncodeToString(hash[:])
	if hashStr != expectedHash {
		t.Errorf("expect hash  %s but got %s", expectedHash, hashStr)
	}*/
}

func TestSigWitness(t *testing.T) {
	hashBytes, _ := hex.DecodeString("ac1bb95455cdfb89b6e977568744e09b6b80e08cab9477936a09c4ca07f5b8ab")
	privBytes, _ := hex.DecodeString("e79f3207ea4980b7fed79956d5934249ceac4751a4fae01a0f7c4a96884bc4e3")
	priv, _ := secp256k1.PrivKeyFromScalar(privBytes)
	sig := crypto.SignMesage(hashBytes, priv)

	sigStr := hex.EncodeToString(sig)
	expectedSig := "2c643579e47045be050d3842ed9270151af8885e33954bddad0e53e81d1c2dbe2dc637877a8302110846ebc6a16d9148c106e25f945063ad1c4d4db2b695240800"
	if sigStr != expectedSig {
		t.Errorf("expect sig %s but got %s", expectedSig, sigStr)
	}
}



func Test_Script(t *testing.T) {
		arg, _ := hex.DecodeString("59a27ef3ba84f061517d13f42cf44ed020610061")
		builder := NewTransactionBuilder()
		builder.
			SetVersion(0).
			AppendCellDeps(core.CellDep{
				Out_point: core.OutPoint{
					*core.StringToHash("c12386705b5cbb312b693874f3edf45c43a274482e27b8df0fd80c8d3f5feb8b"),
					0,
				},
				Dep_type: core.DepGroup,
			}).
			AppendCellDeps(core.CellDep{
				Out_point: core.OutPoint{
					*core.StringToHash("0fb4945d52baf91e0dee2a686cdd9d84cad95b566a1d7409b970ee0a0f364f60"),
					2,
				},
				Dep_type: core.Code,
			}).
			AppendOutput(core.CellOutput{
				Capacity: 100000000000, //5000 * 00000000
				Lock: core.Script{
					Args:      [][]byte{arg},
					Code_hash: *core.StringToHash("68d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e88"),
					Hash_type: core.Type,
				},
				Type_: &core.Script{
					Args:      [][]byte{},
					Code_hash: *core.StringToHash("0xece45e0979030e2f8909f76258631c42333b1e906fd9701ec3600a464a90b8f6"),
					Hash_type: core.Data,
				},
			}).
			AppendOutput(core.CellOutput{
				Capacity: 98824000000000, //5000 * 00000000
				Lock: core.Script{
					Args:      [][]byte{arg},
					Code_hash: *core.StringToHash("68d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e88"),
					Hash_type: core.Type,
				},
				Type_: nil,
			}).
			AppendInput(core.CellInput{
				Previous_output: core.OutPoint{
					Tx_hash: *core.StringToHash("31f695263423a4b05045dd25ce6692bb55d7bba2965d8be16b036e138e72cc65"),
					Index:   1,
				},
				Since: 0,
			}).
			AppendOutputData([]byte{}).
			AppendOutputData([]byte{})
		expectSerize := "5100000010000000300000003100000068d5438ac952d2f584abf879527946a537e82c7f3c1cbf6d8ebf9767437d8e880120000000080000001400000059a27ef3ba84f061517d13f42cf44ed020610061"
		serilizeByte, _ := hex.DecodeString(expectSerize)
		tx := builder.Build()
		rawTx := tx.RawTransaction()

		bytes2 := ckbserize.SerializeScript(rawTx.Outputs[0].Lock)
		fmt.Println(hex.EncodeToString(bytes2))
		if !bytes.Equal(bytes2, serilizeByte)  {
			fmt.Println(bytes2)
			fmt.Println(serilizeByte)
			t.Error("serize error")
		}

}
