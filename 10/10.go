package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/oizgagin/adventofcode2019/common"
)

func main() {
	solver := common.NewSolver(parse, solve1, solve2)
	fmt.Println(solver.Solve())
}

type point struct {
	x, y int
}

type pointsmap struct {
	points        []point
	width, height int
}

func parse(lines []string) (interface{}, error) {
	m := pointsmap{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		m.width = len(line)
		for j := 0; j < len(line); j++ {
			if line[j] == '#' {
				m.points = append(m.points, point{x: j, y: m.height})
			}
		}
		m.height++
	}
	return m, nil
}

func solve1(v interface{}) interface{} {
	m := v.(pointsmap)

	abs := func(p1, p2 point) int {
		return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
	}

	eq := func(a, b point) bool {
		return a.x == b.x && a.y == b.y
	}

	sign := func(x, y int) bool {
		return (x <= 0 && y <= 0) || (x > 0 && y > 0)
	}

	max := 0

	for _, p := range m.points {
		ps := make([]point, len(m.points)-1)
		copy(ps, m.points)

		blocked := make(map[point]bool)
		visible := 1

		sort.Slice(ps, func(i, j int) bool { return abs(p, ps[i]) < abs(p, ps[j]) })

		for i, neigh := range ps {
			if eq(p, neigh) {
				continue
			}
			if blocked[neigh] {
				continue
			}

			fmt.Println(p, neigh)
			visible++

			dx, dy := neigh.x-p.x, neigh.y-p.y

			for j := i + 1; j < len(ps); j++ {
				ddx, ddy := ps[j].x-p.x, ps[j].y-p.y

				if !sign(dx, ddx) {
					continue
				}
				if !sign(dy, ddy) {
					continue
				}
				if dx == 0 && ddx != 0 || dx != 0 && ddx == 0 {
					continue
				}
				if dy == 0 && ddy != 0 || dy != 0 && ddy == 0 {
					continue
				}
				if (dx == 0 && ddx == 0 || ddx%dx == 0) && (dy == 0 && ddy == 0 || ddy%dy == 0) {
					blocked[ps[j]] = true
				}
			}
		}

		if visible > max {
			max = visible
		}
	}

	return max
}

func solve2(v interface{}) interface{} {
	return 0
}
