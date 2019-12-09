package main

import (
	"fmt"
	"strings"

	"github.com/oizgagin/adventofcode2019/common"
)

func main() {
	solver := common.NewSolver(parsemap, solve1, solve2)
	fmt.Println(solver.Solve())
}

type Map map[string][]string

func parsemap(links []string) (interface{}, error) {
	m := make(Map)
	for _, link := range links {
		if len(link) == 0 {
			continue
		}

		splitted := strings.Split(link, ")")
		from, to := splitted[0], splitted[1]

		m[from] = append(m[from], to)
		m[to] = append(m[to], from)
	}
	return interface{}(m), nil
}

func solve1(input interface{}) int {
	return countOrbits(input.(Map))
}

func solve2(input interface{}) int {
	return findMinDistance(input.(Map), "YOU", "SAN")
}

func findMinDistance(m Map, from, to string) int {
	type dist struct {
		init bool
		dist int
	}

	shorts := map[string]dist{
		from: dist{
			init: true,
			dist: 0,
		},
	}

	visited := make(map[string]bool)

	next := func() string {
		n, min := "", 0
		for node := range m {
			if !visited[node] && shorts[node].init && (n == "" || shorts[node].dist < min) {
				n, min = node, shorts[node].dist
			}
		}
		return n
	}

	for curr := from; curr != ""; curr = next() {
		for _, neighbor := range m[curr] {
			d := shorts[curr].dist + 1
			if !shorts[neighbor].init || d < shorts[neighbor].dist {
				shorts[neighbor] = dist{
					init: true,
					dist: d,
				}
			}
		}
		visited[curr] = true
	}

	return shorts[to].dist - 2
}

const com = "COM"

func countOrbits(m Map) int {
	type s struct {
		name string
		d    int
	}

	stack := []*s{
		&s{name: com, d: 0},
	}

	visited := make(map[string]bool)

	total := 0
	for len(stack) > 0 {
		n := stack[0]
		visited[n.name] = true

		total += n.d

		for _, child := range m[n.name] {
			if !visited[child] {
				stack = append(stack, &s{name: child, d: n.d + 1})
			}
		}

		stack = stack[1:]
	}

	return total
}
