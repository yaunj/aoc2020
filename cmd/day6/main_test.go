package main

import (
	"testing"

	"github.com/yaunj/aoc2020"
)

var (
	simpleCase = []string{"abcx", "abcy", "abcz"}
	longerCase = [][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
	}
	day6chunks = [][]string{}
)

func init() {
	day6chunks, _ = aoc2020.ChunksFromFile("testdata/day6")
}

func TestUniqueAnswers(t *testing.T) {
	expected := 6
	actual := uniqueAnswers(simpleCase)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}

	longerExpectations := []int{3, 3, 3, 1, 1}
	for i, tc := range longerCase {
		expected = longerExpectations[i]
		actual = uniqueAnswers(tc)

		if actual != expected {
			t.Error("Expected", expected, "got", actual)
		}
	}
}

func TestPart1(t *testing.T) {
	expected := 6596
	actual := Part1(day6chunks)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}

func TestCommonAnswers(t *testing.T) {
	expectations := []int{3, 0, 1, 1, 1}

	for i, tc := range longerCase {
		expected := expectations[i]
		actual := commonAnswers(tc)

		if actual != expected {
			t.Error("Expected", expected, "got", actual)
		}
	}
}

func TestPart2(t *testing.T) {
	expected := 3219
	actual := Part2(day6chunks)

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}
}
