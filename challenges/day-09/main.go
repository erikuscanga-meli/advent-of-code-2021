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
	fmt.Println("sum of all risk levels of all low points", total)

	total = partTwo(input)
	fmt.Println("Multiplied size of three larger basins", total)
}

func partOne(input []string) int {
	var total int
	matrix := make([][]int, 0)
	for _, line := range input {
		row := make([]int, 0)
		scanner := bufio.NewScanner(strings.NewReader(line))

		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			row = append(row, atoi(scanner.Text()))
		}
		matrix = append(matrix, row)
	}

	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}
	rows := len(matrix)
	columns := len(matrix[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			ok := true
			for d := 0; d < 4; d++ {
				cr := r + dr[d]
				cc := c + dc[d]
				if (cr >= 0 && rows > cr) &&
					(cc >= 0 && columns > cc) &&
					matrix[cr][cc] <= matrix[r][c] {
					ok = false
				}
			}
			if ok {
				total += matrix[r][c] + 1
			}
		}
	}

	return total
}

func partTwo(input []string) int {
	matrix := make([][]int, 0)
	for _, line := range input {
		row := make([]int, 0)
		scanner := bufio.NewScanner(strings.NewReader(line))

		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			row = append(row, atoi(scanner.Text()))
		}
		matrix = append(matrix, row)
	}

	rows := len(matrix)
	columns := len(matrix[0])
	counts := make([]int, 0)

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			left, up, right, down := c-1, r-1, c+1, r+1
			isLowPoint := true

			if left >= 0 {
				isLowPoint = matrix[r][c] < matrix[r][left] && isLowPoint
			}
			if up >= 0 {
				isLowPoint = matrix[r][c] < matrix[up][c] && isLowPoint
			}
			if right < columns {
				isLowPoint = matrix[r][c] < matrix[r][right] && isLowPoint
			}
			if down < rows {
				isLowPoint = matrix[r][c] < matrix[down][c] && isLowPoint
			}
			if isLowPoint {
				c := basinCount(matrix, r, c, 1, map[string]struct{}{})
				counts = append(counts, c)
			}
		}
	}

	sort.Ints(counts)
	total := 1
	for _, v := range counts[len(counts)-3:] {
		total *= v
	}
	return total
}

func readInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func basinCount(grid [][]int, row int, col int, count int, visited map[string]struct{}) int {
	key := fmt.Sprintf("%d,%d", row, col)
	if _, ok := visited[key]; ok {
		return count - 1
	} else {
		visited[key] = struct{}{}
	}

	if grid[row][col] == 9 {
		return count - 1
	}

	rows := len(grid)
	columns := len(grid[0])
	left, up, right, down := col-1, row-1, col+1, row+1

	if left >= 0 {
		count = basinCount(grid, row, left, count+1, visited)
	}
	if up >= 0 {
		count = basinCount(grid, up, col, count+1, visited)
	}
	if right < columns {
		count = basinCount(grid, row, right, count+1, visited)
	}
	if down < rows {
		count = basinCount(grid, down, col, count+1, visited)
	}
	return count
}
