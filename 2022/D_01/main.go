// --- Day 1: Calorie Counting ---

package main

import (
    "fmt"
    "strconv"
    "strings"
    "os"
    "sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func sumStrings(array[] string) int {
    fmt.
    result := 0
    for _, v := range array {
        val, _ := strconv.Atoi(v)
        result += val
    }
    return result
}

func sumInts(array[] int) int {
    result := 0
    for _, v := range array {
        result += v
    }
    return result
}

func main() {
    file, err := os.ReadFile("input.txt")
    check(err)
    elfs := strings.Split(string(file), "\n\n")
    maxCalories := 0
    elfsCalories := []int{}
    for _, elf := range elfs {
        items := strings.Split(string(elf), "\n")
        calories := sumStrings(items) 
        elfsCalories = append(elfsCalories, calories)
        if calories >= maxCalories {
            maxCalories = calories
        }
    }
    fmt.Println("part one:", maxCalories)
    sort.Sort(sort.Reverse(sort.IntSlice(elfsCalories)))
    fmt.Println("part two:", (sumInts(elfsCalories[:3])))
}
