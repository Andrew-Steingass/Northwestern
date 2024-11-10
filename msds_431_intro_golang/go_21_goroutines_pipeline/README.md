# Image Processing Pipeline

A Go program that demonstrates concurrent vs sequential image processing using goroutines.

## Features
- Resizes images to 500x500 pixels
- Converts images to grayscale
- Processes images both sequentially and concurrently
- Compares processing performance

## Setup
```bash
go get github.com/nfnt/resize # Required external package for image resizing
```

## Directory Structure
```
.
├── image_processing/          # Core processing functions and tests
├── main.go                   # Main program
├── images/                   # Input images
└── images/output/            # Processed images
```

## Run Processing Tests
```bash
# Compare both processing methods
go run main.go -mode=both

# Test sequential processing only
go run main.go -mode=sequential

# Test concurrent processing only
go run main.go -mode=concurrent
```

Sample Output:
```
Processing 4 images...

=== Sequential Processing ===
Time taken: 2.5s
Successfully processed: 4

=== Concurrent Processing ===
Time taken: 0.8s
Successfully processed: 4

=== Performance Comparison ===
Speedup: 3.13x faster with concurrent processing
```

## Run Unit Tests
```bash
# Run all tests
go test ./image_processing

# Run tests with verbose output
go test -v ./image_processing

# Run a specific test
go test -v ./image_processing -run TestReadImage
```