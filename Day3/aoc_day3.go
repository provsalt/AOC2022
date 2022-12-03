package Day3

import (
	"fmt"
	"strings"
	"unicode"
)

func Run(input string) {
	allLines := strings.Split(input, "\n")
	var rucksack int
	for _, line := range allLines {
		length := len(line) / 2
		first := strings.Clone(line[0:length])
		second := strings.Clone(line[length:])
		// check if first and second have the same character
		for _, r := range first {
			if strings.ContainsRune(second, r) {
				rucksack += int(r - 38)
				if unicode.IsLower(r) {
					rucksack -= 58 // -96
				}
				break
			}
		}
	}
	fmt.Printf("Total rucksack %d\n", rucksack)
	// part two checking two other lines for the same character
	var badges int
	for i := 0; i < len(allLines); i += 3 {
		for _, r := range allLines[i] {
			if strings.ContainsRune(allLines[i+2], r) && strings.ContainsRune(allLines[i+1], r) {
				badges += int(r - 38)
				if unicode.IsLower(r) {
					badges -= 58
				}
				break
			}
		}
	}
	fmt.Printf("Badges %d\n", badges)
}
