package main

import (
	"flag"
	"fmt"
	"github.com/unbyte/er"
	"os"
	"regexp/syntax"
	"strings"
)

var (
	pattern string
	amount  int
)

func init() {
	flag.StringVar(&pattern, "p", "", "pattern string")
	flag.StringVar(&pattern, "pattern", "", "pattern string")
	flag.IntVar(&amount, "a", 1, "amount of strings to be generated. default to 1.")
	flag.IntVar(&amount, "amount", 1, "amount of strings to be generated. default to 1.")
}

func main() {
	flag.Parse()
	if !(amount > 0 && len(pattern) > 0) {
		fmt.Println("PATTERN can't be empty and AMOUNT can't be lower than 1.")
		os.Exit(1)
	}

	generator, err := er.Parse(pattern, syntax.Perl)
	if err != nil {
		fmt.Println("Error when parse regexp: ", err.Error())
		os.Exit(1)
	}

	if amount == 1 {
		s, err := generator.Generate()
		if err != nil {
			fmt.Println("Error when generate: ", err.Error())
			os.Exit(1)
		}
		fmt.Println(s)
	} else {
		s, err := generator.GenerateMultiple(amount)
		if err != nil {
			fmt.Println("Error when generate: ", err.Error())
			os.Exit(1)
		}
		fmt.Println(strings.Join(s, "\n"))
	}
}
