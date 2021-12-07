package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := parseToIntArray(readInput())

	total := partOne(input)
	fmt.Println("minimal amount of fuel consumption (constant):", total)

	total = partTwo(input)
	fmt.Println("minimal amount of fuel consumption (n+1):", total)
}

func partOne(input []int) int {
	var minimal float64
	for i := 1; i < len(input); i++ {
		var consumption float64
		for _, position := range input {
			if position != i {
				consumption += math.Abs(float64(position - i))
			}
		}
		if consumption < minimal || i == 1 {
			minimal = consumption
		}
	}
	return int(minimal)
}

func partTwo(input []int) int {
	var minimal float64
	for i := 1; i < len(input); i++ {
		var consumption float64
		for _, position := range input {
			if position != i {
				consumption += calculate(math.Abs(float64(position - i)))
			}
		}
		if consumption < minimal || i == 1 {
			minimal = consumption
		}
	}
	return int(minimal)
}

func readInput() string {
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
	return str[0]
}

func parseToIntArray(input string) []int {
	var arr []int
	segments := strings.Split(input, ",")
	for _, s := range segments {
		n, _ := strconv.Atoi(s)
		arr = append(arr, n)
	}
	return arr
}

func calculate(x float64) float64 {
	var total float64
	for i := 1; i <= int(x); i++ {
		total += float64(1 * i)
	}
	return total
}
