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
