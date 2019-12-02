package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("../inputs/day1_1")
	if err != nil {
		log.Fatal(err)
	}
	masses := strings.Split(string(bytes), "\n")
	var total float64
	for _, v := range masses {
		if len(v) == 0 {
			continue
		}
		intVal, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		f := math.Floor(float64(intVal/3)) - 2
		total += f
	}
	fmt.Println(int64(total))
	fmt.Println("=============part 2 ==============")
	// println(fuelFuel(1969))
	// println(fuelFuel(100756))
	var total2 int64
	for _, v := range masses {
		if len(v) == 0 {
			continue
		}
		intVal, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		total2 += fuelFuel(intVal)
	}
	fmt.Println(total2)

}

func fuelFuel(module int64) int64 {
	baseVal := module
	var subTotals float64
FUELFUEL:
	fuelBase := math.Floor(float64(baseVal / 3))
	if fuelBase > 0 {
		f := fuelBase - 2
		if f < 0 {
			f = 0
		}
		subTotals += f
		baseVal = int64(f)
		goto FUELFUEL
	}
	return int64(subTotals)
}
