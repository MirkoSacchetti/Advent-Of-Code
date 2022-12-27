// --- Day 6: Tuning Trouble ---

package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func unshift(buffer []rune, c rune) []rune {
	markers := make([]rune, len(buffer))
	buffer = append(buffer, c)
	copy(markers, buffer[1:])
	return markers
}

func parseLine(line string, lenMarker int) int {
	markers := []rune{}
	for index, c := range line {
		// setup
		if len(markers) < lenMarker {
			markers = append(markers, c)
		} else {
			markers = unshift(markers, c)
			noDuplicate := true
			counter := make(map[rune]int)
			for _, r := range markers {
				counter[r]++
				if counter[r] > 1 {
					noDuplicate = false
					break
				}
			}
			if noDuplicate {
				return index + 1
			}
		}

	}
	return 0
}

func main() {
	file, err := os.ReadFile("input.txt")
	check(err)
	lines := strings.Split(string(file), "\n")
	markerIndex := 0
	startOfMessageMarkerIndex := 0
	for _, line := range lines {
		if len(line) > 0 {
			markerIndex = parseLine(line, 4)
			startOfMessageMarkerIndex = parseLine(line, 14)
		}
	}
	fmt.Println("part one:", markerIndex)
	fmt.Println("part two:", startOfMessageMarkerIndex)
}
