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

		{"3,9,8,9,10,9,4,9,99,-1,8", "3,9,8,9,10,9,4,9,99,1,8", []int{8}, []int{1}},
		{"3,9,8,9,10,9,4,9,99,-1,8", "3,9,8,9,10,9,4,9,99,0,8", []int{9}, []int{0}},

		{"3,3,1108,-1,8,3,4,3,99", "3,3,1108,1,8,3,4,3,99", []int{8}, []int{1}},
		{"3,3,1108,-1,8,3,4,3,99", "3,3,1108,0,8,3,4,3,99", []int{7}, []int{0}},

		{"3,3,1107,-1,8,3,4,3,99", "3,3,1107,0,8,3,4,3,99", []int{9}, []int{0}},
		{"3,3,1107,-1,8,3,4,3,99", "3,3,1107,1,8,3,4,3,99", []int{7}, []int{1}},

		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "", []int{0}, []int{0}},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", "", []int{10}, []int{1}},

		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", "", []int{0}, []int{0}},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", "", []int{10}, []int{1}},

		{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "", []int{7}, []int{999}},
		{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "", []int{8}, []int{1000}},
		{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", "", []int{9}, []int{1001}},
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

		outputs, collectOutputs := outputs()

		cpu.Exec(inputs, outputs)

		if got := cpu.Memory().String(); tc.want != "" && got != tc.want {
			t.Fatalf("got %v, want %v", got, tc.want)
		}

		if tc.wantOutputs != nil {
			if outputs := collectOutputs(); !reflect.DeepEqual(outputs, tc.wantOutputs) {
				t.Fatalf("got %v outputs, want %v", outputs, tc.wantOutputs)
			}
		}
	}

}

func outputs() (chan int, func() []int) {
	outsCh, outs := make(chan int), make([]int, 0)

	started := make(chan struct{})
	finished := make(chan struct{})

	go func() {
		close(started)
		defer close(finished)

		for out := range outsCh {
			outs = append(outs, out)
		}
	}()

	<-started

	return outsCh, func() []int {
		<-finished
		return outs
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
