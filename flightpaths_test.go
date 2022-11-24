package main

import (
	"testing"
)

func TestFlightPaths(t *testing.T) {

	var tests = []struct {
		start    int
		end      int
		expected int
	}{
		{0, 1, 1},
		{0, 2, 2},
		{0, 3, 4},
		{1, 2, 1},
		{2, 3, 1},
		{1, 3, 2},
	}

	for _, test := range tests {
		actual := flightPaths(test.start, test.end)

		if len(actual) != test.expected {
			t.Errorf("flightPaths failed for test: %v", test)
		}
	}
}

func TestFlightPaths8x8Matrix(t *testing.T) {

	combinations = nil
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			combinations = append(combinations, []int{j, i})
		}
	}

	actual := flightPaths(0, 7)
	expected := 64

	if len(actual) != expected {
		t.Errorf("flightPaths faild for 8x8 matrix")
	}
}

func TestCostCalculator(t *testing.T) {

	var tests = []struct {
		path     [][]int
		expected []int
	}{
		{[][]int{{0, 1, 1, 2}, {0, 2}}, []int{55, 80}},
		{[][]int{{0, 1, 1, 2, 2, 3}, {0, 3}, {0, 2, 2, 3}}, []int{125, 90, 150}},
		{[][]int{{1, 2, 2, 3}, {1, 3}}, []int{110, 50}},
		{[][]int{{2, 3}}, []int{70}},
		{[][]int{{1, 2}}, []int{40}},
	}

	for _, test := range tests {
		actual := costCalculator(test.path)
		for i := range actual {
			if actual[i] != test.expected[i] {
				t.Errorf("costCalculator incorrectly calculated the cost for test: %v", test)
			}
		}
	}
}

func TestStopCalculator(t *testing.T) {
	var tests = []struct {
		path      [][]int
		departure string
		expected  string
	}{
		{[][]int{{0, 2, 2, 3}}, "Castle Black", "Castle Black -> Riverrun -> King's Landing"},
		{[][]int{{1, 3}}, "Winterfell", "Winterfell -> King's Landing"},
		{[][]int{{1, 2}}, "Winterfell", "Winterfell -> Riverrun"},
		{[][]int{{2, 3}}, "Riverrun", "Riverrun -> King's Landing"},
		{[][]int{{1, 2, 2, 3}}, "Winterfell", "Winterfell -> Riverrun -> King's Landing"},
		{[][]int{{0, 1, 1, 2, 2, 3}}, "Castle Black", "Castle Black -> Winterfell -> Riverrun -> King's Landing"},
	}

	for _, test := range tests {
		actual := stopCalculator(test.path, test.departure)
		if actual[0] != test.expected {
			t.Errorf("stopCalculator failed to output the correct path as strings for test: %v", test)
		}
	}
}
