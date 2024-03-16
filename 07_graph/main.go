package main

import (
	"fmt"
	"math"
)

func main() {
	var mg = make(map[string]map[string]int)
	var parents = make(map[string]string)
	var costs = make(map[string]int)

	mg["start"] = map[string]int{}
	mg["start"]["a"] = 6
	mg["start"]["b"] = 1

	mg["b"] = map[string]int{}
	mg["b"]["a"] = 3
	mg["b"]["end"] = 5

	mg["a"] = map[string]int{}
	mg["a"]["end"] = 1
	mg["end"] = map[string]int{}

	parents["a"] = "start"
	parents["b"] = "start"

	costs["a"] = 6
	costs["b"] = 2
	costs["end"] = int(math.Inf(1))

	fmt.Println(mg, parents, costs)

}
