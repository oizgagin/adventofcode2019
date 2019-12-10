package intcode

import (
	"fmt"
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

func (mem Memory) Get(pos int) int {
	return mem[pos]
}

func (mem Memory) Set(pos, value int) {
	mem[pos] = value
}

func (mem Memory) Copy() Memory {
	m := NewMemory()
	for addr, value := range mem {
		m[addr] = value
	}
	return m
}

type CPU struct {
	ip     int
	mem    Memory
	input  func() int
	output func(int)
}

func NewCPU(mem Memory, input func() int, output func(int)) *CPU {
	return &CPU{ip: 0, mem: mem, input: input, output: output}
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

func (opcode Opcode) String() string {
	strings := map[Opcode]string{
		OpcodeAdd:    "ADD",
		OpcodeMul:    "MUL",
		OpcodeInput:  "INPUT",
		OpcodeOutput: "OUTPUT",
		OpcodeJt:     "JT",
		OpcodeJf:     "JF",
		OpcodeLt:     "LT",
		OpcodeEq:     "EQ",
		OpcodeHalt:   "HALT",
	}
	return strings[opcode]
}

type ParamMode int

const (
	ModePosition  ParamMode = 0
	ModeImmediate ParamMode = 1
)

type ParamModes int

func (modes ParamModes) Get(i int) ParamMode {
	m := int(modes)
	for j := 0; j < i; j++ {
		m /= 10
	}
	return ParamMode(m % 10)
}

type CPUState string

const (
	CPUHalt    CPUState = "HALT"
	CPUSuspend CPUState = "SUSPEND"
)

func (cpu *CPU) Exec() CPUState {
	for {
		switch opcode, modes := parseOpcode(cpu.mem.Get(cpu.ip)); opcode {
		case OpcodeAdd:
			cpu.execAdd(modes)
		case OpcodeMul:
			cpu.execMul(modes)
		case OpcodeInput:
			cpu.execInput(modes)
		case OpcodeOutput:
			cpu.execOutput(modes)
			return CPUSuspend
		case OpcodeJt:
			cpu.execJt(modes)
		case OpcodeJf:
			cpu.execJf(modes)
		case OpcodeLt:
			cpu.execLt(modes)
		case OpcodeEq:
			cpu.execEq(modes)
		case OpcodeHalt:
			return CPUHalt
		}
	}
}

func (cpu *CPU) execAdd(modes ParamModes) {
	var (
		in1 = cpu.param(cpu.ip+1, modes.Get(0))
		in2 = cpu.param(cpu.ip+2, modes.Get(1))
		out = cpu.param(cpu.ip+3, ModeImmediate)
	)
	cpu.mem.Set(out, in1+in2)
	cpu.ip += 4
}

func (cpu *CPU) execMul(modes ParamModes) {
	var (
		in1 = cpu.param(cpu.ip+1, modes.Get(0))
		in2 = cpu.param(cpu.ip+2, modes.Get(1))
		out = cpu.param(cpu.ip+3, ModeImmediate)
	)
	cpu.mem.Set(out, in1*in2)
	cpu.ip += 4
}

func (cpu *CPU) execInput(modes ParamModes) {
	out := cpu.param(cpu.ip+1, ModeImmediate)
	cpu.mem.Set(out, cpu.input())
	cpu.ip += 2
}

func (cpu *CPU) execOutput(modes ParamModes) {
	outval := cpu.param(cpu.ip+1, modes.Get(0))
	cpu.output(outval)
	cpu.ip += 2
}

func (cpu *CPU) execJt(modes ParamModes) {
	var (
		value = cpu.param(cpu.ip+1, modes.Get(0))
		addr  = cpu.param(cpu.ip+2, modes.Get(1))
	)
	if value != 0 {
		cpu.ip = addr
	} else {
		cpu.ip += 3
	}
}

func (cpu *CPU) execJf(modes ParamModes) {
	var (
		value = cpu.param(cpu.ip+1, modes.Get(0))
		addr  = cpu.param(cpu.ip+2, modes.Get(1))
	)
	if value == 0 {
		cpu.ip = addr
	} else {
		cpu.ip += 3
	}
}

func (cpu *CPU) execLt(modes ParamModes) {
	var (
		in1  = cpu.param(cpu.ip+1, modes.Get(0))
		in2  = cpu.param(cpu.ip+2, modes.Get(1))
		addr = cpu.param(cpu.ip+3, ModeImmediate)
	)
	if in1 < in2 {
		cpu.mem.Set(addr, 1)
	} else {
		cpu.mem.Set(addr, 0)
	}
	cpu.ip += 4
}

func (cpu *CPU) execEq(modes ParamModes) {
	var (
		in1  = cpu.param(cpu.ip+1, modes.Get(0))
		in2  = cpu.param(cpu.ip+2, modes.Get(1))
		addr = cpu.param(cpu.ip+3, ModeImmediate)
	)
	if in1 == in2 {
		cpu.mem.Set(addr, 1)
	} else {
		cpu.mem.Set(addr, 0)
	}
	cpu.ip += 4
}

func (cpu *CPU) param(addr int, mode ParamMode) int {
	v := cpu.mem.Get(addr)

	if mode == ModeImmediate {
		return v
	}
	if mode == ModePosition {
		return cpu.mem.Get(v)
	}
	panic(fmt.Sprintf("invalid mode: %v", mode))
}

func parseOpcode(v int) (Opcode, ParamModes) {
	return Opcode(v % 100), ParamModes(v / 100)
}
