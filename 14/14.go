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
			amount: atoi(trim(split(trim(s), " ")[0])),
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
	counts := count(tree, "FUEL")

	fmt.Println(counts)

	var leafs []string
	for to, froms := range tree {
		for _, from := range froms {
			if from.id == "ORE" {
				leafs = append(leafs, to.id)
				break
			}
		}
	}

	ores := 0
	for _, leaf := range leafs {
		c := counts[leaf]
		for to, froms := range tree {
			if to.id == leaf {
				reactions := c / to.amount
				if c%to.amount != 0 {
					reactions += 1
				}
				for _, from := range froms {
					if from.id == "ORE" {
						ores += reactions * from.amount
					}
				}
			}
		}
	}

	return ores
}

func solve2(v interface{}) interface{} {
	return 0
}

type ingredient struct {
	amount int
	id     string
}

func count(tree map[ingredient][]ingredient, node string) map[string]int {
	c := make(map[string]int)
	left := make(map[string]int)

	var visit func(int, string)

	visit = func(need int, id string) {
		if id == "ORE" {
			return
		}

		c[id] += need
		for from, tos := range tree {
			if from.id == id {
				reactions := need / from.amount
				if need%from.amount != 0 {
					left[id] += from.amount - need%from.amount
					reactions += 1
				}
				for _, to := range tos {
					visit(reactions*to.amount, to.id)
				}
			}
		}
	}

	visit(1, node)

	muls := make(map[string]int)
	for from := range tree {
		muls[from.id] = from.amount
	}
	muls["ORE"] = 1
	fmt.Println(muls)

	var del func(string, int)

	del = func(id string, reactions int) {
		c[id] -= reactions

		for from, tos := range tree {
			if from.id == id {
				left[id] -= reactions * from.amount
				for _, to := range tos {
					if to.id == "ORE" {
						continue
					}
					need := (reactions * to.amount) / muls[to.id]
					if (reactions*to.amount)%muls[to.id] != 0 {
						need += 1
					}
					del(to.id, need)
				}
			}
		}
	}

	fmt.Println("LEFT", left)

	for id, leftcount := range left {
		for from, _ := range tree {
			if from.id == id {
				if leftcount >= from.amount {
					del(id, leftcount/from.amount)
				}
			}
		}
	}

	return c
}
