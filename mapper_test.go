package ckb_sdk_go

import (
	"fmt"
	"github.com/hunjixin/automapper"
	"reflect"
	"testing"
)

func Test_BlockNumberConvert(t *testing.T) {
	xx := automapper.MustMapper("1231", reflect.TypeOf(uint32(0)))
	fmt.Println(xx)
}
