package main

import (
	"hack_assembler/symboltable"
	"hack_assembler/code"
	"hack_assembler/parser"
	"os"
	"bufio"

	"fmt"
)

func main() {
	// parse input option, open input file and target file for write
	inputFile := os.Args[1]
	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	// init symboltable and translator
	st := symboltable.SymbolTable{}
	st.InitSymbolTable()
	t := code.Translator{}
	t.InitTranslator(st)

	// read file to string slice, and add label symbol
	scanner := bufio.NewScanner(input)
	pc := 0
	inputRows := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// ignore empty line
		if len(line) == 0 {
			continue
		}
		// ignore line comment
		if line[0] == '/' {
			continue
		}

		targLine := ""
		for _, c := range line {
			// ignore comment
			if c == '/' {
				break
			}

			if c != ' ' {
				targLine += string(c)
			}
		}

		// check if this is a label
		if targLine[0] == '(' {
			st.SetLabelSymbolValue(targLine[1:len(targLine) - 1], pc)
			continue
		}

		inputRows = append(inputRows, targLine)
		pc += 1
	}

	// parse each line, translate and write to targetfile.
	statements := parser.Parse(inputRows)

	mcodes := make([]string, 0)
	for _, statement := range statements {
		mcodes = append(mcodes, t.Translate(statement))
	}

	// test
	fmt.Println(inputRows[127])
	fmt.Println(statements[127])
	fmt.Pritnln(mcodes[127])

	outputFile, err := os.OpenFile("out.hack", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(nil)
	}

	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	for _, mcode := range mcodes {
		_, _ = writer.WriteString(mcode + "\n")
	}

	writer.Flush()
}
