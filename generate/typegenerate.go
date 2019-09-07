package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
)


type TypeGenerate struct {
	PackageName string
	Importer    []string
	Ast         *RustAst
	AssignStmt  []string
	Structs     []string
	TypeBuffer  *bytes.Buffer
	TypeMap     map[string]*TypeConvert
}

func (typeGenerate *TypeGenerate) Walk(ast *RustAst) {
	for _, token := range ast.Tokens {
		if token.Type != nil {
			typeGenerate.WalkStruct(token.Type)
		}
		if token.Trait != nil {
			typeGenerate.WalkTrait(token.Trait)
		}
	}
}

func (typeGenerate *TypeGenerate) WalkStruct(ast *RustStruct) {
	if ast.IsPublic {
		typeGenerate.TypeBuffer = &bytes.Buffer{}
		//write rpc struct
		typeGenerate.AssignStmt = append(typeGenerate.AssignStmt, "var T"+Capitalize(ast.Name)+"=reflect.TypeOf("+Capitalize(ast.Name)+"{})")
		typeGenerate.TypeBuffer.WriteString("type " + Capitalize(ast.Name) + " struct {\n")
		for _, field := range ast.Fields {
			typeGenerate.WalkField(field)
		}
		typeGenerate.TypeBuffer.WriteString("\n}\n")
		typeGenerate.Structs = append(typeGenerate.Structs, typeGenerate.TypeBuffer.String())

	}
	//TODO  convert code betweenn rpc and real struct
}

func (typeGenerate *TypeGenerate) WalkTrait(ast *Trait) {
	typeGenerate.TypeBuffer.WriteString("//" + ast.Name + "\n")
	for _, f := range ast.Funcs {
		typeGenerate.WalkRpcFunc(f)
	}
}

func (typeGenerate *TypeGenerate) WalkRpcFunc(rpcFunc *RpcFunc) {
	//NOT DONE
}

func (typeGenerate *TypeGenerate) WalkFunc(rFunc *Func) {
	//NOT DONE
}
func (typeGenerate *TypeGenerate) getTypeName(t *TypeRef) string {
	if t.Array != nil {
		return typeGenerate.getTypeName(t.Array)
	}

	if t.NullAbleType != nil {
		return typeGenerate.getTypeName(t.NullAbleType)
	}
	return t.Type
}

func (typeGenerate *TypeGenerate) WalkArgument(argument *Argument) {
	//NOT DONE
}

func (typeGenerate *TypeGenerate) WalkPair(pair *Pair) {
	typeGenerate.WalkType(pair.Type)
}

func (typeGenerate *TypeGenerate) WalkField(field *Field) {
	name := field.Pair.Name
	if field.IsPublic {
		typeGenerate.TypeBuffer.WriteString(Capitalize(name) + " ")
		typeGenerate.WalkType(field.Pair.Type)
	} else {
		//NOTICE private
		typeGenerate.TypeBuffer.WriteString(name + " ")
		typeGenerate.WalkType(field.Pair.Type)
	}
	typeGenerate.TypeBuffer.WriteByte('\n')
}

func (typeGenerate *TypeGenerate) WalkType(typeRef *TypeRef) {
	if !typeRef.Void {
		if typeRef.Array != nil {
			typeGenerate.TypeBuffer.WriteString("[]")
			typeGenerate.WalkType(typeRef.Array)
			return
		}

		if typeRef.NullAbleType != nil {
			typeGenerate.TypeBuffer.WriteByte('*')
			typeGenerate.WalkType(typeRef.NullAbleType)
			return
		}
		if redefineType, ok := typeGenerate.TypeMap[typeRef.Type]; ok {
			typeGenerate.TypeBuffer.WriteString(redefineType.ToType)
		} else {
			typeGenerate.TypeBuffer.WriteString(typeRef.Type)
		}
	}
}

func (typeGenerate *TypeGenerate) SaveTo(realTypePath string) {
	typeGenerate.TypeBuffer = &bytes.Buffer{}
	typeGenerate.TypeBuffer.WriteString("package " + typeGenerate.PackageName + "\n")
	typeGenerate.TypeBuffer.WriteString(`import ("reflect")` + "\n")
	for _, assign := range typeGenerate.AssignStmt {
		typeGenerate.TypeBuffer.WriteString(assign)
		typeGenerate.TypeBuffer.WriteString("\n")
	}
	for _, goStruct := range typeGenerate.Structs {
		typeGenerate.TypeBuffer.WriteString(goStruct)
		typeGenerate.TypeBuffer.WriteString("\n")
	}
	fileBytes, _ := format.Source(typeGenerate.TypeBuffer.Bytes())
	ioutil.WriteFile(realTypePath, fileBytes, os.ModePerm)

}
