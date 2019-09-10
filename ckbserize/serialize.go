package ckbserize

import (
	"ckb-sdk-go/core"
	"encoding/binary"
	"encoding/hex"
	"fmt"
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

func serializeStruct(data map[string][]byte) []byte {
	result := []byte{}
	for _, v := range data {
		result = append(result, serializeBytes(v)...)
	}
	return result
}

func serializeTable(table map[string][]byte) []byte {
	body := []byte{}
	var lengths []int
	for _, v := range table {
		d := serializeBytes(v)
		lengths = append(lengths, len(v))
		body = append(body, d...)
	}

	headerLength := fullLengthSize + offsetSize*len(table)
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
	table := map[string][]byte{
		"codeHas": serializedCodeHash,
		"hashTyp": serializedHashType,
		"args":    serializedArgs,
	}
	fmt.Println(hex.EncodeToString(serializedCodeHash))
	fmt.Println(hex.EncodeToString(serializedHashType))
	fmt.Println(hex.EncodeToString(serializedArgs))
	fmt.Println(hex.EncodeToString( serializeTable(table)))
	return serializeTable(table)
}

func SerializeVersion(version uint32) []byte {
	return toHexInLittleEndian(uint64(version), 4)
}

func SerializeOutPoint(outPoint core.OutPoint) []byte {
	data := map[string][]byte{
		"txHash": outPoint.Tx_hash[:],
		"index":  toHexInLittleEndian(outPoint.Index, 4),
	}
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
	data := map[string][]byte{
		"outPoint": serializedOutPoint,
		"depTyp":   serializedDepType,
	}
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
	data := map[string][]byte{
		"since":          serializedSince,
		"previousOutput": serializedOutPoint,
	}
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
	data := map[string][]byte{
		"capacity": serializedCapacity,
		"lock":     serializedLockScript,
		"type":     serialiedTypeScript,
	}
	fmt.Println(hex.EncodeToString(serializedCapacity))
	fmt.Println(hex.EncodeToString(serializedLockScript))
	fmt.Println(hex.EncodeToString(serialiedTypeScript))
	//fmt.Println(hex.EncodeToString(serializeTable(data)))
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

	table := map[string][]byte{
		"version":     serializedVersion,
		"cellDeps":    serializedCellDeps,
		"headerDeps":  serializedHeaderDeps,
		"inputs":      serializedInputs,
		"outputs":     serializedOutputs,
		"outputsData": serializedOutputsData,
	}

	return serializeTable(table)
}
