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
	Intercept        float64
	Slope            float64
	RSquared         float64
	AdjustedRSquared float64
	RMSE             float64
	FStatistic       float64
	TStatistic       float64
	PValue           float64
	InterceptCI      [2]float64
	SlopeCI          [2]float64
	StandardError    float64
	DurbinWatson     float64
}

func main() {
	for i := 1; i <= 4; i++ {
		x := anscombe[fmt.Sprintf("x%d", i)]
		y := anscombe[fmt.Sprintf("y%d", i)]
		results := calculateRegression(x, y)
		printResults(i, results)
		if err := plotData(i, x, y, results); err != nil {
			log.Fatalf("could not plot data: %v", err)
		}
	}
}

// Function to perform the regression calculation
func calculateRegression(x, y []float64) RegressionResults {
	alpha, beta := stat.LinearRegression(x, y, nil, false)
	residuals := calculateResiduals(x, y, alpha, beta)
	rSquared := calculateRSquared(x, y)
	adjustedRSquared := calculateAdjustedRSquared(rSquared, len(x), 1)
	rmse := calculateRMSE(residuals)
	fStat := calculateFStatistic(y, residuals, rmse)
	tStat, pValue := calculateTStatAndPValue(x, beta, rmse)
	interceptCI, slopeCI := calculateConfidenceIntervals(x, y, alpha, beta, rmse)
	standardError := calculateStandardError(residuals)
	durbinWatson := calculateDurbinWatson(residuals)

	return RegressionResults{
		Intercept:        alpha,
		Slope:            beta,
		RSquared:         rSquared,
		AdjustedRSquared: adjustedRSquared,
		RMSE:             rmse,
		FStatistic:       fStat,
		TStatistic:       tStat,
		PValue:           pValue,
		InterceptCI:      interceptCI,
		SlopeCI:          slopeCI,
		StandardError:    standardError,
		DurbinWatson:     durbinWatson,
	}
}

func calculateResiduals(x, y []float64, alpha, beta float64) []float64 {
	residuals := make([]float64, len(x))
	for i := range x {
		fitted := alpha + beta*x[i]
		residuals[i] = y[i] - fitted
	}
	return residuals
}

func calculateRSquared(x, y []float64) float64 {
	r := stat.Correlation(x, y, nil)
	return r * r
}

func calculateAdjustedRSquared(rSquared float64, n, p int) float64 {
	return 1 - (1-rSquared)*(float64(n-1)/float64(n-p-1))
}

func calculateRMSE(residuals []float64) float64 {
	residualVariance := sumOfSquares(residuals)
	mse := residualVariance / float64(len(residuals)-2)
	return math.Sqrt(mse)
}

func calculateFStatistic(y, residuals []float64, rmse float64) float64 {
	totalVariance := sumOfSquares(y)
	explainedVariance := totalVariance - sumOfSquares(residuals)
	return (explainedVariance / 1) / rmse
}

func calculateTStatAndPValue(x []float64, beta, rmse float64) (float64, float64) {
	meanX := stat.Mean(x, nil)
	var sumSquaredXDiff float64
	for _, xi := range x {
		sumSquaredXDiff += math.Pow(xi-meanX, 2)
	}
	standardError := math.Sqrt(rmse / sumSquaredXDiff)
	tStat := beta / standardError
	df := float64(len(x) - 2)
	tDist := distuv.StudentsT{Mu: 0, Sigma: 1, Nu: df}
	pValue := 2 * (1 - tDist.CDF(math.Abs(tStat)))
	return tStat, pValue
}

func calculateConfidenceIntervals(x, y []float64, alpha, beta, rmse float64) ([2]float64, [2]float64) {
	meanX := stat.Mean(x, nil)
	var sumSquaredXDiff float64
	for _, xi := range x {
		sumSquaredXDiff += math.Pow(xi-meanX, 2)
	}
	standardError := math.Sqrt(rmse / sumSquaredXDiff)
	tDist := distuv.StudentsT{Mu: 0, Sigma: 1, Nu: float64(len(x) - 2)}
	tValue := tDist.Quantile(0.975) // 95% confidence interval

	interceptCI := [2]float64{alpha - tValue*standardError, alpha + tValue*standardError}
	slopeCI := [2]float64{beta - tValue*standardError, beta + tValue*standardError}

	return interceptCI, slopeCI
}

func calculateStandardError(residuals []float64) float64 {
	return math.Sqrt(sumOfSquares(residuals) / float64(len(residuals)-2))
}

func calculateDurbinWatson(residuals []float64) float64 {
	var sumDiffs, sumSquares float64
	for i := 1; i < len(residuals); i++ {
		sumDiffs += math.Pow(residuals[i]-residuals[i-1], 2)
		sumSquares += math.Pow(residuals[i], 2)
	}
	return sumDiffs / sumSquares
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
	fmt.Printf("Intercept: %.4f (95%% CI: [%.4f, %.4f])\n", results.Intercept, results.InterceptCI[0], results.InterceptCI[1])
	fmt.Printf("Slope: %.4f (95%% CI: [%.4f, %.4f])\n", results.Slope, results.SlopeCI[0], results.SlopeCI[1])
	fmt.Printf("R-squared: %.4f\n", results.RSquared)
	fmt.Printf("Adjusted R-squared: %.4f\n", results.AdjustedRSquared)
	fmt.Printf("RMSE: %.4f\n", results.RMSE)
	fmt.Printf("F-statistic: %.4f\n", results.FStatistic)
	fmt.Printf("t-statistic: %.4f\n", results.TStatistic)
	fmt.Printf("p-value: %.4f\n", results.PValue)
	fmt.Printf("Standard Error: %.4f\n", results.StandardError)
	fmt.Printf("Durbin-Watson: %.4f\n\n", results.DurbinWatson)
}

// Function to plot the data and regression line
func plotData(setNumber int, x, y []float64, results RegressionResults) error {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Anscombe Dataset %d", setNumber)
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
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
		return fmt.Errorf("could not create scatter plot: %w", err)
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

	if err := p.Save(6*vg.Inch, 6*vg.Inch, fmt.Sprintf("anscombe_%d.png", setNumber)); err != nil {
		return fmt.Errorf("could not save plot: %w", err)
	}
	fmt.Printf("Plot for dataset %d saved as anscombe_%d.png\n", setNumber, setNumber)
	return nil
}
