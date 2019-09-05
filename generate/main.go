package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"os"
	"strconv"
)

type TypeConvert struct {
	Convert func(string) (interface{}, error)
	ToType  string
}

var (
	Uint32Convert = &TypeConvert{
		ToType: "uint32",
		Convert: func(s string) (interface{}, error) {
			n, err := strconv.ParseUint(s, 16, 32)
			if err != nil {
				return nil, err
			}
			return uint32(n), nil
		},
	}
	Uint64Convert = &TypeConvert{
		ToType: "uint64",
		Convert: func(s string) (interface{}, error) {
			n, err := strconv.ParseUint(s, 16, 64)
			if err != nil {
				return nil, err
			}
			return uint64(n), nil
		},
	}
	SliceConvert = &TypeConvert{
		ToType: "[]byte",
		Convert: func(s string) (interface{}, error) {
			sliceBytes, err := hex.DecodeString(s)
			if err != nil {
				return nil, err
			}
			return sliceBytes, nil
		},
	}
	ArrayTenConvert = &TypeConvert{
		ToType: "[10]byte",
		Convert: func(s string) (interface{}, error) {
			sliceBytes, err := hex.DecodeString(s)
			if err != nil {
				return nil, err
			}
			return sliceBytes[0:10], nil
		},
	}
	ArrayConvert = &TypeConvert{
		ToType: "[32]byte",
		Convert: func(s string) (interface{}, error) {
			sliceBytes, err := hex.DecodeString(s)
			if err != nil {
				return nil, err
			}
			return sliceBytes[0:32], nil
		},
	}
	StringConvert = &TypeConvert{
		ToType: "string",
		Convert: func(s string) (interface{}, error) {
			return s, nil
		},
	}
)

func main() {
	ast := &RustAst{}
	r, err := os.Open(`/home/hunjixin/project/ckb-sdk-go/generate/m.rs`)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	err = parser.Parse(r, ast)
	if err != nil {
		log.Fatal(err)
	}

	collector := &TypeCollecor{map[string]string{
		"ScriptHashType":"RpcScriptHashType",
		"Witness":"RpcWitness",
	}}
	collector.Walk(ast)
	collector.newTypeMap["String"] = "string"
	collector.newTypeMap["JsonBytes"] = "string"
	collector.newTypeMap["AlertId"] = "string"
	collector.newTypeMap["Version"] = "string"
	collector.newTypeMap["BlockNumber"] = "string"
	collector.newTypeMap["Capacity"] = "string"
	collector.newTypeMap["ProposalShortId"] = "string"
	collector.newTypeMap["Timestamp"] = "string"
	collector.newTypeMap["Cycle"] = "string"
	collector.newTypeMap["Unsigned"] = "string"
	collector.newTypeMap["AlertPriority"] = "string"
	collector.newTypeMap["EpochNumber"] = "string"
	collector.newTypeMap["U256"] = "string"
	collector.newTypeMap["Hash"] = "string"
	collector.newTypeMap["H256"] = "string"

	simpleType := map[string]bool{
		"int":             true,
		"string":          true,
		"bool":            true,
		"uint":            true,
		"int32":           true,
		"uint32":          true,
		"int64":           true,
		"uint64":          true,
		"AlertId":         true,
		"DAO":             true,
		"H256":            true,
		"UInt32":          true,
		"Index":           true,
		"Version":         true,
		"BlockNumber":     true,
		"Capacity":        true,
		"ProposalShortId": true,
		"Timestamp":       true,
		"Cycle":           true,
		"Unsigned":        true,
		"AlertPriority":   true,
		"EpochNumber":     true,
		"U256":            true,
		"ScriptHashType":            true,
	}

	realType := map[string]*TypeConvert{
		"AlertId":         Uint32Convert,
		"DAO":             SliceConvert,
		"UInt32":          Uint32Convert,
		"Index":           Uint32Convert,
		"Version":         Uint32Convert,
		"BlockNumber":     Uint64Convert,
		"Capacity":        Uint64Convert,
		"ProposalShortId": ArrayTenConvert,
		"Timestamp":       Uint64Convert,
		"Cycle":           Uint64Convert,
		"Unsigned":        Uint64Convert,
		"AlertPriority":   Uint32Convert,
		"EpochNumber":     Uint64Convert,
		//"U256":            StringConvert,
		//"H256":            StringConvert,
		"String":    StringConvert,
		"JsonBytes": SliceConvert,
	}

	codeBuf := bytes.NewBuffer([]byte{})
	g := &RpcGenerate{
		PackageName: "ckb_sdk_go",
		Importer: []string{
			//"encoding/json",
			"github.com/ybbus/jsonrpc",
			//"strconv",
		},
		Structs:     []string{},
		CodeBuffer:  codeBuf,
		Ast:         ast,
		RpcTypeMap:  collector.newTypeMap,
		SimpleType:  simpleType,
		RealTypeMap: realType,
	}
	g.SaveTo("./rpctypes.go", "./client.go", "./types.go")
}
