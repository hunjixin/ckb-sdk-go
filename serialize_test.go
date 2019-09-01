package ckb_sdk_go

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Serialize(t *testing.T) {
	type Entity struct {
		X int32
		Y int32
	}

	world := []Entity{{0, 0}, {10, 20}}
	encoded, err := Marshal(world)
	if err != nil {
		t.Error(err)
	}
	if len(encoded) != 8+4*4 {
		t.Error("len not mathc")
	}
	val, err := UnMarshal(encoded, reflect.TypeOf([]Entity{}))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(val)

}
