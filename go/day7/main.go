package main

import (
	"fmt"
	"sync"

	"github.com/pebrc/aoc-2019/go/day5/intcode"
	"github.com/pebrc/aoc-2019/go/test"
)

var phaseSettingsPart1 = permutations([]int{1, 2, 3, 4, 0}, 5)
var phaseSettingsPart2 = permutations([]int{5, 6, 7, 8, 9}, 5)

type amplifierIO struct {
	out          chan int64
	in           chan int64
	phaseSetting int
	inputPos     int64
}

func (a *amplifierIO) ReadInt() int64 {
	var val int64
	if a.inputPos == 0 {
		val = int64(a.phaseSetting)
	}
	val = <-a.in
	a.inputPos++
	return val
}

func (a *amplifierIO) WriteInt(o int64) {
	a.out <- o
}

var _ intcode.IO = &amplifierIO{}

func runAmplifierProgram(p intcode.Program, phaseSetting []int) int64 {
	chans := []chan int64{make(chan int64, 1), make(chan int64, 0), make(chan int64, 0), make(chan int64, 0), make(chan int64, 0)}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		idx := i
		go func() {
			io := &amplifierIO{
				out:          chans[(idx+1)%5],
				in:           chans[idx],
				phaseSetting: phaseSetting[idx],
			}
			intcode.RunWithInput(p, io, false)
			wg.Done()
		}()
	}
	// start with 0
	chans[0] <- 0
	wg.Wait()
	return <-chans[0]
}

func findPhaseSetting(p intcode.Program, phaseSettings [][]int) (int64, []int) {
	var maxOutput int64
	var bestSetting []int
	for _, phaseSetting := range phaseSettings {
		out := runAmplifierProgram(p, phaseSetting)
		if out > maxOutput {
			maxOutput = out
			bestSetting = phaseSetting
		}
	}
	return maxOutput, bestSetting
}

func permutations(xs []int, n int) [][]int {
	var res [][]int
	if n == 1 {
		cp := make([]int, len(xs))
		copy(cp, xs)
		res = append(res, cp)
	}

	for i := 0; i < n; i++ {
		for _, p := range permutations(xs, n-1) {
			res = append(res, p)
		}
		if n%2 == 1 {
			xs[0], xs[n-1] = xs[n-1], xs[0]
		} else {
			xs[i], xs[n-1] = xs[n-1], xs[i]
		}
	}
	return res
}

func main() {
	test1 := []int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	output, setting := findPhaseSetting(test1, phaseSettingsPart1)
	test.Assert(output, 43210)
	// part1
	program := intcode.ParseProgram("../inputs/day7_1")
	output, setting = findPhaseSetting(program, phaseSettingsPart1)
	fmt.Printf("%d, %v\n", output, setting)
	// part2
	output, setting = findPhaseSetting(program, phaseSettingsPart2)
	fmt.Printf("%d, %v\n", output, setting)

}
