package ckb_sdk_go

import (
	"encoding/binary"
	"errors"
	"math"
	"reflect"
)

type BinCodeSerizlize struct {
	bytes  []byte
	offset int
}

func NewBinCodeSerizlize() *BinCodeSerizlize {
	return &BinCodeSerizlize{
		bytes:  []byte{},
		offset: 0,
	}
}

func (binCode *BinCodeSerizlize) Grow(b int) {
	newBytes := make([]byte, len(binCode.bytes)+b)
	copy(newBytes, binCode.bytes)
	binCode.bytes = newBytes
}

func (binCode *BinCodeSerizlize) Bool(b reflect.Value) error {
	binCode.Grow(1)
	if b.Bool() {
		binCode.bytes[binCode.offset] = byte(1)
	} else {
		binCode.bytes[binCode.offset] = byte(0)
	}
	binCode.offset++
	return nil
}

func (binCode *BinCodeSerizlize) Int8(b reflect.Value) error {
	binCode.Grow(1)
	binCode.bytes[binCode.offset] = byte(int8(b.Int()))
	binCode.offset++
	return nil
}

func (binCode *BinCodeSerizlize) Int16(b reflect.Value) error {
	binCode.Grow(2)
	binary.LittleEndian.PutUint16(binCode.bytes[binCode.offset:], uint16(b.Int()))
	binCode.offset = binCode.offset + 2
	return nil
}

func (binCode *BinCodeSerizlize) Int32(b reflect.Value) error {
	binCode.Grow(4)
	binary.LittleEndian.PutUint32(binCode.bytes[binCode.offset:], uint32(b.Int()))
	binCode.offset = binCode.offset + 4
	return nil
}

func (binCode *BinCodeSerizlize) Int64(b reflect.Value) error {
	binCode.Grow(8)
	binary.LittleEndian.PutUint64(binCode.bytes[binCode.offset:], uint64(b.Int()))
	binCode.offset = binCode.offset + 8
	return nil
}

func (binCode *BinCodeSerizlize) Uint8(b reflect.Value) error {
	binCode.Grow(1)
	binCode.bytes[binCode.offset] = uint8(b.Uint())
	binCode.offset++
	return nil
}

func (binCode *BinCodeSerizlize) Uint16(b reflect.Value) error {
	binCode.Grow(2)
	binary.LittleEndian.PutUint16(binCode.bytes[binCode.offset:], uint16(b.Uint()))
	binCode.offset = binCode.offset + 2
	return nil
}

func (binCode *BinCodeSerizlize) Uint32(b reflect.Value) error {
	binCode.Grow(4)
	binary.LittleEndian.PutUint32(binCode.bytes[binCode.offset:], uint32(b.Uint()))
	binCode.offset = binCode.offset + 4
	return nil
}

func (binCode *BinCodeSerizlize) Uint64(b reflect.Value) error {
	binCode.Grow(8)
	binary.LittleEndian.PutUint64(binCode.bytes[binCode.offset:], b.Uint())
	binCode.offset = binCode.offset + 8
	return nil
}

func (binCode *BinCodeSerizlize) Float32(b reflect.Value) error {
	float32Val := float32(b.Float())
	return binCode.Uint32(reflect.ValueOf(math.Float32bits(float32Val)))
}

func (binCode *BinCodeSerizlize) Float64(b reflect.Value) error {
	return binCode.Uint64(reflect.ValueOf(math.Float64bits(b.Float())))
}

/*func(binCode *BinCodeSerizlize) Int128(b 128) {

}*/

func (binCode *BinCodeSerizlize) String_(b reflect.Value) error {
	binCode.SliceBytes(reflect.ValueOf([]byte(b.String())))
	return nil
}

func (binCode *BinCodeSerizlize) Map(b reflect.Value) error {
	len := b.Len()
	binCode.Uint64(reflect.ValueOf(uint64(len)))
	iter := b.MapRange()
	for iter.Next() {
		binCode.Marshal(iter.Key())
		binCode.Marshal(iter.Value())
	}
	return nil
}

func (binCode *BinCodeSerizlize) ArrayBytes(b reflect.Value) error {
	len := b.Len()
	sliceBytes := b.Slice(0, len)
	return binCode.SliceBytes(sliceBytes)
}
func (binCode *BinCodeSerizlize) SliceBytes(b reflect.Value) error {
	if b.IsNil() {
		binCode.Uint64(reflect.ValueOf(uint64(0)))
	} else {
		len := b.Len()
		binCode.Uint64(reflect.ValueOf(uint64(len)))
		binCode.Grow(len)
		copy(binCode.bytes[binCode.offset:], b.Bytes())
		binCode.offset = binCode.offset + len
	}
	return nil
}

func (binCode *BinCodeSerizlize) Array(b reflect.Value) error {
	len := b.Len()
	binCode.Uint64(reflect.ValueOf(uint64(len)))
	for i := 0; i <= len; i++ {
		err := binCode.Marshal(b.Index(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func (binCode *BinCodeSerizlize) Slice(b reflect.Value) error {
	if b.IsNil() {
		binCode.Uint64(reflect.ValueOf(uint64(0)))
	} else {
		len := b.Len()
		binCode.Uint64(reflect.ValueOf(uint64(len)))
		for i := 0; i < len; i++ {
			err := binCode.Marshal(b.Index(i))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (binCode *BinCodeSerizlize) Null() error {
	binCode.Grow(1)
	binCode.bytes[binCode.offset] = 0
	binCode.offset++
	return nil
}

func (binCode *BinCodeSerizlize) Marshal(a reflect.Value) error {
	var err error

	if a.Kind() != reflect.Ptr && a.Type().Implements(nilMarshal) {
		err = a.Interface().(Marshaler).Marshal(binCode, a)
	} else {
		switch a.Type().Kind() {
		case reflect.String:
			err = binCode.String_(a)
		case reflect.Array:
			if a.Type().Elem().Kind() == reflect.Uint8 {
				err = binCode.ArrayBytes(a)
			} else {
				err = binCode.Array(a)
			}
		case reflect.Slice:
			if a.Type().Elem().Kind() == reflect.Uint8 {
				err = binCode.SliceBytes(a)
			} else {
				err = binCode.Slice(a)
			}
		case reflect.Bool:
			err = binCode.Bool(a)
		case reflect.Float32:
			err = binCode.Float32(a)
		case reflect.Float64:
			err = binCode.Float64(a)
		case reflect.Int8:
			err = binCode.Int8(a)
		case reflect.Int16:
			err = binCode.Int16(a)
		case reflect.Int32:
			err = binCode.Int32(a)
		case reflect.Int64:
			err = binCode.Int64(a)
		case reflect.Uint8:
			err = binCode.Uint8(a)
		case reflect.Uint16:
			err = binCode.Uint16(a)
		case reflect.Uint32:
			err = binCode.Uint32(a)
		case reflect.Uint64:
			err = binCode.Uint64(a)
		case reflect.Struct:
			err = binCode.Struct(a)
		case reflect.Ptr:
			if a.IsNil() {
				err = binCode.Int8(reflect.ValueOf(0))

			} else {
				err = binCode.Int8(reflect.ValueOf(1))
				if err != nil {
					return err
				}
				err = binCode.Marshal(a.Elem())
			}
		default:
			err = errors.New("not supprot type")
		}
	}

	return err
}
func (binCode *BinCodeSerizlize) Struct(a reflect.Value) error {
	var err error
	for i := 0; i < a.NumField(); i++ {
		field := a.Field(i)
		err = binCode.Marshal(field)
		if err != nil {
			return err
		}
	}
	return nil
}

func Marshal(val interface{}) ([]byte, error) {
	seriz := NewBinCodeSerizlize()
	err := seriz.Marshal(reflect.Indirect(reflect.ValueOf(val)))
	if err != nil {
		return nil, err
	}
	return seriz.bytes, nil
}
