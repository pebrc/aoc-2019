package test

import (
	"fmt"
	"log"
)

func Assert(actual, expected int64) {
	if actual != expected {
		log.Fatal(fmt.Errorf("expected %d but was %d", expected, actual))
	}
}

func AssertN(actual, expected []int64) bool {
	if (len(actual) != len(expected)) {
		return false
	}
	for i :=0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}
