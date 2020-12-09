package main

import (
	"fmt"
	"testing"
)

func TestSeatID(t *testing.T) {
	/*
	   BFFFBBFRRR: row 70, column 7, seat ID 567.
	   FFFBBBFRRR: row 14, column 7, seat ID 119.
	   BBFFBBFRLL: row 102, column 4, seat ID 820.
	*/

	cases := []BoardingPass{
		{"FBFBBFFRLR", 44, 5, 357},
		{"BFFFBBFRRR", 70, 7, 567},
		{"FFFBBBFRRR", 14, 7, 119},
		{"BBFFBBFRLL", 102, 4, 820},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Boarding pass %d: %s", i, tc.ID), func(t *testing.T) {
			actual, err := NewBoardingPass(tc.ID)

			if err != nil {
				t.Fatal("Got unexpected error:", err)
			}

			if actual.Row != tc.Row {
				t.Error("Expected Row", tc.Row, "got", actual.Row)
			}
			if actual.Column != tc.Column {
				t.Error("Expected Column", tc.Column, "got", actual.Column)
			}
			if actual.SeatID != tc.SeatID {
				t.Error("Expected SeatID", tc.SeatID, "got", actual.SeatID)
			}
		})
	}
}
