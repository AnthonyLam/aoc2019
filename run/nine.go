package run

import "github.com/AnthonyLam/aoc2019/run/intcode"

func Nine(in chan string, stdin, stdout chan int) {
	program := intcode.NewProgram(in)
	program.Run(stdin, stdout)
}
