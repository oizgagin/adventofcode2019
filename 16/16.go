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

func phase(v []int) []int {
	return nil
}
