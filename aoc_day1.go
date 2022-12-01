package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day1() {
	data, err := os.ReadFile("day1_input.txt")
	if err != nil {
		panic(err)
	}
	data2 := strings.Split(string(data), "\n")

	var groups [][]int
	var group []int
	for _, line := range data2 {
		if line == "" {
			groups = append(groups, group)
			group = []int{}
		} else {
			num, _ := strconv.Atoi(line)
			group = append(group, num)
		}
	}
	var total []int

	for _, grp := range groups {
		var sum int
		for _, num := range grp {
			sum += num
		}
		total = append(total, sum)
	}
	first := highestNum(total)
	fmt.Printf("The highest calories is %d \n", first)
	// part two find three highest numbers in the total array
	second := highestNum(total)
	third := highestNum(total)
	fmt.Printf("The Elves are carrying %d calories", first+second+third)
}

func highestNum(total []int) int {
	var highest [2]int
	for i, num := range total {
		if num > highest[1] {
			highest = [2]int{i, num}
		}
	}
	total[highest[0]] = 0 // now you're the lowest

	return highest[1]
}
