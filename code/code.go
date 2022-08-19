package code

import (
	"strconv"
	"hack_assembler/symboltable"
	"hack_assembler/parser"

	"fmt"
)

const (
	INSTRUCTION_SIZE = 16
)

type Translator struct {
	compMap map[string]string
	destMap map[string]string
	jumpMap map[string]string
	symbolTable symboltable.SymbolTable
}

func (t *Translator)InitTranslator(st symboltable.SymbolTable) {
	t.compMap = make(map[string]string)
	t.compMap["0"] = "0101010"
	t.compMap["1"] = "0111111"
	t.compMap["-1"] = "0111010"
	t.compMap["D"] = "0001100"
	t.compMap["A"] = "0110000"
	t.compMap["M"] = "1110000"
	t.compMap["!D"] = "0001101"
	t.compMap["!A"] = "0110001"
	t.compMap["!M"] = "1110001"
	t.compMap["-D"] = "0001111"
	t.compMap["-A"] = "0110011"
	t.compMap["-M"] = "1110011"
	t.compMap["D+1"] = "0011111"
	t.compMap["A+1"] = "0110111"
	t.compMap["M+1"] = "1110111"
	t.compMap["D-1"] = "0001110"
	t.compMap["A-1"] = "0110010"
	t.compMap["M-1"] = "1110010"
	t.compMap["D+A"] = "0000010"
	t.compMap["D+M"] = "1000010"
	t.compMap["D-A"] = "0010011"
	t.compMap["D-M"] = "1010011"
	t.compMap["A-D"] = "0000111"
	t.compMap["M-D"] = "1000111"
	t.compMap["D&A"] = "0000000"
	t.compMap["D&M"] = "1000000"
	t.compMap["D|A"] = "0010101"
	t.compMap["D|M"] = "1010101"

	t.destMap = make(map[string]string)
	t.destMap[""] = "000"
	t.destMap["M"] = "001"
	t.destMap["D"] = "010"
	t.destMap["DM"] = "011"
	t.destMap["A"] = "100"
	t.destMap["AM"] = "101"
	t.destMap["AD"] = "110"
	t.destMap["ADM"] = "111"

	t.jumpMap = make(map[string]string)
	t.jumpMap[""] = "000"
	t.jumpMap["JGT"] = "001"
	t.jumpMap["JEQ"] = "010"
	t.jumpMap["JGE"] = "011"
	t.jumpMap["JLT"] = "100"
	t.jumpMap["JNE"] = "101"
	t.jumpMap["JLE"] = "110"
	t.jumpMap["JMP"] = "111"

	t.symbolTable = st
}

func (t *Translator)Translate(statement parser.Statement) string {
	if statement.InstructionType == parser.AInstruction {
		return t.translateA(statement)
	} else {
		return t.translateC(statement)
	}
}

func (t *Translator)translateA(statement parser.Statement) string {
	var address int
	var ok bool
	var err error
	if address, ok = t.symbolTable.GetSymbolValue(statement.Address); !ok {
		address, err = strconv.Atoi(statement.Address)
		if err != nil {
			address = t.symbolTable.GetVarSymbolValue(statement.Address)
		}
	}

	binaryRep := strconv.FormatInt(int64(address), 2)

	for len(binaryRep) < INSTRUCTION_SIZE {
		binaryRep = "0" + binaryRep
	}

	return binaryRep
}

func (t *Translator)translateC(statement parser.Statement) string {
	// ----test----
	fmt.Println("before translate, map is")
	fmt.Println("compMap:", t.compMap)
	fmt.Println("destMap:", t.destMap)
	fmt.Println("jumpMap:", t.jumpMap)

	result := "111"

	// ----test-----
	fmt.Println("in translator")
	fmt.Println("statement.Comp is", statement.Comp)
	fmt.Println("comp is", t.compMap[statement.Comp])
	fmt.Println("statement.Dest is", statement.Dest)
	fmt.Println("dest is", t.compMap[statement.Dest])
	fmt.Println("statement.Jump is", statement.Jump)

	result += t.compMap[statement.Comp] + t.destMap[statement.Dest] + t.jumpMap[statement.Jump]
	return result
}
