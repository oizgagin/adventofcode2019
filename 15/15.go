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

type Movement int

const (
	North Movement = 1
	South Movement = 2
	West  Movement = 3
	East  Movement = 4
)

func (movement Movement) String() string {
	m := map[Movement]string{
		North: "N",
		South: "S",
		West:  "W",
		East:  "E",
	}
	return m[movement]
}

type Status int

const (
	Wall  Status = 0
	Moved Status = 1
	Found Status = 2
)

func (status Status) String() string {
	m := map[Status]string{
		Wall:  "WALL",
		Moved: "MOVED",
		Found: "FOUND",
	}
	return m[status]
}

func solve1(v interface{}) interface{} {
	var currMovement Movement

	input := func() int {
		return int(currMovement)
	}

	var currStatus Status

	output := func(v int) {
		currStatus = Status(v)
	}

	cpu := intcode.NewCPU(v.(intcode.Memory), input, output)

	move := func(m Movement) Status {
		currMovement = m
		if state := cpu.Exec(); state == intcode.CPUHalt {
			panic("unexpected halt received")
		}
		return currStatus
	}

	reversed := map[Movement]Movement{
		South: North,
		North: South,
		East:  West,
		West:  East,
	}

	min := 0

	var walk func(int, Movement)

	walk = func(depth int, to Movement) {
		status := move(to)

		if status == Wall {
			return
		}

		defer move(reversed[to])

		if status == Found {
			if depth < min || min == 0 {
				min = depth
			}
			return
		}

		for _, movement := range []Movement{South, North, West, East} {
			if movement != reversed[to] {
				walk(depth+1, movement)
			}
		}
	}

	for _, movement := range []Movement{South, North, West, East} {
		walk(1, movement)
	}

	return min
}

func solve2(v interface{}) interface{} {
	return 0
}
