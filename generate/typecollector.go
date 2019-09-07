package main


type TypeCollector struct {
	newTypeMap map[string]string
}

func (typeCollector *TypeCollector) Walk(ast *RustAst) {
	for _, token := range ast.Tokens {
		if token.Type != nil {
			typeCollector.WalkStruct(token.Type)
		}
		if token.Trait != nil {
			typeCollector.WalkTrait(token.Trait)
		}
	}
}

func (typeCollector *TypeCollector) WalkStruct(ast *RustStruct) {
	if ast.IsPublic {
		typeCollector.newTypeMap[ast.Name] = "Rpc" + Capitalize(ast.Name)
		for _, field := range ast.Fields {
			typeCollector.WalkField(field)
		}
	}
}

func (typeCollector *TypeCollector) WalkTrait(ast *Trait) {
	for _, f := range ast.Funcs {
		typeCollector.WalkRpcFunc(f)
	}
}

func (typeCollector *TypeCollector) WalkRpcFunc(rpcFunc *RpcFunc) {
	typeCollector.WalkFunc(rpcFunc.Func)
}

func (typeCollector *TypeCollector) WalkFunc(rFunc *Func) {
	for _, arg := range rFunc.Args {
		typeCollector.WalkArgument(arg)
	}
	if rFunc.Return != nil {
		typeCollector.WalkType(rFunc.Return)
	}
}

func (typeCollector *TypeCollector) WalkArgument(argument *Argument) {
	if !argument.IsSelf {
		typeCollector.WalkPair(argument.Pair)
	} else {
		// self
	}
}

func (typeCollector *TypeCollector) WalkPair(pair *Pair) {
	typeCollector.WalkType(pair.Type)
}

func (typeCollector *TypeCollector) WalkField(field *Field) {
	if field.IsPublic {
		typeCollector.WalkType(field.Pair.Type)
	} else {
		typeCollector.WalkType(field.Pair.Type)
	}
}

func (typeCollector *TypeCollector) WalkType(typeRef *TypeRef) {
	if !typeRef.Void {
		if typeRef.Array != nil {
			typeCollector.WalkType(typeRef.Array)
			return
		}

		if typeRef.NullAbleType != nil {
			typeCollector.WalkType(typeRef.NullAbleType)
			return
		}
	}
}
