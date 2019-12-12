package run

import (
	"github.com/AnthonyLam/aoc2019/run/intcode"
)

func Seven(in chan string, _, output chan int) {
	perms := permutations([]int{0,1,2,3,4}, 5)
	max := 0
	program := intcode.NewProgram(in)
	program.Backup()
	for _, p := range perms {
		signal := 0
		for _, i := range p {
			stdin := make(chan int, 5)
			stdout := make(chan int, 5)
			stdin <- i
			stdin <- signal
			program.Run(stdin, stdout)
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

func Seven2(in chan string, _, output chan int) {
	const length = 5
	base := intcode.NewProgram(in)
	max := 0
	perms := permutations([]int{5, 6, 7, 8, 9}, length)
	programs := make([]*intcode.IntcodeProgram, length)
	stdins := make([]chan int, length)



	for _, p := range perms {
		for i:=0 ; i<length; i++ {
			stdins[i] = make(chan int, length)
			programs[i] = base.Copy()
			stdins[i] <- p[i]
		}

		stdins[0] <- 0

		go programs[0].Run(stdins[0], stdins[1])
		go programs[1].Run(stdins[1], stdins[2])
		go programs[2].Run(stdins[2], stdins[3])
		go programs[3].Run(stdins[3], stdins[4])
		programs[4].Run(stdins[4], stdins[0])
		v := <- stdins[0]
		if max < v {
			max = v
		}
	}
	output <- max
	close(output)
}

func permutations(seed []int, length int) [][]int {
	output := [][]int{}
	if length == 1 {
		n := make([]int, len(seed))
		copy(n, seed)
		return [][]int{n}
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
