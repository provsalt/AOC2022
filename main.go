package main

import (
	"aoc2022/Day3"
	"os"
)

func main() {
	data, err := os.ReadFile("Day3/day3_input.txt")
	if err != nil {
		panic(err)
	}
	Day3.Run(string(data))
}
