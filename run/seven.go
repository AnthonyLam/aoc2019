package run

import "fmt"

func Seven(in chan string, userinput, output chan int) {
	out := permutations([]int{0,1,2,3,4}, 5)
	fmt.Println(out)
}

func permutations(seed []int, length int) [][]int {
	output := [][]int{}
	if length == 1 {
		return  [][]int{seed}
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
