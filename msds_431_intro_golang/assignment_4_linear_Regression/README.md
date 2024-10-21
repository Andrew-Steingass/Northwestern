
# Anscombe Regression Analysis in Go

## Overview
This project implements linear regression on Anscombe's Quartet using Go. It compares regression results with Python/R and benchmarks Go's performance.

## Requirements
- Go installed ([download here](https://golang.org/dl/))
- Install dependencies:
  ```bash
        go get -u gonum.org/v1/gonum/...
        go get gonum.org/v1/gonum/stat
        go get gonum.org/v1/plot/...
  ```

## Running the Program
To run the regression on all datasets (`x1`, `x2`, `x3`, `x4`):
```bash
go run main.go
```
This prints results including intercept, slope, R-squared, RMSE, F-statistic, t-statistic, and p-value.

## Running Unit Tests
To verify regression results:
```bash
go test
```

## Running Benchmarks
To benchmark the performance:
```bash
go test -bench=.
```



## Recommendation to Management
Using Go for regression analysis showed comparable results to Python and R, with slight differences due to rounding and floating-point precision. Go's execution time was faster, making it suitable for performance-critical tasks.

However, data scientists may find Go's statistical ecosystem less mature compared to Python's `pandas`, `statsmodels`, and R's vast library of statistical packages. Python and R have richer support for advanced statistical modeling and better community support. For routine statistical tasks, Python or R is recommended, but Go can be an efficient choice for high-performance applications. If there is a need for heavy processing, Go is recommended

## Challenges Faced
During development, some differences were noted in how calculations like regression coefficients, p-values, and rounding were handled between Go and Python/R. These differences stem from how floating-point precision is handled across languages and their respective libraries. Careful handling of these discrepancies was necessary to ensure comparable results across different platforms.