package intcode

import (
	"strconv"
	"strings"
)

const (
	ADD = 1
	MULT = 2
	INP = 3
	OUT = 4
	JIT = 5
	JIF = 6
	LE = 7
	EQ = 8
	RBO = 9 // Relative Base Offset
	EXIT = 99
)

const RAM = 2 * 1024

type IntcodeProgram struct{
	stack        []int
	stored       []int
	ip           int
	relativeBase int
}

func NewProgram(in chan string) *IntcodeProgram {
	program := new(IntcodeProgram)
	program.ip = 0
	program.stack = make([]int, RAM)
	for i, s := range strings.Split(<-in, ",") {
		x, _ := strconv.Atoi(s)
		program.stack[i] = x
	}
	return program
}

func (pr *IntcodeProgram) Copy() *IntcodeProgram {
	program := new(IntcodeProgram)
	program.ip = pr.ip
	program.stack = make([]int, RAM)
	copy(program.stack, pr.stack)
	program.stored = make([]int, RAM)
	copy(program.stored, pr.stored)

	return program
}

func (pr *IntcodeProgram) Backup() {
	pr.stored = make([]int, RAM)
	copy(pr.stored, pr.stack)
}

func (pr *IntcodeProgram) Restore() {
	pr.ip = 0
	copy(pr.stack, pr.stored)
}

func (pr *IntcodeProgram) Modify(ip int, val int) {
	pr.stack[ip] = val
}


func (pr *IntcodeProgram) Run(stdin, stdout chan int) []int {
	for {
		ins := pr.getInstruction()
		switch ins.opcode {
		case ADD:
			pr.increment()
			x := ins.modes[Param1].read(pr)
			pr.increment()
			y := ins.modes[Param2].read(pr)
			pr.increment()
			ins.modes[Param3].write(pr, x+y)
			pr.increment()
		case MULT:
			pr.increment()
			x := ins.modes[Param1].read(pr)
			pr.increment()
			y := ins.modes[Param2].read(pr)
			pr.increment()
			ins.modes[Param3].write(pr, x*y)
			pr.increment()
		case INP:
			pr.increment()
			ins.modes[Param1].write(pr, <-stdin)
			pr.increment()
		case OUT:
			pr.increment()
			stdout <- ins.modes[Param1].read(pr)
			pr.increment()
		case JIT:
			pr.increment()
			v := ins.modes[Param1].read(pr)
			pr.increment()
			if v != 0 {
				pr.jump(ins.modes[Param2].read(pr))
			} else {
				pr.increment()
			}
		case JIF:
			pr.increment()
			v := ins.modes[Param1].read(pr)
			pr.increment()
			if v == 0 {
				pr.jump(ins.modes[Param2].read(pr))
			} else {
				pr.increment()
			}
		case LE:
			pr.increment()
			v := ins.modes[Param1].read(pr)
			pr.increment()
			w := ins.modes[Param2].read(pr)
			pr.increment()
			if v < w {
				ins.modes[Param3].write(pr, 1)
			} else {
				ins.modes[Param3].write(pr, 0)
			}
			pr.increment()
		case EQ:
			pr.increment()
			v := ins.modes[Param1].read(pr)
			pr.increment()
			w := ins.modes[Param2].read(pr)
			pr.increment()
			if v == w {
				ins.modes[Param3].write(pr, 1)
			} else {
				ins.modes[Param3].write(pr, 0)
			}
			pr.increment()
		case RBO:
			pr.increment()
			pr.relativeBase += ins.modes[Param1].read(pr)
			pr.increment()
		default:
		case EXIT:
			if stdout != nil {
				close(stdout)
			}
			return pr.stack
		}
	}
}

func (pr *IntcodeProgram) increment() {
	pr.ip++
}

func (pr *IntcodeProgram) jump(i int) {
	pr.ip = i
}

func (pr *IntcodeProgram) getInstruction() *Instruction {
	i := Immediate{}.read(pr)
	ins := new(Instruction)
	// Get first 2, shift right 2
	ins.opcode = i % 100
	i = i/100
	for m,_ := range ins.modes {
		switch i % 10 {
		case PositionMode:
			ins.modes[m] = Position{}
		case ImmediateMode:
			ins.modes[m] = Immediate{}
		case RelativeMode:
			ins.modes[m] = Relative{}
		}
		i = i/10
	}
	return ins
}
