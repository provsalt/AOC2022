package Day4

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(input string) {
	lines := strings.Split(input, "\r\n")
	var contains, overlaps int
	for _, line := range lines {
		p := strings.Split(line, ",")
		sec, sec2 := strings.Split(p[0], "-"), strings.Split(p[1], "-")
		min, max := toInt(sec[0]), toInt(sec[1])
		min2, max2 := toInt(sec2[0]), toInt(sec2[1])
		if min >= min2 && min <= max2 && max <= max2 || min2 >= min && min2 <= max && max2 <= max {
			contains++
		}
		if min <= max2 && min2 <= max {
			overlaps++
		}
	}
	fmt.Printf("Day 4 Part 1: %d", contains)
	fmt.Printf("Day 4 Part 2: %d", overlaps)
}

func toInt(strInt string) int {
	conv, err := strconv.Atoi(strInt)
	if err != nil {
		panic(err)
	}
	return conv
}
