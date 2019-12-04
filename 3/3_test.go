package main

import (
	"reflect"
	"testing"
)

func TestUnmarshalPath(t *testing.T) {

	testCases := []struct {
		s    string
		want []point
	}{
		{
			s: "U1,R1,D1,L1",
			want: []point{
				point{0, 0},
				point{0, 1},
				point{1, 1},
				point{1, 0},
				point{0, 0},
			},
		},
		{
			s: "U1,R1,D1,L1,U2,R2,D2,L2",
			want: []point{
				point{0, 0},
				point{0, 1},
				point{1, 1},
				point{1, 0},
				point{0, 0},
				point{0, 2},
				point{2, 2},
				point{2, 0},
				point{0, 0},
			},
		},
	}

	for _, tc := range testCases {
		if got := parsePath(tc.s); !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("got %#v, want %#v", got, tc.want)
		}
	}

}

func TestCross(t *testing.T) {

	testCases := []struct {
		a1, b1, a2, b2 point
		wantPoint      point
		wantCross      bool
	}{
		{
			a1:        point{0, 0},
			b1:        point{10, 0},
			a2:        point{0, 0},
			b2:        point{0, 10},
			wantPoint: point{0, 0},
			wantCross: true,
		},
		{
			a1:        point{1, 1},
			b1:        point{1, 10},
			a2:        point{0, 5},
			b2:        point{10, 5},
			wantPoint: point{1, 5},
			wantCross: true,
		},
	}

	for _, tc := range testCases {
		if gotPoint, gotCross := cross(tc.a1, tc.b1, tc.a2, tc.b2); gotCross != tc.wantCross || !reflect.DeepEqual(gotPoint, tc.wantPoint) {
			t.Fatalf("cross(%v, %v, %v, %v) = (%v, %v), want (%v, %v)", tc.a1, tc.b1, tc.a2, tc.b2, gotPoint, gotCross, tc.wantPoint, tc.wantCross)
		}
	}

}

func TestSolve1(t *testing.T) {

	testCases := []struct {
		p1, p2 string
		want   int
	}{
		{
			p1:   "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			p2:   "U62,R66,U55,R34,D71,R55,D58,R83",
			want: 159,
		},
		{
			p1:   "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			p2:   "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			want: 135,
		},
	}

	for _, tc := range testCases {
		if got := solve1(parsePath(tc.p1), parsePath(tc.p2)); got != tc.want {
			t.Fatalf("solve1(%q, %q) = %d, want %d", tc.p1, tc.p2, got, tc.want)
		}
	}

}

func TestSolve2(t *testing.T) {

	testCases := []struct {
		p1, p2 string
		want   int
	}{
		{
			p1:   "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			p2:   "U62,R66,U55,R34,D71,R55,D58,R83",
			want: 610,
		},
		{
			p1:   "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			p2:   "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			want: 410,
		},
	}

	for _, tc := range testCases {
		if got := solve2(parsePath(tc.p1), parsePath(tc.p2)); got != tc.want {
			t.Fatalf("solve2(%q, %q) = %d, want %d", tc.p1, tc.p2, got, tc.want)
		}
	}

}
