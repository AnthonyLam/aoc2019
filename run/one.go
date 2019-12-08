package run

import (
	"math"
	"strconv"
)

func One(in chan string) interface{} {
	var sum int
	for k := range in {
		v, _ := strconv.Atoi(k)
		sum += int(math.Floor(float64(v)/3.0) - 2)
	}
	return sum
}

func One2(in chan string) interface{} {
	var sum int
	for k := range in {
		v, _ := strconv.Atoi(k)
		sum += sub2(v, 0)
	}
	return sum
}
func sub2(v int, total int) int {

	fuel := int(math.Floor(float64(v)/3.0) - 2)
	if fuel <= 0 {
		return total
	}
	return sub2(fuel, total+fuel)
}
