package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	template, polymerMap := readInput()

	total := solve(template, polymerMap, 10)
	fmt.Println("total after 10 cycles:", total)

	total = solve(template, polymerMap, 40)
	fmt.Println("total after 40 cycles:", total)
}

func solve(tmpl string, seq map[string]string, t int) int {
	p := make(map[string]int)
	for i := 0; i < len(tmpl)-1; i++ {
		p[tmpl[i:i+2]] += 1
	}
	p[tmpl[len(tmpl)-1:]] += 1

	for i := 0; i < t; i++ {
		np := make(map[string]int)
		for k, v := range p {
			if p, ok := seq[k]; ok {
				np[k[0:1]+p] += v
				np[p+k[1:2]] += v
			} else {
				np[k] += v
			}
		}
		p = np
	}

	c := make(map[string]int)
	for k, v := range p {
		c[k[0:1]] += v
	}

	min, max := math.MaxInt, math.MinInt
	for _, val := range c {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return max - min
}

func readInput() (string, map[string]string) {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	tmpl := ""
	p := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			continue
		}

		if strings.Contains(line, "->") {
			segments := strings.Split(line, " -> ")
			p[segments[0]] = segments[1]
		} else {
			tmpl = line
		}
	}

	return tmpl, p
}
