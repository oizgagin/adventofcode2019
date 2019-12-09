package main

import (
	"testing"

	"github.com/oizgagin/adventofcode2019/intcode"
)

func TestSolve1(t *testing.T) {

	testCases := []struct {
		prog string
		want int
	}{
		{
			prog: "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			want: 43210,
		},
		{
			prog: "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			want: 54321,
		},
		{
			prog: "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			want: 65210,
		},
	}

	for _, tc := range testCases {
		mem := intcode.NewMemoryFromString(tc.prog)
		if got := solve1(interface{}(mem)); got != tc.want {
			t.Fatalf("got %d, want %d", got, tc.want)
		}
	}

}
