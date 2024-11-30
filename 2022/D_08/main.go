// --- Day 8: Treetop Tree House ---
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkVisibility(x, y int, trees [][]int) bool {
	isVisible := true
	// check perimeter
	if x == 0 || y == 0 || x == len(trees)-1 || y == len(trees[0])-1 {
		return isVisible
	}
	// check left->right
	for i := 0; i < x; i++ {
		if trees[i][y] >= trees[x][y] {
			isVisible = false
			break
		}
	}
	if isVisible {
		return isVisible
	}
	// check right->left
	isVisible = true
	for i := len(trees) - 1; i > x; i-- {
		if trees[i][y] >= trees[x][y] {
			isVisible = false
			break
		}
	}
	if isVisible {
		return isVisible
	}
	// check top->top
	isVisible = true
	for i := 0; i < y; i++ {
		if trees[x][i] >= trees[x][y] {
			isVisible = false
			break
		}
	}
	if isVisible {
		return isVisible
	}
	// check bottom->top
	isVisible = true
	for i := len(trees[x]) - 1; i > y; i-- {
		if trees[x][i] >= trees[x][y] {
			isVisible = false
			break
		}
	}
	return isVisible
}

func getScenicScore(x, y int, trees [][]int) int {
	// check to right
	rightScore := 0
	if x+1 < len(trees) {
		rightScore = 1
		for i := x + 1; i < len(trees); i++ {
			if trees[i][y] <= trees[x][y] {
				rightScore++
			} else {
				break
			}
		}
	}
	// // check left
	leftScore := 1
	if x-1 > 0 {
		leftScore = 1
		for i := x - 1; i >= 0; i-- {
			if trees[i][y] <= trees[x][y] {
				leftScore++
			} else {
				break
			}
		}
	}
	// // // check top
	topScore := 1
	if y-1 > 0 {
		topScore = 1
		for i := y; i >= 0; i-- {
			if trees[x][i] <= trees[x][y] {
				topScore++
			} else {
				break
			}
		}
	}
	// // // check bottom
	bottomScore := 1
	if y+1 < len(trees[0]) {
		bottomScore = 1
		for i := len(trees[x]) - 1; i > y; i-- {
			if trees[x][i] >= trees[x][y] {
				bottomScore++
			} else {
				break
			}
		}
	}
	return topScore * leftScore * rightScore * bottomScore
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

func maxScenicScore(treeGrid [][]int) int {
	maxScore := 0
	for x := 0; x < len(treeGrid); x++ {
		for y := 0; y < len(treeGrid[x]); y++ {
			score := getScenicScore(x, y, treeGrid)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func countVisibleTrees(treeGrid [][]int) int {
	counter := 0
	for x := 0; x < len(treeGrid); x++ {
		for y := 0; y < len(treeGrid[x]); y++ {
			if checkVisibility(x, y, treeGrid) {
				counter++
			}
		}
	}
	return counter
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	// convert the file in a matrix of integers
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
	fmt.Println("part two:", maxScenicScore(treeGrid))
}
