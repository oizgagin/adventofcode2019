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
)

func init() {
	flag.Parse()
}

func main() {
	input, err := readInts(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SOLUTION: %d\n", solve(input))
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

func solve(input []int) int {
	result := 0
	for i := 0; i < len(input); i++ {
		result += calcFuel(input[i])
	}
	return result
}

func calcFuel(module int) int {
	return (module / 3) - 2
}
