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

	row, column, seatID := 0, 0, 0
	min, max := 0, 127

	for _, part := range ID[:7] {
		pivot := ((max - min) / 2) + min
		// fmt.Printf("min: %3d  pivot: %3d  max: %3d  part: %c\n", min, pivot, max, part)

		switch part {
		case 'F': // lower
			max = pivot
			break
		case 'B': // upper
			min = pivot + 1
			break
		default:
			return nil, fmt.Errorf("invalid format: %s", ID)
		}

		row = max
	}

	min, max = 0, 7
	for _, part := range ID[7:] {
		pivot := ((max - min) / 2) + min
		// fmt.Printf("min: %3d  pivot: %3d  max: %3d  part: %c\n", min, pivot, max, part)

		switch part {
		case 'L': // lower
			max = pivot
			break
		case 'R': // upper
			min = pivot + 1
			break
		default:
			return nil, fmt.Errorf("invalid format: %s", ID)
		}

		column = max
	}

	seatID = row*8 + column

	return &BoardingPass{ID: ID, Row: row, Column: column, SeatID: seatID}, nil
}

// Part1 solves part 1 of the challenge
func Part1(passes []string) int {
	max := 0

	for _, pass := range passes {
		bp, err := NewBoardingPass(pass)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing boarding pass:", err)
			continue
		}

		if bp.SeatID > max {
			max = bp.SeatID
		}
	}

	return max
}

// Part2 solves part 2 of the challenge
func Part2(passes []string) int {
	ids := []int{}

	for _, pass := range passes {
		bp, _ := NewBoardingPass(pass)
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

	fmt.Println("Part 1:", Part1(passes))
	fmt.Println("Part 2:", Part2(passes))
}
