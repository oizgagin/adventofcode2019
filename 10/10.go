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
	points := sortPoints(m.points, point{x: 17, y: 22})
	return points[200]
}

func sortPoints(points []point, p point) []point {
	abs := func(p1, p2 point) int {
		return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
	}

	angle := func(p1, p2 point) float64 {
		x, y := (p2.x - p1.x), -(p2.y - p1.y)
		if x == 0 && y == 0 {
			return 0
		}
		a := math.Abs(math.Atan2(float64(y), float64(x))*180/math.Pi - 90)
		if x < 0 && y >= 0 {
			a = 360 - a
		}
		return a
	}

	groups := make(map[float64][]point)
	angles := make([]float64, 0, len(points))

	for _, point := range points {
		a := angle(p, point)
		groups[a] = append(groups[a], point)
		angles = append(angles, a)
	}

	sort.Slice(angles, func(i, j int) bool { return angles[i] < angles[j] })

	set := func(arr []float64) []float64 {
		seen := make(map[float64]bool)

		var r []float64
		for _, elem := range arr {
			if !seen[elem] {
				r = append(r, elem)
				seen[elem] = true
			}
		}

		return r
	}

	angles = set(angles)

	for _, ps := range groups {
		sort.Slice(ps, func(i, j int) bool { return abs(p, ps[i]) < abs(p, ps[j]) })
	}

	var result []point

	result = append(result, groups[0][0])
	groups[0] = groups[0][1:]

	for {
		for _, a := range angles {
			if len(groups[a]) == 0 {
				continue
			}

			result = append(result, groups[a][0])
			groups[a] = groups[a][1:]
			continue
		}

		if len(result) == len(points) {
			break
		}
	}

	return result
}
