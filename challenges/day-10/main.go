package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	slice := readInput()
	partOne, partTwo := solve(slice)
	fmt.Println("Part One:", partOne, ", Part Two:", partTwo)
}

var plut = map[rune]int{
	'(': 3,
	'[': 57,
	'{': 1197,
	'<': 25137,
}

var alut = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func solve(in []string) (int, int) {
	pscore := 0
	var ascores []int
Lines:
	for _, l := range in {
		var s []rune
		for _, c := range l {
			switch c {
			case '(', '[', '{', '<':
				s = append(s, c)
			case ']', '}', '>':
				c--
				fallthrough
			case ')':
				c--
				if s[len(s)-1] != c {
					pscore += plut[c]
					continue Lines // invalid line
				}
				s = s[:len(s)-1]
			}
		}

		as := 0
		for i := len(s) - 1; i >= 0; i-- {
			as = as*5 + alut[s[i]]
		}
		ascores = append(ascores, as)
	}

	sort.Ints(ascores)
	return pscore, ascores[len(ascores)/2]
}

func readInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
