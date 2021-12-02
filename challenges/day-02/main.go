package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	commands := readInput()

	total := partOne(commands)
	fmt.Println("Part one: horizontal * depth =", total)

	total = partTwo(commands)
	fmt.Println("Part two: horizontal * depth =", total)
}

func partOne(commands []string) int {
	horizontal, depth := 0, 0
	for _, cmd := range commands {
		split := strings.Split(cmd, " ")
		units, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			horizontal += units
		case "down":
			depth += units
		case "up":
			depth -= units
		}
	}
	return horizontal * depth
}

func partTwo(commands []string) int {
	horizontal, depth, aim := 0, 0, 0
	for _, cmd := range commands {
		split := strings.Split(cmd, " ")
		units, _ := strconv.Atoi(split[1])
		switch split[0] {
		case "forward":
			horizontal += units
			depth += aim * units
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	return horizontal * depth
}

func readInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic("error opening file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var commands []string
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
	return commands
}
