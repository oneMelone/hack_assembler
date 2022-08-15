package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	input := []string{
		"@17",
		"D=M+1",
		"@key",
		"M;JLT",
		"@3",
		"0",
	}
	statements := Parse(input)
	fmt.Printf("len of statements is %+v\n", len(statements))
	for _, s := range statements {
		fmt.Printf("%+v\n", s)
	}
}
