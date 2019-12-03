package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/pebrc/aoc-2019/go/test"
)

func parseCmd(c string, x, y int) [][2]int {
	parsed, err := strconv.ParseInt(c[1:len(c)], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	offset := int(parsed)
	op := string(c[0])
	var result [][2]int
	switch op {
	case "R":
		to := x + offset
		for x++; x <= to; x++ {
			result = append(result, [2]int{x, y})
		}
	case "U":
		to := y + offset
		for y++; y <= to; y++ {
			result = append(result, [2]int{x, y})
		}
	case "L":
		to := x - offset
		for x--; x >= to; x-- {
			result = append(result, [2]int{x, y})
		}
	case "D":
		to := y - offset
		for y--; y >= to; y-- {
			result = append(result, [2]int{x, y})
		}
	}
	return result
}

func parseInput(file string) [][]string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(bytes), "\n")
	instructions := make([][]string, len(lines))
	for i, l := range lines {
		instructions[i] = strings.Split(l, ",")
	}
	return instructions
}

func findCrossedWires(input [][]string) (int64, int64) {
	grid := make(map[[2]int][2]int)
	var min int64 = math.MaxInt64
	minSteps := math.MaxInt32
	for i, l := range input {
		var x, y, steps int
		for _, cmd := range l {
			otherWire := (i + 1) % len(input)
			for _, p := range parseCmd(cmd, x, y) {
				x, y = p[0], p[1]
				steps++
				dist := int64(math.Abs(float64(x)) + math.Abs(float64(y)))
				wires, cross := grid[p]
				if cross && wires[otherWire] != 0 {
					println(fmt.Sprintf("crossing at x %d, y %d, dist %d, steps %d other %d", x, y, dist, steps, wires[otherWire]))
					if dist < min {
						min = dist
					}
					curSteps := steps + wires[otherWire]
					if curSteps < minSteps {
						minSteps = curSteps
					}
				}

				if !cross {
					wires = [2]int{}
				}
				wires[i] = steps
				grid[p] = wires
			}
		}
	}
	return min, int64(minSteps)
}

func main() {
	minDist, minSteps := findCrossedWires(parseInput("../inputs/day3_test0"))
	test.Assert(minDist, 6)
	test.Assert(minSteps, 30)

	minDist, minSteps = findCrossedWires(parseInput("../inputs/day3_test1"))
	test.Assert(minDist, 159)
	test.Assert(minSteps, 610)

	minDist, minSteps = findCrossedWires(parseInput("../inputs/day3_1"))
	fmt.Printf("minDist %d minSteps %d", minDist, minSteps)
}
