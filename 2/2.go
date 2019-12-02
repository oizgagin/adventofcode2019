package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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
	memory, err := readMemory(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PART #%d, SOLUTION: %d\n", *part, solve(memory, *part))
}

func solve(memory map[int]int, part int) int {
	if part == 1 {
		return solve1(memory)
	}
	panic(fmt.Sprintf("unexpected part %d", part))
}

func solve1(memory map[int]int) int {
	memory[1] = 12
	memory[2] = 2
	exec(memory)
	return memory[0]
}

const (
	OpCodeAdd  = 1
	OpCodeMul  = 2
	OpCodeHalt = 99
)

func exec(memory map[int]int) {
	for pos := 0; memory[pos] != OpCodeHalt; pos += 4 {
		switch memory[pos] {
		case OpCodeAdd:
			execBinary(memory, pos, add)
		case OpCodeMul:
			execBinary(memory, pos, mul)
		}
	}
}

func execBinary(memory map[int]int, pos int, f func(int, int) int) {
	p1, p2, p3 := memory[pos+1], memory[pos+2], memory[pos+3]
	memory[p3] = f(memory[p1], memory[p2])
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func readMemory(filename string) (map[int]int, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ioutil.Readfile(%q) = %v", filename, err)
	}
	return parseMemory(string(content))
}

func parseMemory(raw string) (map[int]int, error) {
	memory := make(map[int]int)
	for i, s := range strings.Split(raw, ",") {
		value, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, fmt.Errorf("invalid int at pos #%d: %v", i, s)
		}
		memory[i] = value
	}
	return memory, nil
}

func marshalMem(mem map[int]int) string {
	keys := make([]int, 0, len(mem))
	for k := range mem {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	parts := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		parts[i] = strconv.Itoa(mem[keys[i]])
	}

	return strings.Join(parts, ",")
}
