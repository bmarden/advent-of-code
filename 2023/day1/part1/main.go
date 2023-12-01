package main

import (
	"fmt"
	"os"
	"strconv"

	"advent-of-code/util"
)

func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
	fileLines := util.ReadFile("2023/day1/part1/input.txt")
	sum := 0

	for _, line := range fileLines {

		lineDigits, err := getDigits(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum += lineDigits

	}

	fmt.Println(sum)
}

func getDigits(s string) (int, error) {
	frontStrNum, backStrNum := "", ""

	backPtr := len(s) - 1
	frontPtr := 0
	for frontPtr <= backPtr {
		if frontStrNum != "" && backStrNum != "" {
			break
		}

		if frontStrNum == "" {
			frontNum, err := util.CheckStrIsNumeric(string(s[frontPtr]))
			if err == nil {
				frontStrNum = frontNum
			} else {
				frontPtr++
			}
		}

		if backStrNum == "" {
			backNum, err := util.CheckStrIsNumeric(string(s[backPtr]))
			if err == nil {
				backStrNum = backNum
			} else {
				backPtr--
			}
		}
	}

	finalDigits, err := strconv.Atoi(frontStrNum + backStrNum)
	if err != nil {
		return 0, err
	}
	return finalDigits, nil
}
