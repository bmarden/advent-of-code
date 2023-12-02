package main

import (
	"advent-of-code/util"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fileLines := util.ReadFile("2023/day2/input.txt")
	rx := regexp.MustCompile(`([0-9]+)\s([a-z]+)`)

	gameProductSum := 0
	for _, line := range fileLines {
		parts := strings.Split(line, ":")
		rounds := strings.Split(parts[1], ";")

		cubeCountMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, round := range rounds {

			roundParts := strings.Split(round, ",")
			for _, roundPart := range roundParts {
				matches := rx.FindAllStringSubmatch(roundPart, -1)
				for _, match := range matches {
					curCubeCount := cubeCountMap[match[2]]
					cubeCountMap[match[2]] = max(util.MustAtoi(match[1]), curCubeCount)
				}
			}
		}
		gameProduct := cubeCountMap["red"] * cubeCountMap["green"] * cubeCountMap["blue"]
		gameProductSum += gameProduct
	}

	fmt.Println(gameProductSum)
}
