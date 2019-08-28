package main

import (
	"bytes"
	"log"
	"os"

)



func main() {
	ast := &RustAst{}
	r, err := os.Open(`C:\Users\Drep\Desktop\goproject\src\github.com\hunjixin\ckb-sdk-go\generate\m.rs`)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	err = parser.Parse(r, ast)
	if err != nil {
		log.Fatal(err)
	}

	collector := &TypeCollecor{map[string]string{}}
	collector.Walk(ast)
	collector.newTypeMap["String"] = "string"

	codeBuf := bytes.NewBuffer([]byte{})
	g := &GoGenerate{
		PackageName: "ckb_sdk_go",
		Importer:[]string{
			"encoding/json",
			"github.com/ybbus/jsonrpc",
			"strconv",
		},
		Structs:[]string{`
type CkbClient struct {
	url string
	client jsonrpc.RPCClient
}
		`},
		CodeBuffer:codeBuf,
		Ast:ast,
		TypeMap: collector.newTypeMap,
	}
	g.SaveTo("./xxxx.go")
}
