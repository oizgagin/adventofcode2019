package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oizgagin/adventofcode2019/common"
)

func main() {
	solver := common.NewSolver(parse, solve1, solve2)
	fmt.Println(solver.Solve())
}

type ingredient struct {
	amount uint64
	id     string
}

func parse(lines []string) (interface{}, error) {
	tree := make(map[ingredient][]ingredient)

	trim := func(s string) string {
		return strings.TrimSpace(s)
	}

	split := func(s, sep string) []string {
		return strings.Split(s, sep)
	}

	atoi := func(s string) (n int) {
		n, _ = strconv.Atoi(s)
		return
	}

	parseIng := func(s string) ingredient {
		return ingredient{
			amount: uint64(atoi(trim(split(trim(s), " ")[0]))),
			id:     trim(split(trim(s), " ")[1]),
		}
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		var froms []ingredient
		for _, from := range split(split(line, "=>")[0], ",") {
			froms = append(froms, parseIng(from))
		}

		to := parseIng(split(line, "=>")[1])

		tree[to] = froms
	}

	return tree, nil
}

func solve1(v interface{}) interface{} {
	tree := v.(map[ingredient][]ingredient)
	return int(countores(tree, 1))
}

func countores(tree map[ingredient][]ingredient, fuels uint64) uint64 {
	free := make(map[string]uint64)

	var ores uint64

	var visit func(uint64, string)

	visit = func(need uint64, id string) {
		if id == "ORE" {
			ores += need
			return
		}

		if free[id] >= need {
			free[id] -= need
			return
		}

		need -= free[id]
		free[id] = 0

		for from, tos := range tree {
			if from.id == id {
				mul := need / from.amount
				if need%from.amount != 0 {
					mul += 1
					free[id] += from.amount - need%from.amount
				}

				for _, to := range tos {
					visit(mul*to.amount, to.id)
				}
			}
		}
	}

	visit(fuels, "FUEL")

	return ores
}

func solve2(v interface{}) interface{} {
	tree := v.(map[ingredient][]ingredient)

	maxOres := uint64(1000000000000)
	orePerFuel := countores(tree, 1)

	min, max := maxOres/orePerFuel, 2*maxOres/orePerFuel

	for min != max-1 {
		m := (max + min) / 2
		if ores := countores(tree, m); ores >= maxOres {
			max = m
		} else {
			min = m
		}
	}
	return int(min)
}
