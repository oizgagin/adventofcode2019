package main

import (
	"fmt"
	"math"
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

	absi := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	max := 0
	maxp := point{}

	for _, p := range m.points {
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

			visible++

			dx, dy := neigh.x-p.x, neigh.y-p.y

			for j := 0; j < len(ps); j++ {
				if eq(neigh, ps[j]) {
					continue
				}

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
				if dx == 0 && ddx == 0 {
					if absi(dy) > absi(ddy) {
						continue
					}
					blocked[ps[j]] = true
					continue
				}
				if dy == 0 && ddy == 0 {
					if absi(dx) > absi(ddx) {
						continue
					}
					blocked[ps[j]] = true
					continue
				}
				if dx != 0 && dy != 0 && dy*ddx != dx*ddy {
					continue
				}
				blocked[ps[j]] = true
			}
		}

		if visible > max {
			max = visible
			maxp = p
		}
	}

	fmt.Println(maxp)

	return max
}

func solve2(v interface{}) interface{} {
	m := v.(pointsmap)

	sortPoints(m.points, point{x: 17, y: 22})

	return 0
}

func sortPoints(points []point, p point) {
	abs := func(p1, p2 point) int {
		return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
	}

	angle := func(p1, p2 point) float64 {
		x, y := (p2.x - p1.x), (p2.y - p1.y)
		a := math.Atan2(float64(y), float64(x))*180/math.Pi - 90
		if x < 0 && a > 0 {
			a -= 360
		}
		return math.Abs(a)
	}

	sort.Slice(points, func(i, j int) bool {
		p1, p2 := points[i], points[j]

		a1, a2 := angle(p, p1), angle(p, p2)
		fmt.Println(p, p1, a1, p2, a2)
		if a1 != a2 {
			return a1 < a2
		}
		return abs(p, p1) < abs(p, p2)
	})
}
