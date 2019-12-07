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
	ReadInt()  int64
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
	ip int64
	op      int64
	modArg1 int64
	modArg2 int64
	modArg3 int64
}

func parse(ip int64, op int64) cmd {
	return cmd{
		ip: ip,
		op:      op % 100,
		modArg1: (op / 100) % 10,
		modArg2: (op / 1000) % 10,
		modArg3: (op / 10000) % 10,
	}
}

func (c cmd) arg(pos, mode int64, p Program) int64 {
	arg := p[c.ip+pos]
	if mode == 0 {
		return p[arg]
	}
	return arg
}

func (c cmd) arg1(p Program) int64 {
	return c.arg(1, c.modArg1, p)
}

func (c cmd) arg2(p Program) int64 {
	return c.arg(2, c.modArg2, p)
}

func (c cmd) arg3(p Program) int64 {
	return c.arg(3, c.modArg3, p)
}

func intcode(prgr Program, io IO, debug bool)  {
	var ip int64
	var exit bool
	for !exit {
		cmd := parse(ip, prgr[ip])
		switch cmd.op {
		case 1:
			prgr[prgr[ip+3]] = cmd.arg1(prgr) + cmd.arg2(prgr)
			ip += 4
		case 2:
			prgr[prgr[ip+3]] = cmd.arg1(prgr) * cmd.arg2(prgr)
			ip += 4
		case 3:
			prgr[prgr[ip+1]] = io.ReadInt()
			ip += 2
		case 4:
			io.WriteInt(prgr[prgr[ip+1]])
			 ip += 2
		case 5:
			if cmd.arg1(prgr) != 0 {
				ip = cmd.arg2(prgr)
			} else {
				ip += 3
			}
		case 6:
			if cmd.arg1(prgr) == 0 {
				ip = cmd.arg2(prgr)
			} else {
				ip += 3
			}
		case 7:
			var x int64 = 0
			if cmd.arg1(prgr) < cmd.arg2(prgr){
				x = 1
			}
			prgr[prgr[ip+3]] = x
			ip += 4
		case 8:
			var x int64 = 0
			if cmd.arg1(prgr) == cmd.arg2(prgr){
				x = 1
			}
			prgr[prgr[ip+3]] = x
			ip += 4
		case 99:
			exit = true
		}
		if debug {
			fmt.Printf("%+v: %v\n", cmd, prgr)
		}
	}
}

func RunWithInput(prgr Program, input IO, debug bool)  {
	executable := make([]int64, len(prgr))
	copy(executable, prgr)
	intcode(executable, input, debug)
}

func ParseProgram(fileName string) Program{
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