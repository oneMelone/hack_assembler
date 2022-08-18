package code

import (
	"strconv"
)

const (
	INSTRUCTION_SIZE = 16
)

func translate(statement Statement) string {
	if statement.InstructionType == parser.AInstruction {
		return translateA(statement)
	} else {
		return translateC(statement)
	}
}

func translateA(statement Statement) string {
	var address int
	if address, ok := GetSymbolValue(statement.Address); !ok {
		address, err := strconv.Atoi(statement.Address)
		if err != nil {
			panic("unknown symbol")
		}
	}

	binaryRep := strconv.FormatInt(address, 2)

	for len(binaryRep) < INSTRUCTION_SIZE {
		binaryRep = "0" + binaryRep
	}

	return binaryRep
}
