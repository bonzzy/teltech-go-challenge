package core

import "math"

func Round(value float64, precision float64) float64 {
	return math.Round(value*math.Pow(10, precision)) / math.Pow(10, precision)
}
