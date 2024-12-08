package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile = flag.String("memprofile", "", "write memory profile to file")
)

// Load CSV and skip the first row (header)
func loadCSV(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", fileName, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file %s: %w", fileName, err)
	}

	// Skip header row if exists
	if len(records) > 0 && len(records[0]) > 0 {
		first := records[0][0]
		_, err := strconv.ParseFloat(first, 64)
		if err != nil {
			return records[1:], nil
		}
	}
	return records, nil
}

// Load data
func loadData(fileName string) []float64 {
	records, err := loadCSV(fileName)
	if err != nil {
		ErrorLogger.Fatalf("Error loading data from %s: %v", fileName, err)
	}

	data := make([]float64, len(records))
	for i, row := range records {
		val, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			ErrorLogger.Fatalf("Failed to parse float in %s: %v", fileName, err)
		}
		data[i] = val
	}
	DebugLogger.Printf("Loaded %d data points from %s", len(data), fileName)
	return data
}

// Load resampling indices
func loadIndices(fileName string) [][]int {
	records, err := loadCSV(fileName)
	if err != nil {
		ErrorLogger.Fatalf("Error loading indices: %v", err)
	}

	numRows := len(records)

	// Create indices in column-wise format to match R
	indices := make([][]int, 100) // B=100 bootstrap samples
	for i := 0; i < 100; i++ {
		indices[i] = make([]int, numRows)
		for j := 0; j < numRows; j++ {
			val, err := strconv.Atoi(records[j][i]) // Note the [j][i] order
			if err != nil {
				ErrorLogger.Fatalf("Failed to parse int: %v", err)
			}
			indices[i][j] = val - 1 // Convert to 0-based
		}
	}
	DebugLogger.Printf("Loaded %dx%d indices from %s", len(indices), numRows, fileName)
	return indices
}

// Compute median
func Median(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sortedData := append([]float64(nil), data...)
	sort.Float64s(sortedData)
	n := len(sortedData)

	// R's specific median calculation
	if n%2 == 0 {
		j := n / 2
		i := j - 1
		return (sortedData[i] + sortedData[j]) / 2.0
	} else {
		i := n / 2
		return sortedData[i]
	}
}

// Bootstrap medians
func bootstrapMedians(data []float64, indices [][]int) []float64 {
	DebugLogger.Printf("Starting bootstrap median calculation with %d samples", len(indices))
	medians := make([]float64, len(indices))
	for i, resampleIndices := range indices {
		resample := make([]float64, len(resampleIndices))
		for j, idx := range resampleIndices {
			resample[j] = data[idx]
		}
		medians[i] = Median(resample)
	}
	return medians
}

// Calculate standard error
func calculateSE(medians []float64) float64 {
	if len(medians) == 0 {
		return 0
	}
	mean := 0.0
	nf := float64(len(medians))
	for _, v := range medians {
		mean += v
	}
	mean /= nf

	sumSqDiff := 0.0
	for _, v := range medians {
		diff := v - mean
		sumSqDiff += diff * diff
	}
	return math.Sqrt(sumSqDiff / (nf - 1.0))
}

// Save to CSV
func saveToCSV(fileName string, header []string, rows [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", fileName, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header to file %s: %w", fileName, err)
	}

	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("failed to write row to file %s: %w", fileName, err)
		}
	}
	return nil
}

// Load R medians for comparison
func loadRMedians(fileName string) []float64 {
	records, err := loadCSV(fileName)
	if err != nil {
		ErrorLogger.Fatalf("Error loading R medians: %v", err)
	}

	medians := make([]float64, len(records))
	for i, row := range records {
		val, err := strconv.ParseFloat(row[0], 64)
		if err != nil {
			ErrorLogger.Fatalf("Failed to parse R median: %v", err)
		}
		medians[i] = val
	}
	return medians
}

// Compare medians function
func compareMedians(goMedians, rMedians []float64) {
	fmt.Println("\nComparing first 5 bootstrap medians:")
	fmt.Printf("%-15s %-15s %-15s %-15s\n", "Index", "Go Median", "R Median", "Diff")
	for i := 0; i < 5 && i < len(goMedians); i++ {
		diff := math.Abs(goMedians[i] - rMedians[i])
		fmt.Printf("%-15d %-15.6f %-15.6f %-15.6f\n", i, goMedians[i], rMedians[i], diff)
	}
}

func debugMedian(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	DebugLogger.Printf("Debug Median Calculation for %d values", len(data))
	fmt.Println("\nDebug Median Calculation:")
	fmt.Printf("Input data length: %d\n", len(data))

	sortedData := append([]float64(nil), data...)
	sort.Float64s(sortedData)
	n := len(sortedData)

	fmt.Printf("First 5 sorted values: %v\n", sortedData[:5])
	fmt.Printf("Last 5 sorted values: %v\n", sortedData[len(sortedData)-5:])

	var result float64
	if n%2 == 0 {
		j := n / 2
		i := j - 1
		fmt.Printf("Even length %d, using average of positions %d (%.6f) and %d (%.6f)\n",
			n, i, sortedData[i], j, sortedData[j])
		result = (sortedData[i] + sortedData[j]) / 2.0
	} else {
		i := n / 2
		fmt.Printf("Odd length %d, using position %d: %.6f\n",
			n, i, sortedData[i])
		result = sortedData[i]
	}

	fmt.Printf("Calculated median: %.6f\n", result)
	return result
}

func debugBootstrapMedians(data []float64, indices [][]int) []float64 {
	medians := make([]float64, len(indices))
	debugLimit := 2

	for i := 0; i < len(indices); i++ {
		resample := make([]float64, len(indices[i]))

		if i < debugLimit {
			fmt.Printf("\n=== Bootstrap Sample %d ===\n", i)
			fmt.Printf("Using indices column %d\n", i)
			fmt.Printf("First 5 indices: %v\n", indices[i][:5])

			for j := 0; j < len(indices[i]); j++ {
				resample[j] = data[indices[i][j]]
				if j < 5 {
					fmt.Printf("Value at index %d: %.6f\n", indices[i][j], resample[j])
				}
			}
			medians[i] = debugMedian(resample)
		} else {
			for j := 0; j < len(indices[i]); j++ {
				resample[j] = data[indices[i][j]]
			}
			medians[i] = Median(resample)
		}
	}
	return medians
}

func main() {
	// Initialize flags and logging
	flag.Parse()
	initLoggers()

	// CPU profiling
	// CPU profiling setup with more explicit error handling
	if *cpuprofile != "" {
		InfoLogger.Printf("Attempting to create CPU profile at: %s", *cpuprofile)
		f, err := os.Create(*cpuprofile)
		if err != nil {
			ErrorLogger.Fatalf("Could not create CPU profile: %v", err)
		}
		defer f.Close()
		InfoLogger.Println("CPU profile file created successfully")

		if err := pprof.StartCPUProfile(f); err != nil {
			ErrorLogger.Fatalf("Could not start CPU profile: %v", err)
		}
		defer pprof.StopCPUProfile()
		InfoLogger.Println("CPU profiling started successfully")
	}

	InfoLogger.Println("Starting bootstrap analysis")

	sampleSize := 25
	distributions := []string{"symmetric", "positively_skewed", "negatively_skewed"}

	var seRows [][]string
	seRows = append(seRows, []string{"Distribution", "SampleSize", "SEMedian"})

	for _, dist := range distributions {
		InfoLogger.Printf("Processing %s distribution", dist)
		fmt.Printf("\n=== Processing %s distribution ===\n", dist)

		dataFile := fmt.Sprintf("%s_data_%d.csv", dist, sampleSize)
		data := loadData(dataFile)
		fmt.Printf("Raw data (first 5): %v\n", data[:5])

		indicesFile := fmt.Sprintf("resampling_indices_%d.csv", sampleSize)
		indices := loadIndices(indicesFile)

		goMedians := debugBootstrapMedians(data, indices)
		rMedianFile := fmt.Sprintf("%s_medians_%d.csv", dist, sampleSize)
		rMedians := loadRMedians(rMedianFile)

		compareMedians(goMedians, rMedians)

		goSE := calculateSE(goMedians)
		rSE := calculateSE(rMedians)

		fmt.Printf("\nStandard Error Comparison:\n")
		fmt.Printf("Go SE: %.6f\n", goSE)
		fmt.Printf("R SE: %.6f\n", rSE)
		fmt.Printf("Difference: %.6f\n", math.Abs(goSE-rSE))

		seRows = append(seRows, []string{
			dist,
			fmt.Sprintf("%d", sampleSize),
			fmt.Sprintf("%.6f", goSE),
		})

		InfoLogger.Printf("Completed %s distribution analysis", dist)
	}

	for _, row := range seRows {
		fmt.Printf("%-20s %-10s %-15s\n", row[0], row[1], row[2])
	}

	InfoLogger.Println("Analysis completed successfully")

	// Memory profiling
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		InfoLogger.Printf("Writing memory profile to %s", *memprofile)
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	InfoLogger.Println("Analysis completed successfully")
}
