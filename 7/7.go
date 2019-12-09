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

	phases := common.Permutations(amplifiers)

	maxSignal := 0
	for _, phase := range phases {
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

			outputs := cpus[i].Exec(in)
			input = outputs[0]
		}

		if input > maxSignal {
			maxSignal = input
		}
	}

	return maxSignal
}

func solve2(mem interface{}) int {
	cpu := intcode.NewCPU()
	cpu.LoadMemory(mem.(intcode.Memory))
	return 0
}
