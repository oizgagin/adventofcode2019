package intcode_test

import (
	"reflect"
	"testing"

	"github.com/oizgagin/adventofcode2019/intcode"
)

func TestIntcode_Exec(t *testing.T) {

	testCases := []struct {
		mem         string
		want        string
		inputs      []int
		wantOutputs []int
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50", nil, nil},
		{"1,0,0,0,99", "2,0,0,0,99", nil, nil},
		{"2,3,0,3,99", "2,3,0,6,99", nil, nil},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801", nil, nil},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99", nil, nil},
		{"1002,4,3,4,33", "1002,4,3,4,99", nil, nil},
		{"1101,100,-1,4,0", "1101,100,-1,4,99", nil, nil},
		{"3,2", "3,2,99", []int{99}, nil},
		{"4,2,99", "4,2,99", nil, []int{99}},
	}

	for _, tc := range testCases {
		mem := intcode.NewMemoryFromString(tc.mem)

		cpu := intcode.NewCPU()
		cpu.LoadMemory(mem)

		inputs := make(chan int)
		go func() {
			for _, input := range tc.inputs {
				inputs <- input
			}
			close(inputs)
		}()

		outputs, outputsCh := make([]int, 0), make(chan int)
		go func() {
			for v := range outputsCh {
				outputs = append(outputs, v)
			}
		}()

		cpu.Exec(inputs, outputsCh)

		if got := cpu.Memory().String(); got != tc.want {
			t.Fatalf("got %v, want %v", got, tc.want)
		}

		if tc.wantOutputs != nil {
			if !reflect.DeepEqual(outputs, tc.wantOutputs) {
				t.Fatalf("got %v outputs, want %v", outputs, tc.wantOutputs)
			}
		}
	}

}

func TestIntcode_ParseOpcode(t *testing.T) {

	testCases := []struct {
		v          int
		wantOpcode intcode.Opcode
		wantModes  []intcode.ParamMode
	}{
		{1002, intcode.OpcodeMul, []intcode.ParamMode{intcode.ModePosition, intcode.ModeImmediate, intcode.ModePosition}},
	}

	for _, tc := range testCases {
		if gotOpcode, gotModes := intcode.ParseOpcode(tc.v); gotOpcode != tc.wantOpcode || !reflect.DeepEqual(gotModes, tc.wantModes) {
			t.Fatalf("ParseOpcode(%d) = (%v, %v), want (%v, %v)", tc.v, gotOpcode, gotModes, tc.wantOpcode, tc.wantModes)
		}
	}

}
