package day5

import (
	"fmt"
	"github.com/samber/lo"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Run(input string) {
	lines := strings.Split(input, "\r\n")
	crateRegexp := regexp.MustCompile("\\[(.*?)\\]")
	var rows [][]string
	var cols [][]string
	for i, s := range lines {
		if i < 10 {
			if i >= 8 {
				cols = lo.Map[string, []string](rows[0], func(_ string, i int) []string {
					sl := lo.Map[[]string, string](rows, func(row []string, _ int) string {
						if row[i] == "[0]" {
							return ""
						}
						return strings.Replace(strings.Replace(row[i], "[", "", 1), "]", "", 1)
					})

					sort.SliceStable(sl, func(i, j int) bool {
						return i > j
					})
					return deleteEmpty(sl)
				})
				continue
			}
			for j, char := range s {
				if string(char) == " " {
					if s[j+1:j+4] == "   " {
						s = s[:j+1] + "[0]" + s[j+4:]
					}
				}
			}
			if i == 0 {
				// hacky solution since goland keeps removing the whitespaces
				s = s + "[0][0][0]"
			}
			rows = append(rows, crateRegexp.FindAllString(strings.ReplaceAll(s, " ", ""), -1))
			continue
		}
		// move actions
		parseMoves := regexp.MustCompile("[0-9]+").FindAllString(s, -1)
		move := lo.Map[string, int](parseMoves, func(s string, _ int) int {
			m, _ := strconv.Atoi(s)
			return m
		})
		times, from, to := move[0], move[1]-1, move[2]-1
		for j := 0; j < times; j++ {
			fmt.Println(i, from, to, times)
			cols[to] = append(cols[to], cols[from][len(cols[from])-1])
			cols[from] = cols[from][:len(cols[from])-1]
		}
	}
	lo.ForEach(cols, func(col []string, i int) {
		fmt.Println(col[0])
	})
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

//package day5
//
//import (
//	"fmt"
//	"github.com/samber/lo"
//	"golang.org/x/exp/maps"
//	"golang.org/x/exp/slices"
//	"regexp"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//func Run(input string) {
//	letterRegex := regexp.MustCompile(`\[([A-Z])\]`)
//	moveRegex := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
//
//	sections := strings.Split(strings.ReplaceAll(string(input), "\r", ""), "\n\n")
//	graphSection, moveSection := strings.Split(sections[0], "\n"), strings.Split(sections[1], "\n")
//
//	initialGraph := make(map[int][]string)
//	for _, line := range graphSection[:len(graphSection)-1] {
//		for _, entry := range strings.Fields(line) {
//			letter := letterRegex.FindStringSubmatch(entry)[1]
//
//			ind := strings.Index(line, letter)
//			line = strings.Replace(line, letter, " ", 1)
//
//			column := (ind + 3) / 4
//			if _, ok := initialGraph[column]; !ok {
//				initialGraph[column] = make([]string, 0)
//			}
//			initialGraph[column] = append(initialGraph[column], letter)
//		}
//	}
//
//	partOneGraph := maps.Clone(initialGraph)
//	partTwoGraph := maps.Clone(initialGraph)
//	for _, rawMove := range moveSection {
//		move := lo.Map(moveRegex.FindStringSubmatch(rawMove)[1:], func(s string, _ int) int {
//			return lo.Must(strconv.Atoi(s))
//		})
//
//		amount, from, to := move[0], move[1], move[2]
//		for i := 0; i < amount; i++ {
//			crate := partOneGraph[from][0]
//			partOneGraph[from] = partOneGraph[from][1:]
//			partOneGraph[to] = append([]string{crate}, partOneGraph[to]...)
//		}
//
//		crates := slices.Clone(partTwoGraph[from][:amount])
//		partTwoGraph[from] = partTwoGraph[from][amount:]
//		partTwoGraph[to] = append(crates, partTwoGraph[to]...)
//	}
//
//	indexes := lo.Keys(initialGraph)
//	sort.Ints(indexes)
//
//	partOne, partTwo := &strings.Builder{}, &strings.Builder{}
//	for _, index := range indexes {
//		partOne.WriteString(partOneGraph[index][0])
//		partTwo.WriteString(partTwoGraph[index][0])
//	}
//
//	fmt.Printf("Part One: %s\n", partOne.String())
//	fmt.Printf("Part Two: %s\n", partTwo.String())
//}
