package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/oizgagin/adventofcode2019/intcode"
)

var (
	filename = flag.String("filename", "input", "input file")
	part     = flag.Int("part", 1, "part no")
)

func init() {
}

func main() {
	flag.Parse()
	fmt.Printf("task #5, part #%d, solution: %d", *part, solve(*filename, *part))
}

func solve(filename string, part int) int {
	mem, err := readline(filename)
	if err != nil {
		log.Fatalf("readline(%q): %v", filename, err)
		return 0
	}

	if part == 1 {
		return solve1(mem)
	}
	return solve2(mem)
}

func solve1(mem string) (output int) {
	cpu := intcode.NewCPU()
	cpu.LoadMemory(intcode.NewMemoryFromString(mem))

	inputs := make(chan int)
	go func() {
		inputs <- 1
		close(inputs)
	}()

	outputs := make(chan int)
	go func() {
		for {
			output = <-outputs
		}
	}()

	cpu.Exec(inputs, outputs)

	return
}

func solve2(mem string) (output int) {
	cpu := intcode.NewCPU()
	cpu.LoadMemory(intcode.NewMemoryFromString(mem))

	inputs := make(chan int)
	go func() {
		inputs <- 5
		close(inputs)
	}()

	collected, outputs := make(chan struct{}), make(chan int)
	go func() {
		defer close(collected)
		for elem := range outputs {
			output = elem
		}
	}()

	cpu.Exec(inputs, outputs)

	<-collected

	return
}

func readline(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll: %v", err)
	}
	return strings.TrimSpace(string(content)), nil
}
