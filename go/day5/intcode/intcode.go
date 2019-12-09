package intcode

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Program []int64

type IO interface {
	ReadInt() int64
	WriteInt(int64)
}

func Input(i int64) IO {
	return sideEffects{input: i}
}

type TestIO struct {
	Input  int64
	Output []int64
}

func (t *TestIO) ReadInt() int64 {
	return t.Input
}

func (t *TestIO) WriteInt(i int64) {
	t.Output = append(t.Output, i)
}

var _ IO = &TestIO{}

type sideEffects struct {
	input int64
}

func (s sideEffects) ReadInt() int64 {
	return s.input
}

func (s sideEffects) WriteInt(i int64) {
	fmt.Println(i)
}

var _ IO = &sideEffects{}

type cmd struct {
	ip      int64
	rb      int64
	op      int64
	opName  string
	numArgs int
	modes   []int64
}

func parse(ip, rb int64, op int64) cmd {
	opCode := op % 100
	var numArgs int
	var opStr string
	switch opCode {
	case 1:
		opStr = "+"
		numArgs = 3
	case 2:
		opStr = "*"
		numArgs = 3
	case 7:
		opStr = "<"
		numArgs = 3
	case 8:
		opStr = "="
		numArgs = 3
	case 3:
		opStr = "read"
		numArgs = 1
	case 4:
		opStr = "write"
		numArgs = 1
	case 9:
		opStr = "add_rb"
		numArgs = 1
	case 5:
		opStr = "jmp_true"
		numArgs = 2
	case 6:
		opStr = "jmp_false"
		numArgs = 2
	}
	return cmd{
		ip:      ip,
		rb:      rb,
		op:      opCode,
		opName:  opStr,
		numArgs: numArgs,
		modes:   []int64{0, (op / 100) % 10, (op / 1000) % 10, (op / 10000) % 10},
	}
}

func (c cmd) offset() int64 {
	return int64(c.numArgs + 1)
}

func (c cmd) ptr(pos, mode int64, p Program) int64 {
	idx := c.ip + pos
	ptr := p[idx]
	switch mode {
	case 0:
		return ptr
	case 1:
		return idx
	case 2:
		return c.rb + ptr
	}
	panic("not a supported arg mode")
}

func (c cmd) arg(pos, mode int64, p Program) int64 {
	return p[c.ptr(pos, mode, p)]
}

func (c cmd) arg1(p Program) int64 {
	return c.arg(1, c.modes[1], p)
}

func (c cmd) arg2(p Program) int64 {
	return c.arg(2, c.modes[2], p)
}

func (c cmd) arg3(p Program) int64 {
	return c.arg(3, c.modes[3], p)
}

func (c cmd) ptr3(p Program) int64 {
	return c.ptr(3, c.modes[3], p)
}

func (c cmd) maxPtr(p Program) int64 {
	var max, i int64
	for i = 1; i <= int64(c.numArgs); i++ {
		ptr := c.ptr(i, c.modes[i], p)
		if ptr > max {
			max = ptr
		}
	}
	return max
}

func (c cmd) String(p Program) string {
	str := fmt.Sprintf("ip: %d rb: %d op: %s", c.ip, c.rb, c.opName)
	for i := 1; i <= c.numArgs; i++ {
		str = str + fmt.Sprintf(" arg_%d: %d", i, c.arg(int64(i), c.modes[i], p))
	}
	return str
}

func alloc(src Program, c cmd, debug bool) Program {
	allocated := src
	diff := (c.maxPtr(src) + 1) - int64(len(src))
	if diff > 0 {
		if debug {
			println(fmt.Sprintf("allocating adding %d to %d new len %d", diff, len(src), len(src)+int(diff)))
		}
		allocated = make(Program, len(src)+int(diff))
		copy(allocated, src)
	}
	return allocated
}

func intcode(prgr Program, io IO, debug bool) {
	var ip, rb int64
	var exit bool
	for !exit {
		cmd := parse(ip, rb, prgr[ip])
		prgr = alloc(prgr, cmd, debug)
		if debug {
			fmt.Printf("{%s} %v\n", cmd.String(prgr), prgr[ip:ip+cmd.offset()])
		}
		switch cmd.op {
		case 1: // addition
			prgr[cmd.ptr3(prgr)] = cmd.arg1(prgr) + cmd.arg2(prgr)
			ip += cmd.offset()
		case 2: // multiplication
			prgr[cmd.ptr3(prgr)] = cmd.arg1(prgr) * cmd.arg2(prgr)
			ip += cmd.offset()
		case 3: // read
			prgr[cmd.ptr3(prgr)] = io.ReadInt()
			ip += cmd.offset()
		case 4: // write
			io.WriteInt(cmd.arg1(prgr))
			ip += cmd.offset()
		case 5: // jmp if true
			if cmd.arg1(prgr) != 0 {
				ip = cmd.arg2(prgr)
			} else {
				ip += cmd.offset()
			}
		case 6: // jmp if false
			if cmd.arg1(prgr) == 0 {
				ip = cmd.arg2(prgr)
			} else {
				ip += cmd.offset()
			}
		case 7: // cmp less than
			var x int64 = 0
			if cmd.arg1(prgr) < cmd.arg2(prgr) {
				x = 1
			}
			prgr[cmd.ptr3(prgr)] = x
			ip += cmd.offset()
		case 8: // cmp equals
			var x int64 = 0
			if cmd.arg1(prgr) == cmd.arg2(prgr) {
				x = 1
			}
			prgr[cmd.ptr3(prgr)] = x
			ip += cmd.offset()
		case 9: // set rb
			rb += cmd.arg1(prgr)
			ip += cmd.offset()
		case 99:
			exit = true
		}

	}
}

func RunWithInput(prgr Program, input IO, debug bool) {
	executable := make([]int64, len(prgr))
	copy(executable, prgr)
	intcode(executable, input, debug)
}

func ParseProgram(fileName string) Program {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(string(bytes), ",")
	baseProgram := make([]int64, len(strs))
	for i, s := range strs {
		intVal, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		baseProgram[i] = intVal
	}
	return baseProgram
}
