// --- Day 8: Treetop Tree House ---
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkVisibility(trees []int) []int {
	visibles := []int{trees[0]} // the first is always visible
	for i := 1; i < len(trees)-1; i++ {
		if trees[i] > visibles[len(visibles)-1] {
			visibles = append(visibles, trees[i])
		}
	}
	return visibles
}

func reverseInts(array []int) []int {
	result := []int{}
	for i := len(array) - 1; i >= 0; i-- {
		result = append(result, array[i])
	}
	return result
}

func sumInts(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func countVisibleTrees(treeGrid [][]int) int {
	visibleTrees := 0
	for y := 0; y < len(treeGrid); y++ {
		for x := 0; x < len(treeGrid[y]); x++ {
			if x != 0 && y != 0 {
				visibleTrees++
			}

		}
	}
	return visibleTrees
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	treeGrid := [][]int{}
	for row, line := range lines {
		if len(line) > 0 {
			treeGrid = append(treeGrid, []int{})
			for _, c := range line {
				tree, _ := strconv.Atoi(string(c))
				treeGrid[row] = append(treeGrid[row], tree)
			}
		}
	}
	fmt.Println("part one:", countVisibleTrees(treeGrid))
}
