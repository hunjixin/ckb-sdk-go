package bincode

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

func (serizlize *BinCodeSerizlize) Grow(b int) {
	newBytes := make([]byte, len(serizlize.bytes)+b)
	copy(newBytes, serizlize.bytes)
	serizlize.bytes = newBytes
}

func (serizlize *BinCodeSerizlize) Bool(b reflect.Value) error {
	serizlize.Grow(1)
	if b.Bool() {
		serizlize.bytes[serizlize.offset] = byte(1)
	} else {
		serizlize.bytes[serizlize.offset] = byte(0)
	}
	serizlize.offset++
	return nil
}

func (serizlize *BinCodeSerizlize) Int8(b reflect.Value) error {
	serizlize.Grow(1)
	serizlize.bytes[serizlize.offset] = byte(int8(b.Int()))
	serizlize.offset++
	return nil
}

func (serizlize *BinCodeSerizlize) Int16(b reflect.Value) error {
	serizlize.Grow(2)
	binary.LittleEndian.PutUint16(serizlize.bytes[serizlize.offset:], uint16(b.Int()))
	serizlize.offset = serizlize.offset + 2
	return nil
}

func (serizlize *BinCodeSerizlize) Int32(b reflect.Value) error {
	serizlize.Grow(4)
	binary.LittleEndian.PutUint32(serizlize.bytes[serizlize.offset:], uint32(b.Int()))
	serizlize.offset = serizlize.offset + 4
	return nil
}

func (serizlize *BinCodeSerizlize) Int64(b reflect.Value) error {
	serizlize.Grow(8)
	binary.LittleEndian.PutUint64(serizlize.bytes[serizlize.offset:], uint64(b.Int()))
	serizlize.offset = serizlize.offset + 8
	return nil
}

func (serizlize *BinCodeSerizlize) Uint8(b reflect.Value) error {
	serizlize.Grow(1)
	serizlize.bytes[serizlize.offset] = uint8(b.Uint())
	serizlize.offset++
	return nil
}

func (serizlize *BinCodeSerizlize) Uint16(b reflect.Value) error {
	serizlize.Grow(2)
	binary.LittleEndian.PutUint16(serizlize.bytes[serizlize.offset:], uint16(b.Uint()))
	serizlize.offset = serizlize.offset + 2
	return nil
}

func (serizlize *BinCodeSerizlize) Uint32(b reflect.Value) error {
	serizlize.Grow(4)
	binary.LittleEndian.PutUint32(serizlize.bytes[serizlize.offset:], uint32(b.Uint()))
	serizlize.offset = serizlize.offset + 4
	return nil
}

func (serizlize *BinCodeSerizlize) Uint64(b reflect.Value) error {
	serizlize.Grow(8)
	binary.LittleEndian.PutUint64(serizlize.bytes[serizlize.offset:], b.Uint())
	serizlize.offset = serizlize.offset + 8
	return nil
}

func (serizlize *BinCodeSerizlize) Float32(b reflect.Value) error {
	float32Val := float32(b.Float())
	return serizlize.Uint32(reflect.ValueOf(math.Float32bits(float32Val)))
}

func (serizlize *BinCodeSerizlize) Float64(b reflect.Value) error {
	return serizlize.Uint64(reflect.ValueOf(math.Float64bits(b.Float())))
}

/*func(binCode *BinCodeSerizlize) Int128(b 128) {

}*/

func (serizlize *BinCodeSerizlize) String_(b reflect.Value) error {
	serizlize.SliceBytes(reflect.ValueOf([]byte(b.String())))
	return nil
}

func (serizlize *BinCodeSerizlize) Map(b reflect.Value) error {
	len := b.Len()
	serizlize.Uint64(reflect.ValueOf(uint64(len)))
	iter := b.MapRange()
	for iter.Next() {
		serizlize.Marshal(iter.Key())
		serizlize.Marshal(iter.Value())
	}
	return nil
}

func (serizlize *BinCodeSerizlize) ArrayBytes(b reflect.Value) error {
	len := b.Len()
	sliceBytes := b.Slice(0, len)
	return serizlize.SliceBytes(sliceBytes)
}
func (serizlize *BinCodeSerizlize) SliceBytes(b reflect.Value) error {
	if b.IsNil() {
		serizlize.Uint64(reflect.ValueOf(uint64(0)))
	} else {
		len := b.Len()
		serizlize.Uint64(reflect.ValueOf(uint64(len)))
		serizlize.Grow(len)
		copy(serizlize.bytes[serizlize.offset:], b.Bytes())
		serizlize.offset = serizlize.offset + len
	}
	return nil
}

func (serizlize *BinCodeSerizlize) Array(b reflect.Value) error {
	len := b.Len()
	serizlize.Uint64(reflect.ValueOf(uint64(len)))
	for i := 0; i <= len; i++ {
		err := serizlize.Marshal(b.Index(i))
		if err != nil {
			return err
		}
	}
	return nil
}

func (serizlize *BinCodeSerizlize) Slice(b reflect.Value) error {
	if b.IsNil() {
		serizlize.Uint64(reflect.ValueOf(uint64(0)))
	} else {
		len := b.Len()
		serizlize.Uint64(reflect.ValueOf(uint64(len)))
		for i := 0; i < len; i++ {
			err := serizlize.Marshal(b.Index(i))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (serizlize *BinCodeSerizlize) Null() error {
	serizlize.Grow(1)
	serizlize.bytes[serizlize.offset] = 0
	serizlize.offset++
	return nil
}

func (serizlize *BinCodeSerizlize) Marshal(a reflect.Value) error {
	var err error

	if a.Kind() != reflect.Ptr && a.Type().Implements(nilMarshal) {
		err =  a.Interface().(Marshaler).Marshal(serizlize, a)
	} else {
		switch a.Type().Kind() {
		case reflect.String:
			err = serizlize.String_(a)
		case reflect.Array:
			if a.Type().Elem().Kind() == reflect.Uint8 {
				err = serizlize.ArrayBytes(a)
			} else {
				err = serizlize.Array(a)
			}
		case reflect.Slice:
			if a.Type().Elem().Kind() == reflect.Uint8 {
				err = serizlize.SliceBytes(a)
			} else {
				err = serizlize.Slice(a)
			}
		case reflect.Bool:
			err = serizlize.Bool(a)
		case reflect.Float32:
			err = serizlize.Float32(a)
		case reflect.Float64:
			err = serizlize.Float64(a)
		case reflect.Int8:
			err = serizlize.Int8(a)
		case reflect.Int16:
			err = serizlize.Int16(a)
		case reflect.Int32:
			err = serizlize.Int32(a)
		case reflect.Int64:
			err = serizlize.Int64(a)
		case reflect.Uint8:
			err = serizlize.Uint8(a)
		case reflect.Uint16:
			err = serizlize.Uint16(a)
		case reflect.Uint32:
			err = serizlize.Uint32(a)
		case reflect.Uint64:
			err = serizlize.Uint64(a)
		case reflect.Struct:
			err = serizlize.Struct(a)
		case reflect.Ptr:
			if a.IsNil() {
				err = serizlize.Int8(reflect.ValueOf(0))

			} else {
				err = serizlize.Int8(reflect.ValueOf(1))
				if err != nil {
					return err
				}
				err = serizlize.Marshal(a.Elem())
			}
		default:
			err = errors.New("not supprot type")
		}
	}

	return err
}
func (serizlize *BinCodeSerizlize) Struct(a reflect.Value) error {
	var err error
	for i := 0; i < a.NumField(); i++ {
		field := a.Field(i)
		err = serizlize.Marshal(field)
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
