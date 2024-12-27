package engine

import "math"

func Lerp(start, end, t float64) float64 {
	return start + (end-start)*Clamp01(t)
}

func Clamp01(t float64) float64 {
	return math.Min(1, math.Max(0, t))
}
