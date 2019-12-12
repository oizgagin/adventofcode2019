package common

import (
	"flag"
	"log"
)

type Solver struct {
	parse  func([]string) (interface{}, error)
	solve1 func(interface{}) interface{}
	solve2 func(interface{}) interface{}
}

func NewSolver(
	parse func([]string) (interface{}, error),
	solve1 func(interface{}) interface{},
	solve2 func(interface{}) interface{},
) *Solver {
	return &Solver{
		parse:  parse,
		solve1: solve1,
		solve2: solve2,
	}
}

func (solver *Solver) Solve() interface{} {
	filename := flag.String("filename", "input", "input file")
	part := flag.Int("part", 1, "part no")

	flag.Parse()

	lines, err := ReadLines(*filename)
	if err != nil {
		log.Fatalf("ReadLines(%q) = %v", *filename, err)
	}

	v, err := solver.parse(lines)
	if err != nil {
		log.Fatalf("parse(%v) = %v", lines, err)
	}

	if *part == 1 {
		return solver.solve1(v)
	}

	if *part == 2 {
		return solver.solve2(v)
	}

	log.Fatalf("invalid part: %d", *part)
	return 0
}
