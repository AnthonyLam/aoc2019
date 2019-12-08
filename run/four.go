package run

import (
	"regexp"
	"strconv"
)

func Four(in chan string) interface{} {
	return dofour(in, func(i int) bool {
		s := strconv.Itoa(i)
		var prev rune
		var hasDouble bool
		for _, c := range s {
			if prev == c {
				hasDouble = true
			} else if prev > c {
				return false
			}
			prev = c
		}
		return hasDouble
	})
}


func Four2(in chan string) interface{} {
	return dofour(in, func(i int) bool {
		s := strconv.Itoa(i)
		var prev rune
		var runlength []rune
		var hasDouble bool
		for _, c := range s {
			if prev == c {
				runlength = append(runlength, c)
			} else if prev > c {
				return false
			} else {
				if len(runlength) == 2 {
					hasDouble = true
				}
				runlength = []rune{prev}
			}
			prev = c
		}
		if len(runlength) == 2 {
			hasDouble = true
		}
		return hasDouble
	})
}

func dofour(in chan string, fn func(int) bool) int {
	re := regexp.MustCompile(`(\d{6})-(\d{6})`)
	var count int
	for k := range in {
		m := re.FindStringSubmatch(k)
		from, err := strconv.Atoi(m[1])
		if err != nil {
			return -1
		}
		to, err := strconv.Atoi(m[2])
		if err != nil {
			return -1
		}

		count = 0
		for i := from; i <= to; i++ {
			if fn(i) {
				count += 1
			}
		}
	}
	return count
}
