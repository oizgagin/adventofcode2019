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

type Tile int

const (
	EmptyTile      Tile = 0
	WallTile       Tile = 1
	BlockTile      Tile = 2
	HorizontalTile Tile = 3
	BallTile            = 4
)

type Grid map[int]map[int]Tile

func NewGrid() Grid {
	return make(map[int]map[int]Tile)
}

func (grid Grid) ensure(x, y int) {
	if _, has := grid[x]; !has {
		grid[x] = make(map[int]Tile)
	}
}

func (grid Grid) Get(x, y int) Tile {
	grid.ensure(x, y)
	return grid[x][y]
}

func (grid Grid) Set(x, y int, tile Tile) {
	grid.ensure(x, y)
	grid[x][y] = tile
}

func solve1(v interface{}) interface{} {
	input := func() int {
		panic("unexpected input call")
		return 0
	}

	grid := NewGrid()

	output := func() func(int) {
		x, y := 0, 0

		setX := func(v int) { x = v }
		setY := func(v int) { y = v }
		setTile := func(v int) { grid.Set(x, y, Tile(v)) }

		calls := 0

		return func(v int) {
			switch calls % 3 {
			case 0:
				setX(v)
			case 1:
				setY(v)
			case 2:
				setTile(v)
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

	blocks := 0
	for x := range grid {
		for y := range grid[x] {
			if grid.Get(x, y) == BlockTile {
				blocks++
			}
		}
	}

	return blocks
}

func solve2(v interface{}) interface{} {
	return 0
}
