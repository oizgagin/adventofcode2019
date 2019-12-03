package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
	paths, err := readPaths(*filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PART #%d, SOLUTION: %d\n", *part, solve(paths, *part))
}

func readPaths(filename string) ([]*Path, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ioutil.Readfile(%q) = %v", filename, err)
	}
	var paths []*Path
	for _, line := range strings.Split(string(content), "\n") {
		if len(line) > 0 {
			paths = append(paths, unmarshalPath(line))
		}
	}
	return paths, nil
}

func solve(paths []*Path, part int) int {
	if part == 1 {
		return solve1(paths)
	}
	return 0
}

func solve1(paths []*Path) int {
	p1, p2 := paths[0], paths[1]

	var crosses []*Point
	for _, seg1 := range p1.Segments {
		for _, seg2 := range p2.Segments {
			if cross, ok := seg1.Cross(seg2); ok {
				crosses = append(crosses, cross)
			}
		}
	}

	minAbs := 0
	for _, cross := range crosses {
		if minAbs == 0 || cross.abs() < minAbs {
			minAbs = cross.abs()
		}
	}
	return minAbs
}

type Point struct {
	X, Y int
}

func makePoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func addPoint(a, b *Point) *Point {
	return &Point{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Point) abs() int {
	return abs(p.X) + abs(p.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Segment struct {
	A, B *Point
}

func makeSegment(x1, y1, x2, y2 int) *Segment {
	return &Segment{
		A: makePoint(x1, y1),
		B: makePoint(x2, y2),
	}
}

func (s *Segment) String() string {
	return fmt.Sprintf("%s-%s", s.A, s.B)
}

func (s *Segment) IsVertical() bool {
	return s.A.X == s.B.X
}

func (s *Segment) IsHorizontal() bool {
	return s.A.Y == s.B.Y
}

func (s *Segment) Cross(other *Segment) (*Point, bool) {
	if s.IsHorizontal() && other.IsHorizontal() {
		return nil, false
	}
	if s.IsVertical() && other.IsVertical() {
		return nil, false
	}

	var h, v *Segment
	if s.IsHorizontal() {
		h, v = s, other
	} else {
		h, v = other, s
	}

	if !in(h.A.X, h.B.X, v.A.X) {
		return nil, false
	}
	if !in(v.A.Y, v.B.Y, h.A.Y) {
		return nil, false
	}

	return &Point{X: v.A.X, Y: h.A.Y}, true
}

func in(a, b, c int) bool {
	return (a <= c && c <= b || b <= c && c <= a)
}

type Path struct {
	Segments []*Segment
}

func unmarshalPath(path string) *Path {
	parts := strings.Split(path, ",")

	var segments []*Segment

	p := makePoint(0, 0)
	for _, part := range parts {
		p2 := addPoint(p, pointFromPathPart(part))
		segments = append(segments, &Segment{A: p, B: p2})
		p = p2
	}

	return &Path{Segments: segments}
}

func pointFromPathPart(part string) *Point {
	direction := part[0]
	offset, _ := strconv.Atoi(part[1:])

	switch direction {
	case 'L':
		return &Point{-offset, 0}
	case 'R':
		return &Point{offset, 0}
	case 'U':
		return &Point{0, offset}
	case 'D':
		return &Point{0, -offset}
	}
	panic(fmt.Sprintf("invalid part: %q", part))
}
