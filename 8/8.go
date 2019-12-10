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
	return 0
}
