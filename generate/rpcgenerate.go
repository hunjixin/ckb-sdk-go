package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

var (
	t = IWalk(&RpcGenerate{})
)

type IWalk interface {
	Walk(ast *RustAst)
	WalkStruct(ast *RustStruct)
	WalkTrait(ast *Trait)
	WalkRpcFunc(rpcFunc *RpcFunc)
	WalkFunc(rFunc *Func)
	WalkArgument(argument *Argument)
	WalkType(typeRef *TypeRef)
}

type RpcGenerate struct {
	PackageName string
	Importer    []string
	AssignStmt  []string
	Structs     []string
	Functions   []string
	CodeBuffer  *bytes.Buffer
	Ast         *RustAst
	RpcTypeMap  map[string]string
	SimpleType  map[string]bool
}

func (rpcGenerate *RpcGenerate) Walk(ast *RustAst) {
	for _, token := range ast.Tokens {
		if token.Type != nil {
			rpcGenerate.WalkStruct(token.Type)
		}
		if token.Trait != nil {
			rpcGenerate.WalkTrait(token.Trait)
		}
	}
}

func (rpcGenerate *RpcGenerate) WalkStruct(ast *RustStruct) {
	if ast.IsPublic {
		rpcGenerate.CodeBuffer = &bytes.Buffer{}
		//write rpc struct
		rpcName := "Rpc" + Capitalize(ast.Name)
		rpcGenerate.AssignStmt = append(rpcGenerate.AssignStmt, "var T"+rpcName+"=reflect.TypeOf(Rpc"+Capitalize(ast.Name)+"{})")
		rpcGenerate.CodeBuffer.WriteString("type " + rpcName + " struct {\n")
		for _, field := range ast.Fields {
			rpcGenerate.WalkField(field)
		}
		rpcGenerate.CodeBuffer.WriteString("\n}\n")
		rpcGenerate.Structs = append(rpcGenerate.Structs, rpcGenerate.CodeBuffer.String())

	}
	//TODO  convert code betweenn rpc and real struct
}

func (rpcGenerate *RpcGenerate) WalkTrait(ast *Trait) {
	rpcGenerate.CodeBuffer.WriteString("//" + ast.Name + "\n")
	for _, f := range ast.Funcs {
		rpcGenerate.WalkRpcFunc(f)
	}
}

func (rpcGenerate *RpcGenerate) WalkRpcFunc(rpcFunc *RpcFunc) {
	rpcGenerate.CodeBuffer = &bytes.Buffer{}
	rpcGenerate.CodeBuffer.WriteString("func (ckbClient *CkbClient) ")
	rpcGenerate.CodeBuffer.WriteString(LowToHeigh(rpcFunc.RpcName))
	rpcGenerate.WalkFunc(rpcFunc.Func)

	argLen := len(rpcFunc.Func.Args) - 1
	rpcGenerate.CodeBuffer.WriteString("(")
	argsName := []string{}
	for index, arg := range rpcFunc.Func.Args {
		rpcGenerate.WalkArgument(arg)
		if !arg.IsSelf {
			if argLen != index {
				rpcGenerate.CodeBuffer.WriteString(", ")
			}
			argsName = append(argsName, arg.Pair.Name)
		} else {
			// not done
		}
	}
	rpcGenerate.CodeBuffer.WriteString(")")
	if rpcFunc.Func.Return != nil && !rpcFunc.Func.Return.Void {
		returnName := rpcGenerate.getTypeName(rpcFunc.Func.Return)
		rpcGenerate.CodeBuffer.WriteString(" ( *" + returnName + " , error){\n")
		variable := strings.ToLower(returnName)
		_, isSimple := rpcGenerate.SimpleType[returnName]
		if isSimple {
			rpcGenerate.CodeBuffer.WriteString(variable + ":= " + returnName + "(\"\")\n")
			rpcGenerate.CodeBuffer.WriteString("err := ckbClient.client.CallFor(&" + variable + ",\"" + rpcFunc.RpcName + "\"")
		} else {
			rpcGenerate.CodeBuffer.WriteString(variable + ":= &" + returnName + "{}\n")
			rpcGenerate.CodeBuffer.WriteString("err := ckbClient.client.CallFor(" + variable + ",\"" + rpcFunc.RpcName + "\"")
		}

		for _, arg := range argsName {
			rpcGenerate.CodeBuffer.WriteString("," + arg)
		}
		rpcGenerate.CodeBuffer.WriteString(")\n")
		rpcGenerate.CodeBuffer.WriteString(`
		if err != nil {
			return nil, err
		}`)
		if isSimple {
			rpcGenerate.CodeBuffer.WriteString("\nreturn &" + variable + ", nil\n")
		} else {
			rpcGenerate.CodeBuffer.WriteString("\nreturn " + variable + ", nil\n")
		}

	} else {
		rpcGenerate.CodeBuffer.WriteString(" error {\n")

		rpcGenerate.CodeBuffer.WriteString(`res, err := ckbClient.client.Call("` + rpcFunc.RpcName + `"`)
		for _, arg := range argsName {
			rpcGenerate.CodeBuffer.WriteString("," + arg)
		}
		rpcGenerate.CodeBuffer.WriteString(")\n")
		rpcGenerate.CodeBuffer.WriteString(
			`
		if err != nil {
			return err
		}
		if res.Error != nil {
			return res.Error
		}
		return nil
	`)
	}

	rpcGenerate.CodeBuffer.WriteString("\n}\n")
	rpcGenerate.Functions = append(rpcGenerate.Functions, rpcGenerate.CodeBuffer.String())
}

func (rpcGenerate *RpcGenerate) WalkFunc(rFunc *Func) {
	//NOT DONE
}
func (rpcGenerate *RpcGenerate) getTypeName(t *TypeRef) string {
	if t.Array != nil {
		return rpcGenerate.getTypeName(t.Array)
	}

	if t.NullAbleType != nil {
		return rpcGenerate.getTypeName(t.NullAbleType)
	}
	if redefineType, ok := rpcGenerate.RpcTypeMap[t.Type]; ok {
		return redefineType
	} else {
		return t.Type
	}
}

func (rpcGenerate *RpcGenerate) WalkArgument(argument *Argument) {
	if !argument.IsSelf {
		rpcGenerate.WalkPair(argument.Pair)
	} else {
		// self
	}
}

func (rpcGenerate *RpcGenerate) WalkPair(pair *Pair) {
	rpcGenerate.CodeBuffer.WriteString(pair.Name)
	rpcGenerate.CodeBuffer.WriteByte(' ')
	rpcGenerate.WalkType(pair.Type)
}

func (rpcGenerate *RpcGenerate) WalkField(field *Field) {
	name := field.Pair.Name
	if field.IsPublic {
		rpcGenerate.CodeBuffer.WriteString(Capitalize(name) + " ")
		rpcGenerate.WalkType(field.Pair.Type)
		rpcGenerate.CodeBuffer.WriteString( "`json:\""+name+"\"`")
	} else {
		//NOTICE private
		rpcGenerate.CodeBuffer.WriteString(name + " ")
		rpcGenerate.WalkType(field.Pair.Type)
	}
	rpcGenerate.CodeBuffer.WriteByte('\n')
}

func (rpcGenerate *RpcGenerate) WalkType(typeRef *TypeRef) {
	if !typeRef.Void {
		if typeRef.Array != nil {
			rpcGenerate.CodeBuffer.WriteString("[]")
			rpcGenerate.WalkType(typeRef.Array)
			return
		}

		if typeRef.NullAbleType != nil {
			rpcGenerate.CodeBuffer.WriteByte('*')
			rpcGenerate.WalkType(typeRef.NullAbleType)
			return
		}
		if redefineType, ok := rpcGenerate.RpcTypeMap[typeRef.Type]; ok {
			rpcGenerate.CodeBuffer.WriteString(redefineType)
		} else {
			rpcGenerate.CodeBuffer.WriteString(typeRef.Type)
		}
	}
}

func (rpcGenerate *RpcGenerate) SaveTo(rpcTypePath string, rpcClientPath string) {
	rpcGenerate.CodeBuffer = &bytes.Buffer{}
	rpcGenerate.CodeBuffer.WriteString("package " + rpcGenerate.PackageName + "\n")
	rpcGenerate.CodeBuffer.WriteString("import (\n")
	for _, im := range rpcGenerate.Importer {
		rpcGenerate.CodeBuffer.WriteString("\"" + im + "\"\n")
	}
	rpcGenerate.CodeBuffer.WriteString(")\n")
	rpcGenerate.CodeBuffer.WriteString(
		`
type CkbClient struct {
	url    string
	client jsonrpc.RPCClient
}

func NewCkbClient(url string) *CkbClient {
	return &CkbClient{
		url:    url,
		client: jsonrpc.NewClient(url),
	}
}
`,
	)

	for _, function := range rpcGenerate.Functions {
		rpcGenerate.CodeBuffer.WriteString(function)
		rpcGenerate.CodeBuffer.WriteString("\n")
	}
	fileBytes, _ := format.Source(rpcGenerate.CodeBuffer.Bytes())
	ioutil.WriteFile(rpcClientPath, fileBytes, os.ModePerm)

	rpcGenerate.CodeBuffer = &bytes.Buffer{}
	rpcGenerate.CodeBuffer.WriteString("package " + rpcGenerate.PackageName + "\n")
	for _, assign := range rpcGenerate.AssignStmt {
		rpcGenerate.CodeBuffer.WriteString(assign)
		rpcGenerate.CodeBuffer.WriteString("\n")
	}
	for _, goStruct := range rpcGenerate.Structs {
		rpcGenerate.CodeBuffer.WriteString(goStruct)
		rpcGenerate.CodeBuffer.WriteString("\n")
	}
	fileBytes, _ = format.Source(rpcGenerate.CodeBuffer.Bytes())
	ioutil.WriteFile(rpcTypePath, fileBytes, os.ModePerm)
}
