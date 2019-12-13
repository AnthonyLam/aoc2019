package intcode

const (
	ParamCount = 3
	Param1 = 0
	Param2 = 1
	Param3 = 2

)

type Instruction struct {
	modes  [ParamCount]Mode
	opcode int
}
