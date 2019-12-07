package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

var pairs [][]string
var paths [][]string

func contains(xs []string, s string) bool {
	for _, x := range xs {
		if x == s {
			return true
		}
	}
	return false
}

func remove(xs []string, s string) []string {
	var result []string
	for _, x := range xs {
		if x != s {
			result = append(result, x)
		}
	}
	return result
}

func findSegments(s string) [][]string {
	var results [][]string
	for _, x := range pairs {
		if contains(x, s) {
			results = append(results, x)
		}
	}
	return results
}

func followSinglePath(path []string, s string) {
	if contains(path, s) {
		// exit
		if contains(path, "SAN") {
			paths = append(paths, path)
			println(fmt.Sprintf("%d paths: added %v", len(paths), path))
		}
	} else {
		// linear recursion
		findPaths(path, s)
	}
}

func findPaths(path []string, s string) {
	// find all segments
	path = append(path, s)
	segments := findSegments(s)
	for _, segment := range segments {
		subpath := make([]string, len(path))
		copy(subpath, path)
		newKey := remove(segment, s)[0]
		followSinglePath(subpath, newKey)
	}
}

func main() {
	bytes, err := ioutil.ReadFile("../inputs/day6_1")
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range strings.Split(string(bytes), "\n") {
		pairs = append(pairs, strings.Split(s, ")"))
	}

	findPaths(nil, "YOU")
	shortest := math.MaxInt32
	for _, p := range paths {
		if len(p) < shortest {
			shortest = len(p) -3 // removing SAN and YOU and counting only the 'transfers'
		}
	}
	println(shortest)
}
