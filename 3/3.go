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
	fmt.Printf("task #3, part #%d, solution: %d", *part, solve(*filename, *part))
}

func solve(filename string, part int) int {
	lines, err := readlines(filename)
	if err != nil {
		log.Fatalf("%v: readlines: %v", filename, err)
	}
	if len(lines) != 2 {
		log.Fatalf("%v: want 2 lines, got %v", filename, len(lines))
	}

	p1, p2 := parsePath(lines[0]), parsePath(lines[1])

	if part == 1 {
		return solve1(p1, p2)
	}
	if part == 2 {
		return solve2(p1, p2)
	}

	log.Fatalf("invalid part: %d", part)
	return 0
}

func solve1(p1, p2 []point) int {
	min := 0

	for i := 0; i+1 < len(p1); i++ {
		for j := 0; j+1 < len(p2); j++ {
			if p, ok := cross(p1[i], p1[i+1], p2[j], p2[j+1]); ok && (min == 0 || absPoint(p) < min) {
				min = absPoint(p)
			}
		}
	}

	return min
}

func solve2(p1, p2 []point) int {
	min := 0

	d1, d2 := distances(p1), distances(p2)

	for i := 0; i+1 < len(p1); i++ {
		for j := 0; j+1 < len(p2); j++ {
			if cross, ok := cross(p1[i], p1[i+1], p2[j], p2[j+1]); ok {
				if d := d1[p1[i]] + d2[p2[j]] + absPoint(subPoints(cross, p1[i])) + absPoint(subPoints(cross, p2[j])); min == 0 || d < min {
					min = d
				}
			}
		}
	}

	return min
}

func distances(p []point) map[point]int {
	if len(p) == 0 {
		return nil
	}

	d := make(map[point]int)

	d[p[0]] = absPoint(p[0])
	for i := 1; i < len(p); i++ {
		if _, has := d[p[i]]; !has {
			d[p[i]] = d[p[i-1]] + absPoint(subPoints(p[i], p[i-1]))
		}
	}
	return d
}

func absPoint(p point) int {
	return abs(p.x) + abs(p.y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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

type point struct {
	x, y int
}

func parsePath(s string) []point {
	segments := strings.Split(s, ",")

	var path []point

	p := point{0, 0}
	path = append(path, p)

	for _, seg := range segments {
		p = sumPoints(p, seg2offset(seg))
		path = append(path, p)
	}

	return path
}

func subPoints(p1, p2 point) point {
	return point{p1.x - p2.x, p1.y - p2.y}
}

func sumPoints(p1, p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y}
}

func cross(a1, b1, a2, b2 point) (point, bool) {
	vertical := func(a, b point) bool {
		return a.x == b.x
	}

	horizontal := func(a, b point) bool {
		return a.y == b.y
	}

	if vertical(a1, b1) && vertical(a2, b2) {
		return point{}, false
	}

	if horizontal(a1, b1) && horizontal(a2, b2) {
		return point{}, false
	}

	var ha, hb point
	if horizontal(a1, b1) {
		ha, hb = a1, b1
	} else {
		ha, hb = a2, b2
	}

	var va, vb point
	if vertical(a1, b1) {
		va, vb = a1, b1
	} else {
		va, vb = a2, b2
	}

	if in(ha.x, hb.x, va.x) && in(va.y, vb.y, ha.y) {
		return point{va.x, ha.y}, true
	}

	return point{}, false
}

func in(a, b, c int) bool {
	if a > b {
		a, b = b, a
	}
	return a <= c && c <= b
}

var validDirections = map[byte]bool{
	'L': true,
	'R': true,
	'U': true,
	'D': true,
}

func seg2offset(seg string) point {
	if len(seg) < 2 {
		log.Fatalf("invalid segment: %v", seg)
	}

	direction := seg[0]
	if !validDirections[direction] {
		log.Fatalf("invalid direction: %v", direction)
	}

	offset, err := strconv.Atoi(seg[1:])
	if err != nil {
		log.Fatalf("invalid offset: %v", seg[1:])
	}

	switch direction {
	case 'L':
		return point{-offset, 0}
	case 'R':
		return point{offset, 0}
	case 'U':
		return point{0, offset}
	case 'D':
		return point{0, -offset}
	}

	log.Fatalf("unexpected direction: %v", direction)
	return point{}
}
