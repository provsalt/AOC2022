package main

import (
	"aoc2022/day5"
	"os"
)

func main() {
	data, err := os.ReadFile("day5/day5_input.txt")
	if err != nil {
		panic(err)
	}
	day5.Run(string(data))
}
