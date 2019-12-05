package intcode_test

import (
	"testing"

	"github.com/oizgagin/adventofcode2019/intcode"
)

func TestIntcode_Exec(t *testing.T) {

	testCases := []struct {
		mem  string
		want string
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50"},
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}

	for _, tc := range testCases {
		mem := intcode.NewMemoryFromString(tc.mem)

		cpu := intcode.NewCPU()
		cpu.LoadMemory(mem)

		cpu.Exec()

		if got := cpu.Memory().String(); got != tc.want {
			t.Fatalf("got %v, want %v", got, tc.want)
		}
	}

}
