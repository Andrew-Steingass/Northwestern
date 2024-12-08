package main

import (
	"math"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup
	initLoggers()

	// Run tests
	code := m.Run()

	// Exit
	os.Exit(code)
}

func TestMedian(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Empty slice", []float64{}, 0},
		{"Single value", []float64{1}, 1},
		{"Odd length", []float64{1, 2, 3}, 2},
		{"Even length", []float64{1, 2, 3, 4}, 2.5},
		{"Unordered odd", []float64{3, 1, 2}, 2},
		{"Unordered even", []float64{4, 1, 3, 2}, 2.5},
		{"With duplicates", []float64{1, 2, 2, 3, 3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Median(tt.input)
			if math.Abs(got-tt.expected) > 1e-6 {
				t.Errorf("Median() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCalculateSE(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Empty slice", []float64{}, 0},
		{"Single value", []float64{1}, 0},
		{"Multiple values", []float64{1, 2, 3}, 1},
		{"Same values", []float64{2, 2, 2}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateSE(tt.input)
			if math.Abs(got-tt.expected) > 1e-6 {
				t.Errorf("calculateSE() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBootstrapMedians(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	indices := [][]int{
		{0, 1, 2, 2, 1},
		{1, 1, 3, 4, 0},
	}

	medians := bootstrapMedians(data, indices)

	// Check length
	if len(medians) != len(indices) {
		t.Errorf("Expected %d medians, got %d", len(indices), len(medians))
	}

	// Check some expected values
	expectedMedians := []float64{2, 2} // These values should be calculated based on your resampling
	for i, expected := range expectedMedians {
		if math.Abs(medians[i]-expected) > 1e-6 {
			t.Errorf("Median[%d] = %v, want %v", i, medians[i], expected)
		}
	}
}

func BenchmarkMedian(b *testing.B) {
	data := []float64{1, 3, 5, 2, 4, 6, 8, 7, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Median(data)
	}
}

func BenchmarkCalculateSE(b *testing.B) {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculateSE(data)
	}
}
