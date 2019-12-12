package run

import "github.com/AnthonyLam/aoc2019/run/intcode"

func Five(in chan string, userinput chan int, output chan int)  {
	var program = intcode.NewProgram(in)
	program.Run(userinput, output)
}
