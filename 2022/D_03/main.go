//--- Day 3: Rucksack Reorganization ---

// 1 rucksack
// 2 compartments with equal number of items
// 1 item common 
// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.

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

func sumStrings(array[] string) int {
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

func find_badge(a string, b string, c string) rune {
    var common rune 
    for _,char := range a {
        if strings.ContainsRune(b, char) {
            if strings.ContainsRune(c, char) {
                common = char
                break
            }
        }
    }
    return common
}

func find_common(a string, b string) rune {
    var char rune 
    for _,c := range a {
        if strings.ContainsRune(b, c) {
            char = c
            break
        }
    }
    return char
}

func get_priority(c rune) int {
    if int(c) >= 97 {
        // lowercase start from ASCII code 97 decoded as 'a'
        // add 1 because we are programmers
        return int(c) - 97 + 1
    } else {
        // uppercase start from ASCII code 65 decoded as 'A'
        // the order of the fame score is inverted
        // from ASCII codes, first a-z and A-Z so add 27
        return int(c) - 65 + 27
    }
}


func main() {
    file, err := os.ReadFile("input.txt")
    check(err)
    rucksacks := strings.Split(string(file), "\n")
    scores := []int{}
    group := []string{}
    group_score := []int{}
    for _, r := range rucksacks {
        n := len(r)/2
        if n >1 {
            a := r[:n]
            b := r[n:]
            diffC := find_common(a,b)
            score :=  get_priority(diffC)
            scores = append(scores, score)
        }

        group = append(group, r)
        if len(group) == 3 {
            common_char := find_badge(group[0], group[1], group[2])
            score := get_priority(common_char)
            group_score = append(group_score, score)
            group = []string{}
        } 
    }
    fmt.Println("part one:", sumInts(scores))
    fmt.Println("part two:", sumInts(group_score))
}
