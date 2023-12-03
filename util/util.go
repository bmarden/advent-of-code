package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(path string) []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Open(homeDir + "/code/advent-of-code/" + path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func CheckStrIsNumeric(s string) (string, error) {
	if i, err := strconv.Atoi(s); err == nil {
		return fmt.Sprint(i), nil
	}

	return "", fmt.Errorf("could not convert %s to int", s)
}

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func IsDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
