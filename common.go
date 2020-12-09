package aoc2020

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// IntsFromFile loads a list of ints (one per line) from a file
func IntsFromFile(path string) ([]int, error) {
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

// IntsFromFileViaFields uses strings.Fields to load a list of ints from a file
func IntsFromFileViaFields(path string) ([]int, error) {
	var num int
	ints := []int{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	for _, part := range strings.Fields(string(data)) {
		num, err = strconv.Atoi(part)
		ints = append(ints, num)
	}

	return ints, nil
}

// LinesFromFile returns a slice of strings, one per line in a file
func LinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
