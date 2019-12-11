package main

import "testing"

func TestIsValid(t *testing.T) {

	testCases := []struct {
		n    int
		want bool
	}{
		{122345, true},
		{111123, true},
		{135679, false},
		{111111, true},
		{223450, false},
		{123789, false},
	}

	for _, tc := range testCases {
		if got := isValid(tc.n); got != tc.want {
			t.Fatalf("isValid(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}

func TestIsValid2(t *testing.T) {

	testCases := []struct {
		n    int
		want bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for _, tc := range testCases {
		if got := isValid2(tc.n); got != tc.want {
			t.Fatalf("isValid2(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
