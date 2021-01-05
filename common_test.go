package aoc2020

import (
	"fmt"
	"testing"
)

const (
	day1input = "testdata/day1"
	day4input = "testdata/day4"
)

func TestIntsFromFile(t *testing.T) {
	expectedCount := 200
	ints, err := IntsFromFile(day1input)
	if err != nil {
		t.Fatal("Got unexpected error:", err)
	}
	actualCount := len(ints)

	if actualCount != expectedCount {
		t.Error("Expected", expectedCount, "got", actualCount)
	}
}

func BenchmarkIntsFromFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntsFromFile(day1input)
	}
}

func TestIntsFromFileViaFields(t *testing.T) {
	expectedCount := 200
	ints, err := IntsFromFileViaFields(day1input)
	if err != nil {
		t.Fatal("Got unexpected error:", err)
	}
	actualCount := len(ints)

	if actualCount != expectedCount {
		t.Error("Expected", expectedCount, "got", actualCount)
	}
}

func BenchmarkIntsFromFileViaFields(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntsFromFileViaFields(day1input)
	}
}

func TestIntsFromFileVs(t *testing.T) {
	a, _ := IntsFromFile(day1input)
	b, _ := IntsFromFileViaFields(day1input)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Fatalf("IntsFromFile and IntsFromFileViaFields differs at offset %d: %d vs %d\n", i, a[i], b[i])
		}
	}
}

func TestChunksFromFile(t *testing.T) {
	cases := []struct {
		path     string
		expected int
	}{
		{day1input, 1},
		{day4input, 288},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("ChunksForDay%d", i), func(t *testing.T) {
			chunks, err := ChunksFromFile(tc.path)

			if err != nil {
				t.Error("Got unexpected error:", err)
			}

			if len(chunks) != tc.expected {
				t.Error("Expected", tc.expected, "chunk(s), got", len(chunks))
			}
		})
	}
}
