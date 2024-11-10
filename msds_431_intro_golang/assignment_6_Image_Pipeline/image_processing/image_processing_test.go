// image_processing/image_processing_test.go
package imageprocessing

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"testing"
)

// createTestImage creates a simple test image with a solid color
func createTestImage(width, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{R: 100, G: 200, B: 255, A: 255})
		}
	}
	return img
}

// TestReadImage checks if we can:
// 1. Read a valid image
// 2. Handle an empty path
// 3. Handle a non-existent image
func TestReadImage(t *testing.T) {
	// Create temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "image_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test image file
	testPath := filepath.Join(tmpDir, "test.jpg")
	testImg := createTestImage(100, 100)
	if err := WriteImage(testPath, testImg); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{"Valid image", testPath, false},
		{"Empty path", "", true},
		{"Non-existent image", "nonexistent.jpg", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadImage(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestWriteImage checks if we can:
// 1. Write an image successfully
// 2. Handle nil image
// 3. Handle empty path
func TestWriteImage(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "image_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	testImg := createTestImage(100, 100)

	tests := []struct {
		name    string
		path    string
		img     image.Image
		wantErr bool
	}{
		{"Valid image", filepath.Join(tmpDir, "test.jpg"), testImg, false},
		{"Nil image", filepath.Join(tmpDir, "nil.jpg"), nil, true},
		{"Empty path", "", testImg, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteImage(tt.path, tt.img)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestGrayscale checks if we can:
// 1. Convert an image to grayscale
// 2. Handle nil image
func TestGrayscale(t *testing.T) {
	tests := []struct {
		name    string
		img     image.Image
		wantErr bool
	}{
		{"Valid image", createTestImage(100, 100), false},
		{"Nil image", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gray, err := Grayscale(tt.img)
			if (err != nil) != tt.wantErr {
				t.Errorf("Grayscale() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && gray == nil {
				t.Error("Grayscale() returned nil image for valid input")
			}
		})
	}
}

// TestResize checks if we can:
// 1. Resize an image
// 2. Handle nil image
func TestResize(t *testing.T) {
	tests := []struct {
		name    string
		img     image.Image
		wantErr bool
	}{
		{"Valid image", createTestImage(100, 100), false},
		{"Nil image", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resized, err := Resize(tt.img)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resize() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && resized == nil {
				t.Error("Resize() returned nil image for valid input")
			}
		})
	}
}
