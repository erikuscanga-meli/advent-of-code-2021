package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	slice := readInput()

	power := partOne(slice)
	fmt.Println("the power consumption of the submarine is", power)

	fn1 := func(x, y []string) bool { return len(x) >= len(y) }
	fn2 := func(x, y []string) bool { return len(x) < len(y) }
	lifeSupport := partTwo(slice, fn2) * partTwo(slice, fn1)
	fmt.Println("the life support rating of the submarine is", lifeSupport)
}

func partOne(input []string) int64 {
	var gamma, epsilon string
	for i := 0; i < len(input[0]); i++ {
		var zeros, ones int
		for _, str := range input {
			if string(str[i]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if ones > zeros {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	x, _ := strconv.ParseInt(gamma, 2, 64)
	y, _ := strconv.ParseInt(epsilon, 2, 64)

	return x * y
}

func partTwo(input []string, predicate func([]string, []string) bool) int64 {
	for i := 0; i < len(input[0]); i++ {
		zeroes := make([]string, 0)
		ones := make([]string, 0)

		for _, str := range input {
			if string(str[i]) == "1" {
				ones = append(ones, str)
			} else {
				zeroes = append(zeroes, str)
			}
		}

		if predicate(ones, zeroes) {
			input = ones
		} else {
			input = zeroes
		}

		if len(input) == 1 {
			break
		}
	}

	x, _ := strconv.ParseInt(input[0], 2, 64)
	return x
}

func readInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic("error opening file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var str []string
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	return str
}
