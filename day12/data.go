package day12

import (
	_ "embed"

	"github.com/fastcat/aoc2021/util"
)

//go:embed example-inputs.txt
var exampleInputs string
var examplesInputStanzas = util.Stanzas(exampleInputs)

//go:embed example-paths.txt
var examplePathsInput string
var examplesPaths = util.Stanzas(examplePathsInput)

//go:embed challenge.txt
var challengeInput string
