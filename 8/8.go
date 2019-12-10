package main

import (
	"fmt"
	"strings"

	"github.com/oizgagin/adventofcode2019/common"
)

func main() {
	solver := common.NewSolver(parse, solve1, solve2)
	fmt.Println(solver.Solve())
}

func parse(lines []string) (interface{}, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines in input")
	}
	return interface{}(lines[0]), nil
}

func solve1(v interface{}) int {
	line := strings.TrimSpace(v.(string))

	const (
		width     = 25
		height    = 6
		layersize = width * height
	)

	min, result := 0, 0
	for layer := 0; layer < len(line)/layersize; layer++ {
		offset := layer * layersize

		zeroes, ones, twos := 0, 0, 0
		for i := 0; i < layersize; i++ {
			switch line[offset+i] {
			case '0':
				zeroes++
			case '1':
				ones++
			case '2':
				twos++
			}
		}

		if zeroes < min || min == 0 {
			min = zeroes
			result = ones * twos
		}
	}

	return result
}

func solve2(v interface{}) int {
	line := strings.TrimSpace(v.(string))

	const (
		width     = 25
		height    = 6
		layersize = width * height
	)

	type color rune

	const (
		black       color = '0'
		white       color = '1'
		transparent color = '2'
	)

	merge := func(base, c color) color {
		if base == transparent {
			return c
		}
		return base
	}

	image := make([][]color, height)
	for i := 0; i < height; i++ {
		image[i] = make([]color, width)
	}

	layers := len(line) / layersize

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			pixel := transparent
			for k := 0; k < layers; k++ {
				pixel = merge(pixel, color(line[k*layersize+i*width+j]))
			}
			image[i][j] = pixel
		}
	}

	printlayer := func(image [][]color, height, width int) {
		fmt.Printf("[")
		for i := 0; i < height; i++ {
			if i == 0 {
				fmt.Printf("[")
			} else {
				fmt.Printf(" [")
			}
			for j := 0; j < width; j++ {
				fmt.Printf(" %s", string(image[i][j]))
			}
			fmt.Printf("]")
			if i == height-1 {
				fmt.Printf("]")
			}
			fmt.Printf("\n")
		}
	}

	printlayer(image, height, width)
	return 0
}
