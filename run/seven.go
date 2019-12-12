package run

import (
	"github.com/AnthonyLam/aoc2019/run/intcode"
)

func Seven(in chan string, userinput, output chan int) {
	perms := permutations([]int{0,1,2,3,4}, 5)
	max := 0
	var stdin chan int = make(chan int, 5)
	var stdout chan int = make(chan int, 5)
	program := intcode.NewProgram(in, stdin, stdout)
	program.Backup()
	for _, p := range perms {
		signal := 0
		for _, i := range p {
			stdin <- i
			stdin <- signal
			program.Run()
			program.Restore()
			signal = <-stdout
		}
		if max < signal {
			max = signal
		}
	}
	output <- max
	close(output)
}

func Seven2(in chan string, userinput, output chan int) {
	const length = 5
	max := 0
	base := intcode.NewProgram(in, nil, nil)

	perms := permutations([]int{5, 6, 7, 8, 9}, length)
	programs := make([]*intcode.IntcodeProgram, length)
	stdins := make([]chan int, length)
	stdouts := make([]chan int, length)

	for i:=0 ; i<length; i++ {
		stdins[i] = make(chan int, length)
		stdouts[i] = make(chan int, length)
		programs[i] = base.Copy(stdins[i], stdouts[i])
	}

	for _, p := range perms {
		stdins[0] <- 0

	}
}

func permutations(seed []int, length int) [][]int {
	output := [][]int{}
	if length == 1 {
		n := make([]int, len(seed))
		copy(n, seed)
		return  [][]int{n}
	}
	output = append(output, permutations(seed, length-1)...)

	for i := 0; i < length-1; i++ {
		if length % 2 == 0 {
			swap(&seed, i, length-1)
		} else {
			swap(&seed, 0, length-1)
		}
		output = append(output, permutations(seed, length-1)...)
	}
	return output
}

func swap(seed *[]int, k,v int) {
	tmp := (*seed)[k]
	(*seed)[k] = (*seed)[v]
	(*seed)[v] = tmp
}
