package main

import (
	"github.com/pebrc/aoc-2019/go/day5/intcode"
	"github.com/pebrc/aoc-2019/go/test"
)



func main() {

	baseProgram := intcode.ParseProgram("../inputs/day5_1")
	// part 1
	intcode.RunWithInput(baseProgram, intcode.Input(1), false)

	// cmpTest
	io := &intcode.TestIO{Input: 8}
	cmpTest := []int64{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	intcode.RunWithInput(cmpTest, io , false) //
	test.Assert(1, io.Output[0])
	io = &intcode.TestIO{Input: 7}
	intcode.RunWithInput(cmpTest, io, false)
	test.Assert(0, io.Output[0])

	// jump test
	io = &intcode.TestIO{Input:42}
	intcode.RunWithInput([]int64{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9}, io, false)
	test.Assert(1, io.Output[0])

	//part 2
	intcode.RunWithInput(baseProgram, intcode.Input(5), false)
}
