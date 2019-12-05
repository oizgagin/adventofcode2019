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

const (
	OpCodeAdd  = 1
	OpCodeMul  = 2
	OpCodeHalt = 99
)

func (cpu *CPU) Exec() {
	opcode := cpu.mem[cpu.ip]
	for opcode != OpCodeHalt {
		switch opcode {
		case OpCodeAdd:
			cpu.execAdd()
		case OpCodeMul:
			cpu.execMul()
		}

		cpu.ip += 4
		opcode = cpu.mem[cpu.ip]
	}
}

func (cpu *CPU) execAdd() {
	in1, in2, out1 := cpu.mem[cpu.ip+1], cpu.mem[cpu.ip+2], cpu.mem[cpu.ip+3]
	cpu.mem[out1] = cpu.mem[in1] + cpu.mem[in2]
}

func (cpu *CPU) execMul() {
	in1, in2, out1 := cpu.mem[cpu.ip+1], cpu.mem[cpu.ip+2], cpu.mem[cpu.ip+3]
	cpu.mem[out1] = cpu.mem[in1] * cpu.mem[in2]
}

func (cpu *CPU) Memory() Memory {
	return cpu.mem
}
