package main

import (
	"aoc2022/Day2"
	"os"
)

func main() {
	data, err := os.ReadFile("Day2/day2_input.txt")
	if err != nil {
		panic(err)
	}
	Day2.Run(string(data))
}
