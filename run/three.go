package run

import (
	"regexp"
	"strconv"
)

type Segment struct {
	Vertical bool
	StartX, EndX int
	StartY, EndY int
}

func parse(in chan string) ([]Segment,[]Segment) {
	re := regexp.MustCompile(`([UDLR])(\d+),?`)
	lines := make([][][]string, 0)
	for k := range in {
		vals := re.FindAllStringSubmatch(k, -1)
		lines = append(lines, vals) // [["R98", "R", "98"], ...]
	}
	return CreateSegments(lines[0]), CreateSegments(lines[1])
}

func Three(in chan string) interface{} {
	wire1, wire2 := parse(in)
	x := 9999999
	for _, w1 := range wire1 {
		for _, w2 := range wire2 {

			// Check intersection based on which wire is vertical
			var m int
			if w1.Vertical {
				d,b := checkIntersects(w1, w2)
				m = abs(d) + abs(b)
			} else {
				d,b := checkIntersects(w2, w1)
				m = abs(d) + abs(b)
			}

			// Assign the minimum
			if m != 0 && m < x {
				x = m
			}
		}
	}
	return x
}

func Three2(in chan string) interface{} {
	wire1, wire2 := parse(in)
	length1, length2 := 0, 0
	for _, w1 := range wire1 {
		for _, w2 := range wire2 {

			// Check intersection based on which wire is vertical
			var d int
			var b int
			if w1.Vertical {
				d,b = checkIntersects(w1, w2)
			} else {
				d,b = checkIntersects(w2, w1)
			}

			// Finding the running length!!
			if d+b != 0  {
				return length(Segment{
					Vertical: w1.Vertical,
					StartY: w1.StartY,
					EndY: b,
					StartX: w1.StartX,
					EndX: d,
				}) +
					length(Segment{
						Vertical: w2.Vertical,
						StartY: w2.StartY,
						EndY: b,
						StartX: w2.StartX,
						EndX: d,
					}) + length1 + length2

			}
			length2 += length(w2) // Keep track of current wire length
		}
		length1 += length(w1) // Keep track of current wire length
		length2 = 0
	}
	return 0
}

// Build out all possible line segments
func CreateSegments(wire [][]string) []Segment {
	xPos, yPos := 0, 0
	segments := make([]Segment, len(wire))
	for x, seg := range wire {

		currentSegment := Segment{}
		currentSegment.StartX = xPos
		currentSegment.StartY = yPos
		distance, _ := strconv.Atoi(seg[2])
		switch seg[1] {
		case "R":
			currentSegment.Vertical = false
			xPos += distance
		case "L":
			currentSegment.Vertical = false
			xPos -= distance
		case "U":
			currentSegment.Vertical = true
			yPos += distance
		case "D":
			currentSegment.Vertical = true
			yPos -= distance
		}
		currentSegment.EndX = xPos
		currentSegment.EndY = yPos
		segments[x] = currentSegment
	}
	return segments
}

// TODO optimize this to not need to create a h and v type
func checkIntersects(a,b Segment) (int,int) {
	if a.Vertical != b.Vertical  {
		if between(a.StartX, b.StartX, b.EndX) && between(b.StartY, a.StartY, a.EndY) {
			return a.StartX, b.StartY
		}
	}
	return 0,0
}

func between(a int, x,y int) bool {
	if x > y {
		return a <= x && a >= y
	} else {
		return a >= x && a <= y
	}
}

func abs(f int) int {
	if f <0 {
		return -1 * f
	}
	return f
}

func length(a Segment) int {
	s,e := 0,0
	if a.Vertical {
		s = a.StartY
		e = a.EndY
	} else {
		s = a.StartX
		e = a.EndX
	}
	if s > e {
		return abs(s - e)
	} else {
		return abs(e - s)
	}
}