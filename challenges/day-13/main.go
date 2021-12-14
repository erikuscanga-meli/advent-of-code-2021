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
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	ls := bufio.NewScanner(file)
	resolve(ls)
}

func resolve(ls *bufio.Scanner) {
	pts, folds, sizeX, sizeY := readInput(ls)
	sheet := makeSheet(sizeX, sizeY, folds)
	for _, pt := range pts {
		x, y := pt.x, pt.y
		for _, fold := range folds {
			x, y = fold.apply(x, y)
		}
		sheet[x][y] = 1
	}

	var total int
	for y := 0; y < len(sheet); y++ {
		for x := 0; x < len(sheet[y]); x++ {
			if sheet[y][x] == 1 {
				total++
			}
		}
	}

	for j := 0; j < len(sheet[0]); j++ {
		for i := 0; i < len(sheet); i++ {
			if sheet[i][j] == 1 {
				fmt.Printf(" # ")
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Println("")
	}
}

type Pt struct {
	x int
	y int
}

type Fold struct {
	x int
	y int
}

func (f Fold) apply(x int, y int) (int, int) {
	return abs(f.x - abs(f.x-x)), abs(f.y - abs(f.y-y))
}

func makeSheet(sizeX int, sizeY int, folds []Fold) [][]int {
	for _, f := range folds {
		if f.x != 0 {
			sizeX = max(f.x, sizeX-f.x-1)
		} else {
			sizeY = max(f.y, sizeY-f.y-1)
		}
	}

	sheet := make([][]int, sizeX)
	for idx := range sheet {
		sheet[idx] = make([]int, sizeY)
	}

	return sheet
}

func readInput(ls *bufio.Scanner) ([]Pt, []Fold, int, int) {
	line, ok := read(ls)

	var (
		x    int
		y    int
		maxX = 0
		maxY = 0
		pts  = make([]Pt, 0)
	)

	for ok && line != "" {
		coords := strings.Split(line, ",")
		x, _ = strconv.Atoi(coords[0])
		y, _ = strconv.Atoi(coords[1])
		pts = append(pts, Pt{x, y})

		maxX = max(maxX, x)
		maxY = max(maxY, y)

		line, ok = read(ls)
	}

	line, ok = read(ls)
	folds := make([]Fold, 0)
	for ok {
		x, y = 0, 0
		if strings.Contains(line, "fold along x=") {
			x, _ = strconv.Atoi(strings.Split(line, "=")[1])
		} else {
			y, _ = strconv.Atoi(strings.Split(line, "=")[1])
		}

		folds = append(folds, Fold{x, y})
		line, ok = read(ls)
	}

	return pts, folds, maxX + 1, maxY + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func read(scanner *bufio.Scanner) (string, bool) {
	if ok := scanner.Scan(); ok {
		return scanner.Text(), ok
	}
	return "", false
}
