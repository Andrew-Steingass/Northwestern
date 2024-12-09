
# GoLang_Package

A Go package for calculating the trimmed mean of a slice of floats. This package provides flexible trimming capabilities, supporting symmetric and asymmetric trimming proportions.

## Installation

To use this package in your Go program, follow these steps:

1. **Install the Package**:
   Run the following command to add the package to your project:
   ```bash
   go get github.com/Andrew-Steingass/GoLang_Package/trimmean
   ```

2. **Import the Package**:
   Add the following import statement to your Go program:
   ```go
   import "github.com/Andrew-Steingass/GoLang_Package/trimmean"
   ```

3. **Use the TrimmedMean Function**:
   Example usage:
   ```go
   package main

   import (
       "fmt"
       "github.com/Andrew-Steingass/GoLang_Package/trimmean"
   )

   func main() {
       data := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1}
       result := trimmean.TrimmedMean(data, 0.1)
       fmt.Printf("Trimmed Mean: %.2f\n", result)
   }
   ```

## Comparing Results with R

I chose to compare hard coded samples between R and Go. On the last assignment, I had problems because I went a different route. This will allow a straight comparison between both methods, R and Go. Both methods produce the same results.
