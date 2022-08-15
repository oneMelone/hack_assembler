package parser

const (
	AInstruction = 0
	CInstruction = 1
)

type Statement struct {
	InstructionType int
	Address         string
	Dest            string
	Comp            string
	Jump            string
}
