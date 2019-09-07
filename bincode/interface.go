package bincode

import "reflect"

type Marshaler interface {
	Marshal(serilize *BinCodeSerizlize, val reflect.Value) error
}

type UnMarshaler interface {
	UnMarshal(deserizlize *BinCodeDeSerizlize) (reflect.Value, error)
}

