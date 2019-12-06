package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	filename = flag.String("filename", "input", "input file")
	part     = flag.Int("part", 1, "part no")
)

func main() {
	flag.Parse()
	fmt.Printf("task #6, part #%d, solution: %d", *part, solve(*filename, *part))
}

func solve(filename string, part int) int {
	m, err := readlines(filename)
	if err != nil {
		log.Fatalf("readlines(%q) = %v", filename, err)
		return 0
	}

	if part == 1 {
		return solve1(parsemap(m))
	}
	return solve2(parsemap(m))
}

func solve1(m map[string][]string) int {
	return countOrbits(m)
}

func solve2(m map[string][]string) int {
	return findMinDistance(m, "YOU", "SAN")
}

func findMinDistance(m map[string][]string, from, to string) int {
	type dist struct {
		init bool
		dist int
	}

	shorts := map[string]dist{
		from: dist{
			init: true,
			dist: 0,
		},
	}

	visited := make(map[string]bool)

	next := func() string {
		n, min := "", 0
		for node := range m {
			if !visited[node] && shorts[node].init && (n == "" || shorts[node].dist < min) {
				n, min = node, shorts[node].dist
			}
		}
		return n
	}

	for curr := from; curr != ""; curr = next() {
		for _, neighbor := range m[curr] {
			d := shorts[curr].dist + 1
			if !shorts[neighbor].init || d < shorts[neighbor].dist {
				shorts[neighbor] = dist{
					init: true,
					dist: d,
				}
			}
		}
		visited[curr] = true
	}

	return shorts[to].dist - 2
}

func parsemap(links []string) map[string][]string {
	m := make(map[string][]string)
	for _, link := range links {
		splitted := strings.Split(link, ")")
		from, to := splitted[0], splitted[1]

		m[from] = append(m[from], to)
		m[to] = append(m[to], from)
	}
	return m
}

const com = "COM"

func countOrbits(m map[string][]string) int {
	type s struct {
		name string
		d    int
	}

	stack := []*s{
		&s{name: com, d: 0},
	}

	visited := make(map[string]bool)

	total := 0
	for len(stack) > 0 {
		n := stack[0]
		visited[n.name] = true

		total += n.d

		for _, child := range m[n.name] {
			if !visited[child] {
				stack = append(stack, &s{name: child, d: n.d + 1})
			}
		}

		stack = stack[1:]
	}

	return total
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
