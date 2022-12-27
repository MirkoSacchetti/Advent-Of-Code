// --- Day 5: Supply Stacks ---

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadStacks(stacks [][]rune, line string) [][]rune {
	// the stacks used started from index 1
	index := 1
	// every step is 4 characters forward + 1 delta
	// new step => next stack
	for i := 1; i < len(line); i = i + 4 {
		// create a new layer in stack
		if len(stacks) <= index {
			stacks = append(stacks, []rune{})
		}
		if line[i] != ' ' {
			// append reversed (FIFO)
			stacks[index] = append([]rune{rune(line[i])}, stacks[index]...)
		}
		index++
	}
	return stacks
}

func parseMove(line string) Move {
	var numbers []int
	parts := strings.Split(line, " ")
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		if num != 0 {
			numbers = append(numbers, num)
		}
	}
	return Move{
		From:   numbers[1],
		To:     numbers[2],
		Crates: numbers[0],
	}
}

func moveCrates(stacks [][]rune, crates int, from int, to int) [][]rune {
	for i := 0; i < crates; i++ {
		s := stacks[from]
		stacks[from] = s[:len(s)-1]
		stacks[to] = append(stacks[to], s[len(s)-1])
	}
	return stacks
}

func moveCrates9001(stacks [][]rune, crates int, from int, to int) [][]rune {
	s := stacks[from]
	stacks[from] = s[:len(s)-crates]
	stacks[to] = append(stacks[to], s[len(s)-crates:]...)
	return stacks
}

func getToppest(stacks [][]rune) string {
	toppest := []rune{}
	for _, s := range stacks {
		if len(s) > 0 {
			toppest = append(toppest, s[len(s)-1])
		} else {
			toppest = append(toppest, ' ')
		}
	}
	return string(toppest)
}

func printRuneSlices(slices [][]rune) {
	for _, slice := range slices {
		for _, rn := range slice {
			fmt.Printf("%c", rn)
		}
		fmt.Println()
	}
}

type Move struct {
	From   int
	To     int
	Crates int
}

func main() {
	file, err := os.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(file), "\n")
	// index 0 is inititalized empty
	stacks := [][]rune{{' '}}
	stacks9001 := [][]rune{{' '}}
	for _, line := range lines {
		if len(line) > 0 {
			if line[:4] == "move" {
				move := parseMove(line)
				stacks = moveCrates(stacks, move.Crates, move.From, move.To)
				stacks9001 = moveCrates9001(stacks9001, move.Crates, move.From, move.To)
			} else {
				stacks = loadStacks(stacks, line)
				stacks9001 = loadStacks(stacks9001, line)
			}
		}
	}
	fmt.Println("part one:", getToppest(stacks))
	fmt.Println("part two:", getToppest(stacks9001))
}
