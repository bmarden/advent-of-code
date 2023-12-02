package main

import (
	"advent-of-code/util"
	"fmt"
	"os"
	"strconv"
)

var numMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	fileLines := util.ReadFile("2023/day1/part2/input.txt")

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
		// We got both the front and back numbers, so we can break
		if frontStrNum != "" && backStrNum != "" {
			break
		}

		if frontStrNum == "" {
			frontStringVal := string(s[frontPtr])

			if val, ok := util.CheckStrIsNumeric(frontStringVal); ok == nil {
				frontStrNum = val
			} else if val, ok := checkForStringNum(s[frontPtr:], false); ok == nil {
				frontStrNum = val
			} else {
				frontPtr++
			}
		}

		if backStrNum == "" {
			backStringVal := string(s[backPtr])

			if val, ok := util.CheckStrIsNumeric(backStringVal); ok == nil {
				backStrNum = val
			} else if val, ok := checkForStringNum(s[:backPtr+1], true); ok == nil {

				backStrNum = val
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

func checkForStringNum(s string, reverse bool) (string, error) {
	// All numbers from numMap are 3-5 characters long
	numberLengths := []int{3, 4, 5}

	for _, length := range numberLengths {
		if len(s) < length {
			continue
		}

		potentialMatch := ""
		// If reverse, splice from the end of the string
		if reverse {
			potentialMatch = s[len(s)-length:]
		} else {
			potentialMatch = s[0:length]
		}

		if val, ok := numMap[potentialMatch]; ok {
			return val, nil
		}
	}

	return "", fmt.Errorf("could not convert %s to int", s)
}
