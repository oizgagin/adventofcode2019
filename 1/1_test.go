package main

import "testing"

func TestCalcFuel(t *testing.T) {

	testCases := []struct {
		mass     int
		wantFuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, tc := range testCases {
		if got, want := calcFuel(tc.mass), tc.wantFuel; got != want {
			t.Fatalf("calcFuel(%d) = %d, want %d", tc.mass, got, want)
		}
	}

}

func TestSolve(t *testing.T) {

	testCases := []struct {
		input []int
		want  int
	}{
		{
			input: []int{12, 14, 1969, 100756},
			want:  2 + 2 + 654 + 33583,
		},
	}

	for _, tc := range testCases {
		if got, want := solve(tc.input), tc.want; got != want {
			t.Fatalf("solve(%v) = %d, want %d", tc.input, got, want)
		}
	}

}
