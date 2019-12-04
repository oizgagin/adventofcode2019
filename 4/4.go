package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	fmt.Printf("task #4, part #%d, solution: %d", *part, solve(*filename, *part))
}

func solve(filename string, part int) int {
	if part == 1 {
		return solve1(filename)
	}
	if part == 2 {
		return solve2(filename)
	}
	log.Fatalf("invalid part: %d", part)
	return 0
}

func solve1(filename string) int {
	line, err := readline(filename)
	if err != nil {
		log.Fatalf("readline(%q) = %v", filename, err)
	}

	ran := strings.Split(line, "-")
	if len(ran) != 2 {
		log.Fatalf("invalid range: %v", line)
	}

	start, err := strconv.Atoi(ran[0])
	if err != nil {
		log.Fatalf("invalid start: %v", err)
	}

	end, err := strconv.Atoi(ran[1])
	if err != nil {
		log.Fatalf("invalid end: %v", err)
	}

	valids := 0
	for i := start; i <= end; i++ {
		if isValid(i) {
			valids++
		}
	}
	return valids
}

func solve2(filename string) int {
	line, err := readline(filename)
	if err != nil {
		log.Fatalf("readline(%q) = %v", filename, err)
	}

	ran := strings.Split(line, "-")
	if len(ran) != 2 {
		log.Fatalf("invalid range: %v", line)
	}

	start, err := strconv.Atoi(ran[0])
	if err != nil {
		log.Fatalf("invalid start: %v", err)
	}

	end, err := strconv.Atoi(ran[1])
	if err != nil {
		log.Fatalf("invalid end: %v", err)
	}

	valids := 0
	for i := start; i <= end; i++ {
		if isValid(i) {
			valids++
		}
	}
	return valids
}

func solve2(filename string) int {
	return 0
}

func isValid(i int) bool {
	dups, decreasing := false, true

	prev := i % 10
	for i /= 10; i > 0; i /= 10 {
		curr := i % 10
		if curr == prev {
			dups = true
		}
		if prev < curr {
			decreasing = false
		}
		prev = curr
	}
	return decreasing && dups
}

func readline(filename string) (string, error) {
	lines, err := readlines(filename)
	if err != nil {
		return "", err
	}
	if len(lines) == 0 {
		return "", fmt.Errorf("no lines")
	}
	return lines[0], nil
}

func readlines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	res := lines[:0]
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) != 0 {
			res = append(res, line)
		}
	}

	return res, nil
}
