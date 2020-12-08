package main

import "testing"

var testcase = []passwordWithPolicy{
	{"a", 1, 3, "abcde"},
	{"b", 1, 3, "cdefg"},
	{"c", 2, 9, "ccccccccc"},
}

func TestPart1(t *testing.T) {
	expected := 2
	actual := Part1(testcase)

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 1
	actual := Part2(testcase)

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
