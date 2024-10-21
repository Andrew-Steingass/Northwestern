package main

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Anscombe dataset for four different sets of (x, y)
var anscombe = map[string][]float64{
	"x1": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x2": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x3": {10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	"x4": {8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	"y1": {8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	"y2": {9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
	"y3": {7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	"y4": {6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
}

// Perform linear regression and return slope and intercept
func linearRegression(x, y []float64) (float64, float64) {
	slope, intercept := stat.LinearRegression(x, y, nil, false)
	return slope, intercept
}

func main() {
	// Perform linear regression for each dataset
	datasets := []struct {
		x, y string
		name string
	}{
		{"x1", "y1", "Dataset 1"},
		{"x2", "y2", "Dataset 2"},
		{"x3", "y3", "Dataset 3"},
		{"x4", "y4", "Dataset 4"},
	}

	for _, data := range datasets {
		slope, intercept := linearRegression(anscombe[data.x], anscombe[data.y])
		fmt.Printf("%s: Slope = %.3f, Intercept = %.3f\n", data.name, slope, intercept)
	}

	// Generate scatter plots for each dataset
	plotAnscombe()
}

// Generate scatter plots for the Anscombe dataset
func plotAnscombe() {
	datasets := []struct {
		x, y  string
		title string
	}{
		{"x1", "y1", "Set I"},
		{"x2", "y2", "Set II"},
		{"x3", "y3", "Set III"},
		{"x4", "y4", "Set IV"},
	}

	for _, data := range datasets {
		pts := make(plotter.XYs, len(anscombe[data.x]))
		for i := range pts {
			pts[i].X = anscombe[data.x][i]
			pts[i].Y = anscombe[data.y][i]
		}

		plotAnscombeData(data.title, pts, fmt.Sprintf("anscombe_%s.png", data.title))
	}
}

// Plot data and save it as an image
func plotAnscombeData(title string, pts plotter.XYs, filename string) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	s, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	p.Add(s)

	p.X.Min = 2
	p.X.Max = 20
	p.Y.Min = 2
	p.Y.Max = 14

	if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
		panic(err)
	}
}

// Benchmark for linear regression on dataset 1
func BenchmarkLinearRegressionDataset1(b *testing.B) {
	x := anscombe["x1"]
	y := anscombe["y1"]
	for i := 0; i < b.N; i++ {
		linearRegression(x, y)
	}
}

// Benchmark for linear regression on dataset 2
func BenchmarkLinearRegressionDataset2(b *testing.B) {
	x := anscombe["x2"]
	y := anscombe["y2"]
	for i := 0; i < b.N; i++ {
		linearRegression(x, y)
	}
}

// Benchmark for linear regression on dataset 3
func BenchmarkLinearRegressionDataset3(b *testing.B) {
	x := anscombe["x3"]
	y := anscombe["y3"]
	for i := 0; i < b.N; i++ {
		linearRegression(x, y)
	}
}

// Benchmark for linear regression on dataset 4
func BenchmarkLinearRegressionDataset4(b *testing.B) {
	x := anscombe["x4"]
	y := anscombe["y4"]
	for i := 0; i < b.N; i++ {
		linearRegression(x, y)
	}
}
