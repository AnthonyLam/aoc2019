package run

import "github.com/AnthonyLam/aoc2019/run/intcode"

func Two(in chan string) interface{} {
	program := intcode.NewProgram(in, nil, nil)
	return program.Run()
}

func Two2(in chan string) interface{} {
	program := intcode.NewProgram(in, nil, nil)

	// initialize our machine
	program.Backup()
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			program.Modify(1, noun)
			program.Modify(2, verb)
			v := program.Run()
			if v[0] == 19690720 {
				return []int{noun, verb}
			}
			program.Restore()
		}
	}
	return nil
}


