package day11

import (
	"strings"
	"testing"

	"github.com/fastcat/aoc2021/util"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      Board
		assertion assert.ErrorAssertionFunc
	}{
		{
			"1x1",
			"1",
			Board{[]int{1}, 1, 1},
			nil,
		},
		{
			"2x2",
			"12\n34\n",
			Board{[]int{1, 2, 3, 4}, 2, 2},
			nil,
		},
		{
			"bogus",
			"12\n3a",
			Board{[]int{1, 2, 3, 0}, 2, 2},
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if tt.assertion == nil {
				tt.assertion = assert.NoError
			}
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBoard_Step(t *testing.T) {
	tests := []struct {
		name        string
		before      Board
		wantFlashes int
		after       Board
	}{
		{
			"1x1",
			MustParse("1"),
			0,
			MustParse("2"),
		},
		{
			"3x3",
			MustParse("000\n090\n000"),
			1,
			MustParse("222\n202\n222"),
		},
		{
			"2x2c",
			MustParse("88\n99"),
			4,
			MustParse("00\n00"),
		},
		{
			"e1",
			MustParse("11111\n19991\n19191\n19991\n11111"),
			9,
			MustParse("34543\n40004\n50005\n40004\n34543"),
		},
		{
			"e2",
			MustParse("34543\n40004\n50005\n40004\n34543"),
			0,
			MustParse("45654\n51115\n61116\n51115\n45654"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{}
			*b = tt.before
			assert.Equal(t, tt.wantFlashes, b.Step())
			assert.Equal(t, tt.after, *b)
		})
	}
}

func assertBoard(t *testing.T, expected, actual Board) bool {
	t.Helper()
	ok := assert.Equal(t, expected, actual)
	if !ok {
		t.Logf("Expected:\n%s\nActual:\n%s\n", &expected, &actual)
	}
	return ok
}

func TestPart1Example(t *testing.T) {
	b, err := Parse(exampleInput)
	assert.NoError(t, err)
	assert.Equal(t, strings.TrimSpace(exampleInput), b.String())
	steps := util.Stanzas(exampleStepsInput)
	tf := 0
	for i := 1; i <= 100; i++ {
		f := b.Step()
		tf += f
		if i <= 10 {
			assertBoard(t, MustParse(steps[i-1]), b)
		} else if i%10 == 0 {
			assertBoard(t, MustParse(steps[10+(i-20)/10]), b)
		}
	}
	assert.Equal(t, 1656, tf)
}

func TestPart1Challenge(t *testing.T) {
	b, err := Parse(challengeInput)
	assert.NoError(t, err)
	tf := 0
	for i := 1; i <= 100; i++ {
		tf += b.Step()
	}
	t.Logf("challenge total flashes = %d", tf)
}
