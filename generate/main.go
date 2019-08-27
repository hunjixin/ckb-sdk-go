package main

import (
	"github.com/alecthomas/repr"
	"log"
	"os"

	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
)
type FuncOrSturcts struct {
	FuncOrSturct []*FuncOrSturct `@@*`
}

type FuncOrSturct struct {
	Type   *RustStruct   `@@`
	//Func *Func `| @@`
}

type RustStruct struct {
	IsPublic bool  	`@"pub"?`
	Name string    	`"struct" @Ident`
	Fields []*Field `"{" @@* "}"`
}

type Field struct {
	IsPublic bool   	`@"pub"?`
	Name     string      `@Ident`
	Type     TypeRef	 `":" @@ ","`
}

type TypeRef struct {
	Array       *TypeRef 	`( "Vec" "<" @@ ">"`
	NullAbleType *TypeRef     `  | "Option" "<" @@ ">"`
	Type         string   	`  | @Ident)`
}


var (
	graphQLLexer = lexer.Must(ebnf.New(`
    Comment = ("#" | "//") { "\u0000"…"\uffff"-"\n" } .
    Ident = (alpha | "_") { "_" | alpha | digit } .
    Number = ("." | digit) {"." | digit} .
    Whitespace = " " | "\t" | "\n" | "\r" .
    Punct = "!"…"/" | ":"…"@" | "["…`+"\"`\""+` | "{"…"~" .

    alpha = "a"…"z" | "A"…"Z" .
    digit = "0"…"9" .
`))

	parser = participle.MustBuild(&FuncOrSturcts{},
		participle.Lexer(graphQLLexer),
		participle.Elide("Comment", "Whitespace"),
	)

	cli struct {
		Files []string `arg:"" type:"existingfile" required:"" help:"GraphQL schema files to parse."`
	}
)

func main() {
		ast := &FuncOrSturcts{}
		r, err := os.Open("/home/hunjixin/project/ckb-sdk-go/generate/m.rs")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()

		err = parser.Parse(r, ast)
		if err != nil {
			log.Fatal(err)
		}

	repr.Println(ast)

}