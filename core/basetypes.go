package core

import (
	"ckb-sdk-go/bincode"
	"encoding/hex"
	"reflect"
	"strings"
)

type Status string

const (
	Pending   Status = "pending"
	Proposed  Status = "proposed"
	Committed Status = "committed"
)

type RpcScriptHashType string

const (
	RpcData RpcScriptHashType = "Data"
	RpcType RpcScriptHashType = "Type"
)

type ScriptHashType uint32

const (
	Data ScriptHashType = 0
	Type ScriptHashType = 1
)

type DepType string

/*const (
	Code     ScriptHashType = "code"
	DepGroup ScriptHashType = "depGroup"
)*/

type CellStatus string

const (
	Live    CellStatus = "live"
	Unknown CellStatus = "unknown"
)

type CapacityUnit int

const (
	Shannon CapacityUnit = 1
	Byte    CapacityUnit = 100000000
)


type Witness [][]byte
type H256 [32]byte
type RpcH256 string
type U256 [32]byte
type RpcU256 string
var (
	ZeroH256 = H256{}
)

func (h256 H256) Bytes() []byte {
	return h256[:]
}
func (h256 *H256) SetBytes(h256Bytes []byte) {
	copy(h256[:], h256Bytes)
}

func StringToHash(str string) *H256 {
	str = strings.ReplaceAll(str,"0x","")
	hbytes,_ := hex.DecodeString(str)
	h256 := &H256{}
	h256.SetBytes(hbytes)
	return  h256
}

func (_ H256) UnMarshal(binCode *bincode.BinCodeDeSerizlize) (reflect.Value, error) {
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

func (_ H256) Marshal(binCode *bincode.BinCodeSerizlize, val reflect.Value) error {
	hexH256 := "0x" + hex.EncodeToString(val.Interface().(H256).Bytes())
	return binCode.SliceBytes(reflect.ValueOf([]byte(hexH256)))
}