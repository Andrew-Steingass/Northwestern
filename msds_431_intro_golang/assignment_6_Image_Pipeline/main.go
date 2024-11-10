// main.go
package main

import (
	"flag"
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"log"
	"path/filepath"
	"sync"
	"time"
)

type Result struct {
	Path    string
	Success bool
	Error   error
}

// processOneImage handles a single image processing pipeline
func processOneImage(path string) Result {
	result := Result{
		Path:    path,
		Success: false,
	}

	// Load image
	img, err := imageprocessing.ReadImage(path)
	if err != nil {
		result.Error = err
		return result
	}

	// Resize
	img, err = imageprocessing.Resize(img)
	if err != nil {
		result.Error = err
		return result
	}

	// Convert to grayscale
	img, err = imageprocessing.Grayscale(img)
	if err != nil {
		result.Error = err
		return result
	}

	// Save processed image
	outPath := filepath.Join("images/output", filepath.Base(path))
	err = imageprocessing.WriteImage(outPath, img)
	if err != nil {
		result.Error = err
		return result
	}

	result.Success = true
	return result
}

// Sequential Processing
func processImagesSequential(paths []string) ([]Result, time.Duration) {
	startTime := time.Now()
	results := make([]Result, len(paths))

	for i, path := range paths {
		results[i] = processOneImage(path)
	}

	return results, time.Since(startTime)
}

// Concurrent Processing
func processImagesConcurrent(paths []string) ([]Result, time.Duration) {
	startTime := time.Now()
	var wg sync.WaitGroup
	results := make([]Result, len(paths))

	// Process each image in its own goroutine
	for i, path := range paths {
		wg.Add(1)
		go func(index int, imagePath string) {
			defer wg.Done()
			results[index] = processOneImage(imagePath)
		}(i, path)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return results, time.Since(startTime)
}

func main() {
	// Flag for choosing processing mode
	mode := flag.String("mode", "both", "Processing mode: concurrent, sequential, or both")
	flag.Parse()

	// List of images to process
	imagePaths := []string{
		"images/cat1.jpg",
		"images/cat2.jpg",
		"images/cat3.jpg",
		"images/cat4.jpg",
	}

	fmt.Printf("\nProcessing %d images...\n\n", len(imagePaths))

	switch *mode {
	case "concurrent":
		results, duration := processImagesConcurrent(imagePaths)
		fmt.Printf("Concurrent Processing Time: %v\n", duration)
		printResults("Concurrent", results)

	case "sequential":
		results, duration := processImagesSequential(imagePaths)
		fmt.Printf("Sequential Processing Time: %v\n", duration)
		printResults("Sequential", results)

	case "both":
		// Run sequential first
		seqResults, seqDuration := processImagesSequential(imagePaths)

		fmt.Println("\n=== Sequential Processing ===")
		fmt.Printf("Time taken: %v\n", seqDuration)
		printResults("Sequential", seqResults)

		fmt.Println("\n=== Concurrent Processing ===")
		concResults, concDuration := processImagesConcurrent(imagePaths)
		fmt.Printf("Time taken: %v\n", concDuration)
		printResults("Concurrent", concResults)

		// Show comparison
		fmt.Println("\n=== Performance Comparison ===")
		fmt.Printf("Sequential Time: %v\n", seqDuration)
		fmt.Printf("Concurrent Time: %v\n", concDuration)
		speedup := float64(seqDuration) / float64(concDuration)
		fmt.Printf("Speedup: %.2fx faster with concurrent processing\n", speedup)

	default:
		log.Fatalf("Invalid mode: %s. Use 'concurrent', 'sequential', or 'both'", *mode)
	}
}

func printResults(mode string, results []Result) {
	successful := 0
	failed := 0
	for _, result := range results {
		if result.Success {
			successful++
		} else {
			failed++
			fmt.Printf("Failed to process %s: %v\n", result.Path, result.Error)
		}
	}
	fmt.Printf("Successfully processed: %d\n", successful)
	fmt.Printf("Failed to process: %d\n", failed)
}
