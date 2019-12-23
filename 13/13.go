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

type Tile int

const (
	EmptyTile      Tile = 0
	WallTile       Tile = 1
	BlockTile      Tile = 2
	HorizontalTile Tile = 3
	BallTile       Tile = 4
)

func (tile Tile) String() string {
	switch tile {
	case EmptyTile:
		return " "
	case WallTile:
		return "W"
	case BlockTile:
		return "B"
	case HorizontalTile:
		return "="
	case BallTile:
		return "o"
	default:
		return "?"
	}
}

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
			row += grid.Get(x, y).String()
		}
		rows = append(rows, row)
	}
	return strings.Join(rows, "\n")
}

func solve1(v interface{}) interface{} {
	input := func() int {
		panic("unexpected input call")
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

/*
func solve2(v interface{}) interface{} {
	keys := make(chan int, 1)

	go func() {
		err := keyboard.Open()
		if err != nil {
			panic(err)
		}
		defer keyboard.Close()

		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			switch key {
			case keyboard.KeyArrowLeft:
				keys <- -1
			case keyboard.KeyArrowRight:
				keys <- 1
			}
		}
	}()

	calledInput := false

	input := func() int {
		calledInput = true
		select {
		case key := <-keys:
			return key
		default:
			return 0
		}
	}

	grid := NewGrid()

	score := 0

	output := func() func(int) {
		x, y := 0, 0

		setX := func(v int) { x = v }
		setY := func(v int) { y = v }
		setTile := func(v int) {
			if x == -1 && y == 0 {
				score = v
			} else {
				grid.Set(x, y, Tile(v))
			}
		}

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

	clean := func() {
		fmt.Println("\033[2J")
	}

	mem := v.(intcode.Memory)
	mem.Set(0, 2)

	cpu := intcode.NewCPU(mem, input, output)
	for {
		state := cpu.Exec()
		if state == intcode.CPUHalt {
			break
		}
		if !calledInput {
			continue
		}
		clean()
		fmt.Println("SCORE:", score)
		fmt.Println(grid)
		time.Sleep(time.Second / 10)
	}

	return score
}
*/

func solve2(v interface{}) interface{} {
	padX, ballX := 0, 0

	input := func() int {
		if padX < ballX {
			return 1
		}
		if padX > ballX {
			return -1
		}
		return 0
	}

	//grid := NewGrid()

	score := 0

	output := func() func(int) {
		x, y := 0, 0

		setX := func(v int) { x = v }
		setY := func(v int) { y = v }
		setTile := func(v int) {
			if x == -1 && y == 0 {
				score = v
				return
			}
			if Tile(v) == BallTile {
				ballX = x
			}
			if Tile(v) == HorizontalTile {
				padX = x
			}
		}

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

	mem := v.(intcode.Memory)
	mem.Set(0, 2)

	cpu := intcode.NewCPU(mem, input, output)
	for {
		state := cpu.Exec()
		if state == intcode.CPUHalt {
			break
		}
	}

	return score
}
