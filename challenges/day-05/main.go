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
	total1, total2 := solve(slice)
	fmt.Println("Overlapped points (h/v):", total1)
	fmt.Println("Overlapped points (h/v/d):", total2)
}

func solve(input []string) (int, int) {
	positions1, positions2 := map[string]int{}, map[string]int{}
	for _, line := range input {
		p := strings.Split(line, "->")
		p1, p2 := p[0], p[1]

		xy1 := strings.Split(p1, ",")
		xy2 := strings.Split(p2, ",")
		x1, y1 := toInt(xy1[0]), toInt(xy1[1])
		x2, y2 := toInt(xy2[0]), toInt(xy2[1])

		if x1 == x2 {
			for y := min(y1, y2); y < max(y1, y2)+1; y++ {
				positions1[fmt.Sprint(x1, y)] += 1
				positions2[fmt.Sprint(x1, y)] += 1
			}
		} else if y1 == y2 {
			for x := min(x1, x2); x < max(x1, x2)+1; x++ {
				positions1[fmt.Sprint(x, y1)] += 1
				positions2[fmt.Sprint(x, y1)] += 1
			}
		} else {
			x, y := x1, y1
			dx := 1
			if x2 < x1 {
				dx = -1
			}

			dy := 1
			if y2 < y1 {
				dy = -1
			}

			for {
				if x == x2 && y == y2 {
					break
				}
				positions2[fmt.Sprint(x, y)] += 1
				x += dx
				y += dy
			}
			positions2[fmt.Sprint(x, y)] += 1
		}
	}

	return compute(positions1), compute(positions2)
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

func compute(m map[string]int) int {
	total := 0
	for _, v := range m {
		if v > 1 {
			total++
		}
	}
	return total
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func toInt(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return n
}
