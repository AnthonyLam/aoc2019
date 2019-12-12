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
	EXIT = 99
)


type IntcodeProgram struct{
	stack []int
	stored []int
	ip int
	stdin chan int
	stdout chan int
}

func NewProgram(in chan string, stdin chan int, stdout chan int) *IntcodeProgram {
	program := new(IntcodeProgram)
	program.ip = 0
	for k := range in {
		program.stack = append(program.stack, split(k)...)
	}
	program.stdin = stdin
	program.stdout = stdout
	return program
}

func split(k string) []int {
	var stack []int
	strs := strings.Split(k, ",")

	// Convert all strings to int
	for _, s := range strs {
		x, _ := strconv.Atoi(s)
		stack = append(stack, x)
	}
	return stack
}

func (pr *IntcodeProgram) Backup() {
	pr.stored = make([]int, len(pr.stack))
	copy(pr.stored, pr.stack)
}

func (pr *IntcodeProgram) Restore() {
	pr.ip = 0
	copy(pr.stack, pr.stored)
}

func (pr *IntcodeProgram) Modify(ip int, val int) {
	pr.stack[ip] = val
}


func (pr *IntcodeProgram) Run() []int {
	for {
		ins := pr.getInstruction()
		switch ins.opcode {
		case ADD:
			pr.increment()
			x := ins.param1.read(pr)
			pr.increment()
			y := ins.param2.read(pr)
			pr.increment()
			ins.param3.write(pr, x+y)
			pr.increment()
		case MULT:
			pr.increment()
			x := ins.param1.read(pr)
			pr.increment()
			y := ins.param2.read(pr)
			pr.increment()
			ins.param3.write(pr, x*y)
			pr.increment()
		case INP:
			pr.increment()
			ins.param1.write(pr, <-pr.stdin)
			pr.increment()
		case OUT:
			pr.increment()
			pr.stdout <- ins.param1.read(pr)
			pr.increment()
		case JIT:
			pr.increment()
			v := ins.param1.read(pr)
			pr.increment()
			if v != 0 {
				pr.jump(ins.param2.read(pr))
			} else {
				pr.increment()
			}
		case JIF:
			pr.increment()
			v := ins.param1.read(pr)
			pr.increment()
			if v == 0 {
				pr.jump(ins.param2.read(pr))
			} else {
				pr.increment()
			}
		case LE:
			pr.increment()
			v := ins.param1.read(pr)
			pr.increment()
			w := ins.param2.read(pr)
			pr.increment()
			if v < w {
				ins.param3.write(pr, 1)
			} else {
				ins.param3.write(pr, 0)
			}
			pr.increment()
		case EQ:
			pr.increment()
			v := ins.param1.read(pr)
			pr.increment()
			w := ins.param2.read(pr)
			pr.increment()
			if v == w {
				ins.param3.write(pr, 1)
			} else {
				ins.param3.write(pr, 0)
			}
			pr.increment()
		default:
		case EXIT:
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
	if i % 10 == 1 {ins.param1 = Immediate{}} else {ins.param1 =Position{}}
	i = i/10
	if i % 10 == 1 {ins.param2 = Immediate{}} else {ins.param2 =Position{}}
	i = i/10
	if i % 10 == 1 {ins.param3 = Immediate{}} else {ins.param3 =Position{}}
	return ins
}
