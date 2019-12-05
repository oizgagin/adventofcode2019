package intcode

import (
	"strconv"
	"strings"
)

type Memory map[int]int

func NewMemory() Memory {
	return Memory(make(map[int]int))
}

func NewMemoryFromString(s string) Memory {
	mem := NewMemory()
	for i, ss := range strings.Split(s, ",") {
		v, _ := strconv.Atoi(ss)
		mem[i] = v
	}
	return mem
}

func (mem Memory) String() string {
	ss := make([]string, len(mem))
	for i := 0; i < len(mem); i++ {
		ss[i] = strconv.Itoa(mem[i])
	}
	return strings.Join(ss, ",")
}

type CPU struct {
	ip  int
	mem Memory
}

func NewCPU() *CPU {
	return &CPU{ip: 0}
}

func (cpu *CPU) LoadMemory(mem Memory) {
	cpu.mem = mem
}

type Opcode int

const (
	OpcodeAdd    Opcode = 1
	OpcodeMul    Opcode = 2
	OpcodeInput  Opcode = 3
	OpcodeOutput Opcode = 4
	OpcodeHalt   Opcode = 99
)

var OpcodeLens = map[Opcode]int{
	OpcodeAdd:    4,
	OpcodeMul:    4,
	OpcodeHalt:   1,
	OpcodeInput:  2,
	OpcodeOutput: 2,
}

type ParamMode int

const (
	ModePosition  ParamMode = 0
	ModeImmediate ParamMode = 1
)

func (cpu *CPU) Exec(input <-chan int, output chan<- int) {
	for {
		opcode, modes := ParseOpcode(cpu.mem[cpu.ip])
		if opcode == OpcodeHalt {
			return
		}

		switch opcode {
		case OpcodeAdd:
			cpu.execAdd(modes)
		case OpcodeMul:
			cpu.execMul(modes)
		case OpcodeInput:
			cpu.execInput(modes, input)
		case OpcodeOutput:
			cpu.execOutput(modes, output)
		}

		cpu.ip += OpcodeLens[opcode]
	}
}

func ParseOpcode(v int) (opcode Opcode, modes []ParamMode) {
	opcode = Opcode(v % 100)
	v /= 100

	total := OpcodeLens[opcode] - 1

	modes = make([]ParamMode, total)
	for i := 0; i < total; i++ {
		modes[i] = ParamMode(v % 10)
		v /= 10
	}
	return
}

func (cpu *CPU) getParam(v int, mode ParamMode) int {
	if mode == ModeImmediate {
		return v
	}
	return cpu.mem[v]
}

func (cpu *CPU) execAdd(modes []ParamMode) {
	in1, in2, out := cpu.getParam(cpu.mem[cpu.ip+1], modes[0]), cpu.getParam(cpu.mem[cpu.ip+2], modes[1]), cpu.mem[cpu.ip+3]
	cpu.mem[out] = in1 + in2
}

func (cpu *CPU) execMul(modes []ParamMode) {
	in1, in2, out := cpu.getParam(cpu.mem[cpu.ip+1], modes[0]), cpu.getParam(cpu.mem[cpu.ip+2], modes[1]), cpu.mem[cpu.ip+3]
	cpu.mem[out] = in1 * in2
}

func (cpu *CPU) execInput(modes []ParamMode, input <-chan int) {
	out := cpu.mem[cpu.ip+1]
	cpu.mem[out] = <-input
}

func (cpu *CPU) execOutput(modes []ParamMode, output chan<- int) {
	output <- cpu.getParam(cpu.mem[cpu.ip+1], modes[0])
}

func (cpu *CPU) Memory() Memory {
	return cpu.mem
}
