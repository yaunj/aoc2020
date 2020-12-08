package main

import (
	"fmt"
	"strings"
	"testing"
)

var testmap = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestPart1(t *testing.T) {
	treemap, err := mapFromReader(strings.NewReader(testmap))
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	expected := 7
	actual := Part1(treemap)

	if actual != expected {
		t.Error("Expected", expected, "collisions, got", actual)
	}
}

func TestPart2(t *testing.T) {
	treemap, _ := mapFromReader(strings.NewReader(testmap))

	expected := 336
	actual := Part2(treemap)

	if actual != expected {
		t.Error("Expected", expected, "collisions, got", actual)
	}
}

func TestTraverseSlope(t *testing.T) {
	treemap, _ := mapFromReader(strings.NewReader(testmap))
	cases := []struct {
		Right int
		Down  int
		Exp   int
	}{
		{1, 1, 2},
		{3, 1, 7},
		{5, 1, 3},
		{7, 1, 4},
		{1, 2, 2},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Slope %d", i), func(t *testing.T) {
			actual := traverseSlope(treemap, tc.Right, tc.Down)
			if actual != tc.Exp {
				t.Error("Expected", tc.Exp, "collisions, got", actual)
			}
		})
	}
}
