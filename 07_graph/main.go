package main

import (
	"fmt"
	"slices"
)

type (
	constType  map[string]int
	parentType map[string]string
	mgType     map[string]map[string]int
)

const infValue = 999999999

func main() {
	var mg = make(map[string]map[string]int)
	var parents = make(map[string]string)
	var costs = make(map[string]int)

	mg["start"] = map[string]int{}
	mg["start"]["a"] = 1
	mg["start"]["b"] = 2

	mg["b"] = map[string]int{}
	mg["b"]["a"] = 3
	mg["b"]["end"] = 5

	mg["a"] = map[string]int{}
	mg["a"]["end"] = 1
	mg["end"] = map[string]int{}

	parents["a"] = "start"
	parents["b"] = "start"

	costs["a"] = 1
	costs["b"] = 2
	costs["end"] = infValue

	var proccessedNodes = []string{}
	fmt.Println(mg, parents, costs, proccessedNodes)

	//start    find short way to point
	node := findLowestCostNode(costs, proccessedNodes)
	for node != "" {
		neigthbors := mg[node]
		if len(neigthbors) == 0 {
			break
		}
		cost := costs[node]
		fmt.Println("cost:", cost, "neithbors", neigthbors, "node", node)
		for n, _ := range neigthbors {
			newCost := cost + mg[node][n]
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
				proccessedNodes = append(proccessedNodes, node)
			}
		}
		node = findLowestCostNode(costs, proccessedNodes)
	}

	fmt.Println("parents map:", parents)
	fmt.Println(getWay(parents))
	fmt.Println("finish")
}
func findLowestCostNode(costs map[string]int, processed []string) string {
	minVal := infValue
	var res = ""
	for k, val := range costs {
		if val < minVal && !slices.Contains(processed, k) {
			minVal = val
			res = k
		}
	}
	return res
}

func getWay(parents map[string]string) string {
	keys := []string{"b", "a", "end"}
	var way string
	var arrow string
	for i, key := range keys {
		if i != 0 {
			arrow = "->"
		}
		way += fmt.Sprintf("%s%s", arrow, parents[key])
	}
	if len(way) == 0 {
		return "way not found"
	}
	return way + "->end"
}
