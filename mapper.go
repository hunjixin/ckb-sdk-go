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
}