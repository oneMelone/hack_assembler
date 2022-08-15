package parser

func Parse(input []string) []Statement {
	result := make([]Statement, len(input))

	for row, content := range input {
		if content[0] == '@' {
			result[row] = parseAIns(content)
		} else {
			result[row] = parseCIns(content)
		}
	}

	return result
}

func parseAIns(content string) Statement {
	return Statement{
		InstructionType: AInstruction,
		Address:         content[1:],
	}
}

func parseCIns(content string) Statement {
	equalPos := -1
	semicolonPos := -1
	for i, c := range content {
		if c == '=' {
			equalPos = i
		} else if c == ';' {
			semicolonPos = i
		}
	}

	destEndPos := 0
	compStartPos := 0
	compEndPos := len(content)
	jumpStartPos := len(content)
	jumpEndPos := len(content)
	if equalPos != -1 {
		destEndPos = equalPos
		compStartPos = equalPos + 1
	}
	if semicolonPos != -1 {
		compEndPos = semicolonPos
		jumpStartPos = semicolonPos + 1
	}
	dest := content[0:destEndPos]
	comp := content[compStartPos:compEndPos]
	jump := content[jumpStartPos:jumpEndPos]
	return Statement{
		InstructionType: CInstruction,
		Dest:            dest,
		Comp:            comp,
		Jump:            jump,
	}
}
