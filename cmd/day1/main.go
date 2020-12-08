package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/yaunj/aoc2020"
)

// Part1 finds two ints that sums to 2020 and return their product
func Part1(input []int) int {
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}

	return 0
}

// Part2 finds three ints that sums to 2020 and return their product
func Part2(input []int) int {
	max := len(input)

	for a := 0; a < max-2; a++ {
		for b := a + 1; b < max-1; b++ {
			for c := b + 1; c < max; c++ {
				if input[a]+input[b]+input[c] == 2020 {
					return input[a] * input[b] * input[c]
				}
			}
		}
	}

	return 0
}

func intsFromFile(path string) ([]int, error) {
	var num int
	ints := []int{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	for {
		_, err = fmt.Fscanf(reader, "%d\n", &num)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		ints = append(ints, num)
	}

	return ints, nil
}

func main() {
	fmt.Println("Day 1")
	fmt.Println("=====")

	day1, err := aoc2020.IntsFromFileViaFields("testdata/day1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse input", err)
		os.Exit(1)
	}

	fmt.Println("Part1:", Part1(day1))
	fmt.Println("Part2:", Part2(day1))
}
