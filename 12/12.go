package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/oizgagin/adventofcode2019/common"
)

func main() {
	solver := common.NewSolver(parse, solve1, solve2)
	fmt.Println(solver.Solve())
}

func parse(lines []string) (interface{}, error) {
	return parseCoords(strings.Join(lines, "\n")), nil
}

func solve1(v interface{}) interface{} {
	system := NewSystem(v.([][3]int))
	for i := 0; i < 1000; i++ {
		system.Tick()
	}
	return system.Energy()
}

func parseCoords(s string) (coords [][3]int) {
	re := regexp.MustCompile(`\<x\=([^,]+).*?y\=([^,]+).*?z=([^,]+)\>`)

	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		raws := re.FindAllStringSubmatch(line, -1)[0][1:]

		var coord [3]int
		for i, raw := range raws {
			coord[i], _ = strconv.Atoi(raw)
		}
		coords = append(coords, coord)
	}

	return
}

type System struct {
	coords [][3]int
	speeds [][3]int
}

func NewSystem(coords [][3]int) *System {
	return &System{
		coords: coords,
		speeds: make([][3]int, len(coords)),
	}
}

func (system *System) applyGravity() {
	for i := 0; i < len(system.coords); i++ {
		for j := i + 1; j < len(system.coords); j++ {
			for k := 0; k < 3; k++ {
				if system.coords[i][k] < system.coords[j][k] {
					system.speeds[i][k] += 1
					system.speeds[j][k] -= 1
				}
				if system.coords[i][k] > system.coords[j][k] {
					system.speeds[i][k] -= 1
					system.speeds[j][k] += 1
				}
			}
		}
	}
}

func (system *System) applyVelocity() {
	for i := 0; i < len(system.coords); i++ {
		for k := 0; k < 3; k++ {
			system.coords[i][k] += system.speeds[i][k]
		}
	}
}

func (system *System) Energy() (energy int) {
	for i := 0; i < len(system.coords); i++ {
		kin, pot := 0, 0
		for k := 0; k < 3; k++ {
			kin += abs(system.speeds[i][k])
			pot += abs(system.coords[i][k])
		}
		energy += kin * pot
	}
	return
}

func (system *System) Tick() {
	system.applyGravity()
	system.applyVelocity()
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func solve2(v interface{}) interface{} {
	coords := v.([][3]int)

	var (
		xs = make([]int, len(coords))
		ys = make([]int, len(coords))
		zs = make([]int, len(coords))
	)

	for i := 0; i < len(coords); i++ {
		xs[i], ys[i], zs[i] = coords[i][0], coords[i][1], coords[i][2]
	}

	xp, yp, zp := period(xs), period(ys), period(zs)
	return lcm(xp, yp, zp)
}

func period(coords []int) (n uint64) {
	orig := make([]int, len(coords))
	copy(orig, coords)

	sign := func(i int) int {
		if i < 0 {
			return -1
		}
		if i == 0 {
			return 0
		}
		return 1
	}

	eqs := func(a, b []int) bool {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	speeds := make([]int, len(coords))
	zeroes := make([]int, len(coords))

	n = 1
	for {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				speeds[i] += sign(coords[j] - coords[i])
				speeds[j] += sign(coords[i] - coords[j])
			}
		}
		for i := 0; i < len(coords); i++ {
			coords[i] += speeds[i]
		}
		n++
		if eqs(coords, orig) && eqs(speeds, zeroes) {
			return
		}
	}
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b uint64, ints ...uint64) (r uint64) {
	r = a * b / gcd(a, b)
	for i := 0; i < len(ints); i++ {
		r = lcm(r, ints[i])
	}
	return
}
