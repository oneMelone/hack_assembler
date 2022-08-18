package code

import (
	"strconv"
	"hack_assembler/symboltable"
	"hack_assembler/parser"
)

const (
	INSTRUCTION_SIZE = 16
)

var (
	compMap map[string]string
	destMap map[string]string
	jumpMap map[string]string
	symbolTable symboltable.SymbolTable
)

func InitTranslator(st symboltable.SymbolTable) {
	compMap := make(map[string]string)
	compMap["0"] = "0101010"
	compMap["1"] = "0111111"
	compMap["-1"] = "0111010"
	compMap["D"] = "0001100"
	compMap["A"] = "0110000"
	compMap["M"] = "1110000"
	compMap["!D"] = "0001101"
	compMap["!A"] = "0110001"
	compMap["!M"] = "1110001"
	compMap["-D"] = "0001111"
	compMap["-A"] = "0110011"
	compMap["-M"] = "1110011"
	compMap["D+1"] = "0011111"
	compMap["A+1"] = "0110111"
	compMap["M+1"] = "1110111"
	compMap["D-1"] = "0001110"
	compMap["A-1"] = "0110010"
	compMap["M-1"] = "1110010"
	compMap["D+A"] = "0000010"
	compMap["D+M"] = "1000010"
	compMap["D-A"] = "0010011"
	compMap["D-M"] = "1010011"
	compMap["A-D"] = "0000111"
	compMap["M-D"] = "1000111"
	compMap["D&A"] = "0000000"
	compMap["D&M"] = "1000000"
	compMap["D|A"] = "0010101"
	compMap["D|M"] = "1010101"

	destMap := make(map[string]string)
	destMap[""] = "000"
	destMap["M"] = "001"
	destMap["D"] = "010"
	destMap["DM"] = "011"
	destMap["A"] = "100"
	destMap["AM"] = "101"
	destMap["AD"] = "110"
	destMap["ADM"] = "111"

	jumpMap := make(map[string]string)
	jumpMap[""] = "000"
	jumpMap["JGT"] = "001"
	jumpMap["JEQ"] = "010"
	jumpMap["JGE"] = "011"
	jumpMap["JLT"] = "100"
	jumpMap["JNE"] = "101"
	jumpMap["JLE"] = "110"
	jumpMap["JMP"] = "111"

	symbolTable = st
}

func translate(statement parser.Statement) string {
	if statement.InstructionType == parser.AInstruction {
		return translateA(statement)
	} else {
		return translateC(statement)
	}
}

func translateA(statement parser.Statement) string {
	var address int
	var ok bool
	var err error
	if address, ok = symbolTable.GetSymbolValue(statement.Address); !ok {
		address, err = strconv.Atoi(statement.Address)
		if err != nil {
			address = symbolTable.GetVarSymbolValue(statement.Address)
		}
	}

	binaryRep := strconv.FormatInt(int64(address), 2)

	for len(binaryRep) < INSTRUCTION_SIZE {
		binaryRep = "0" + binaryRep
	}

	return binaryRep
}

func translateC(statement parser.Statement) string {
	result := "111"
	result += compMap[statement.Comp] + destMap[statement.Dest] + jumpMap[statement.Jump]
	return result
}
