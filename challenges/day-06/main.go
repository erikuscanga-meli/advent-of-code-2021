package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice := readInput()
	total := solve(slice[0], 80)
	fmt.Println("total lanterfishes after 80 days:", total)

	total = solve(slice[0], 256)
	fmt.Println("total lanterfishes after 256 days:", total)
}

func solve(input string, days int) int {
	m := make([]int, 9)
	for _, x := range strings.Split(input, ",") {
		n, _ := strconv.ParseInt(x, 10, 64)
		m[n] += 1
	}

	for i := 0; i < days; i++ {
		y := make([]int, 9)
		for x, v := range m {
			if x == 0 {
				y[6] += v
				y[8] += v
			} else {
				y[x-1] += v
			}
		}
		m = y
	}
	return sum(m)
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

func sum(m []int) int {
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}
