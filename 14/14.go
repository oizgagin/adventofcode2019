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
	return 0, nil
}

func solve1(v interface{}) interface{} {
	return 0
}

func solve2(v interface{}) interface{} {
	return 0
}

type Ingridient struct {
	Amount int
	ID     string
}

type Reaction struct {
	To   *Ingridient
	From []*Ingridient
}

func parsereactions(s string) []Reaction {
	return nil
}
