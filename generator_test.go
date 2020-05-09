package er

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"regexp/syntax"
	"testing"
)

type parseUseCase struct {
	Pattern string
	IsError bool
}

var (
	parseUseCases = []parseUseCase{
		{
			"(?<=<li><input[^>]*>).*?(?=</li>)",
			true,
		}, {
			"(?<=<li><input[^>]{0,5}>).*?(?=</li>)",
			true,
		}, {
			"^[👍]$",
			false,
		}, {
			`(((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))\.){3}((\d{1,2})|(1\d{1,2})|(2[0-4]\d)|(25[0-5]))`,
			false,
		}, {
			`[2-9][0-9]{2}\.[0-9]{3}\.[0-9]{4}`,
			false,
		}, {
			`^([1-9][0-9]*)+(.[0-9]{1,2})?$`,
			false,
		}, {
			`^[0-9]+(.[0-9]{1,3})?$`,
			false,
		}, {
			`^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*))$`,
			false,
		}, {
			`^\p{Han}{0,}$`,
			false,
		}, {
			`^[\p{Han}A-Za-z0-9_]+$`,
			false,
		}, {
			`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`,
			false,
		}, {
			`[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(/.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+/.?`,
			false,
		}, {
			`^http://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?$`,
			false,
		}, {
			`^\d{15}|\d{18}$`,
			false,
		}, {
			`^\d{4}-\d{1,2}-\d{1,2}`,
			false,
		}, {
			`((?:(?:25[0-5]|2[0-4]\d|[01]?\d?\d)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d?\d))`,
			false,
		}, {
			`\d+\.\d+\.\d+\.\d+`,
			false,
		}, {
			`^[1-9a-m]{16,19}$`,
			false,
		}, {
			`^.+$`,
			false,
		}, {
			`\bnice\b`,
			false,
		}, {
			`(?s:.)`,
			false,
		}, {
			`$`,
			false,
		},
	}
)

func TestParse(t *testing.T) {
	a := assert.New(t)
	for _, useCase := range parseUseCases {
		_, err := Parse(useCase.Pattern, syntax.Perl)
		a.Equal(useCase.IsError, err != nil, useCase.Pattern, err)
	}
}

var (
	generateUseCases = []string{
		`[2-9][0-9]{2}\.[0-9]{3}\.[0-9]{4}`,
		`^([1-9][0-9]*)+(.[0-9]{1,2})?$`,
		`^[0-9]+(.[0-9]{1,3})?$`,
		`^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*))$`,
		`^[\p{Han}]{0,}$`,
		`^[\p{Han}A-Za-z0-9_]+$`,
		`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`,
		`[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(/.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+/.?`,
		`^http://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?$`,
		`^\d{15}|\d{18}$`,
		`^\d{4}-\d{1,2}-\d{1,2}`,
		`((?:(?:25[0-5]|2[0-4]\d|[01]?\d?\d)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d?\d))`,
		`\d+\.\d+\.\d+\.\d+`,
		`^[1-9a-m]{16,19}$`,
		`^.+$`,
		`\bnice\b`,
		`(?s:.)`,
		`$`,
	}
)

func TestGenerate(t *testing.T) {
	a := assert.New(t)
	for _, useCase := range generateUseCases {
		g, err := Parse(useCase, syntax.Perl)
		a.Nil(err, useCase, err)
		s, err := g.Generate()
		a.Nil(err, useCase, err)
		match, err := regexp.MatchString(useCase, s)
		a.Nil(err, useCase, err)
		a.True(match, useCase)
	}
}

func TestGenerateMultiple(t *testing.T) {
	a := assert.New(t)
	for _, useCase := range generateUseCases {
		g, err := Parse(useCase, syntax.Perl)
		a.Nil(err, useCase, err)
		_, err = g.GenerateMultiple(0)
		a.NotNil(err)
		ss, err := g.GenerateMultiple(10)
		a.Nil(err, useCase, err)
		for _, s := range ss {
			match, err := regexp.MatchString(useCase, s)
			a.Nil(err, useCase, err)
			a.True(match, useCase)
		}
	}
}
