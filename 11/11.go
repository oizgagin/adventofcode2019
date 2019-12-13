package main

import (
	"fmt"

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

type Color int

const (
	Black Color = 0
	White Color = 1
)

type Direction int

const (
	Left  Direction = 0
	Right Direction = 1

	Up   Direction = 2
	Down Direction = 3
)

type Grid map[int]map[int]Color

func NewGrid() Grid {
	return make(map[int]map[int]Color)
}

func (grid Grid) ensure(x, y int) {
	if _, has := grid[x]; !has {
		grid[x] = make(map[int]Color)
	}
}

func (grid Grid) Get(x, y int) Color {
	grid.ensure(x, y)
	return grid[x][y]
}

func (grid Grid) Set(x, y int, color Color) {
	grid.ensure(x, y)
	grid[x][y] = color
}

/*
func (grid Grid) String() string {
	minX, maxX, minY, maxY := 0, 0, 0, 0

	for x := range grid {
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		for y := range grid[x] {
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	var rows []string
	for i := maxY; i >= minY; i-- {
		row := ""
		for j := minX; j <= maxX; j++ {
			if grid.Get(i, j) == Black {
				row += " ."
			} else {
				row += " #"
			}
		}
		rows = append(rows, row)
	}

	return strings.Join(rows, "\n")
}
*/

func solve1(v interface{}) interface{} {
	var (
		grid          = NewGrid()
		currX, currY  = 0, 0
		currDirection = Up
	)

	adjust := func(x, y int, curr, next Direction) (int, int, Direction) {
		switch {
		case curr == Left && next == Left:
			x, y, curr = x, y+1, Down
			return x, y, curr

		case curr == Left && next == Right:
			x, y, curr = x, y-1, Up
			return x, y, curr

		case curr == Right && next == Left:
			x, y, curr = x, y-1, Up
			return x, y, curr

		case curr == Right && next == Right:
			x, y, curr = x, y+1, Down
			return x, y, curr

		case curr == Up && next == Left:
			x, y, curr = x-1, y, Left
			return x, y, curr

		case curr == Up && next == Right:
			x, y, curr = x+1, y, Right
			return x, y, curr

		case curr == Down && next == Left:
			x, y, curr = x+1, y, Right
			return x, y, curr

		case curr == Down && next == Right:
			x, y, curr = x-1, y, Left
			return x, y, curr
		}

		panic(fmt.Sprintf("INVALID STATE: x = %v, y = %v, curr = %v, next = %v", x, y, curr, next))
	}

	input := func() int {
		return int(grid.Get(currX, currY))
	}

	visited := make(map[[2]int]bool)

	output := func() func(int) {
		setColor := func(v int) {
			visited[[2]int{currX, currY}] = true
			grid.Set(currX, currY, Color(v))
		}

		setDirection := func(v int) {
			currX, currY, currDirection = adjust(currX, currY, currDirection, Direction(v))
		}

		calls := 0

		return func(v int) {
			if calls%2 == 0 {
				setColor(v)
			} else {
				setDirection(v)
			}
			calls++
		}
	}()

	cpu := intcode.NewCPU(v.(intcode.Memory), input, output)
	for {
		state := cpu.Exec()
		if state == intcode.CPUHalt {
			break
		}
	}

	return len(visited)
}

func solve2(v interface{}) interface{} {
	var (
		grid          = NewGrid()
		currX, currY  = 0, 0
		currDirection = Up
	)

	grid.Set(currX, currY, White)

	adjust := func(x, y int, curr, next Direction) (int, int, Direction) {
		switch {
		case curr == Left && next == Left:
			x, y, curr = x, y+1, Down
			return x, y, curr

		case curr == Left && next == Right:
			x, y, curr = x, y-1, Up
			return x, y, curr

		case curr == Right && next == Left:
			x, y, curr = x, y-1, Up
			return x, y, curr

		case curr == Right && next == Right:
			x, y, curr = x, y+1, Down
			return x, y, curr

		case curr == Up && next == Left:
			x, y, curr = x-1, y, Left
			return x, y, curr

		case curr == Up && next == Right:
			x, y, curr = x+1, y, Right
			return x, y, curr

		case curr == Down && next == Left:
			x, y, curr = x+1, y, Right
			return x, y, curr

		case curr == Down && next == Right:
			x, y, curr = x-1, y, Left
			return x, y, curr
		}

		panic(fmt.Sprintf("INVALID STATE: x = %v, y = %v, curr = %v, next = %v", x, y, curr, next))
	}

	input := func() int {
		return int(grid.Get(currX, currY))
	}

	output := func() func(int) {
		setColor := func(v int) {
			grid.Set(currX, currY, Color(v))
		}

		setDirection := func(v int) {
			currX, currY, currDirection = adjust(currX, currY, currDirection, Direction(v))
		}

		calls := 0

		return func(v int) {
			if calls%2 == 0 {
				setColor(v)
			} else {
				setDirection(v)
			}
			calls++
		}
	}()

	cpu := intcode.NewCPU(v.(intcode.Memory), input, output)
	for {
		state := cpu.Exec()
		if state == intcode.CPUHalt {
			break
		}
	}

	for y := 0; y <= 10; y++ {
		for x := 0; x <= 42; x++ {
			if grid.Get(x, y) == Black {
				fmt.Printf(" .")
			} else {
				fmt.Printf(" #")
			}
		}
		fmt.Printf("\n")
	}

	return 0
}
