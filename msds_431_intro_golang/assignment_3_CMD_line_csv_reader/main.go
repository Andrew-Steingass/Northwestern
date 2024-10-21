package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func main() {
	// Define flags for the input CSV file and the output JSON lines file
	// NOTE flag is a structured argument.
	//		(defined variable in cmd line, user input, user help message if needed)
	inputFilePath := flag.String("input", "", "Path to the input CSV file")
	outputFilePath := flag.String("output", "", "Path to the output JSON lines file")
	fmt.Println("")
	// Needed to pretty much load the input variable correctly. NOTE is good for all flag's above
	flag.Parse()

	// Check the inptus are valid
	err := valid_inputs(*inputFilePath, *outputFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("User Inputs are valid, procceeding to load CSV")

	records, err := readCSV(*inputFilePath)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}
	fmt.Println("CSV successfully loaded, length = ", len(records))

	err = writeJSONL(records, *outputFilePath)
	if err != nil {
		fmt.Println("Error writing JSONL:", err)
		os.Exit(1)
	}
	fmt.Println("Conversion completed successfully.")
	fmt.Println("New file saved to:", *outputFilePath)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func valid_inputs(inputFilePath string, outputFilePath string) error {
	// GOAL:
	//       1. If no input for one of the fields, print usage data and stop exe
	//		 2. If no input file, throw an error and stop exe
	//		 3. If no output directory, throw and error and stop exe

	// Capture usage information
	var usageInfo strings.Builder
	flag.CommandLine.SetOutput(&usageInfo)
	flag.Usage()
	usage := usageInfo.String()

	// Validate that the user has some sort of input for both flags
	if inputFilePath == "" || outputFilePath == "" {
		return fmt.Errorf("Error: both input and output file paths are required.\n\n%s", usage)
	}

	// Check if the input file exists
	_, err := os.Stat(inputFilePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("Error: input file '%s' does not exist.", inputFilePath)
	}

	// Check if the output file directory exists
	dir := filepath.Dir(outputFilePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("Error: directory for output file does not exist: %s", dir)
	}

	return nil
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func readCSV(filePath string) ([][]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV: %v", err)
	}

	return records, nil
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func writeJSONL(records [][]string, outputPath string) error {
	// Create the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer file.Close()

	// Get the headers from the first row
	headers := records[0]
	fmt.Println("Found Headers:", headers)

	// Create a buffered writer for better performance
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Process each record (skipping the header row)
	for _, record := range records[1:] {
		// Create a map to hold the data for this record
		obj := make(map[string]string)

		// Populate the map with header:value pairs
		for i, value := range record {
			if i < len(headers) { // Ensure we don't go out of bounds
				obj[strings.TrimSpace(headers[i])] = strings.TrimSpace(value)
			}
		}

		// Marshal the map to JSON
		jsonData, err := json.Marshal(obj)
		if err != nil {
			return fmt.Errorf("error marshaling JSON: %v", err)
		}

		// Extra validation step (optional, but provides additional certainty)
		var tempObj map[string]interface{}
		if err := json.Unmarshal(jsonData, &tempObj); err != nil {
			return fmt.Errorf("invalid JSON produced: %v", err)
		}

		// Write the JSON data to the file
		_, err = writer.Write(jsonData)
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}

		// Write a newline character
		_, err = writer.WriteString("\n")
		if err != nil {
			return fmt.Errorf("error writing newline to file: %v", err)
		}
	}

	return nil
}
