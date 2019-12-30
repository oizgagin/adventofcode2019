package main

import (
	"fmt"

	"github.com/oizgagin/adventofcode2019/common"
)

func main() {
	solver := common.NewSolver(parse, solve1, solve2)
	fmt.Println(solver.Solve())
}

func parse(lines []string) (interface{}, error) {
	return nil, nil
}

func solve1(v interface{}) interface{} {
	return 0
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
		sum += digit(pos, i, vs[i])
	}
	return ones(sum)
}

func digit(pos, i, elem int) int {
	return elem * pattern(pos, i)
}

func pattern(pos, i int) int {
	return 0
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
