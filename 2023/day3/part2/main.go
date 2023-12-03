package main

import (
	"advent-of-code/util"
	"fmt"
	"strings"
)

type Number struct {
	Start int
	End   int
	Row   int
	Final int
}

func main() {
	fileLines := util.ReadFile("2023/day3/input.txt")
	matrix := readFileIntoMatrix(fileLines)

	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "*" {
				if num, err := checkForAdjacentNumbers(matrix, i, j); err == nil {
					sum += num
				}
			}
		}
	}
	fmt.Println(sum)
}

func getNumIndices(row []string, idx int) (int, int) {
	start, end := idx, idx

	// Find the start of the number
	for start > 0 && util.IsDigit(row[start-1]) {
		start--
	}

	// Find the end of the number
	for end < len(row) && util.IsDigit(row[end]) {
		end++
	}

	return start, end
}

func checkForAdjacentNumbers(matrix [][]string, i int, j int) (int, error) {
	directions := []struct{ di, dj int }{
		{-1, -1}, {-1, 0}, {-1, 1}, // Diagonals and above
		{0, -1}, {0, 1}, // Left and right
		{1, -1}, {1, 0}, {1, 1}, // Diagonals and below
	}

	numbers := make([]Number, 0)

	for _, dir := range directions {
		newI, newJ := i+dir.di, j+dir.dj
		// Check if the new position is within the matrix and is a digit
		if newI >= 0 && newI < len(matrix) && newJ >= 0 && newJ < len(matrix[i]) && util.IsDigit(matrix[newI][newJ]) {
			start, end := getNumIndices(matrix[newI], newJ)
			add := true
			// Check if the number is already in the list
			for _, num := range numbers {
				if num.Start == start && num.End == end && num.Row == newI {
					add = false
					break
				}
			}
			if add {
				numbers = append(numbers, Number{start, end, newI, util.MustAtoi(strings.Join(matrix[newI][start:end], ""))})
			}
		}
	}

	if len(numbers) == 2 {
		return numbers[0].Final * numbers[1].Final, nil
	}
	return 0, fmt.Errorf("no adjacent numbers found")
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
