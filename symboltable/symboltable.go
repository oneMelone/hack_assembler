package symboltable

import (
    "strconv"
)

// from string symbol to value
type SymbolTable struct {
	table map[string]int
	currentPtr int
}

func (s *SymbolTable) InitSymbolTable {
    s.table = make(map[string]int)
    for i := 0; i <= 15; i++ {
        symbol := "R" + strconv.Itoa(i)
        s.table[symbol] = i
    }
    s.table["SCREEN"] = 16384
    s.table["KBD"] = 24576
    s.table["SP"] = 0
    s.table["LCL"] = 1
    s.table["ARG"] = 2
    s.table["THIS"] = 3
    s.table["THAT"] = 4
    s.currentPtr = 16
}
