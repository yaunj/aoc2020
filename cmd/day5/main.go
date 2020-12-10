package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/yaunj/aoc2020"
)

// BoardingPass represents the boarding passes
type BoardingPass struct {
	ID     string
	Row    int
	Column int
	SeatID int
}

// NewBoardingPass creates a new BoardingPass from an ID
func NewBoardingPass(ID string) (*BoardingPass, error) {
	if len(ID) != 10 {
		return nil, fmt.Errorf("length of ID must == 10")
	}

	row, err := bisect(0, 127, ID[:7], 'F', 'B')
	if err != nil {
		return nil, err
	}

	column, err := bisect(0, 7, ID[7:], 'L', 'R')
	if err != nil {
		return nil, err
	}

	seatID := row*8 + column

	return &BoardingPass{ID: ID, Row: row, Column: column, SeatID: seatID}, nil
}

func bisect(min, max int, sequence string, low, high rune) (int, error) {
	res := 0

	for _, part := range sequence {
		pivot := ((max - min) / 2) + min

		switch part {
		case low:
			max = pivot
		case high:
			min = pivot + 1
			break
		default:
			return -1, fmt.Errorf("invalid format (%c not in [%c, %c])", part, low, high)
		}

		res = max
	}

	return res, nil
}

// Part1 solves part 1 of the challenge
func Part1(passes []*BoardingPass) int {
	max := 0

	for _, bp := range passes {
		if bp.SeatID > max {
			max = bp.SeatID
		}
	}

	return max
}

// Part2 solves part 2 of the challenge
func Part2(passes []*BoardingPass) int {
	ids := []int{}

	for _, bp := range passes {
		ids = append(ids, bp.SeatID)
	}

	sort.Ints(ids)
	prev := ids[0]

	for i := 1; i < len(ids); i++ {
		if ids[i] == prev+2 {
			return prev + 1
		}

		prev = ids[i]
	}

	return 0
}

func main() {
	fmt.Println("Day 5")
	fmt.Println("=====")

	passes, err := aoc2020.LinesFromFile("testdata/day5")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading boarding passes:", err)
		os.Exit(1)
	}

	boardingPasses := []*BoardingPass{}
	for _, pass := range passes {
		bp, err := NewBoardingPass(pass)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing boarding pass:", err)
			continue
		}

		boardingPasses = append(boardingPasses, bp)
	}

	fmt.Println("Part 1:", Part1(boardingPasses))
	fmt.Println("Part 2:", Part2(boardingPasses))
}
