package main

import (
	"reflect"
	"testing"
)

func TestSegment_Cross(t *testing.T) {

	testCases := []struct {
		a, b      *Segment
		wantPoint *Point
		wantCross bool
	}{
		{
			a:         makeSegment(0, 0, 0, 2),
			b:         makeSegment(0, 0, 2, 0),
			wantPoint: makePoint(0, 0),
			wantCross: true,
		},
		{
			a:         makeSegment(1, 1, 1, 2),
			b:         makeSegment(0, 0, 1, 0),
			wantPoint: nil,
			wantCross: false,
		},
		{
			a:         makeSegment(1, 1, 10, 1),
			b:         makeSegment(5, 0, 5, 5),
			wantPoint: makePoint(5, 1),
			wantCross: true,
		},
		{
			a:         makeSegment(100, 200, 100, 50),
			b:         makeSegment(50, 100, 200, 100),
			wantPoint: makePoint(100, 100),
			wantCross: true,
		},
		{
			a:         makeSegment(100, 200, 100, 50),
			b:         makeSegment(50, 0, 200, 0),
			wantPoint: nil,
			wantCross: false,
		},
	}

	for _, tc := range testCases {
		gotPoint, gotCross := tc.a.Cross(tc.b)

		if gotCross != tc.wantCross {
			t.Fatalf("cross(%v, %v) = %v, want %v", tc.a, tc.b, gotCross, tc.wantCross)
		}
		if !reflect.DeepEqual(gotPoint, tc.wantPoint) {
			t.Fatalf("got %v, want %v", gotPoint, tc.wantPoint)
		}
	}

}

func TestSolve1(t *testing.T) {

	testCases := []struct {
		p1, p2   string
		wantDist int
	}{
		{
			p1:       "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			p2:       "U62,R66,U55,R34,D71,R55,D58,R83",
			wantDist: 159,
		},
		{
			p1:       "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			p2:       "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			wantDist: 135,
		},
	}

	for _, tc := range testCases {
		gotDist := solve1([]*Path{
			unmarshalPath(tc.p1),
			unmarshalPath(tc.p2),
		})

		if gotDist != tc.wantDist {
			t.Fatalf("got %d, want %d", gotDist, tc.wantDist)
		}
	}

}
