package main

import (
	"hack_assembler/symboltable"
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

	// init symboltable
	st := symboltable.SymbolTable{}
	st.InitSymbolTable()

	// read file to string slice, and add label symbol
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// parse each line, write to targetfile.
}
