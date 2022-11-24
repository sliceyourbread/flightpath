package main

import "testing"

func BenchmarkFlightPaths4x4(b *testing.B) {
	start := 0
	end := 3

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		flightPaths(start, end)
	}
}

func BenchmarkFlightPaths8x8(b *testing.B) {
	start := 0
	end := 7

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		flightPaths(start, end)
	}
}

func BenchmarkCostcalculator4x4(b *testing.B) {
	paths := [][]int{{0, 1, 1, 2, 2, 3}, {0, 3}, {0, 2, 2, 3}}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		costCalculator(paths)
	}
}

func BenchmarkStopCalculator4x4(b *testing.B) {
	paths := [][]int{{0, 1, 1, 2, 2, 3}, {0, 3}, {0, 2, 2, 3}}
	dep := "Castle Black"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		stopCalculator(paths, dep)
	}
}
