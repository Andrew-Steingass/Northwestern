package main

import (
	"fmt"

	"github.com/Andrew-Steingass/GoLang_Package/trimmean" // Import your package
)

func main() {
	// Hardcoded datasets
	dataIntegers := []float64{
		23, 56, 87, 12, 44, 98, 65, 32, 78, 90, 34, 45, 67, 89, 12, 34, 56, 78, 90, 123,
		145, 167, 189, 201, 223, 245, 267, 289, 310, 332, 354, 376, 398, 420, 442, 464,
		486, 508, 530, 552, 574, 596, 618, 640, 662, 684, 706, 728, 750, 772, 794, 816,
		838, 860, 882, 904, 926, 948, 970, 992, 1014, 1036, 1058, 1080, 1102, 1124,
		1146, 1168, 1190, 1212, 1234, 1256, 1278, 1300, 1322, 1344, 1366, 1388, 1410,
		1432, 1454, 1476, 1498, 1520, 1542, 1564, 1586, 1608, 1630, 1652, 1674, 1696,
		1718, 1740, 1762, 1784, 1806, 1828, 1850,
	}

	dataFloats := []float64{
		1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.2, 12.3, 13.4, 14.5, 15.6,
		16.7, 17.8, 18.9, 19.0, 20.1, 21.2, 22.3, 23.4, 24.5, 25.6, 26.7, 27.8, 28.9,
		29.0, 30.1, 31.2, 32.3, 33.4, 34.5, 35.6, 36.7, 37.8, 38.9, 39.0, 40.1, 41.2,
		42.3, 43.4, 44.5, 45.6, 46.7, 47.8, 48.9, 49.0, 50.1, 51.2, 52.3, 53.4, 54.5,
		55.6, 56.7, 57.8, 58.9, 59.0, 60.1, 61.2, 62.3, 63.4, 64.5, 65.6, 66.7, 67.8,
		68.9, 69.0, 70.1, 71.2, 72.3, 73.4, 74.5, 75.6, 76.7, 77.8, 78.9, 79.0, 80.1,
		81.2, 82.3, 83.4, 84.5, 85.6, 86.7, 87.8, 88.9, 89.0, 90.1, 91.2, 92.3, 93.4,
		94.5, 95.6, 96.7, 97.8, 98.9, 99.0,
	}

	// Calculate the trimmed mean with symmetric trimming (5%)
	trimmedMeanIntegers := trimmean.TrimmedMean(dataIntegers, 0.05)
	trimmedMeanFloats := trimmean.TrimmedMean(dataFloats, 0.05)

	// Round the results
	roundedIntegers := trimmean.Round(trimmedMeanIntegers, 4)
	roundedFloats := trimmean.Round(trimmedMeanFloats, 5)

	// Print results
	fmt.Printf("Trimmed Mean of Integers: %.4f\n", roundedIntegers)
	fmt.Printf("Trimmed Mean of Floats: %.5f\n", roundedFloats)
}
