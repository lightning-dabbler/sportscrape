package util

import "math"

// Round
//
// Parameter:
//   - val: The number to round
//   - x: Decimal places
//
// Returns a rounded float64 number
func Round(val float64, x int) float64 {
	factor := math.Pow(10, float64(x))
	return math.Round(val*factor) / factor
}
