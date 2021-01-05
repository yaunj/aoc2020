package main

import (
	"fmt"
	"os"

	"github.com/yaunj/aoc2020"
)

func uniqueAnswers(answers []string) int {
	seen := map[rune]bool{}

	for _, answer := range answers {
		for _, c := range answer {
			seen[c] = true
		}
	}

	count := 0

	for range seen {
		count++
	}

	return count
}

func commonAnswers(answers []string) int {
	seen := map[rune]int{}

	for _, answer := range answers {
		for _, c := range answer {
			seen[c]++
		}
	}

	count := 0

	for _, value := range seen {
		if value == len(answers) {
			count++
		}
	}

	return count
}

// Part1 solves the first part of day6 challenge
func Part1(groups [][]string) int {
	sum := 0

	for _, group := range groups {
		sum += uniqueAnswers(group)
	}

	return sum
}

// Part2 solves the second part of day6 challenge
func Part2(groups [][]string) int {
	sum := 0

	for _, group := range groups {
		sum += commonAnswers(group)
	}

	return sum
}

func main() {
	fmt.Println("Day 6")
	fmt.Println("=====")

	groups, err := aoc2020.ChunksFromFile("testdata/day6")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading customs forms:", err)
		os.Exit(1)
	}

	fmt.Println("Part 1:", Part1(groups))
	fmt.Println("Part 2:", Part2(groups))
}
