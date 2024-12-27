package engine

import "math"

func Deg2Rad(degree float64) float64 {
	return degree * math.Pi / 180
}

func Rad2Deg(radian float64) float64 {
	return radian * 180 / math.Pi
}
