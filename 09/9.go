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

func solve1(v interface{}) int {
	output := 0

	cpu := intcode.NewCPU(v.(intcode.Memory), func() int { return 1 }, func(out int) { output = out })
	for {
		state := cpu.Exec()

		if state == intcode.CPUHalt {
			break
		}
	}

	return output
}

func solve2(v interface{}) int {
	output := 0

	cpu := intcode.NewCPU(v.(intcode.Memory), func() int { return 2 }, func(out int) { output = out })
	for {
		state := cpu.Exec()

		if state == intcode.CPUHalt {
			break
		}
	}

	return output
}
