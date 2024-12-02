package main

import (
	"advent-of-code/util"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	input := util.ReadFile("2024/day1/input.txt")

	allListTotal := 0

	list1 := []int{}
	list2 := []int{}

	for _, i := range input {
		num1, num2 := getLocationIds(i)
		list1 = append(list1, util.MustAtoi(num1))
		list2 = append(list2, util.MustAtoi(num2))
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		allListTotal += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println(allListTotal)
}

func getLocationIds(lines string) (string, string) {
	lists := strings.Fields(lines)
	list1 := lists[0]
	list2 := lists[1]

	return list1, list2
}
