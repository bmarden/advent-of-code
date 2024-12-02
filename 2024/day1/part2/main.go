package main

import (
	"advent-of-code/util"
	"fmt"
	"strings"
)

func main() {
	input := util.ReadFile("2024/day1/input.txt")

	list1 := []int{}
	list2 := []int{}

	for _, i := range input {
		num1, num2 := getLocationIds(i)
		list1 = append(list1, util.MustAtoi(num1))
		list2 = append(list2, util.MustAtoi(num2))
	}

	list2Map := make(map[int]int)

	for i := 0; i < len(list2); i++ {
		list2Map[list2[i]]++
	}

	totalSimilarity := 0
	for _, num := range list1 {
		if val, ok := list2Map[num]; ok {
			simScore := val * num
			totalSimilarity += simScore
		}
	}

	fmt.Println(totalSimilarity)
}

func getLocationIds(lines string) (string, string) {
	lists := strings.Fields(lines)
	list1 := lists[0]
	list2 := lists[1]

	return list1, list2
}
