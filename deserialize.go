package ckb_sdk_go

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"reflect"
)

type BinCodeDeSerizlize struct {
	buf *bytes.Buffer
}

func NewBinCodeDeSerizlize(b []byte) *BinCodeDeSerizlize {
	return &BinCodeDeSerizlize{
		buf: bytes.NewBuffer(b),
	}
}

func (binCode *BinCodeDeSerizlize) Bool() (reflect.Value, error) {
	b, err := binCode.buf.ReadByte()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if b == 0 {
		return reflect.ValueOf(false), nil
	} else {
		return reflect.ValueOf(true), nil
	}
}

func (binCode *BinCodeDeSerizlize) Int8() (reflect.Value, error) {
	b, err := binCode.buf.ReadByte()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(int8(b)), nil
}

func (binCode *BinCodeDeSerizlize) Int16() (reflect.Value, error) {
	val := binary.LittleEndian.Uint16(binCode.buf.Next(2))
	return reflect.ValueOf(int16(val)), nil
}

func (binCode *BinCodeDeSerizlize) Int32() (reflect.Value, error) {
	val := binary.LittleEndian.Uint32(binCode.buf.Next(4))
	return reflect.ValueOf(int32(val)), nil
}

func (binCode *BinCodeDeSerizlize) Int64() (reflect.Value, error) {
	val := binary.LittleEndian.Uint64(binCode.buf.Next(8))
	return reflect.ValueOf(int64(val)), nil
}

func (binCode *BinCodeDeSerizlize) Uint8() (reflect.Value, error) {
	b, err := binCode.buf.ReadByte()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(b), nil
}

func (binCode *BinCodeDeSerizlize) Uint16() (reflect.Value, error) {
	val := binary.LittleEndian.Uint16(binCode.buf.Next(2))
	return reflect.ValueOf(val), nil
}

func (binCode *BinCodeDeSerizlize) Uint32() (reflect.Value, error) {
	val := binary.LittleEndian.Uint32(binCode.buf.Next(4))
	return reflect.ValueOf(val), nil
}

func (binCode *BinCodeDeSerizlize) Uint64() (reflect.Value, error) {
	val := binary.LittleEndian.Uint64(binCode.buf.Next(8))
	return reflect.ValueOf(val), nil
}

func (binCode *BinCodeDeSerizlize) Float32() (reflect.Value, error) {
	val, err := binCode.Uint32()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	float32Val := math.Float32frombits(uint32(val.Uint()))
	return reflect.ValueOf(float32Val), nil
}

func (binCode *BinCodeDeSerizlize) Float64() (reflect.Value, error) {
	val, err := binCode.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	float64Val := math.Float64frombits(uint64(val.Uint()))
	return reflect.ValueOf(float64Val), nil
}

/*func(binCode *BinCodeSerizlize) Int128(b 128) {

}*/

func (binCode *BinCodeDeSerizlize) String_() (reflect.Value, error) {
	val, err := binCode.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := val.Uint()
	strBytes := make([]byte, len)
	_, err = binCode.buf.Read(strBytes)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(string(strBytes)), nil
}

func (binCode *BinCodeDeSerizlize) Array(t reflect.Type) (reflect.Value, error) {
	val, err := binCode.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := int(val.Uint())
	arrayType := reflect.ArrayOf(len, t)
	arrayVal := reflect.New(arrayType).Elem()
	for i := 0; i < len; i++ {
		val, err = binCode.UnMarshal(t)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		arrayVal.Index(i).Set(val)
	}

	return arrayVal, nil
}

func (binCode *BinCodeDeSerizlize) Slice(t reflect.Type) (reflect.Value, error) {
	val, err := binCode.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := int(val.Uint())
	sliceVal := reflect.MakeSlice(t, len, len)
	for i := 0; i < len; i++ {
		val, err = binCode.UnMarshal(t.Elem())
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		sliceVal.Index(i).Set(val)
	}

	return sliceVal, nil
}

func (binCode *BinCodeDeSerizlize) ArrayBytes() (reflect.Value, error) {
	val, err := binCode.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := val.Uint()

	arrayType := reflect.ArrayOf(int(len), reflect.TypeOf(byte(0)))
	arrayVal := reflect.New(arrayType).Elem()

	bytes := make([]byte, len)
	_, err = binCode.buf.Read(bytes)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	reflect.Copy(arrayVal, reflect.ValueOf(bytes))
	return arrayVal, nil
}

func (binCode *BinCodeDeSerizlize) SliceBytes() (reflect.Value, error) {
	val, err := binCode.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := val.Uint()
	bytes := make([]byte, len)
	_, err = binCode.buf.Read(bytes)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(bytes), nil
}

func (binCode *BinCodeDeSerizlize) Null() (reflect.Value, error) {
	return binCode.Bool()
}

func (binCode *BinCodeDeSerizlize) Struct(t reflect.Type) (reflect.Value, error) {
	valT := reflect.New(t).Elem()
	for i := 0; i < valT.NumField(); i++ {
		field := valT.Field(i)
		fmt.Println(t.Field(i).Name)
		if "Hash_type" == t.Field(i).Name{
			fmt.Println()
		}
		val, err := binCode.UnMarshal(field.Type())
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		if val.IsValid() {
			fmt.Println(val.Interface())
			field.Set(val)
		}
	}
	return valT, nil
}
func (binCode *BinCodeDeSerizlize) Map(t reflect.Type) (reflect.Value, error) {
	val, err := binCode.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := val.Uint()
	newMap := reflect.MakeMapWithSize(t, int(len))
	keyType := t.Key()
	valType := t.Elem()
	for i := 0; i < int(len); i++ {
		key, err := binCode.UnMarshal(keyType)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		val, err := binCode.UnMarshal(valType)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		newMap.SetMapIndex(key, val)
	}
	return newMap, nil
}

var (
	nilUnmarshal = reflect.TypeOf((*UnMarshaler)(nil)).Elem()
)

func (binCode *BinCodeDeSerizlize) UnMarshal(t reflect.Type) (reflect.Value, error) {
	var err error
	var val reflect.Value
	if t.Kind()!= reflect.Ptr&&t.Implements(nilUnmarshal) {
		val, err = reflect.New(t).Elem().Interface().(UnMarshaler).UnMarshal(binCode)
	} else {
		switch t.Kind() {
		case reflect.String:
			val, err = binCode.String_()
		case reflect.Array:
			if t.Elem().Kind() == reflect.Uint8 {
				val, err = binCode.ArrayBytes()
			} else {
				val, err = binCode.Array(t.Elem())
			}
		case reflect.Slice:
			if t.Elem().Kind() == reflect.Uint8 {
				val, err = binCode.SliceBytes()
			} else {
				val, err = binCode.Slice(t)
			}
		case reflect.Bool:
			val, err = binCode.Bool()
		case reflect.Float32:
			val, err = binCode.Float32()
		case reflect.Float64:
			val, err = binCode.Float64()
		case reflect.Int8:
			val, err = binCode.Int8()
		case reflect.Int16:
			val, err = binCode.Int16()
		case reflect.Int32:
			val, err = binCode.Int32()
		case reflect.Int64:
			val, err = binCode.Int64()
		case reflect.Uint8:
			val, err = binCode.Uint8()
		case reflect.Uint16:
			val, err = binCode.Uint16()
		case reflect.Uint32:
			val, err = binCode.Uint32()
		case reflect.Uint64:
			val, err = binCode.Uint64()
		case reflect.Struct:
			val, err = binCode.Struct(t)
		case reflect.Ptr:
			val, err = binCode.Uint8()
			if err == nil {
				if val.Uint() == 0 {
					val = reflect.ValueOf(nil)
				} else {
					elem := t.Elem()
					fmt.Println(t.Elem().Name())
					val, err = binCode.UnMarshal(elem)
					val = val.Addr()
					fmt.Println(val.Interface())
				}
			}
		default:
			err = errors.New("not supprot type")
		}
	}
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	//fmt.Println(val.Interface())
	return val, nil
}

func UnMarshal(data []byte, t reflect.Type) (interface{}, error) {
	des := NewBinCodeDeSerizlize(data)
	return des.UnMarshal(t)
}
