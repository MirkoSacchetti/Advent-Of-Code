// --- Day 4: Camp Cleanup --- 

package main

import (
    "fmt"
    "strconv"
    "strings"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func sumInts(array[] int) int {
    result := 0
    for _, v := range array {
        result += v
    }
    return result
}

func check_if_overlap(pairA string, pairB string) bool{
    a := strings.Split(pairA, "-")
    b := strings.Split(pairB, "-")
    a0, _ := strconv.Atoi(a[0])
    a1, _ := strconv.Atoi(a[1])
    b0, _ := strconv.Atoi(b[0])
    b1, _ := strconv.Atoi(b[1])
    if a0 <= b0 && a1 <= b0 {
        return true // a overlap b
    }
    if b0 <= a0 && b1 <= a0 {
        return true // b overlap a
    }
    return false
}

func check_if_contained(pairA string, pairB string) bool{
    a := strings.Split(pairA, "-")
    b := strings.Split(pairB, "-")
    a0, _ := strconv.Atoi(a[0])
    a1, _ := strconv.Atoi(a[1])
    b0, _ := strconv.Atoi(b[0])
    b1, _ := strconv.Atoi(b[1])
    if a0 <= b0 && a1 >= b1 {
        return true // a contain b
    }
    if b0 <= a0 && b1 >= a1 {
        return true // b contain a
    }
    return false
}

func main() {
    file, err := os.ReadFile("input.txt")
    check(err)
    rows := strings.Split(string(file), "\n")
    counter_overlaps := 0
    counter_contained := 0
    for _, row := range rows {
        pairs := strings.Split(string(row), ",")
        if len(pairs) == 2 && check_if_contained(pairs[0], pairs[1]) {
            counter_contained ++
        }
        if len(pairs) == 2 && check_if_overlap(pairs[0], pairs[1]) {
            counter_overlaps ++
        }
    }
    fmt.Println("part one:", counter_contained)
    fmt.Println("part one:", counter_overlaps)
}
