package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Anscombe dataset
var anscombe = map[string][]float64{
	"x1": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x2": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x3": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x4": {8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	"y1": {8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	"y2": {9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
	"y3": {7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	"y4": {6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.50, 5.56, 7.91, 6.89},
}

// RegressionResults holds the results of the regression analysis
type RegressionResults struct {
	Intercept          float64
	Slope              float64
	RSquared           float64
	AdjustedRSquared   float64
	RMSE               float64
	MAE                float64
	FStatistic         float64
	TStatistic         float64
	PValue             float64
	ConfidenceInterval [2]float64
	PredictionInterval [2]float64
	StandardError      float64
}

func main() {
	for i := 1; i <= 4; i++ {
		x := anscombe[fmt.Sprintf("x%d", i)]
		y := anscombe[fmt.Sprintf("y%d", i)]
		results := calculateRegression(x, y)
		printResults(i, results)
		plotData(i, x, y, results)
	}
}

// Function to perform the regression calculation
func calculateRegression(x, y []float64) RegressionResults {
	alpha, beta := stat.LinearRegression(x, y, nil, false)

	// Calculate fitted values and residuals
	fitted := make([]float64, len(x))
	residuals := make([]float64, len(x))
	for i := range x {
		fitted[i] = alpha + beta*x[i]
		residuals[i] = y[i] - fitted[i]
	}

	// Calculate R-squared
	r := stat.Correlation(x, y, nil)
	rSquared := r * r

	// Calculate RMSE
	residualVariance := sumOfSquares(residuals)
	mse := residualVariance / float64(len(x)-2)
	rmse := math.Sqrt(mse)

	// Calculate Mean Absolute Error (MAE)
	var maeSum float64
	for _, res := range residuals {
		maeSum += math.Abs(res)
	}
	mae := maeSum / float64(len(residuals))

	// Calculate Adjusted R-squared
	n := float64(len(x))
	adjustedRSquared := 1 - ((1-rSquared)*(n-1))/(n-1-1)

	// Calculate F-statistic
	totalVariance := sumOfSquares(y)
	explainedVariance := totalVariance - residualVariance
	fStat := (explainedVariance / 1) / mse

	// Calculate standard error and t-statistic for slope
	meanX := stat.Mean(x, nil)
	var sumSquaredXDiff float64
	for _, xi := range x {
		sumSquaredXDiff += math.Pow(xi-meanX, 2)
	}
	standardError := math.Sqrt(mse / sumSquaredXDiff)
	tStat := beta / standardError

	// Compute p-value for slope using the t-distribution
	df := float64(len(x) - 2)
	tDist := distuv.StudentsT{Mu: 0, Sigma: 1, Nu: df}
	pValue := 2 * (1 - tDist.CDF(math.Abs(tStat)))

	// Confidence interval for slope
	confInterval := [2]float64{beta - 1.96*standardError, beta + 1.96*standardError}

	// Prediction interval for new observations (assuming x-bar ± 2 standard deviations)
	predictionInterval := [2]float64{alpha + beta*(meanX-2*standardError), alpha + beta*(meanX+2*standardError)}

	// Return the results
	return RegressionResults{
		Intercept:          alpha,
		Slope:              beta,
		RSquared:           rSquared,
		AdjustedRSquared:   adjustedRSquared,
		RMSE:               rmse,
		MAE:                mae,
		FStatistic:         fStat,
		TStatistic:         tStat,
		PValue:             pValue,
		ConfidenceInterval: confInterval,
		PredictionInterval: predictionInterval,
		StandardError:      standardError,
	}
}

// Utility function to sum squares of differences
func sumOfSquares(values []float64) float64 {
	mean := stat.Mean(values, nil)
	var ss float64
	for _, v := range values {
		ss += math.Pow(v-mean, 2)
	}
	return ss
}

// Function to print the regression results
func printResults(setNumber int, results RegressionResults) {
	fmt.Printf("Set %d Regression Results:\n", setNumber)
	fmt.Printf("Intercept: %.4f\n", results.Intercept)
	fmt.Printf("Slope: %.4f\n", results.Slope)
	fmt.Printf("R-squared: %.4f\n", results.RSquared)
	fmt.Printf("Adjusted R-squared: %.4f\n", results.AdjustedRSquared)
	fmt.Printf("RMSE: %.4f\n", results.RMSE)
	fmt.Printf("MAE: %.4f\n", results.MAE)
	fmt.Printf("F-statistic: %.4f\n", results.FStatistic)
	fmt.Printf("t-statistic: %.4f\n", results.TStatistic)
	fmt.Printf("p-value: %.4f\n", results.PValue)
	fmt.Printf("95%% Confidence Interval for Slope: [%.4f, %.4f]\n", results.ConfidenceInterval[0], results.ConfidenceInterval[1])
	fmt.Printf("Prediction Interval: [%.4f, %.4f]\n\n", results.PredictionInterval[0], results.PredictionInterval[1])
}

func plotData(setNumber int, x, y []float64, results RegressionResults) {
	// Create a new plot
	p := plot.New()

	// Set plot titles
	p.Title.Text = fmt.Sprintf("Anscombe Dataset %d", setNumber)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Set axis limits: x-axis from 0 to 20, y-axis from 0 to 14
	p.X.Min = 0
	p.X.Max = 20
	p.Y.Min = 0
	p.Y.Max = 14

	// Create points for scatter plot
	pts := make(plotter.XYs, len(x))
	for i := range x {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}

	// Create scatter plot
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatalf("could not create scatter plot: %v", err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red dots

	// Add scatter plot to the plot
	p.Add(s)

	// Plot the regression line
	regressionLine := plotter.NewFunction(func(x float64) float64 {
		return results.Intercept + results.Slope*x
	})
	regressionLine.Color = color.RGBA{B: 255, A: 255} // Blue line

	// Add regression line to the plot
	p.Add(regressionLine)

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 6*vg.Inch, fmt.Sprintf("anscombe_%d.png", setNumber)); err != nil {
		log.Fatalf("could not save plot: %v", err)
	}

	fmt.Printf("Plot for dataset %d saved as anscombe_%d.png\n", setNumber, setNumber)
}
