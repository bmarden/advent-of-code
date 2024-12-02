package main

import (
	"advent-of-code/util"
	"fmt"
	"strings"
)

// func standardizeSpaces(s string) string {
// 	fields := strings.Fields(s)
// 	return strings.Join(fields, " ")
// }

func main() {
	fileLines := util.ReadFile("2023/day4/input.txt")

	sum := 0
	for _, line := range fileLines {
		numbers := strings.Split(line, ":")[1]
		scores := strings.Split(numbers, "|")
		winningNumbers := strings.Fields(scores[0])
		ourNumbers := strings.Fields(scores[1])

		scoreMap := make(map[string]int)
		for _, num := range winningNumbers {
			scoreMap[num] = 0
		}

		score := 0
		for _, num := range ourNumbers {
			if _, ok := scoreMap[num]; ok {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		sum += score
	}

	fmt.Println(sum)
}
