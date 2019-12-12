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

func log(args ...interface{}) {
	//fmt.Println(args...)
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

	absi := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	max := 0

	for _, p := range m.points {
		log("POINT", p)
		ps := make([]point, len(m.points))
		copy(ps, m.points)

		blocked := make(map[point]bool)
		visible := 0

		sort.Slice(ps, func(i, j int) bool { return abs(p, ps[i]) < abs(p, ps[j]) })

		for _, neigh := range ps {
			if eq(p, neigh) {
				continue
			}
			if blocked[neigh] {
				continue
			}

			log("\t", neigh, "IS VISIBLE")

			visible++

			dx, dy := neigh.x-p.x, neigh.y-p.y

			for j := 0; j < len(ps); j++ {
				if eq(neigh, ps[j]) {
					continue
				}

				ddx, ddy := ps[j].x-p.x, ps[j].y-p.y
				log("\t\t", ps[j], "CHECK", neigh, "dx", dx, "dy", dy, "ddx", ddx, "ddy", ddy)

				if !sign(dx, ddx) {
					log("\t\t\tNOT SIGN", "dx", dx, "ddx", ddx)
					continue
				}
				if !sign(dy, ddy) {
					log("\t\t\tNOT SIGN", "dy", dx, "ddy", ddx)
					continue
				}
				if dx == 0 && ddx != 0 || dx != 0 && ddx == 0 {
					log("\t\t\tNOT ZERO", "dx", dx, "ddx", ddx)
					continue
				}
				if dy == 0 && ddy != 0 || dy != 0 && ddy == 0 {
					log("\t\t\tNOT ZERO", "dy", dy, "ddy", ddy)
					continue
				}
				if dx == 0 && ddx == 0 {
					if absi(dy) > absi(ddy) {
						log("\t\t\t", "dy > ddy", "dy", dy, "ddy", ddy)
						continue
					}
					log("\t\t\t", ps[j], "IS BLOCKED BY DX ZERO", neigh, "dx", dx, "dy", dy, "ddx", ddx, "ddy")
					blocked[ps[j]] = true
					continue
				}
				if dy == 0 && ddy == 0 {
					if absi(dx) > absi(ddx) {
						log("\t\t\t", "dx > ddx", "dx", dx, "ddx", ddx)
						continue
					}
					log("\t\t\t", ps[j], "IS BLOCKED BY DY ZERO", neigh, "dx", dx, "dy", dy, "ddx", ddx, "ddy", ddy)
					blocked[ps[j]] = true
					continue
				}
				if dx != 0 && dy != 0 && dy*ddx != dx*ddy {
					continue
				}
				log("\t\t\t", ps[j], "IS BLOCKED BY", neigh, "dx", dx, "dy", dy, "ddx", ddx, "ddy", ddy, "ddx/dx", ddx/dx, "ddy/dy", ddy/dy)
				blocked[ps[j]] = true
			}
		}

		log("VISIBLE", visible)

		if visible > max {
			max = visible
		}
	}

	return max
}

func solve2(v interface{}) interface{} {
	return 0
}
