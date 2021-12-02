package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	slice := readInput()

	total := partOne(slice)
	fmt.Printf("Part one: there are %d measurements that are larger than the previous\n", total)

	total = partTwo(slice)
	fmt.Printf("Part two: there are %d sums that are larger than the previous sum\n", total)
}

func partOne(input []int) int {
	var prev, total int
	for _, n := range input {
		if n > prev {
			total++
		}
		prev = n
	}
	return total
}

func partTwo(input []int) int {
	var x []int
	for i := 0; i < len(input)-2; i++ {
		s := input[i] + input[i+1] + input[i+2]
		x = append(x, s)
	}
	return partOne(x)
}

func readInput() []int {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic("error opening file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var slice []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("error parsing input")
		}
		slice = append(slice, n)
	}
	return slice
}
