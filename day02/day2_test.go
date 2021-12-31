package day02

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Pos struct {
	H, D, A int
}

type verbTable = map[string]func(Pos, int) Pos

var verbsPart1 = verbTable{
	"forward": func(p Pos, i int) Pos {
		p.H += i
		return p
	},
	"down": func(p Pos, i int) Pos {
		p.D += i
		return p
	},
	"up": func(p Pos, i int) Pos {
		p.D -= i
		return p
	},
}

func UpdatePos(p Pos, cmd string, table verbTable) (Pos, error) {
	var verb string
	var delta int
	if n, err := fmt.Sscanf(cmd, "%s %d", &verb, &delta); err != nil {
		return p, fmt.Errorf("malformed command '%s': %w", cmd, err)
	} else if n != 2 {
		return p, fmt.Errorf("incomplete command '%s'", cmd)
	}
	if u, ok := table[verb]; !ok {
		return p, fmt.Errorf("invalid verb '%s'", verb)
	} else {
		return u(p, delta), nil
	}
}

func TestUpdatePosPart1(t *testing.T) {
	type args struct {
		p   Pos
		cmd string
	}
	tests := []struct {
		args      args
		want      Pos
		assertion assert.ErrorAssertionFunc
	}{
		{
			args{Pos{}, "up 5"},
			Pos{0, -5, 0},
			assert.NoError,
		},
		{
			args{Pos{}, "down 5"},
			Pos{0, 5, 0},
			assert.NoError,
		},
		{
			args{Pos{}, "forward 5"},
			Pos{5, 0, 0},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.cmd, func(t *testing.T) {
			got, err := UpdatePos(tt.args.p, tt.args.cmd, verbsPart1)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExamplePart1(t *testing.T) {
	input := "forward 5\n" +
		"down 5\n" +
		"forward 8\n" +
		"up 3\n" +
		"down 8\n" +
		"forward 2\n"
	p := applyCommands(t, input, verbsPart1)
	assert.Equal(t, Pos{15, 10, 0}, p)
}

func applyCommands(t *testing.T, input string, table verbTable) Pos {
	lines := util.Lines(input)
	p := Pos{}
	for i, l := range lines {
		pp, err := UpdatePos(p, l, table)
		assert.NoError(t, err, "line", i, l)
		p = pp
	}
	t.Logf("pos %#v multiplies as %d", p, p.H*p.D)
	return p
}

//go:embed input.txt
var input string

func TestChallenge1(t *testing.T) {
	require.NotEmpty(t, input)
	applyCommands(t, input, verbsPart1)
}

var verbsPart2 = verbTable{
	"forward": func(p Pos, i int) Pos {
		p.H += i
		p.D += i * p.A
		return p
	},
	"down": func(p Pos, i int) Pos {
		p.A += i
		return p
	},
	"up": func(p Pos, i int) Pos {
		p.A -= i
		return p
	},
}

func TestExamplePart2(t *testing.T) {
	input := "forward 5\n" +
		"down 5\n" +
		"forward 8\n" +
		"up 3\n" +
		"down 8\n" +
		"forward 2\n"
	p := applyCommands(t, input, verbsPart2)
	assert.Equal(t, Pos{15, 60, 10}, p)
}

func TestChallenge2(t *testing.T) {
	require.NotEmpty(t, input)
	applyCommands(t, input, verbsPart2)
}
