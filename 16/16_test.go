package main

import (
	"reflect"
	"testing"
)

func TestPhases(t *testing.T) {

	testCases := []struct {
		input  []int
		phases int
		want   []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			phases: 4,
			want:   []int{0, 1, 0, 2, 9, 4, 9, 8},
		},
		{
			input:  []int{8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5},
			phases: 100,
			want:   []int{2, 4, 1, 7, 6, 1, 7, 6},
		},
		{
			input:  []int{1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7},
			phases: 100,
			want:   []int{7, 3, 7, 4, 5, 4, 1, 8},
		},
		{
			input:  []int{6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3},
			phases: 100,
			want:   []int{5, 2, 4, 3, 2, 1, 3, 3},
		},
	}

	for i, tc := range testCases {
		input := tc.input
		for i := 0; i < tc.phases; i++ {
			input = phase(input)
		}
		if got, want := input[:8], tc.want; !reflect.DeepEqual(got, want) {
			t.Fatalf("#%d: got %v, want %v", i, got, want)
		}
	}

}

func TestSolve2(t *testing.T) {

	testCases := []struct {
		input string
		want  string
	}{
		{
			input: "03036732577212944063491565474664",
			want:  "84462026",
		},
		{
			input: "02935109699940807407585447034323",
			want:  "78725270",
		},
		{
			input: "03081770884921959731165446850517",
			want:  "53553731",
		},
	}

	for i, tc := range testCases {
		input, _ := parse([]string{tc.input})
		if got, want := solve2(input), tc.want; got != want {
			t.Fatalf("#%d: got %v, want %v", i, got, want)
		}
	}

}
