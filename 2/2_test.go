package main

import "testing"

func TestExec(t *testing.T) {

	testCases := []struct {
		rawMem  string
		wantMem string
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50"},
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}

	for _, tc := range testCases {
		mem, err := parseMemory(tc.rawMem)
		if err != nil {
			t.Fatalf("parseMemory(%q) = %v", tc.rawMem, err)
		}

		exec(mem)

		if got, want := marshalMem(mem), tc.wantMem; got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	}

}
