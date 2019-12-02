package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func intcode(program []int64) int64 {
	var ip int
	var exit bool
	for !exit {
		switch program[ip] {
		case 1:
			program[program[ip+3]] = program[program[ip+1]] + program[program[ip+2]]
		case 2:
			program[program[ip+3]] = program[program[ip+1]] * program[program[ip+2]]
		case 99:
			exit = true
		}
		ip = ip + 4
	}
	return program[0]
}

func intcodeWithInput(program []int64, noun, verb int64) int64 {
	executable := make([]int64, len(program))
	copy(executable, program)
	executable[1] = noun // noun
	executable[2] = verb  // verb
	return intcode(executable)
}

func assert(actual, expected int64) {
	if actual  != expected {
		log.Fatal(fmt.Errorf("expected %d but was %d", expected, actual))
	}
}

func main() {
	testProgram1 := []int64{
		1, 0, 0, 0, 99,
	}
	assert(intcode(testProgram1), 2)
	testProgram2 := []int64 {
		1,1,1,4,99,5,6,0,99,
	}
	assert(intcode(testProgram2), 30)
	bytes, err := ioutil.ReadFile("../inputs/day2_1")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(bytes), ",")
	baseProgram := make([]int64, len(input))
	for i, s := range input {
		intVal, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		baseProgram[i] = intVal
	}

	fmt.Println(intcodeWithInput(baseProgram, 12, 2))

	// part 2
	// inputs are 0-99
	var noun, verb int64
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			if intcodeWithInput(baseProgram, noun, verb) == 19690720 {
				fmt.Println(100*noun + verb)
				os.Exit(0)
			}
		}
	}
}
