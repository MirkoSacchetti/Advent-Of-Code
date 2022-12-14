// --- Day 5: Supply Stacks ---

//                 [B]     [L]     [S]
//         [Q] [J] [C]     [W]     [F]
//     [F] [T] [B] [D]     [P]     [P]
//     [S] [J] [Z] [T]     [B] [C] [H]
//     [L] [H] [H] [Z] [G] [Z] [G] [R]
// [R] [H] [D] [R] [F] [C] [V] [Q] [T]
// [C] [J] [M] [G] [P] [H] [N] [J] [D]
// [H] [B] [R] [S] [R] [T] [S] [R] [L]
//  1   2   3   4   5   6   7   8   9

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

func parseLine(line string) []int {
	var numbers []int
	parts := strings.Split(line, " ")
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		if num != 0 {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func moveCrates(stacks [][]rune, crates int, from int, to int) [][]rune {
	for i := 0; i < crates; i++ {
		s := stacks[from]
		stacks[to] = append(stacks[to], s[len(s)-1])
		stacks[from] = s[:len(s)-1]
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

func main() {
	stacks := [][]rune{
		[]rune("HCR"),
		[]rune("BJHLSF"),
		[]rune("RMDHJTQ"),
		[]rune("SGRHZBJ"),
		[]rune("RPFZTDCB"),
		[]rune("THCG"),
		[]rune("SNVZBPWL"),
		[]rune("RJQGC"),
		[]rune("LDTRHPFS"),
	}
	stacks9001 := [][]rune{
		[]rune("HCR"),
		[]rune("BJHLSF"),
		[]rune("RMDHJTQ"),
		[]rune("SGRHZBJ"),
		[]rune("RPFZTDCB"),
		[]rune("THCG"),
		[]rune("SNVZBPWL"),
		[]rune("RJQGC"),
		[]rune("LDTRHPFS"),
	}
	file, err := os.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if len(line) > 0 {
			vals := parseLine(line)
			// off by one because we are programmer
			stacks = moveCrates(stacks, vals[0], vals[1]-1, vals[2]-1)
			stacks9001 = moveCrates9001(stacks9001, vals[0], vals[1]-1, vals[2]-1)
		}
	}
	fmt.Println("part one:", getToppest(stacks))
	fmt.Println("part two:", getToppest(stacks9001))
}
