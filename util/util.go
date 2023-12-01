package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile(path string) []string {
	curPath, _ := os.Getwd()
	fmt.Println(curPath)
	file, err := os.Open(path)
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
