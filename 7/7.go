package main

import (
	"fmt"

	"github.com/oizgagin/adventofcode2019/common"
	"github.com/oizgagin/adventofcode2019/intcode"
)

func main() {
	solver := common.NewSolver(parse, solve1, solve2)
	fmt.Println(solver.Solve())
}

func parse(lines []string) (interface{}, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines in input")
	}
	return interface{}(intcode.NewMemoryFromString(lines[0])), nil
}

func solve1(mem interface{}) int {
	amplifiers := 5

	max := 0
	for _, phase := range common.Permutations(amplifiers) {
		cpus := make([]*intcode.CPU, amplifiers)
		for i := 0; i < amplifiers; i++ {
			cpus[i] = intcode.NewCPU()
			cpus[i].LoadMemory(mem.(intcode.Memory))
		}

		input := 0
		for i := 0; i < amplifiers; i++ {
			in := make(chan int, 2)
			in <- phase[i]
			in <- input

			cpus[i].Exec(func() int { return <-in }, func(out int) { input = out })
		}

		if input > max {
			max = input
		}
	}
	return max
}

func solve2(mem interface{}) int {
	amplifiers := 5

	max := 0
	for _, _ = range common.Permutations(amplifiers) {
		var (
			cpus = make([]*intcode.CPU, amplifiers)
		)
		for i := 0; i < amplifiers; i++ {
			cpus[i] = intcode.NewCPU()
			cpus[i].LoadMemory(mem.(intcode.Memory))

			/*
				in := func() int {

				}

				out := func(int) {

				}
			*/
		}
	}
	return max
}
