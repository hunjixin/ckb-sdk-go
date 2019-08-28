package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
)

var (
	t = IWalk(&GoGenerate{})
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

type GoGenerate struct {
	PackageName string
	Importer    []string
	Structs     []string
	Functions   []string
	CodeBuffer  *bytes.Buffer
	Ast         *RustAst
	TypeMap     map[string]string
}

func (goGenerate *GoGenerate) Walk(ast *RustAst) {
	for _, token := range ast.Tokens {
		if token.Type != nil {
			goGenerate.WalkStruct(token.Type)
		}
		if token.Trait != nil {
			goGenerate.WalkTrait(token.Trait)
		}
	}
}

func (goGenerate *GoGenerate) WalkStruct(ast *RustStruct) {
	if ast.IsPublic {
		if ast.Name == "Witness" {
			fmt.Println("")
		}
		//write rpc struct
		goGenerate.CodeBuffer.WriteString("type " + "Rpc" + Capitalize(ast.Name) + " struct {\n")
		for _, field := range ast.Fields {
			goGenerate.WalkField(field)
		}
		goGenerate.CodeBuffer.WriteString("\n}\n")

		//TODO write real struct
	}
	//TODO  convert code betweenn rpc and real struct
}

func (goGenerate *GoGenerate) WalkTrait(ast *Trait) {
	goGenerate.CodeBuffer.WriteString("//" + ast.Name + "\n")
	for _, f := range ast.Funcs {
		goGenerate.WalkRpcFunc(f)
	}
}

func (goGenerate *GoGenerate) WalkRpcFunc(rpcFunc *RpcFunc) {
	goGenerate.CodeBuffer.WriteString("func (ckbClient *CkbClient) ")
	goGenerate.CodeBuffer.WriteString(LowToHeigh(rpcFunc.RpcName))
	goGenerate.WalkFunc(rpcFunc.Func)

	argLen := len(rpcFunc.Func.Args)
	goGenerate.CodeBuffer.WriteString("(")
	argsName := []string{}
	for index, arg := range rpcFunc.Func.Args {
		goGenerate.WalkArgument(arg)
		if !arg.IsSelf {
			if argLen != index {
				goGenerate.CodeBuffer.WriteString(", ")
				goGenerate.CodeBuffer.WriteString(" ")
			}
			argsName = append(argsName,arg.Pair.Name)
		}else{
			// not done
		}
	}
	goGenerate.CodeBuffer.WriteString(")")
	if rpcFunc.Func.Return != nil&&!rpcFunc.Func.Return.Void {
		goGenerate.CodeBuffer.WriteString(" (")
		goGenerate.WalkType(rpcFunc.Func.Return)
		goGenerate.CodeBuffer.WriteString(" , error){\n")
		//generate rpc call code
		/*rpcBlock := &RpcBlock{}
		err := ckbClient.client.CallFor(rpcBlock, "get_block_by_number", strconv.Itoa(number))
		if err != nil {
			return nil, err
		}
		return rpcBlock, nil*/
		returnName := goGenerate.getTypeName(rpcFunc.Func.Return)
		variable := strings.ToLower(returnName)
		goGenerate.CodeBuffer.WriteString(variable+ ":= &"+returnName+"{}\n")
		goGenerate.CodeBuffer.WriteString("err := ckbClient.client.CallFor("+variable +",\""+rpcFunc.RpcName+"\"")
		for _, arg := range argsName{
			goGenerate.CodeBuffer.WriteString("," + arg)
		}
		goGenerate.CodeBuffer.WriteString(")\n")
		goGenerate.CodeBuffer.WriteString(`
		if err != nil {
			return nil, err
		}`)
		goGenerate.CodeBuffer.WriteString("\nreturn "+ variable +", nil\n")
	} else {
		goGenerate.CodeBuffer.WriteString(" error {\n")
	}

	goGenerate.CodeBuffer.WriteString("\n}\n")
}

func (goGenerate *GoGenerate) WalkFunc(rFunc *Func) {
	//NOT DONE
}
func (goGenerate *GoGenerate) getTypeName (t *TypeRef) string {
	if t.Array != nil {
		return goGenerate.getTypeName(t.Array)
	}

	if t.NullAbleType != nil {
		return goGenerate.getTypeName(t.NullAbleType)
	}
	if redefineType, ok := goGenerate.TypeMap[t.Type]; ok {
		return redefineType
	} else {
		return t.Type
	}
}
func (goGenerate *GoGenerate) WalkArgument(argument *Argument) {
	if !argument.IsSelf {
		goGenerate.WalkPair(argument.Pair)
	} else {
		// self
	}
}

func (goGenerate *GoGenerate) WalkPair(pair *Pair) {
	goGenerate.CodeBuffer.WriteString(pair.Name)
	goGenerate.CodeBuffer.WriteByte(' ')
	goGenerate.WalkType(pair.Type)
}

func (goGenerate *GoGenerate) WalkField(field *Field) {
	name := field.Pair.Name
	if field.IsPublic {
		goGenerate.CodeBuffer.WriteString(Capitalize(name))
		goGenerate.CodeBuffer.WriteString("  ")
		goGenerate.WalkType(field.Pair.Type)
	} else {
		//NOTICE private
		goGenerate.CodeBuffer.WriteString(name)
		goGenerate.CodeBuffer.WriteString("  ")
		goGenerate.WalkType(field.Pair.Type)
	}
	goGenerate.CodeBuffer.WriteByte('\n')
}

func (goGenerate *GoGenerate) WalkType(typeRef *TypeRef) {
	if !typeRef.Void {
		if typeRef.Array != nil {
			goGenerate.CodeBuffer.WriteByte('[')
			goGenerate.CodeBuffer.WriteByte(']')
			goGenerate.WalkType(typeRef.Array)
			return
		}

		if typeRef.NullAbleType != nil {
			goGenerate.CodeBuffer.WriteByte('*')
			goGenerate.WalkType(typeRef.NullAbleType)
			return
		}
		if redefineType, ok := goGenerate.TypeMap[typeRef.Type]; ok {
			goGenerate.CodeBuffer.WriteString(redefineType)
		} else {
			goGenerate.CodeBuffer.WriteString(typeRef.Type)
		}
	}
}

func (g *GoGenerate) SaveTo(path string) {
	g.CodeBuffer.WriteString("package " + g.PackageName + "\n")
	g.CodeBuffer.WriteString("import (\n")
	for _, im := range g.Importer {
		g.CodeBuffer.WriteString("\"" + im + "\"\n")
	}
	g.CodeBuffer.WriteString(")\n")

	for _, goStruct := range g.Structs {
		g.CodeBuffer.WriteString(goStruct)
		g.CodeBuffer.WriteString("\n")
	}

	for _, function := range g.Functions {
		g.CodeBuffer.WriteString(function)
		g.CodeBuffer.WriteString("\n")
	}
	fmt.Println(g.CodeBuffer.String())
	g.Walk(g.Ast)
	_, err := format.Source(g.CodeBuffer.Bytes())
	fmt.Println(err)
	ioutil.WriteFile(path, g.CodeBuffer.Bytes(), os.ModePerm)
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
