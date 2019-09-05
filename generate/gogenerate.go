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
	Structs     []string
	Functions   []string
	CodeBuffer  *bytes.Buffer
	Ast         *RustAst
	RpcTypeMap  map[string]string
	SimpleType  map[string]bool

	RealStructs    []string
	RealTypeBuffer *bytes.Buffer
	RealTypeMap    map[string]*TypeConvert
}

func (goGenerate *RpcGenerate) Walk(ast *RustAst) {
	for _, token := range ast.Tokens {
		if token.Type != nil {
			goGenerate.WalkStruct(token.Type)
		}
		if token.Trait != nil {
			goGenerate.WalkTrait(token.Trait)
		}
	}
}

func (goGenerate *RpcGenerate) WalkStruct(ast *RustStruct) {
	if ast.IsPublic {
		goGenerate.CodeBuffer = &bytes.Buffer{}
		goGenerate.RealTypeBuffer = &bytes.Buffer{}
		//write rpc struct
		goGenerate.RealStructs = append(goGenerate.RealStructs, "var T"+Capitalize(ast.Name)+"=reflect.TypeOf("+Capitalize(ast.Name)+"{})")
		rpcName := "Rpc" + Capitalize(ast.Name)
		goGenerate.CodeBuffer.WriteString("type " + rpcName + " struct {\n")
		goGenerate.RealTypeBuffer.WriteString("type " + Capitalize(ast.Name) + " struct {\n")
		for _, field := range ast.Fields {
			goGenerate.WalkField(field)
		}
		goGenerate.CodeBuffer.WriteString("\n}\n")
		goGenerate.RealTypeBuffer.WriteString("\n}\n")
		goGenerate.Structs = append(goGenerate.Structs, goGenerate.CodeBuffer.String())
		goGenerate.RealStructs = append(goGenerate.RealStructs, goGenerate.RealTypeBuffer.String())

	}
	//TODO  convert code betweenn rpc and real struct
}

func (goGenerate *RpcGenerate) WalkTrait(ast *Trait) {
	goGenerate.CodeBuffer.WriteString("//" + ast.Name + "\n")
	for _, f := range ast.Funcs {
		goGenerate.WalkRpcFunc(f)
	}
}

func (goGenerate *RpcGenerate) WalkRpcFunc(rpcFunc *RpcFunc) {
	goGenerate.CodeBuffer = &bytes.Buffer{}
	goGenerate.CodeBuffer.WriteString("func (ckbClient *CkbClient) ")
	goGenerate.CodeBuffer.WriteString(LowToHeigh(rpcFunc.RpcName))
	goGenerate.WalkFunc(rpcFunc.Func)

	argLen := len(rpcFunc.Func.Args) - 1
	goGenerate.CodeBuffer.WriteString("(")
	argsName := []string{}
	for index, arg := range rpcFunc.Func.Args {
		goGenerate.WalkArgument(arg)
		if !arg.IsSelf {
			if argLen != index {
				goGenerate.CodeBuffer.WriteString(", ")
			}
			argsName = append(argsName, arg.Pair.Name)
		} else {
			// not done
		}
	}
	goGenerate.CodeBuffer.WriteString(")")
	if rpcFunc.Func.Return != nil && !rpcFunc.Func.Return.Void {
		returnName := goGenerate.getTypeName(rpcFunc.Func.Return)
		goGenerate.CodeBuffer.WriteString(" ( *" + returnName + " , error){\n")
		//generate rpc call code
		/*rpcBlock := &RpcBlock{}
		err := ckbClient.client.CallFor(rpcBlock, "get_block_by_number", strconv.Itoa(number))
		if err != nil {
			return nil, err
		}
		return rpcBlock, nil*/
		variable := strings.ToLower(returnName)
		_, isSimple := goGenerate.SimpleType[returnName]
		if isSimple {
			goGenerate.CodeBuffer.WriteString(variable + ":= " + returnName + "(\"\")\n")
			goGenerate.CodeBuffer.WriteString("err := ckbClient.client.CallFor(&" + variable + ",\"" + rpcFunc.RpcName + "\"")
		} else {
			goGenerate.CodeBuffer.WriteString(variable + ":= &" + returnName + "{}\n")
			goGenerate.CodeBuffer.WriteString("err := ckbClient.client.CallFor(" + variable + ",\"" + rpcFunc.RpcName + "\"")
		}

		for _, arg := range argsName {
			goGenerate.CodeBuffer.WriteString("," + arg)
		}
		goGenerate.CodeBuffer.WriteString(")\n")
		goGenerate.CodeBuffer.WriteString(`
		if err != nil {
			return nil, err
		}`)
		if isSimple {
			goGenerate.CodeBuffer.WriteString("\nreturn &" + variable + ", nil\n")
		} else {
			goGenerate.CodeBuffer.WriteString("\nreturn " + variable + ", nil\n")
		}

	} else {
		goGenerate.CodeBuffer.WriteString(" error {\n")

		goGenerate.CodeBuffer.WriteString(`res, err := ckbClient.client.Call("` + rpcFunc.RpcName + `"`)
		for _, arg := range argsName {
			goGenerate.CodeBuffer.WriteString("," + arg)
		}
		goGenerate.CodeBuffer.WriteString(")\n")
		goGenerate.CodeBuffer.WriteString(
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

	goGenerate.CodeBuffer.WriteString("\n}\n")
	goGenerate.Functions = append(goGenerate.Functions, goGenerate.CodeBuffer.String())
}

func (goGenerate *RpcGenerate) WalkFunc(rFunc *Func) {
	//NOT DONE
}
func (goGenerate *RpcGenerate) getTypeName(t *TypeRef) string {
	if t.Array != nil {
		return goGenerate.getTypeName(t.Array)
	}

	if t.NullAbleType != nil {
		return goGenerate.getTypeName(t.NullAbleType)
	}
	if redefineType, ok := goGenerate.RpcTypeMap[t.Type]; ok {
		return redefineType
	} else {
		return t.Type
	}
}

func (goGenerate *RpcGenerate) WalkArgument(argument *Argument) {
	if !argument.IsSelf {
		goGenerate.WalkPair(argument.Pair)
	} else {
		// self
	}
}

func (goGenerate *RpcGenerate) WalkPair(pair *Pair) {
	goGenerate.CodeBuffer.WriteString(pair.Name)
	goGenerate.CodeBuffer.WriteByte(' ')
	goGenerate.WalkType(pair.Type)
}

func (goGenerate *RpcGenerate) WalkField(field *Field) {
	name := field.Pair.Name
	if field.IsPublic {
		goGenerate.CodeBuffer.WriteString(Capitalize(name) + " ")
		goGenerate.RealTypeBuffer.WriteString(Capitalize(name) + " ")
		goGenerate.WalkType(field.Pair.Type)
		goGenerate.CodeBuffer.WriteString( "`json:\""+name+"\"`")
	} else {
		//NOTICE private
		goGenerate.CodeBuffer.WriteString(name + " ")
		goGenerate.RealTypeBuffer.WriteString(name + " ")
		goGenerate.WalkType(field.Pair.Type)
	}
	goGenerate.CodeBuffer.WriteByte('\n')
	goGenerate.RealTypeBuffer.WriteByte('\n')
}

func (goGenerate *RpcGenerate) WalkType(typeRef *TypeRef) {
	if !typeRef.Void {
		if typeRef.Array != nil {
			goGenerate.CodeBuffer.WriteString("[]")
			goGenerate.RealTypeBuffer.WriteString("[]")
			goGenerate.WalkType(typeRef.Array)
			return
		}

		if typeRef.NullAbleType != nil {
			goGenerate.CodeBuffer.WriteByte('*')
			goGenerate.RealTypeBuffer.WriteByte('*')
			goGenerate.WalkType(typeRef.NullAbleType)
			return
		}
		if redefineType, ok := goGenerate.RpcTypeMap[typeRef.Type]; ok {
			goGenerate.CodeBuffer.WriteString(redefineType)
		} else {
			goGenerate.CodeBuffer.WriteString(typeRef.Type)
		}
		if redefineType, ok := goGenerate.RealTypeMap[typeRef.Type]; ok {
			goGenerate.RealTypeBuffer.WriteString(redefineType.ToType)
		} else {
			goGenerate.RealTypeBuffer.WriteString(typeRef.Type)
		}
	}
}

func (g *RpcGenerate) SaveTo(rpcTypePath string, rpcClientPath string, realTypePath string) {
	g.Walk(g.Ast)

	g.CodeBuffer = &bytes.Buffer{}
	g.CodeBuffer.WriteString("package " + g.PackageName + "\n")
	g.CodeBuffer.WriteString("import (\n")
	for _, im := range g.Importer {
		g.CodeBuffer.WriteString("\"" + im + "\"\n")
	}
	g.CodeBuffer.WriteString(")\n")
	g.CodeBuffer.WriteString(
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

	for _, function := range g.Functions {
		g.CodeBuffer.WriteString(function)
		g.CodeBuffer.WriteString("\n")
	}
	fileBytes, _ := format.Source(g.CodeBuffer.Bytes())
	ioutil.WriteFile(rpcClientPath, fileBytes, os.ModePerm)

	g.CodeBuffer = &bytes.Buffer{}
	g.CodeBuffer.WriteString("package " + g.PackageName + "\n")
	for _, goStruct := range g.Structs {
		g.CodeBuffer.WriteString(goStruct)
		g.CodeBuffer.WriteString("\n")
	}
	fileBytes, _ = format.Source(g.CodeBuffer.Bytes())
	ioutil.WriteFile(rpcTypePath, fileBytes, os.ModePerm)

	g.RealTypeBuffer = &bytes.Buffer{}
	g.RealTypeBuffer.WriteString("package " + g.PackageName + "\n")
	g.RealTypeBuffer.WriteString(`import ("reflect")` + "\n")
	for _, goStruct := range g.RealStructs {
		g.RealTypeBuffer.WriteString(goStruct)
		g.RealTypeBuffer.WriteString("\n")
	}
	fileBytes, _ = format.Source(g.RealTypeBuffer.Bytes())
	ioutil.WriteFile(realTypePath, fileBytes, os.ModePerm)

}

type TypeCollecor struct {
	newTypeMap map[string]string
}

func (typeCollecor *TypeCollecor) Walk(ast *RustAst) {
	for _, token := range ast.Tokens {
		if token.Type != nil {
			typeCollecor.WalkStruct(token.Type)
		}
		if token.Trait != nil {
			typeCollecor.WalkTrait(token.Trait)
		}
	}
}

func (typeCollecor *TypeCollecor) WalkStruct(ast *RustStruct) {
	if ast.IsPublic {
		typeCollecor.newTypeMap[ast.Name] = "Rpc" + Capitalize(ast.Name)
		for _, field := range ast.Fields {
			typeCollecor.WalkField(field)
		}
	}
}

func (typeCollecor *TypeCollecor) WalkTrait(ast *Trait) {
	for _, f := range ast.Funcs {
		typeCollecor.WalkRpcFunc(f)
	}
}

func (typeCollecor *TypeCollecor) WalkRpcFunc(rpcFunc *RpcFunc) {
	typeCollecor.WalkFunc(rpcFunc.Func)
}

func (typeCollecor *TypeCollecor) WalkFunc(rFunc *Func) {
	for _, arg := range rFunc.Args {
		typeCollecor.WalkArgument(arg)
	}
	if rFunc.Return != nil {
		typeCollecor.WalkType(rFunc.Return)
	}
}

func (typeCollecor *TypeCollecor) WalkArgument(argument *Argument) {
	if !argument.IsSelf {
		typeCollecor.WalkPair(argument.Pair)
	} else {
		// self
	}
}

func (typeCollecor *TypeCollecor) WalkPair(pair *Pair) {
	typeCollecor.WalkType(pair.Type)
}

func (typeCollecor *TypeCollecor) WalkField(field *Field) {
	if field.IsPublic {
		typeCollecor.WalkType(field.Pair.Type)
	} else {
		typeCollecor.WalkType(field.Pair.Type)
	}
}

func (typeCollecor *TypeCollecor) WalkType(typeRef *TypeRef) {
	if !typeRef.Void {
		if typeRef.Array != nil {
			typeCollecor.WalkType(typeRef.Array)
			return
		}

		if typeRef.NullAbleType != nil {
			typeCollecor.WalkType(typeRef.NullAbleType)
			return
		}
	}
}
