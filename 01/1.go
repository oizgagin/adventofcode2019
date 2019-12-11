package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	filename = flag.String("filename", "input", "input file")
	part     = flag.Int("part", 1, "part no")
)

func init() {
	flag.Parse()
}

func main() {
	input, err := readInts(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PART #%d, SOLUTION: %d\n", *part, solve(input, *part))
}

func solve(input []int, part int) int {
	if part == 1 {
		return solve1(input)
	}
	if part == 2 {
		return solve2(input)
	}
	panic(fmt.Sprintf("unknown part: %d", part))
}

func solve1(input []int) int {
	result := 0
	for i := 0; i < len(input); i++ {
		result += calcFuel(input[i])
	}
	return result
}

func solve2(input []int) int {
	result := 0
	for i := 0; i < len(input); i++ {
		result += calcFuelRecursively(input[i])
	}
	return result
}

func calcFuelRecursively(module int) int {
	result := 0
	for mass := calcFuel(module); mass > 0; mass = calcFuel(mass) {
		result += mass
	}
	return result
}

func calcFuel(module int) int {
	return (module / 3) - 2
}

func readInts(filename string) ([]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%q) = %v", filename, err)
	}

	r := bufio.NewReader(f)

	var ints []int
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, err
		}

		ints = append(ints, i)
	}
	return ints, nil
}
