package Day2

import (
	"fmt"
	"strings"
)

var charToChoice = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,

	"X": 1,
	"Y": 2,
	"Z": 3,
}

const (
	lose = 0
	draw = 3
	win  = 6
)

func Run(input string) {
	rounds := strings.Split(input, "\n")
	var totalScore int
	for _, round := range rounds {
		choices := strings.Split(round, " ")
		score := rps(charToChoice[choices[0]], charToChoice[choices[1]])
		totalScore += score
	}
	fmt.Printf("Total score: %d\n", totalScore)
	// part two
	totalScore = 0
	deciperChoice := map[string]int{
		"X": lose,
		"Y": draw,
		"Z": win,
	}
	for _, round := range rounds {
		choices := strings.Split(round, " ")
		score := func() int {
			switch deciperChoice[choices[1]] {
			case lose:
				switch charToChoice[choices[0]] {
				case 1:
					return rps(1, 3)
				case 2:
					return rps(2, 1)
				case 3:
					return rps(3, 2)
				}
			case draw:
				return rps(charToChoice[choices[0]], charToChoice[choices[0]])
			case win:
				switch charToChoice[choices[0]] {
				case 1:
					return rps(1, 2)
				case 2:
					return rps(2, 3)
				case 3:
					return rps(3, 1)
				}
			default:
				panic("invalid choice")
			}
			return 0
		}()
		totalScore += score
	}
	fmt.Printf("Total score: %d", totalScore)
}

func rps(opponent int, player int) int {
	if opponent == player {
		return player + draw // draw
	}

	// r + p = 3
	// r + s = 4
	// p + s = 5
	switch opponent + player {
	case 3:
		if player == 1 {
			return player // we lost
		}
	case 4:
		if player == 3 {
			return player // we lost
		}
	case 5:
		if player == 2 {
			return player // we lost
		}
	}
	return player + win // we won
}
