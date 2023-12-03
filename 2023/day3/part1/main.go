package main

import (
	"advent-of-code/util"
	"fmt"
	"regexp"
	"strings"
)

var symbol = regexp.MustCompile(`[^\d\.\s]`)

func main() {
	fileLines := util.ReadFile("2023/day3/input.txt")
	matrix := readFileIntoMatrix(fileLines)

	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			if util.IsDigit(matrix[i][j]) {
				// Get the full number to check
				fullNumber := getFullNumber(matrix[i], j)
				if num, err := checkNumberSurroundings(matrix, i, j, fullNumber); err == nil {
					sum += num
				}
				j = j + len(fullNumber) - 1
			}
		}
	}
	fmt.Println(sum)
}

func getFullNumber(row []string, idx int) string {
	start, end := idx, idx

	// Find the start of the number
	for start > 0 && util.IsDigit(row[start-1]) {
		start--
	}

	// Find the end of the number
	for end < len(row) && util.IsDigit(row[end]) {
		end++
	}

	// Combine the digits to form the full number
	return strings.Join(row[start:end], "")
}

func checkNumberSurroundings(matrix [][]string, i int, j int, fullNumber string) (int, error) {
	directions := []struct{ i, j int }{
		{-1, -1}, {-1, 0}, {-1, 1}, // diagonals and above
		{0, -1}, {0, 1}, // left and right
		{1, -1}, {1, 0}, {1, 1}, // diagonals and below
	}

	for j != len(matrix[i])-1 && util.IsDigit(matrix[i][j]) {
		for _, dir := range directions {
			newI, newJ := i+dir.i, j+dir.j
			if newI >= 0 && newI < len(matrix) && newJ >= 0 && newJ < len(matrix[i]) {
				if symbol.Match([]byte(matrix[newI][newJ])) {
					return util.MustAtoi(fullNumber), nil
				}
			}
		}
		j++
	}

	return 0, nil
}

func readFileIntoMatrix(lines []string) [][]string {
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = make([]string, len(line))
		for j, char := range line {
			matrix[i][j] = string(char)
		}
	}
	return matrix
}
