package main

import (
	"advent-of-code/util"
	"fmt"
	"regexp"
	"strings"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
)

func main() {
	fileLines := util.ReadFile("2023/day2/input.txt")
	rx := regexp.MustCompile(`([0-9]+)\s([a-z]+)`)
	gameIdSum := 0

	for _, line := range fileLines {
		parts := strings.Split(line, ":")
		gameNumber := strings.Split(parts[0], " ")[1]
		rounds := strings.Split(parts[1], ";")

		gamePossible := true
		for _, round := range rounds {
			cubeCountMap := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}

			roundParts := strings.Split(round, ",")
			for _, roundPart := range roundParts {
				matches := rx.FindAllStringSubmatch(roundPart, -1)
				for _, match := range matches {
					cubeCountMap[match[2]] = util.MustAtoi(match[1])
				}
			}

			if cubeCountMap["red"] > maxRed || cubeCountMap["green"] > maxGreen || cubeCountMap["blue"] > maxBlue {
				gamePossible = false
				break
			}
		}

		if gamePossible {
			gameNumberInt := util.MustAtoi(gameNumber)
			gameIdSum += gameNumberInt
		}
	}

	fmt.Println(gameIdSum)
}
