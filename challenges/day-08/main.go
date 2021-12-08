package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	total := partOne(input)
	fmt.Println(total)

	total = partTwo(input)
	fmt.Println(total)
}

var numbersMap = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func partOne(input []string) int {
	digits := make(map[int]int, 0)
	for _, line := range input {
		segments := strings.Split(line, " | ")
		segments = strings.Split(segments[1], " ")
		for _, v := range segments {
			for i, n := range numbersMap {
				if len(v) == len(i) {
					digits[n] += 1
				}
			}
		}
	}
	return sumMap(digits, 1, 4, 7, 8)
}

func partTwo(input []string) int {
	var total int
	for _, line := range input {
		digitsMap := make(map[string]string)
		parts := strings.Split(line, "|")
		segments := strings.Split(parts[0], " ")
		segments = append(segments, strings.Split(parts[1], " ")...)
		var x, y, z string
		for _, seq := range segments {
			d := sorting(seq)
			if _, ok := digitsMap[d]; ok {
				continue
			}

			switch len(seq) {
			case 2:
				digitsMap[d] = "1"
				x = d
			case 3:
				digitsMap[d] = "7"
				y = d
			case 4:
				digitsMap[d] = "4"
				z = d
			case 7:
				digitsMap[d] = "8"
			}
		}

		for _, seq := range segments {
			d := sorting(seq)
			co := count(d, x)
			cs := count(d, y)
			cf := count(d, z)

			switch len(seq) {
			case 5:
				switch {
				case co == 2 && cs == 3 && cf == 3:
					digitsMap[d] = "3"
				case co == 1 && cs == 2 && cf == 3:
					digitsMap[d] = "5"
				case co == 1 && cs == 2 && cf == 2:
					digitsMap[d] = "2"
				}
			case 6:
				switch {
				case co == 1 && cs == 2 && cf == 3:
					digitsMap[d] = "6"
				case co == 2 && cs == 3 && cf == 4:
					digitsMap[d] = "9"
				case co == 2 && cs == 3 && cf == 3:
					digitsMap[d] = "0"
				}
			}
		}

		values := strings.Split(parts[1], " ")
		var sb strings.Builder
		for _, digit := range values {
			sb.WriteString(digitsMap[sorting(digit)])
		}
		n, _ := strconv.Atoi(sb.String())
		total += n
	}

	return total
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

func getSegments(input []string) []string {
	var lines []string
	for _, line := range input {
		lines = append(lines, strings.Split(line, "|")...)
	}
	return lines
}

func sumMap(m map[int]int, args ...int) int {
	var t int
	if len(args) > 0 {
		for _, arg := range args {
			t += m[arg]
		}
	} else {
		for _, n := range m {
			t += n
		}
	}
	return t
}

func sumSlice(s []int) int {
	var t int
	for _, v := range s {
		t = v
	}
	return t
}

func sorting(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func count(x, y string) int {
	var z int
	for i := 0; i < len(x); i++ {
		if strings.Contains(y, string(x[i])) {
			z++
		}
	}
	return z
}
