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
}

func main() {
	flag.Parse()
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
		if isValid2(i) {
			valids++
		}
	}
	return valids
}

func isValid(n int) bool {
	dups, decreasing := false, true

	prev := n % 10
	for n /= 10; n > 0; n /= 10 {
		curr := n % 10
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

func isValid2(n int) bool {
	count, dups := 0, false

	prev, count := n%10, 1
	for n /= 10; n > 0; n /= 10 {
		curr := n % 10
		if curr == prev {
			count++
		} else {
			if count%2 == 0 {
				dups = true
			}
			count = 1
		}
		if prev < curr {
			return false
		}
		prev = curr
	}
	return dups
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
