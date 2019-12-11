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
	var m [][]int

	i := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		m = append(m, make([]int, len(line)))
		for j := 0; j < len(line); j++ {
			if line[j] == '.' {
				m[i][j] = 0
			} else if line[j] == '#' {
				m[i][j] = 1
			} else {
				return nil, fmt.Errorf("unexpected symbol: %v", line[j])
			}
		}

		i++
	}

	return m, nil
}

func solve1(v interface{}) int {
	m := v.([][]int)

	max := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 0 {
				continue
			}
			if detected := search(m, i, j); max == 0 || detected > max {
				max = detected
			}
		}
	}
	return max
}

func search(m [][]int, i, j int) int {
	return 0
}

func solve2(v interface{}) int {
	m := v.([][]int)
	fmt.Println(m)
	return 0
}
