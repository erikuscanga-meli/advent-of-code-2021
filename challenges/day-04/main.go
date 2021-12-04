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

	total := partOne(slice)
	fmt.Println("score for first completed board", total)

	total = partTwo(slice)
	fmt.Println("score for last completed board", total)
}

type bingoBoard [][]string

func partOne(input []string) int {
	drawNumbers := strings.Split(input[0], ",")
	boards := getBoards(input[2:])

	for _, drawNumber := range drawNumbers {
		for n := 0; n < len(boards); n++ {
			boards[n] = markNumber(boards[n], drawNumber)
			if checkBoard(boards[n]) {
				return shoutBingo(boards[n], drawNumber)
			}
		}
	}

	return 0
}

func partTwo(input []string) int {
	drawNumbers := strings.Split(input[0], ",")
	boards := getBoards(input[2:])
	var winnerOrder []int
	var lastPickerNumber string

	for _, drawNumber := range drawNumbers {
		for n := 0; n < len(boards); n++ {
			if exists(winnerOrder, n) {
				continue
			}

			boards[n] = markNumber(boards[n], drawNumber)
			if checkBoard(boards[n]) {
				winnerOrder = append(winnerOrder, n)
				lastPickerNumber = drawNumber
			}
		}

		if len(winnerOrder) == len(boards) {
			break
		}
	}

	lastWinner := winnerOrder[len(winnerOrder)-1]
	return shoutBingo(boards[lastWinner], lastPickerNumber)
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

func getBoards(input []string) []bingoBoard {
	totalBoards := len(input) / 5
	boards := make([]bingoBoard, totalBoards)

	var n, m int
	for y := 0; y < len(input); y++ {
		if input[y] == "" {
			continue
		}

		numbers := []string{}
		for _, v := range strings.Split(input[y], " ") {
			if strings.TrimSpace(v) == "" {
				continue
			}
			numbers = append(numbers, v)
		}

		if boards[n] == nil {
			boards[n] = make([][]string, 5)
		}
		boards[n][m] = numbers

		m++
		if m > 4 {
			n++
			m = 0
		}

		if n >= totalBoards {
			break
		}
	}
	return boards
}

func markNumber(board bingoBoard, n string) bingoBoard {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == n {
				board[x][y] = "x"
				return board
			}
		}
	}
	return board
}

func checkBoard(board bingoBoard) bool {
	for x := 0; x < len(board); x++ {
		isMarked := true

		// rows
		for y := 0; y < len(board[x]); y++ {
			isMarked = isMarked && board[x][y] == "x"
		}
		if isMarked {
			return true
		}

		// columns
		isMarked = true
		for y := 0; y < len(board[x]); y++ {
			isMarked = isMarked && board[y][x] == "x"
		}
		if isMarked {
			return true
		}
	}
	return false
}

func shoutBingo(board bingoBoard, lastNumber string) int {
	var total int
	for _, row := range board {
		for _, value := range row {
			if value == "x" {
				continue
			}

			n, _ := strconv.Atoi(value)
			total += n
		}
	}

	n, _ := strconv.Atoi(lastNumber)
	return total * n
}

func exists(slice []int, n int) bool {
	for _, v := range slice {
		if n == v {
			return true
		}
	}
	return false
}
