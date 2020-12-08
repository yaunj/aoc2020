package main

import "testing"

var (
	descriptionCase = []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
)

func TestPart1(t *testing.T) {
	expected := 514579
	actual := Part1(descriptionCase)

	if actual != expected {
		t.Fatal("Expected", expected, "got", actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 241861950
	actual := Part2(descriptionCase)

	if actual != expected {
		t.Fatal("Expected", expected, "got", actual)
	}
}
