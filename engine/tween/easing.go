package tween

import "math"

func EaseIn(t float64) float64 {
	return t * t
}

func EaseInSine(t float64) float64 {
	return 1 - math.Cos((t*math.Pi)/2)
}

func EaseInOutCubic(t float64) float64 {
	if t < 0.5 {
		return 4 * t * t * t
	}

	return 1 - math.Pow(-2*t+2, 4)/2
}
