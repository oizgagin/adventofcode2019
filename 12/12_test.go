package main

import "testing"

func TestTick(t *testing.T) {

	testCases := []struct {
		coords string
		steps  int
		want   int
	}{
		{
			coords: `
				<x=-1, y=0, z=2>
				<x=2, y=-10, z=-7>
				<x=4, y=-8, z=8>
				<x=3, y=5, z=-1>
			`,
			steps: 10,
			want:  179,
		},
		{
			coords: `
				<x=-8, y=-10, z=0>
				<x=5, y=5, z=10>
				<x=2, y=-7, z=3>
				<x=9, y=-8, z=-3>
			`,
			steps: 100,
			want:  1940,
		},
	}

	for i, tc := range testCases {
		system := NewSystem(parseCoords(tc.coords))
		for i := 0; i < tc.steps; i++ {
			system.Tick()
		}

		if got := system.Energy(); got != tc.want {
			t.Fatalf("#%d: got %d, want %d", i, got, tc.want)
		}
	}

}

func TestSolve2(t *testing.T) {

	testCases := []struct {
		coords string
		want   int
	}{
		{
			coords: `
				<x=-1, y=0, z=2>
				<x=2, y=-10, z=-7>
				<x=4, y=-8, z=8>
				<x=3, y=5, z=-1>
			`,
			want: 2772,
		},
		{
			coords: `
				<x=-8, y=-10, z=0>
				<x=5, y=5, z=10>
				<x=2, y=-7, z=3>
				<x=9, y=-8, z=-3>
			`,
			want: 4686774924,
		},
	}

	for i, tc := range testCases {
		if got := solve2(parseCoords(tc.coords)); got != tc.want {
			t.Fatalf("#%d: got %d, want %d", i, got, tc.want)
		}
	}

}
