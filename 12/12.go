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

func solve2(v interface{}) interface{} {
	return 0
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
