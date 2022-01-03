package day12

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/fastcat/aoc2021/util"
)

type Vertex string

const (
	Start Vertex = "start"
	End   Vertex = "end"
)

func (v Vertex) Valid() bool {
	if len(v) < 1 {
		return false
	}
	upper := false
	for p, r := range v {
		if r >= utf8.RuneSelf || !unicode.IsLetter(r) {
			return false
		}
		if p == 0 {
			upper = unicode.IsUpper(r)
		} else if unicode.IsUpper(r) != upper {
			return false
		}
	}
	return true
}

func (v Vertex) Big() bool {
	return unicode.IsUpper(rune(v[0]))
}

type Graph map[Vertex]map[Vertex]bool

func (g Graph) AddVertex(v Vertex) {
	if !v.Valid() {
		panic(fmt.Errorf("invalid vertex '%s'", v))
	}
	if g[v] == nil {
		g[v] = map[Vertex]bool{}
	}
}

func (g Graph) AddEdge(a, b Vertex) {
	g.AddVertex(a)
	g.AddVertex(b)
	g[a][b] = true
	g[b][a] = true
}

func Parse(input string) (Graph, error) {
	g := Graph{}
	for _, l := range util.Lines(input) {
		ab := strings.SplitN(l, "-", 2)
		if len(ab) != 2 {
			return g, fmt.Errorf("invalid edge '%s'", l)
		}
		g.AddEdge(Vertex(ab[0]), Vertex(ab[1]))
	}
	if g[Start] == nil || g[End] == nil {
		return g, fmt.Errorf("missing start/end edges")
	}
	return g, nil
}

func MustParse(input string) Graph {
	g, err := Parse(input)
	if err != nil {
		panic(err)
	}
	return g
}

type Path []Vertex

func (p Path) Equal(other Path) bool {
	if len(p) != len(other) {
		return false
	}
	for i, v := range p {
		if v != other[i] {
			return false
		}
	}
	return true
}

func (p Path) Includes(v Vertex) bool {
	// TODO: this is inefficient
	// could keep a map[Vertex]int counter set and update that as we build new paths
	for _, vv := range p {
		if vv == v {
			return true
		}
	}
	return false
}

func (p Path) String() string {
	var b strings.Builder
	for i, v := range p {
		if i > 0 {
			b.WriteRune(',')
		}
		b.WriteString(string(v))
	}
	return b.String()
}

func (p Path) Append(v Vertex) Path {
	// TODO: concurrent walk would require we make clones or use a pool or something
	return append(p, v)
}

func (p Path) Copy() Path {
	pp := make(Path, len(p))
	copy(pp, p)
	return pp
}

func (g Graph) CountPaths(visitor func(Path)) int {
	p := Path{Start}
	return g.pathsFrom(p, visitor)
}

func (g Graph) pathsFrom(initial Path, visitor func(Path)) int {
	edges := g[initial[len(initial)-1]]
	count := -0
	for v, ok := range edges {
		if !ok {
			continue
		}
		if !v.Big() && initial.Includes(v) {
			continue
		}
		next := initial.Append(v)
		if v == End {
			count++
			if visitor != nil {
				visitor(next)
			}
		} else {
			count += g.pathsFrom(next, visitor)
		}
	}
	return count
}
