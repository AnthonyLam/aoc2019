package intcode

type Mode interface {
	read(*IntcodeProgram) int
	write(*IntcodeProgram, int)
}

type Immediate struct{}

func (Immediate) read(program *IntcodeProgram) int {
	return program.stack[program.ip]
}

// This should never be called as we don't write in Imm mode
func (Immediate) write(program *IntcodeProgram, value int) {
	program.stack[program.ip] = value
	panic("Encountered a faulty write, segfault?")
}

type Position struct{}

func (Position) read(program *IntcodeProgram) int {
	return program.stack[program.stack[program.ip]]
}

func (Position) write(program *IntcodeProgram, value int) {
	program.stack[program.stack[program.ip]] = value
}

