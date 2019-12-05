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

func (cpu *CPU) Memory() Memory {
	return cpu.mem
}

type Opcode int

const (
	OpcodeAdd    Opcode = 1
	OpcodeMul    Opcode = 2
	OpcodeInput  Opcode = 3
	OpcodeOutput Opcode = 4
	OpcodeJt     Opcode = 5
	OpcodeJf     Opcode = 6
	OpcodeLt     Opcode = 7
	OpcodeEq     Opcode = 8
	OpcodeHalt   Opcode = 99
)

type ParamMode int

const (
	ModePosition  ParamMode = 0
	ModeImmediate ParamMode = 1
)

func (cpu *CPU) Exec(input <-chan int, output chan<- int) {
	defer close(output)

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
		case OpcodeJt:
			cpu.execJt(modes)
		case OpcodeJf:
			cpu.execJf(modes)
		case OpcodeLt:
			cpu.execLt(modes)
		case OpcodeEq:
			cpu.execEq(modes)
		}
	}
}

func (cpu *CPU) execAdd(modes []ParamMode) {
	in1, in2, out := cpu.getParam(cpu.mem[cpu.ip+1], modes[0]), cpu.getParam(cpu.mem[cpu.ip+2], modes[1]), cpu.mem[cpu.ip+3]
	cpu.mem[out] = in1 + in2
	cpu.ip += 4
}

func (cpu *CPU) execMul(modes []ParamMode) {
	in1, in2, out := cpu.getParam(cpu.mem[cpu.ip+1], modes[0]), cpu.getParam(cpu.mem[cpu.ip+2], modes[1]), cpu.mem[cpu.ip+3]
	cpu.mem[out] = in1 * in2
	cpu.ip += 4
}

func (cpu *CPU) execInput(modes []ParamMode, input <-chan int) {
	out := cpu.mem[cpu.ip+1]
	cpu.mem[out] = <-input
	cpu.ip += 2
}

func (cpu *CPU) execOutput(modes []ParamMode, output chan<- int) {
	output <- cpu.getParam(cpu.mem[cpu.ip+1], modes[0])
	cpu.ip += 2
}

func (cpu *CPU) execJt(modes []ParamMode) {
	v, addr := cpu.getParam(cpu.mem[cpu.ip+1], modes[0]), cpu.getParam(cpu.mem[cpu.ip+2], modes[1])
	if v != 0 {
		cpu.ip = addr
	} else {
		cpu.ip += 3
	}
}

func (cpu *CPU) execJf(modes []ParamMode) {
	v, addr := cpu.getParam(cpu.mem[cpu.ip+1], modes[0]), cpu.getParam(cpu.mem[cpu.ip+2], modes[1])
	if v == 0 {
		cpu.ip = addr
	} else {
		cpu.ip += 3
	}
}

func (cpu *CPU) execLt(modes []ParamMode) {
	in1, in2 := cpu.getParam(cpu.mem[cpu.ip+1], modes[0]), cpu.getParam(cpu.mem[cpu.ip+2], modes[1])
	addr := cpu.mem[cpu.ip+3]

	if in1 < in2 {
		cpu.mem[addr] = 1
	} else {
		cpu.mem[addr] = 0
	}

	cpu.ip += 4
}

func (cpu *CPU) execEq(modes []ParamMode) {
	in1, in2 := cpu.getParam(cpu.mem[cpu.ip+1], modes[0]), cpu.getParam(cpu.mem[cpu.ip+2], modes[1])
	addr := cpu.mem[cpu.ip+3]

	if in1 == in2 {
		cpu.mem[addr] = 1
	} else {
		cpu.mem[addr] = 0
	}

	cpu.ip += 4
}

func (cpu *CPU) getParam(v int, mode ParamMode) int {
	if mode == ModeImmediate {
		return v
	}
	return cpu.mem[v]
}

var OpcodeArity = map[Opcode]int{
	OpcodeAdd:    3,
	OpcodeMul:    3,
	OpcodeInput:  1,
	OpcodeOutput: 1,
	OpcodeJt:     2,
	OpcodeJf:     2,
	OpcodeLt:     3,
	OpcodeEq:     3,
	OpcodeHalt:   0,
}

func ParseOpcode(v int) (opcode Opcode, modes []ParamMode) {
	opcode = Opcode(v % 100)
	v /= 100

	total := OpcodeArity[opcode]

	modes = make([]ParamMode, total)
	for i := 0; i < total; i++ {
		modes[i] = ParamMode(v % 10)
		v /= 10
	}
	return
}
