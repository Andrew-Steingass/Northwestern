package main

import (
	"math"
	"testing"
)

// Expected values from Python or R for comparison (using dataset 1 for testing)
var expectedResults = RegressionResults{
	Intercept:  3.0001,
	Slope:      0.5001,
	RSquared:   0.6665,
	RMSE:       1.2366,
	FStatistic: 17.9899,
	TStatistic: 4.2415,
	PValue:     0.0022,
}

// Unit test to check if the regression results are correct
func TestCalculateRegression(t *testing.T) {
	x := anscombe["x1"]
	y := anscombe["y1"]

	results := calculateRegression(x, y)

	// Allow some tolerance for floating-point comparisons
	tol := 0.0001

	if math.Abs(results.Intercept-expectedResults.Intercept) > tol {
		t.Errorf("Intercept mismatch: got %f, expected %f", results.Intercept, expectedResults.Intercept)
	}
	if math.Abs(results.Slope-expectedResults.Slope) > tol {
		t.Errorf("Slope mismatch: got %f, expected %f", results.Slope, expectedResults.Slope)
	}
	if math.Abs(results.RSquared-expectedResults.RSquared) > tol {
		t.Errorf("R-squared mismatch: got %f, expected %f", results.RSquared, expectedResults.RSquared)
	}
	if math.Abs(results.RMSE-expectedResults.RMSE) > tol {
		t.Errorf("RMSE mismatch: got %f, expected %f", results.RMSE, expectedResults.RMSE)
	}
	if math.Abs(results.FStatistic-expectedResults.FStatistic) > tol {
		t.Errorf("F-statistic mismatch: got %f, expected %f", results.FStatistic, expectedResults.FStatistic)
	}
	if math.Abs(results.TStatistic-expectedResults.TStatistic) > tol {
		t.Errorf("T-statistic mismatch: got %f, expected %f", results.TStatistic, expectedResults.TStatistic)
	}
	if math.Abs(results.PValue-expectedResults.PValue) > tol {
		t.Errorf("P-value mismatch: got %f, expected %f", results.PValue, expectedResults.PValue)
	}
}

// Benchmark test to measure the performance of the regression calculation
func BenchmarkCalculateRegression(b *testing.B) {
	x := anscombe["x1"]
	y := anscombe["y1"]

	// Run the benchmark for b.N iterations
	for i := 0; i < b.N; i++ {
		_ = calculateRegression(x, y)
	}
}
