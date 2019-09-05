package ckb_sdk_go

import (
	"encoding/hex"
	"github.com/hunjixin/automapper"
	"reflect"
	"strconv"
)
func init(){
	automapper.MustCreateMapper("", uint32(0)).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			n, _ := strconv.ParseUint(str, 10, 32)
			destVal.Elem().SetUint(n)
		})
	automapper.MustCreateMapper("", uint64(0)).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			n, _ := strconv.ParseUint(str, 10, 64)
			destVal.Elem().SetUint(n)
		})
	automapper.MustCreateMapper("", []byte{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			sliceBytes, _ := hex.DecodeString(str)
			destVal.Elem().SetBytes(sliceBytes)
		})

	automapper.MustCreateMapper("", [10]byte{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			sliceBytes, _ := hex.DecodeString(str)
			destVal.Elem().SetBytes(sliceBytes[0:10])
		})
	automapper.MustCreateMapper("", [32]byte{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			sliceBytes, _ := hex.DecodeString(str)
			destVal.Elem().SetBytes(sliceBytes[0:32])
		})
	automapper.MustCreateMapper("", []byte{}).
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			str := sourceVal.(string)
			sliceBytes, _ := hex.DecodeString(str)
			destVal.Elem().SetBytes(sliceBytes)
		})
	automapper.MustCreateMapper(uint32(0), "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			number := sourceVal.(uint32)
			destVal.Elem().SetString(strconv.FormatUint(uint64(number), 10))
		})

	automapper.MustCreateMapper(H256{}, "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			h256 := sourceVal.(H256)
			destVal.Elem().SetString("0x"+hex.EncodeToString(h256[:]))
		})

	automapper.MustCreateMapper(U256{}, "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			h256 := sourceVal.(U256)
			destVal.Elem().SetString("0x"+hex.EncodeToString(h256[:]))
		})

	automapper.MustCreateMapper(uint64(0), "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			number := sourceVal.(uint64)
			destVal.Elem().SetString(strconv.FormatUint(uint64(number), 10))
		})

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
	automapper.MustCreateMapper([]byte{}, "").
		Mapping(func(destVal reflect.Value, sourceVal interface{}) {
			bytes := sourceVal.([]byte)
			destVal.Elem().SetString("0x"+hex.EncodeToString(bytes))
		})
}