// --- Day 2: Rock Paper Scissors ---

// val
// A/X 1p -> Rock / lose ( -1 )
// B/Y 2p -> Paper / draw  ( 0 )
// C/Z 3p -> Scissors  / win ( + 1 )

// score
// 0p lost
// 3p draw
// 6p win

// score = val + res
package main

import (
    "fmt"
    "strings"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func get_cheating(you int, result int) string{
    items := map[int]string{
        0: "X",
        1: "Y",
        2: "Z",
    }
    resp := 0
    if result < 2  { // lose go forward
        resp = (you - 1 + 2) % 3
    }
    if result > 2  { // win go previus
        resp = (you - 1 + 1) % 3
    }
    if result == 2  { // draw same item
        resp = you - 1
    }
    return items[resp]
}

func get_score(you int, me int) int{
    if (you == me) {
        return 3
    }
    if (me + 2 ) % 3 == (you + 1) % 3  {
        return 0
    }
    return 6
}

var values = map[string]int {
    "A": 1,
    "B": 2,
    "C": 3,
    "X": 1,
    "Y": 2,
    "Z": 3,
}

func main() {
    file, err := os.ReadFile("input.txt")
    check(err)
    games := strings.Split(string(file), "\n")
    tot_score := 0
    tot_score_cheat := 0
    for _, game := range games {
        items := strings.Split(string(game), " ")
        if len(items) == 2 {
            pointYou := values[items[0]]
            pointMe := values[items[1]]
            score := get_score(pointYou, pointMe)
            tot_score += pointMe + score

            item_cheat := get_cheating(pointYou, pointMe)
            point_cheat := values[item_cheat]
            score_cheat := get_score(pointYou, point_cheat)
            tot_score_cheat += point_cheat + score_cheat

            fmt.Println(items, item_cheat, score_cheat )
        }
    }
    fmt.Println("part one:", tot_score)
    fmt.Println("part two:", tot_score_cheat)
}
