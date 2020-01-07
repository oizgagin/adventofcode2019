package main

import (
	"fmt"
	"strings"

	"github.com/oizgagin/adventofcode2019/common"
)

func main() {
	solver := common.NewSolver(parse, solve1, solve2)
	fmt.Println(solver.Solve())
}

func parse(lines []string) (interface{}, error) {
	line := strings.TrimSpace(lines[0])
	res := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		res[i] = int(line[i] - '0')
	}
	return res, nil
}

func solve1(v interface{}) interface{} {
	input := v.([]int)
	for i := 0; i < 100; i++ {
		input = phase(input)
	}
	return input[:8]
}

func solve2(v interface{}) interface{} {
	vs := v.([]int)

	input := make([]int, len(vs)*10000)
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(vs); j++ {
			input[i*len(vs)+j] = vs[j]
		}
	}
	for i := 0; i < 100; i++ {
		fmt.Println("PHASE", i)
		input = phase(input)
	}

	offset := arr2num(vs[:8])
	return input[offset : offset+8]
}

func arr2num(a []int) (res int) {
	mul := 1
	for i := len(a) - 1; i >= 0; i-- {
		res += a[i] * mul
		mul *= 10
	}
	return
}

func phase(vs []int) []int {
	out := make([]int, len(vs))
	for i := 0; i < len(vs); i++ {
		out[i] = calc(i, vs)
	}
	return out
}

func calc(pos int, vs []int) int {
	sum := 0
	for i := 0; i < len(vs); i++ {
		sum += next(pos, i, vs[i])
	}
	return ones(sum)
}

func next(pos, i, elem int) int {
	return elem * pattern(pos, i)
}

var pat = []int{0, 1, 0, -1}

func pattern(pos, i int) int {
	return pat[((i+1)/(pos+1))%4]
}

func ones(x int) int {
	return abs(x) % 10
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
