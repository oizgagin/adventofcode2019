package main

import (
	"fmt"
	"strings"

	"github.com/oizgagin/adventofcode2019/common"
	"github.com/oizgagin/adventofcode2019/intcode"
)

func main() {
	solver := common.NewSolver(parse, solve1, solve2)
	fmt.Println(solver.Solve())
}

func parse(lines []string) (interface{}, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines in input")
	}
	return interface{}(intcode.NewMemoryFromString(lines[0])), nil
}

type Movement int

const (
	North Movement = 1
	South Movement = 2
	West  Movement = 3
	East  Movement = 4
)

func (movement Movement) String() string {
	m := map[Movement]string{
		North: "N",
		South: "S",
		West:  "W",
		East:  "E",
	}
	return m[movement]
}

type Status int

const (
	Unknown Status = -1
	Wall    Status = 0
	Moved   Status = 1
	Found   Status = 2
)

func (status Status) String() string {
	m := map[Status]string{
		Wall:  "WALL",
		Moved: "MOVED",
		Found: "FOUND",
	}
	return m[status]
}

func solve1(v interface{}) interface{} {
	var currMovement Movement

	input := func() int {
		return int(currMovement)
	}

	var currStatus Status

	output := func(v int) {
		currStatus = Status(v)
	}

	cpu := intcode.NewCPU(v.(intcode.Memory), input, output)

	move := func(m Movement) Status {
		currMovement = m
		if state := cpu.Exec(); state == intcode.CPUHalt {
			panic("unexpected halt received")
		}
		return currStatus
	}

	reversed := map[Movement]Movement{
		South: North,
		North: South,
		East:  West,
		West:  East,
	}

	min := 0

	var walk func(int, Movement)

	walk = func(depth int, to Movement) {
		status := move(to)

		if status == Wall {
			return
		}

		defer move(reversed[to])

		if status == Found {
			if depth < min || min == 0 {
				min = depth
			}
			return
		}

		for _, movement := range []Movement{South, North, West, East} {
			if movement != reversed[to] {
				walk(depth+1, movement)
			}
		}
	}

	for _, movement := range []Movement{South, North, West, East} {
		walk(1, movement)
	}

	return min
}

type Grid map[int]map[int]Status

func NewGrid() Grid {
	return make(map[int]map[int]Status)
}

func (grid Grid) ensure(x, y int) {
	if _, has := grid[x]; !has {
		grid[x] = make(map[int]Status)
	}
}

func (grid Grid) Get(x, y int) Status {
	grid.ensure(x, y)
	if status, has := grid[x][y]; has {
		return status
	}
	return Unknown
}

func (grid Grid) Set(x, y int, status Status) {
	grid.ensure(x, y)
	grid[x][y] = status
}

func (grid Grid) Has(x, y int) bool {
	if _, has := grid[x]; has {
		return true
	}
	_, has := grid[x][y]
	return has
}

func (grid Grid) String() string {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for x, ys := range grid {
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}

		for y := range ys {
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	var rows []string
	for y := minY; y <= maxY; y++ {
		row := ""
		for x := minX; x <= maxX; x++ {
			switch grid.Get(x, y) {
			case Unknown:
				row += " "
			case Wall:
				row += "#"
			case Moved:
				row += "."
			case Found:
				row += "O"
			default:
				row += "O"
			}
		}
		rows = append(rows, row)
	}
	return strings.Join(rows, "\n")
}

func solve2(v interface{}) interface{} {
	var currMovement Movement

	input := func() int {
		return int(currMovement)
	}

	var currStatus Status

	output := func(v int) {
		currStatus = Status(v)
	}

	cpu := intcode.NewCPU(v.(intcode.Memory), input, output)

	currX, currY := 0, 0

	grid := NewGrid()
	grid.Set(currX, currY, Moved)

	move := func(m Movement) Status {
		currMovement = m
		if state := cpu.Exec(); state == intcode.CPUHalt {
			panic("unexpected halt received")
		}
		return currStatus
	}

	reversed := map[Movement]Movement{
		South: North,
		North: South,
		East:  West,
		West:  East,
	}

	var walk func(Movement)

	walk = func(to Movement) {
		status := move(to)

		if status == Wall {
			switch to {
			case North:
				grid.Set(currX, currY-1, Wall)
			case South:
				grid.Set(currX, currY+1, Wall)
			case West:
				grid.Set(currX-1, currY, Wall)
			case East:
				grid.Set(currX+1, currY, Wall)
			}
			return
		}

		switch to {
		case North:
			currY -= 1
		case South:
			currY += 1
		case West:
			currX -= 1
		case East:
			currX += 1
		}

		fmt.Println("MOVED TO", to, "X", currX, "Y", currY, "STATUS", currStatus)
		grid.Set(currX, currY, currStatus)

		for _, movement := range []Movement{South, North, West, East} {
			if movement != reversed[to] {
				walk(movement)
			}
		}

		move(reversed[to])
		switch reversed[to] {
		case North:
			currY -= 1
		case South:
			currY += 1
		case West:
			currX -= 1
		case East:
			currX += 1
		}
	}

	for _, movement := range []Movement{South, North, West, East} {
		walk(movement)
	}

	fmt.Println(grid)

	minX, maxX, minY, maxY := 0, 0, 0, 0
	for x, ys := range grid {
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}

		for y := range ys {
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	hasOxygen := func(x, y, minutes int) bool {
		isOxy := func(x1, y1 int) bool {
			s := grid.Get(x1, y1)
			return int(Found) <= int(s) && int(s) < int(Found)+minutes+1
		}
		return isOxy(x-1, y) || isOxy(x+1, y) || isOxy(x, y-1) || isOxy(x, y+1)
	}

	minutes := 0
	for {
		fmt.Println(grid)
		fmt.Println("\n\n")

		wasSet := false
		for y := minY; y <= maxY; y++ {
			for x := minX; x <= maxX; x++ {
				if grid.Get(x, y) == Moved && hasOxygen(x, y, minutes) {
					wasSet = true
					grid.Set(x, y, Status(int(Found)+minutes+1))
				}
			}
		}

		if !wasSet {
			break
		}

		minutes++
	}

	return minutes
}
