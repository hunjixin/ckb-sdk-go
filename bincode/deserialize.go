package bincode

import (
	"bytes"
	"encoding/binary"
	"errors"
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

func (deSerizlize *BinCodeDeSerizlize) Bool() (reflect.Value, error) {
	b, err := deSerizlize.buf.ReadByte()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if b == 0 {
		return reflect.ValueOf(false), nil
	} else {
		return reflect.ValueOf(true), nil
	}
}

func (deSerizlize *BinCodeDeSerizlize) Int8() (reflect.Value, error) {
	b, err := deSerizlize.buf.ReadByte()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(int8(b)), nil
}

func (deSerizlize *BinCodeDeSerizlize) Int16() (reflect.Value, error) {
	val := binary.LittleEndian.Uint16(deSerizlize.buf.Next(2))
	return reflect.ValueOf(int16(val)), nil
}

func (deSerizlize *BinCodeDeSerizlize) Int32() (reflect.Value, error) {
	val := binary.LittleEndian.Uint32(deSerizlize.buf.Next(4))
	return reflect.ValueOf(int32(val)), nil
}

func (deSerizlize *BinCodeDeSerizlize) Int64() (reflect.Value, error) {
	val := binary.LittleEndian.Uint64(deSerizlize.buf.Next(8))
	return reflect.ValueOf(int64(val)), nil
}

func (deSerizlize *BinCodeDeSerizlize) Uint8() (reflect.Value, error) {
	b, err := deSerizlize.buf.ReadByte()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(b), nil
}

func (deSerizlize *BinCodeDeSerizlize) Uint16() (reflect.Value, error) {
	val := binary.LittleEndian.Uint16(deSerizlize.buf.Next(2))
	return reflect.ValueOf(val), nil
}

func (deSerizlize *BinCodeDeSerizlize) Uint32() (reflect.Value, error) {
	val := binary.LittleEndian.Uint32(deSerizlize.buf.Next(4))
	return reflect.ValueOf(val), nil
}

func (deSerizlize *BinCodeDeSerizlize) Uint64() (reflect.Value, error) {
	val := binary.LittleEndian.Uint64(deSerizlize.buf.Next(8))
	return reflect.ValueOf(val), nil
}

func (deSerizlize *BinCodeDeSerizlize) Float32() (reflect.Value, error) {
	val, err := deSerizlize.Uint32()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	float32Val := math.Float32frombits(uint32(val.Uint()))
	return reflect.ValueOf(float32Val), nil
}

func (deSerizlize *BinCodeDeSerizlize) Float64() (reflect.Value, error) {
	val, err := deSerizlize.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	float64Val := math.Float64frombits(uint64(val.Uint()))
	return reflect.ValueOf(float64Val), nil
}

/*func(binCode *BinCodeSerizlize) Int128(b 128) {

}*/

func (deSerizlize *BinCodeDeSerizlize) String_() (reflect.Value, error) {
	val, err := deSerizlize.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := val.Uint()
	strBytes := make([]byte, len)
	_, err = deSerizlize.buf.Read(strBytes)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(string(strBytes)), nil
}

func (deSerizlize *BinCodeDeSerizlize) Array(t reflect.Type) (reflect.Value, error) {
	val, err := deSerizlize.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := int(val.Uint())
	arrayType := reflect.ArrayOf(len, t)
	arrayVal := reflect.New(arrayType).Elem()
	for i := 0; i < len; i++ {
		val, err = deSerizlize.UnMarshal(t)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		arrayVal.Index(i).Set(val)
	}

	return arrayVal, nil
}

func (deSerizlize *BinCodeDeSerizlize) Slice(t reflect.Type) (reflect.Value, error) {
	val, err := deSerizlize.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := int(val.Uint())
	sliceVal := reflect.MakeSlice(t, len, len)
	for i := 0; i < len; i++ {
		val, err = deSerizlize.UnMarshal(t.Elem())
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		sliceVal.Index(i).Set(val)
	}

	return sliceVal, nil
}

func (deSerizlize *BinCodeDeSerizlize) ArrayBytes() (reflect.Value, error) {
	val, err := deSerizlize.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := val.Uint()

	arrayType := reflect.ArrayOf(int(len), reflect.TypeOf(byte(0)))
	arrayVal := reflect.New(arrayType).Elem()

	bytes := make([]byte, len)
	_, err = deSerizlize.buf.Read(bytes)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	reflect.Copy(arrayVal, reflect.ValueOf(bytes))
	return arrayVal, nil
}

func (deSerizlize *BinCodeDeSerizlize) SliceBytes() (reflect.Value, error) {
	val, err := deSerizlize.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := val.Uint()
	bytes := make([]byte, len)
	_, err = deSerizlize.buf.Read(bytes)
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return reflect.ValueOf(bytes), nil
}

func (deSerizlize *BinCodeDeSerizlize) Null() (reflect.Value, error) {
	return deSerizlize.Bool()
}

func (deSerizlize *BinCodeDeSerizlize) Struct(t reflect.Type) (reflect.Value, error) {
	valT := reflect.New(t).Elem()
	for i := 0; i < valT.NumField(); i++ {
		field := valT.Field(i)
		val, err := deSerizlize.UnMarshal(field.Type())
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		if val.IsValid() {
			field.Set(val.Convert(field.Type()))
		}
	}
	return valT, nil
}
func (deSerizlize *BinCodeDeSerizlize) Map(t reflect.Type) (reflect.Value, error) {
	val, err := deSerizlize.Uint64()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	len := val.Uint()
	newMap := reflect.MakeMapWithSize(t, int(len))
	keyType := t.Key()
	valType := t.Elem()
	for i := 0; i < int(len); i++ {
		key, err := deSerizlize.UnMarshal(keyType)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		val, err := deSerizlize.UnMarshal(valType)
		if err != nil {
			return reflect.ValueOf(nil), err
		}
		newMap.SetMapIndex(key, val)
	}
	return newMap, nil
}

var (
	nilUnmarshal = reflect.TypeOf((*UnMarshaler)(nil)).Elem()
	nilMarshal   = reflect.TypeOf((*Marshaler)(nil)).Elem()
	TString      = reflect.TypeOf("")
	TUint8       = reflect.TypeOf(uint8(0))
	TUint16      = reflect.TypeOf(uint16(0))
	TUint32      = reflect.TypeOf(uint32(0))
	TUint64      = reflect.TypeOf(uint64(0))

	TInt8    = reflect.TypeOf(int8(0))
	TInt16   = reflect.TypeOf(int16(0))
	TInt32   = reflect.TypeOf(int32(0))
	TInt64   = reflect.TypeOf(int64(0))
	TBool    = reflect.TypeOf(false)
	TFloat32 = reflect.TypeOf(float32(0))
	TFloat64 = reflect.TypeOf(float64(0))
)

func (deSerizlize *BinCodeDeSerizlize) UnMarshal(t reflect.Type) (reflect.Value, error) {
	var err error
	var val reflect.Value
	if t.Kind() != reflect.Ptr && t.Implements(nilUnmarshal) {
		val, err = val.Interface().(UnMarshaler).UnMarshal(deSerizlize)
	} else {
		switch t.Kind() {
		case reflect.String:
			val, err = deSerizlize.String_()
		case reflect.Array:
			if t.Elem().Kind() == reflect.Uint8 {
				val, err = deSerizlize.ArrayBytes()
			} else {
				val, err = deSerizlize.Array(t.Elem())
			}
		case reflect.Slice:
			if t.Elem().Kind() == reflect.Uint8 {
				val, err = deSerizlize.SliceBytes()
			} else {
				val, err = deSerizlize.Slice(t)
			}
		case reflect.Bool:
			val, err = deSerizlize.Bool()
		case reflect.Float32:
			val, err = deSerizlize.Float32()
		case reflect.Float64:
			val, err = deSerizlize.Float64()
		case reflect.Int8:
			val, err = deSerizlize.Int8()
		case reflect.Int16:
			val, err = deSerizlize.Int16()
		case reflect.Int32:
			val, err = deSerizlize.Int32()
		case reflect.Int64:
			val, err = deSerizlize.Int64()
		case reflect.Uint8:
			val, err = deSerizlize.Uint8()
		case reflect.Uint16:
			val, err = deSerizlize.Uint16()
		case reflect.Uint32:
			val, err = deSerizlize.Uint32()
		case reflect.Uint64:
			val, err = deSerizlize.Uint64()
		case reflect.Struct:
			val, err = deSerizlize.Struct(t)
		case reflect.Ptr:
			val, err = deSerizlize.Uint8()
			if err == nil {
				if val.Uint() == 0 {
					val = reflect.ValueOf(nil)
				} else {
					elem := t.Elem()
					val, err = deSerizlize.UnMarshal(elem)
					val = val.Addr()
				}
			}
		default:
			err = errors.New("not supprot type")
		}
	}
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	return val, nil
}

func UnMarshal(data []byte, t reflect.Type) (interface{}, error) {
	des := NewBinCodeDeSerizlize(data)
	refVal, err := des.UnMarshal(t)
	if err != nil {
		return nil, err
	}
	return refVal.Interface(), nil
}
