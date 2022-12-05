package main

import (
	"aoc2022/Day4"
	"os"
)

func main() {
	data, err := os.ReadFile("Day4/day4_input.txt")
	if err != nil {
		panic(err)
	}
	Day4.Run(string(data))
}
