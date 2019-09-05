package main

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
)

type RustAst struct {
	Tokens []*Token `@@*`
}

type Token struct {
	Type  *RustStruct `@@`
	Trait *Trait      `| @@`
}

type Trait struct {
	IsPublic bool       `@"pub"?`
	Name     string     `"trait" @Ident`
	Funcs    []*RpcFunc `"{" @@* "}"`
}

type RpcFunc struct {
	RpcName string `"#" "[" "rpc" "(" "name" "=" "\"" @Ident "\"" ")" "]"`
	Func    *Func  ` @@ ";"`
}

type Func struct {
	Name   string      `"fn" @Ident `
	Args   []*Argument `"(" [ @@ { "," @@ } {","}] ")"`
	Return *TypeRef    `"-" ">" "Result" "<" @@ ">"`
}

type Argument struct {
	IsSelf bool  `( @("&" "self")`
	Pair   *Pair `  | @@ )`
}

type Pair struct {
	Name string   `@Ident`
	Type *TypeRef `":" @@`
}

type RustStruct struct {
	IsPublic bool     `@"pub"`
	Name     string   `"struct" @Ident`
	Fields   []*Field `"{" @@* "}"`
}

type Field struct {
	IsPublic bool  `@"pub"?`
	Pair     *Pair `@@ ","`
}

type TypeRef struct {
	Array        *TypeRef `( "Vec" "<" @@ ">"`
	NullAbleType *TypeRef `  | "Option" "<" @@ ">"`
	Void         bool     `| @( "(" ")" )`
	Type         string   `  | @Ident)`
}

var (
	graphQLLexer = lexer.Must(ebnf.New(`
    Comment = ("//") { "\u0000"…"\uffff"-"\n" } .
    Ident = (alpha | "_") { "_" | alpha | digit } .
    Number = ("." | digit) {"." | digit} .
    Whitespace = " " | "\t" | "\n" | "\r" .
    Punct = "!"…"/" | ":"…"@" | "["…` + "\"`\"" + ` | "{"…"~" .

    alpha = "a"…"z" | "A"…"Z" .
    digit = "0"…"9" .
`))

	parser = participle.MustBuild(&RustAst{},
		participle.Lexer(graphQLLexer),
		participle.Elide("Comment", "Whitespace"),
	)

	cli struct {
		Files []string `arg:"" type:"existingfile" required:"" help:"GraphQL schema files to parse."`
	}
)
