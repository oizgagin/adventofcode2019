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
	line := strings.TrimSpace(lines[0])

	res := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		res[i] = int(line[i] - '0')
	}
	return res, nil
}

func solve1(v interface{}) interface{} {
	input := v.([]int)
	for i := 0; i < 100; i++ {
		input = phase(input)
	}
	return input[:8]
}

func solve2(v interface{}) interface{} {
	return 0
}

func phase(vs []int) []int {
	out := make([]int, len(vs))
	for i := 0; i < len(vs); i++ {
		out[i] = calc(i, vs)
	}
	return out
}

func calc(pos int, vs []int) int {
	sum := 0
	for i := 0; i < len(vs); i++ {
		sum += next(pos, i, vs[i])
	}
	return ones(sum)
}

func next(pos, i, elem int) int {
	return elem * pattern(pos, i)
}

var pat = []int{0, 1, 0, -1}

func pattern(pos, i int) int {
	return pat[((i+1)/(pos+1))%4]
}

func ones(x int) int {
	return abs(x) % 10
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
