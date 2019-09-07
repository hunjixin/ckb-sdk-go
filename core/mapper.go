package core

import (
	"encoding/hex"
	"github.com/hunjixin/automapper"
	"reflect"
	"strconv"
	"strings"
)
func init(){
	//string=>uint32
	automapper.MustCreateMapper("", uint32(0)).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			n, _ := strconv.ParseUint(str, 10, 32)
			destVal.Elem().SetUint(n)
		})
	//string=>uint64
	automapper.MustCreateMapper("", uint64(0)).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			n, _ := strconv.ParseUint(str, 10, 64)
			destVal.Elem().SetUint(n)
		})
	//string=>[]byte
	automapper.MustCreateMapper("", []byte{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			str = strings.ReplaceAll(str, "0x", "")
			sliceBytes, _ := hex.DecodeString(str)
			destVal.Elem().SetBytes(sliceBytes)
		})
	//string=>[10]byte
	automapper.MustCreateMapper("", [10]byte{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			str = strings.ReplaceAll(str, "0x", "")
			sliceBytes, _ := hex.DecodeString(str)
			destVal.Elem().SetBytes(sliceBytes[0:10])
		})
	//string=>[32]byte
	automapper.MustCreateMapper("", [32]byte{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			str = strings.ReplaceAll(str, "0x", "")
			sliceBytes, _ := hex.DecodeString(str)
			destVal.Elem().SetBytes(sliceBytes[0:32])
		})
	//string=>[]byte
	automapper.MustCreateMapper("", []byte{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			str = strings.ReplaceAll(str, "0x", "")
			sliceBytes, _ := hex.DecodeString(str)
			destVal.Elem().SetBytes(sliceBytes)
		})
	//uint32=>string
	automapper.MustCreateMapper(uint32(0), "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			number := sourceVal.(uint32)
			destVal.Elem().SetString(strconv.FormatUint(uint64(number), 10))
		})
	//uint64=>string
	automapper.MustCreateMapper(uint64(0), "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			number := sourceVal.(uint64)
			destVal.Elem().SetString(strconv.FormatUint(uint64(number), 10))
		})
	//H256=>string
	automapper.MustCreateMapper(H256{}, "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			h256 := sourceVal.(H256)
			destVal.Elem().SetString("0x"+hex.EncodeToString(h256[:]))
		})
	//U256=>string
	automapper.MustCreateMapper(U256{}, "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			h256 := sourceVal.(U256)
			destVal.Elem().SetString("0x"+hex.EncodeToString(h256[:]))
		})

	//ScriptHashType<=>RpcScriptHashType
	automapper.MustCreateMapper(ScriptHashType(0), RpcScriptHashType("")).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			hashType := sourceVal.(ScriptHashType)
			if hashType == 0 {
				destVal.Elem().SetString(string(RpcData))
			}else if hashType == 1 {
				destVal.Elem().SetString( string(RpcType))
			}else{
				//panic
			}

		})

	automapper.MustCreateMapper(RpcScriptHashType(""), ScriptHashType(0),).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			hashType := sourceVal.(RpcScriptHashType)
			if hashType == "Data" {
				destVal.Elem().Set(reflect.ValueOf(Data))
			}else if hashType == "Type" {
				destVal.Elem().Set(reflect.ValueOf(Type))
			}else{
				//panic
			}

		})
	//[]byte=>string
	automapper.MustCreateMapper([]byte{}, "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			bytes := sourceVal.([]byte)
			destVal.Elem().SetString("0x"+hex.EncodeToString(bytes))
		})

	//Witness=>RpcWitness
	automapper.MustCreateMapper(Witness{}, RpcWitness{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			witness := sourceVal.(Witness)
			rpcWitness := RpcWitness{
				Data: []string{},
			}
			for _, bytes := range  witness {
				rpcWitness.Data = append(rpcWitness.Data, "0x"+hex.EncodeToString(bytes))
			}
			destVal.Elem().Set(reflect.ValueOf(rpcWitness))
		})

	automapper.MustCreateMapper(RpcWitness{},Witness{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			rpcWitness := sourceVal.(RpcWitness)
			witness := Witness{}
			for _, rpcData := range  rpcWitness.Data {
				str := strings.ReplaceAll(rpcData, "0x", "")
				witnessBytes, _ := hex.DecodeString(str)
				witness = append(witness, witnessBytes)
			}
			destVal.Elem().Set(reflect.ValueOf(witness))
		})
}