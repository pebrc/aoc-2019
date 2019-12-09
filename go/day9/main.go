package main

import (
	"strconv"

	"github.com/pebrc/aoc-2019/go/day5/intcode"
	"github.com/pebrc/aoc-2019/go/test"
)

func main() {
	io := &intcode.TestIO{}
	test1 := []int64{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}
	intcode.RunWithInput(test1, io, false)
	test.AssertN(io.Output, test1)

	io = &intcode.TestIO{}
	test2 := [] int64 {1102,34915192,34915192,7,4,7,99,0}
	intcode.RunWithInput(test2, io, false)
	test.Assert(int64(len(strconv.Itoa(int(io.Output[0])))), 16)

	io = &intcode.TestIO{}
	test3 := []int64{104,1125899906842624,99}
	intcode.RunWithInput(test3, io, false)
	test.Assert(io.Output[0], 1125899906842624)

	p := intcode.ParseProgram("../inputs/day9_1")
	intcode.RunWithInput(p, intcode.Input(1), false)


	intcode.RunWithInput(p, intcode.Input(2), false)
}
