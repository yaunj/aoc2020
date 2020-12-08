package aoc2020

import "testing"

const (
	day1input = "testdata/input1"
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
