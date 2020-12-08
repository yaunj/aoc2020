package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type passwordWithPolicy struct {
	Required string
	Min      int
	Max      int
	Password string
}

func (p *passwordWithPolicy) Valid() bool {
	count := strings.Count(p.Password, p.Required)
	return count >= p.Min && count <= p.Max
}

func (p *passwordWithPolicy) Valid2() bool {
	if len(p.Password) < p.Max {
		// Password is too short
		return false
	}

	offset1 := string(p.Password[p.Min-1])
	offset2 := string(p.Password[p.Max-1])

	if offset1 == p.Required && offset2 != p.Required {
		return true
	}

	if offset1 != p.Required && offset2 == p.Required {
		return true
	}

	return false
}

func passwordWithPolicyFromString(input string) (passwordWithPolicy, error) {
	var min, max int
	var req rune
	var pw string

	_, err := fmt.Sscanf(input, "%d-%d %c: %s", &min, &max, &req, &pw)
	if err != nil {
		return passwordWithPolicy{}, err
	}

	return passwordWithPolicy{string(req), min, max, pw}, nil
}

func passwordsFromFile(path string) ([]passwordWithPolicy, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	var passwords []passwordWithPolicy

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		pwpol, err := passwordWithPolicyFromString(line)
		if err != nil {
			return nil, err
		}

		passwords = append(passwords, pwpol)
	}

	return passwords, nil
}

// Part1 solves part 1 of day 2
func Part1(passwords []passwordWithPolicy) int {
	valid := 0

	for _, pwd := range passwords {
		if pwd.Valid() {
			valid++
		}
	}

	return valid
}

// Part2 solves part 2 of day 2
func Part2(passwords []passwordWithPolicy) int {
	valid := 0

	for _, pwd := range passwords {
		if pwd.Valid2() {
			valid++
		}
	}

	return valid
}

func main() {
	fmt.Println("Day 2")
	fmt.Println("=====")

	passwords, err := passwordsFromFile("testdata/day2")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	fmt.Println("Part 1:", Part1(passwords))
	fmt.Println("Part 2:", Part2(passwords))
}
