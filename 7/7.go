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

	prev := func(i int) int {
		return (i - 1 + amplifiers) % amplifiers
	}

	max := 0
	for _, phase := range common.Permutations(amplifiers) {
		gets := make([]func() int, amplifiers)
		sets := make([]func(int), amplifiers)

		for i := 0; i < amplifiers; i++ {
			inits := []int{phase[prev(i)]}
			if i == amplifiers-1 {
				inits = append(inits, 0)
			}
			gets[i], sets[i] = pipe(inits)
		}

		cpus := make([]*intcode.CPU, amplifiers)
		for i := 0; i < amplifiers; i++ {
			cpus[i] = intcode.NewCPU(mem.(intcode.Memory).Copy(), gets[prev(i)], sets[i])
		}

		for i := 0; i < amplifiers; i++ {
			cpus[i].Exec()
		}

		if signal := gets[amplifiers-1](); signal > max {
			max = signal
		}
	}
	return max
}

func solve2(mem interface{}) int {
	amplifiers := 5

	prev := func(i int) int {
		return (i - 1 + amplifiers) % amplifiers
	}

	max := 0

PHASE_LOOP:
	for _, phase := range common.Permutations(amplifiers) {
		gets := make([]func() int, amplifiers)
		sets := make([]func(int), amplifiers)

		for i := 0; i < amplifiers; i++ {
			inits := []int{phase[prev(i)] + 5}
			if i == amplifiers-1 {
				inits = append(inits, 0)
			}
			gets[i], sets[i] = pipe(inits)
		}

		cpus := make([]*intcode.CPU, amplifiers)
		for i := 0; i < amplifiers; i++ {
			cpus[i] = intcode.NewCPU(mem.(intcode.Memory).Copy(), gets[prev(i)], sets[i])
		}

		for {
			for i := 0; i < amplifiers; i++ {
				if state := cpus[i].Exec(); i == amplifiers-1 && state == intcode.CPUHalt {
					if signal := gets[amplifiers-1](); signal > max {
						max = signal
					}
					continue PHASE_LOOP
				}
			}
		}
	}
	return max
}

func pipe(inits []int) (get func() int, set func(int)) {
	val := 0

	get = func() int {
		if len(inits) > 0 {
			ret := inits[0]
			inits = inits[1:]
			return ret
		}
		return val
	}

	set = func(v int) {
		val = v
	}

	return
}
