package run

import (
	"fmt"
	"strconv"
	"strings"
)

type Color int

const (
	Black Color = 0
	White Color = 1
	Transparent Color = 2
);

type Layer [][]Color

func (l Layer) print() {
	for _, y := range l {
		for _, x := range y {
			switch x {
			case Black:
				fmt.Print("_")
			case White:
				fmt.Print("x")
			}
		}
		fmt.Println()
	}
}

func getLayers(in chan string) []Layer {
	size := strings.Split(<-in, ",")
	input := <-in

	width, _ := strconv.Atoi(size[0])
	height, _ := strconv.Atoi(size[1])

	layers := make([]Layer, 0)
	x := 0
	y := 0
	layer := make(Layer, height)
	layer[0] = make([]Color, width)
	for _,s := range input {
		if x == width {
			x = 0
			y++
			if y == height {
				y = 0
				layers = append(layers, layer)
				layer = make(Layer, height)
			}
			layer[y] = make([]Color, width)
		}
		layer[y][x] = Color(s - 48)
		x++
	}
	layers = append(layers, layer)
	return layers
}

func Eight(in chan string) interface{} {
	const upTo = 3
	layers := getLayers(in)
	maxes := make([]int, upTo, upTo)
	maxes[Black] = len(layers[0]) * len(layers[0][0])
	curr := make([]int, upTo, upTo)

	for _, layer := range layers {
		for _, y := range layer {
			for _, x := range y {
				if x < upTo {
					curr[x] += 1
				}
			}
		}
		if curr[Black] < maxes[Black] {
			copy(maxes, curr)
		}
		curr = make([]int, upTo, upTo)
	}
	return maxes[White] * maxes[Transparent]
}


func Eight2(in chan string) interface{} {
	layers := getLayers(in)
	image := make(Layer, len(layers[0]))
	for i := range image {
		image[i] = make([]Color, len(layers[0][0]))
		copy(image[i], layers[0][i])
	}
	for _, layer := range layers {
		for i, y := range layer {
			for k, x := range y {
				if image[i][k] == Transparent {
					image[i][k] = x
				}
			}
		}
	}
	image.print()
	return -1
}