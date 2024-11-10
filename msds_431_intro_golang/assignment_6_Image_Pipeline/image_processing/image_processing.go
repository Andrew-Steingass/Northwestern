// image_processing.go
package imageprocessing

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

// ReadImage reads an image from the given path and returns the image or an error
func ReadImage(path string) (image.Image, error) {
	if path == "" {
		return nil, errors.New("empty file path provided")
	}

	inputFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %w", err)
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image %s: %w", path, err)
	}

	return img, nil
}

// WriteImage writes the given image to the specified path
func WriteImage(path string, img image.Image) error {
	if img == nil {
		return errors.New("nil image provided")
	}
	if path == "" {
		return errors.New("empty output path provided")
	}

	// Ensure the output directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Get the file extension and convert to lowercase
	ext := strings.ToLower(filepath.Ext(path))

	// Encode based on file extension
	switch ext {
	case ".jpg", ".jpeg":
		if err := jpeg.Encode(outputFile, img, nil); err != nil {
			return fmt.Errorf("failed to encode jpeg: %w", err)
		}
	case ".png":
		if err := png.Encode(outputFile, img); err != nil {
			return fmt.Errorf("failed to encode png: %w", err)
		}
	default:
		return fmt.Errorf("unsupported image format: %s", ext)
	}

	return nil
}

// Grayscale converts the image to grayscale
func Grayscale(img image.Image) (image.Image, error) {
	if img == nil {
		return nil, errors.New("nil image provided")
	}

	bounds := img.Bounds()
	if bounds.Empty() {
		return nil, errors.New("image has no bounds")
	}

	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}

	return grayImg, nil
}

// Resize resizes the image to 500x500 pixels
func Resize(img image.Image) (image.Image, error) {
	if img == nil {
		return nil, errors.New("nil image provided")
	}

	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	if resizedImg == nil {
		return nil, errors.New("resize operation failed")
	}

	return resizedImg, nil
}
