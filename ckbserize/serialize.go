package ckbserize

import (
	"ckb-sdk-go/core"
	"encoding/binary"
)

const (
	offsetSize     = 4
	fullLengthSize = 4
)

func toHexInLittleEndian(number uint64, paddingType int) []byte {
	if paddingType == 4 {
		bytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, uint32(number))
		return bytes
	} else if paddingType == 8 {
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, number)
		return bytes
	} else {
		panic("xxxx")
	}
}

//32
func serializeFixVec(arr [][]byte) []byte {
	result := toHexInLittleEndian(uint64(len(arr)), 4)
	for _, v := range arr {
		result = append(result, serializeBytes(v)...)
	}
	return result
}

func serializeFixVecBytes(arr []byte) []byte {
	result := toHexInLittleEndian(uint64(len(arr)), 4)
	for _, v := range arr {
		result = append(result, serializeByte(v))
	}
	return result
}

func getOffsets(elmLengths []int) []int {
	if elmLengths == nil {
		return []int{}
	}
	headerLength := fullLengthSize + offsetSize*len(elmLengths)
	offsets := []int{headerLength}
	for index, _ := range elmLengths {
		if index > 0 {
			offsets = append(offsets, offsets[len(offsets)-1]+elmLengths[index-1])
		}
	}
	return offsets
}

func serializeByte(b byte) byte {
	/*result := make([]byte, 2*len(arr))
	hex.Encode(arr, result)*/
	return b
}

func serializeBytes(arr []byte) []byte {
	/*result := make([]byte, 2*len(arr))
	hex.Encode(arr, result)*/
	return arr
}

func serializeStruct(data *Seq) []byte {
	result := []byte{}
	data.Range(func(s string, bytes []byte) {
		result = append(result, serializeBytes(bytes)...)
	})
	return result
}

func serializeTable(table *Seq) []byte {
	body := []byte{}
	var lengths []int
	table.Range(func(s string, bytes []byte) {
		d := serializeBytes(bytes)
		lengths = append(lengths, len(bytes))
		body = append(body, d...)
	})
	headerLength := fullLengthSize + offsetSize*table.Len()
	fullLength := toHexInLittleEndian(uint64(headerLength+len(body)), 4)
	offsets := []byte{}
	offsetsValues := getOffsets(lengths)
	for _, offset := range offsetsValues {
		offsets = append(offsets, toHexInLittleEndian(uint64(offset), 4)...)
	}
	fullLength = append(fullLength, offsets...)
	fullLength = append(fullLength, body...)
	return fullLength
}

func serializeDynVec(arr [][]byte) []byte {
	body := []byte{}
	var lengths []int
	for _, v := range arr {
		d := serializeBytes(v)
		lengths = append(lengths, len(v))
		body = append(body, d...)
	}

	headerLength := fullLengthSize + offsetSize*len(arr)
	fullLength := toHexInLittleEndian(uint64(headerLength+len(body)), 4)
	offsets := []byte{}
	offsetsValues := getOffsets(lengths)
	for _, offset := range offsetsValues {
		offsets = append(offsets, toHexInLittleEndian(uint64(offset), 4)...)
	}
	fullLength = append(fullLength, offsets...)
	fullLength = append(fullLength, body...)
	return fullLength
}

func serializeHashType(hashType core.ScriptHashType) []byte {
	if hashType == 0 {
		return []byte{0}
	} else if hashType == 1 {
		return []byte{1}
	}
	panic("")
}

func serializeArgs(args [][]byte) []byte {
	//export const serializeArgs = (args: string[]) => serializeDynVec(args.map(arg => serializeFixVec(arg)))
	serializedCellDepList := [][]byte{}
	for _, arg := range args {
		serializedCellDepList = append(serializedCellDepList, serializeFixVecBytes(arg))
	}
	return serializeDynVec(serializedCellDepList)
}

func SerializeScript(script core.Script) []byte {
/*	d := make([]byte, len(script.Code_hash)/2)
	hex.Decode(script.Code_hash[:], d)*/
	serializedCodeHash := script.Code_hash[:]
	serializedHashType := serializeHashType(script.Hash_type)
	serializedArgs := serializeArgs(script.Args)
	table := NewSeq()
	table.Add("codeHash", serializedCodeHash)
	table.Add("hashType", serializedHashType)
	table.Add("args", serializedArgs)
	return serializeTable(table)
}

func SerializeVersion(version uint32) []byte {
	return toHexInLittleEndian(uint64(version), 4)
}

func SerializeOutPoint(outPoint core.OutPoint) []byte {
	data := NewSeq()
	data.Add("txHash", outPoint.Tx_hash[:])
	data.Add("index", toHexInLittleEndian(outPoint.Index, 4))
	return serializeStruct(data)
}

func SerializeDepType(depType core.DepType) []byte {
	if depType == 0 {
		return []byte{0}
	} else if depType == 1 {
		return []byte{1}
	}
	panic("")
}

func SerializeCellDep(dep core.CellDep) []byte {
	serializedOutPoint := SerializeOutPoint(dep.Out_point)
	serializedDepType := SerializeDepType(dep.Dep_type)
	data := NewSeq()
	data.Add("outPoint", serializedOutPoint)
	data.Add("depTyp", serializedDepType)
	return serializeStruct(data)
}

func SerializeCellDeps(cellDeps []core.CellDep) []byte {
	serializedCellDepList := [][]byte{}
	for _, cellDep := range cellDeps {
		serializedCellDepList = append(serializedCellDepList, SerializeCellDep(cellDep))
	}
	return serializeFixVec(serializedCellDepList)
}

func SerializeHeaderDeps(deps []core.H256) []byte {
	serializedHeaderDepList := [][]byte{}
	for _, dep := range deps {
		serializedHeaderDepList = append(serializedHeaderDepList, serializeBytes(dep[:]))
	}
	return serializeFixVec(serializedHeaderDepList)
}

// TODO: add tests
func SerializeInput(input core.CellInput) []byte {
	serializedOutPoint := SerializeOutPoint(input.Previous_output)
	serializedSince := toHexInLittleEndian(input.Since, 8)
	data := NewSeq()
	data.Add("since", serializedSince)
	data.Add("previousOutput", serializedOutPoint)
	return serializeStruct(data)
}

func SerializeInputs(inputs []core.CellInput) []byte {
	serializedInputList := [][]byte{}
	for _, input := range inputs {
		serializedInputList = append(serializedInputList, SerializeInput(input))
	}
	return serializeFixVec(serializedInputList)
}

func SerializeOutput(output core.CellOutput) []byte {
	serializedCapacity := toHexInLittleEndian(output.Capacity, 8)
	serializedLockScript := SerializeScript(output.Lock)
	serialiedTypeScript := []byte{}
	if output.Type_ != nil {
		serialiedTypeScript = SerializeScript(*output.Type_)
	}
	data := NewSeq()
	data.Add("capacity", serializedCapacity)
	data.Add("lock", serializedLockScript)
	data.Add("type", serialiedTypeScript)
	return serializeTable(data)
}

func SerializeOutputs(outputs []core.CellOutput) []byte {
	serializedOutputList := [][]byte{}
	for _, output := range outputs {
		serializedOutputList = append(serializedOutputList, SerializeOutput(output))
	}
	return serializeDynVec(serializedOutputList)
}

func SerializeOutputsData(outputsData [][]byte) []byte {
	serializedOutputsDatumList := [][]byte{}
	for _, outputdata := range outputsData {
		serializedOutputsDatumList = append(serializedOutputsDatumList, serializeFixVecBytes(outputdata))
	}
	return serializeDynVec(serializedOutputsDatumList)
}

func SerializeRawTransaction(rawTransaction core.RawTransaction) []byte {
	serializedVersion := SerializeVersion(rawTransaction.Version)
	serializedCellDeps := SerializeCellDeps(rawTransaction.CellDeps)
	serializedHeaderDeps := SerializeHeaderDeps(rawTransaction.HeadDeps)
	serializedInputs := SerializeInputs(rawTransaction.Inputs)
	serializedOutputs := SerializeOutputs(rawTransaction.Outputs)
	serializedOutputsData := SerializeOutputsData(rawTransaction.OutputData)

	table := NewSeq()
	table.Add("version", serializedVersion)
	table.Add("cellDeps", serializedCellDeps)
	table.Add("headerDeps", serializedHeaderDeps)
	table.Add("inputs", serializedInputs)
	table.Add("outputs", serializedOutputs)
	table.Add("outputsData", serializedOutputsData)
	return serializeTable(table)
}

type Seq struct {
	Keys []string
	Value map[string][]byte
}

func NewSeq() *Seq {
	return &Seq{
		Keys:  []string{},
		Value: map[string][]byte{},
	}
}

func (seq *Seq) Add(key string, value []byte) {
	seq.Keys = append(seq.Keys, key)
	seq.Value[key] = value
}

func(seq *Seq) Range(f func( string,  []byte)) {
	for _, key := range seq.Keys {
		val, _ := seq.Value[key]
		f(key, val)
	}
}

func (seq *Seq) Len() int {
	return len(seq.Keys)
}