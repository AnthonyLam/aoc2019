package run

import (
	"os"
	"testing"
	"bufio"
	"reflect"
)

const inputDirectory string = "../input/"

func GetInputs(file *os.File, out chan string) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		out <- scanner.Text()
	}
	close(out)
}

var oneTests = []struct {
	file string
	exp  interface{}
	fn   func(chan string) interface{}
}{
// Day 1
	// Part 1
	{"1/one_1", 2, One},
	{"1/one_2", 2, One},
	{"1/one_3", 654, One},
	{"1/one_4", 33583, One},
	{"1/one_5", 3390830, One},
	// Part 2
	{"1/one_1", 2, One2},
	{"1/one_2", 2, One2},
	{"1/one_3", 966, One2},
	{"1/one_4", 50346, One2},
	{"1/one_5", 5083370, One2},

// Day 2
	// Part 1
	{"2/two_1", []int{2,0,0,0,99}, Two},
	{"2/two_2", []int{2,3,0,6,99}, Two},
	{"2/two_3", []int{2,4,4,5,99,9801}, Two},
	{"2/two_4", []int{30,1,1,4,2,5,6,0,99}, Two},

	//{"2/two_5", []int{}, Two}, // Position 0 should be 3931283

	// Part 2
	{"2/two_5", []int{69, 79}, Two2},

// Day 3
	// Part 1
	{"3/three_1", 6, Three},
	{"3/three_2", 159, Three},
	{"3/three_3", 135, Three},
	{"3/three_4", 5357, Three},

	// Part 2
	{"3/three_1", 30, Three2},
	{"3/three_2", 610, Three2},
	{"3/three_3", 410, Three2},
	{"3/three_4", 101956, Three2},

// Day 4
	// Part 1
	{"4/four_1", 1, Four},
	{"4/four_2", 0, Four},
	{"4/four_3", 0, Four},
	{"4/four_5", 1, Four},
	{"4/four_6", 1, Four},
	{"4/four_4", 1660, Four},

	// Part 1
	{"4/four_1", 0, Four2},
	{"4/four_2", 0, Four2},
	{"4/four_3", 0, Four2},
	{"4/four_5", 0, Four2},
	{"4/four_6", 1, Four2},
	{"4/four_4", 1135, Four2},

// Day 6
	// Part 1
	{"6/1", 42, Six},
	{"6/2", 162816, Six},

	// Part 2
	{"6/3", 4, Six2},
	{"6/2", 304, Six2},
}

func Test(t *testing.T) {
	for _, input := range oneTests {
		t.Run(input.file, func(t *testing.T) {
			inp := make(chan string, 20)
			of, err := os.Open(inputDirectory + input.file)
			if err != nil {
				t.Error("Could not open file")
			}
			go GetInputs(of, inp)
			result := input.fn(inp)
			if !reflect.DeepEqual(input.exp, result) {
				t.Errorf("got %d, want %d", result, input.exp)
			}

		})
	}
}


var opcodeTests = []struct{
	file string
	userInput []int
	exp interface{}
	fn func(chan string, chan int, chan int)
} {
// Day 5
	// Part 1
	//{"5/five_1", []int{1}, 13294380, Five},
	//
	//// Part 2
	//// is eq to 8
	//{"5/five_2", []int{0}, 0, Five},
	//{"5/five_2", []int{8}, 1, Five},
	//// is less than 8
	//{"5/five_3", []int{0}, 1, Five},
	//{"5/five_3", []int{7}, 1, Five},
	//{"5/five_3", []int{8}, 0, Five},
	//{"5/five_3", []int{9}, 0, Five},
	//
	//// is eq to 8
	//{"5/five_4", []int{0}, 0, Five},
	//{"5/five_4", []int{8}, 1, Five},
	//// is less than 8
	//{"5/five_5", []int{0}, 1, Five},
	//{"5/five_5", []int{7}, 1, Five},
	//{"5/five_5", []int{8}, 0, Five},
	//{"5/five_5", []int{9}, 0, Five},
	//
	//// is 0, position
	//{"5/five_6", []int{0}, 0, Five},
	//{"5/five_6", []int{8}, 1, Five},
	//
	//// is 0, immediate
	//{"5/five_7", []int{0}, 0, Five},
	//{"5/five_7", []int{8}, 1, Five},
	//
	//// is below 8
	//{"5/five_8", []int{0}, 999, Five},
	//{"5/five_8", []int{8}, 1000, Five},
	//{"5/five_8", []int{9}, 1001, Five},
	//
	//{"5/five_1", []int{5}, 11460760, Five},

// Day 7
	// Part 1
	{"7/1", []int{0}, 43210, Seven},
	{"7/2", []int{0}, 54321, Seven},
	{"7/3", []int{0}, 65210, Seven},
	{"7/4", []int{0}, 75228, Seven},

	// Part 2
	{"7/4", []int{0}, 139629729, Seven},
	{"7/4", []int{0}, 18216, Seven},
	{"7/1", []int{0}, 0, Seven},

}
func TestIntcode(t *testing.T) {

	for _, input := range opcodeTests {
		t.Run(input.file, func(t *testing.T) {

			// File input
			of, err := os.Open(inputDirectory + input.file)
			if err != nil {
				t.Error("Could not open file")
			}

			// "User" input
			user := make(chan int, 10)
			output := make(chan int, 20)
			inp := make(chan string, 20)
			result := make([]int, 0)

			go GetInputs(of, inp)
			go func() {
				for _, u := range input.userInput {
					user <- u
				}
				close(user)
			}()
			go input.fn(inp, user, output)

			// Outputs
			for i := range output {
				result = append(result, i)
			}

			if !reflect.DeepEqual(input.exp, result[len(result)-1]) {
				t.Errorf("got %d, want %d", result[len(result)-1], input.exp)
			}

		})
	}
}