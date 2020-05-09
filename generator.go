package er

import (
	"bytes"
	"errors"
	"math/rand"
	"regexp/syntax"
	"time"
)

// generate
func generate(writer *bytes.Buffer, regexp *syntax.Regexp) error {
	switch regexp.Op {
	case syntax.OpNoMatch,
		syntax.OpEmptyMatch,
		syntax.OpNoWordBoundary,
		syntax.OpBeginLine,
		syntax.OpBeginText,
		syntax.OpEndText:

	case syntax.OpLiteral:
		writer.WriteString(string(regexp.Rune))

	case syntax.OpCharClass:
		if len(regexp.Rune)%2 != 0 || len(regexp.Rune) == 0 {
			return nil
		}
		writer.WriteRune(RandRune(regexp.Rune))

	case syntax.OpAnyCharNotNL:
		writer.WriteRune(RandAscii())

	case syntax.OpAnyChar:
		r := RandInRange(32, 128)
		if r == 127 {
			writer.WriteRune('\n')
		} else {
			writer.WriteRune(rune(r))
		}
	case syntax.OpEndLine:
		writer.WriteRune('\n')

	case syntax.OpWordBoundary:
		writer.WriteRune(0x20)

	case syntax.OpStar:
		return repeat(writer, regexp.Sub, 0, -1)

	case syntax.OpPlus:
		return repeat(writer, regexp.Sub, 1, -1)

	case syntax.OpQuest:
		if RandBool() {
			if err := traversal(writer, regexp.Sub); err != nil {
				return err
			}
		}
	case syntax.OpRepeat:
		return repeat(writer, regexp.Sub, regexp.Min, regexp.Max)

	case syntax.OpConcat, syntax.OpCapture:
		return traversal(writer, regexp.Sub)

	case syntax.OpAlternate:
		return generate(writer, regexp.Sub[RandInt(len(regexp.Sub))])
	}

	return nil
}

// Generate a random string according to *syntax.Regexp ( got by syntax.Parse() )
func Generate(regexp *syntax.Regexp) (string, error) {
	var buf bytes.Buffer
	err := generate(&buf, regexp)
	return buf.String(), err
}

func traversal(writer *bytes.Buffer, children []*syntax.Regexp) error {
	for _, child := range children {
		if err := generate(writer, child); err != nil {
			return err
		}
	}
	return nil
}

// repeat for [min, max) times
func repeat(writer *bytes.Buffer, list []*syntax.Regexp, min, max int) error {
	for count := RandRepeat(min, max); count > 0; count-- {
		if err := traversal(writer, list); err != nil {
			return err
		}
	}
	return nil
}

// Generator
type Generator interface {
	// Generator a random string
	Generate() (string, error)

	// Generator multiple strings
	GenerateMultiple(amount int) ([]string, error)
}

type generator struct {
	regexp *syntax.Regexp
}

// generate one string
func (g *generator) Generate() (string, error) {
	rand.Seed(time.Now().UnixNano())
	return Generate(g.regexp)
}

// generate multiple strings
func (g *generator) GenerateMultiple(amount int) ([]string, error) {
	rand.Seed(time.Now().UnixNano())
	if amount < 1 {
		return nil, errors.New("amount must be greater than or equal to 1")
	}
	result := make([]string, amount)
	var err error
	for i := 0; i < amount; i++ {
		if result[i], err = Generate(g.regexp); err != nil {
			return nil, err
		}
	}
	return result, nil
}

// Parse a pattern string and return a Generator if success
func Parse(pattern string, flags syntax.Flags) (Generator, error) {
	reg, err := syntax.Parse(pattern, flags)
	if err != nil {
		return nil, err
	}
	g := &generator{
		regexp: reg.Simplify(),
	}
	return g, nil
}
