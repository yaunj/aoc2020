package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Treemap is a map of trees on a slope
type Treemap [][]bool

func mapFromFile(path string) (Treemap, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return mapFromReader(file)
}

func mapFromReader(r io.Reader) (Treemap, error) {
	reader := bufio.NewReader(r)
	var treemap Treemap
	var line []bool

	for c, err := reader.ReadByte(); err != io.EOF; c, err = reader.ReadByte() {
		switch c {
		case '.':
			line = append(line, false)
			break
		case '#':
			line = append(line, true)
			break
		case '\n':
			treemap = append(treemap, line)
			line = []bool{}
			break
		default:
			return nil, fmt.Errorf("unexpected map point: %c", c)
		}
	}

	return treemap, nil
}

func traverseSlope(t Treemap, right, down int) int {
	var posX, posY int
	var collisions int

	width := len(t[0])
	height := len(t)

	for {
		// step
		posX, posY = (posX+right)%width, posY+down

		if posY > height-1 {
			// we're out
			break
		}

		if t[posY][posX] {
			collisions++
		}
	}

	return collisions
}

// Part1 solves part 1 of day 3 challenge
func Part1(t Treemap) int {
	return traverseSlope(t, 3, 1)
}

// Part2 solves part 2 of day 3 challenge
func Part2(t Treemap) int {
	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	result := 1

	for _, slope := range slopes {
		hits := traverseSlope(t, slope[0], slope[1])
		result *= hits
	}

	return result
}

func main() {
	fmt.Println("Day 3")
	fmt.Println("=====")

	treemap, err := mapFromFile("testdata/day3")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing input:", err)
		os.Exit(1)
	}

	fmt.Println("Part 1:", Part1(treemap))
	fmt.Println("Part 2:", Part2(treemap))
}
