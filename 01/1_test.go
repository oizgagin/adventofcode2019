package main

import "testing"

func TestSolve(t *testing.T) {

	t.Run("pt1", func(t *testing.T) {
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
			if got, want := solve(tc.input, 1), tc.want; got != want {
				t.Fatalf("solve(%v) = %d, want %d", tc.input, got, want)
			}
		}
	})

	t.Run("pt2", func(t *testing.T) {
		testCases := []struct {
			input []int
			want  int
		}{
			{
				input: []int{14, 1969, 100756},
				want:  2 + 966 + 50346,
			},
		}

		for _, tc := range testCases {
			if got, want := solve(tc.input, 2), tc.want; got != want {
				t.Fatalf("solve(%v) = %d, want %d", tc.input, got, want)
			}
		}
	})

}

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

func TestCalcFuelRecursively(t *testing.T) {

	testCases := []struct {
		mass     int
		wantFuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, tc := range testCases {
		if got, want := calcFuelRecursively(tc.mass), tc.wantFuel; got != want {
			t.Fatalf("calcFuel(%d) = %d, want %d", tc.mass, got, want)
		}
	}

}
