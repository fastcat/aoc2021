package day10

import (
	"errors"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var chunkRules = lexer.Rules{
	"Expr": {
		{Name: "POpen", Pattern: `\(`, Action: lexer.Push("PExpr")},
		{Name: "SOpen", Pattern: `\[`, Action: lexer.Push("SExpr")},
		{Name: "COpen", Pattern: `\{`, Action: lexer.Push("CExpr")},
		{Name: "AOpen", Pattern: `\<`, Action: lexer.Push("AExpr")},
	},
	"PExpr": {
		{Name: "PClose", Pattern: `\)`, Action: lexer.Pop()},
		lexer.Include("Expr"),
	},
	"SExpr": {
		{Name: "SClose", Pattern: `\]`, Action: lexer.Pop()},
		lexer.Include("Expr"),
	},
	"CExpr": {
		{Name: "CClose", Pattern: `\}`, Action: lexer.Pop()},
		lexer.Include("Expr"),
	},
	"AExpr": {
		{Name: "AClose", Pattern: `\>`, Action: lexer.Pop()},
		lexer.Include("Expr"),
	},
}

var chunkLexer = lexer.MustStateful(chunkRules, lexer.InitialState("Expr"))

type Chunk struct {
	Open      string `parser:"(@POpen | @SOpen | @COpen | @AOpen)"`
	ChunkList `parser:"@@"`
	Close     string `parser:"(@PClose | @SClose | @CClose | @AClose)"`
}

type ChunkList struct {
	Contents []Chunk `parser:"@@*"`
}

func Chunks(contents ...Chunk) ChunkList {
	return ChunkList{contents}
}

func ChunkP(contents ...Chunk) Chunk {
	return Chunk{"(", Chunks(contents...), ")"}
}
func ChunkS(contents ...Chunk) Chunk {
	return Chunk{"[", Chunks(contents...), "]"}
}
func ChunkC(contents ...Chunk) Chunk {
	return Chunk{"{", Chunks(contents...), "}"}
}
func ChunkA(contents ...Chunk) Chunk {
	return Chunk{"<", Chunks(contents...), ">"}
}

func (c Chunk) String() string {
	b := strings.Builder{}
	b.WriteString(c.Open)
	for _, i := range c.Contents {
		b.WriteString(i.String())
	}
	b.WriteString(c.Close)
	return b.String()
}

func (c ChunkList) String() string {
	b := strings.Builder{}
	for _, i := range c.Contents {
		b.WriteString(i.String())
	}
	return b.String()
}

var chunksParser = participle.MustBuild(&ChunkList{}, participle.Lexer(chunkLexer))

func IsIncomplete(err error) (bool, participle.Error) {
	var ut participle.UnexpectedTokenError
	if !errors.As(err, &ut) {
		return false, nil
	}
	return true, ut
}

func IsInvalid(err error) (bool, participle.Error) {
	var pe participle.Error
	if !errors.As(err, &pe) {
		return false, nil
	}
	return strings.Contains(pe.Message(), "invalid input text"), pe
}
