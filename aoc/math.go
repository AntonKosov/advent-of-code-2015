package aoc

import "math"

func Max(a ...int) int {
	max := a[0]
	for i := 1; i < len(a); i++ {
		v := a[i]
		if v > max {
			max = v
		}
	}
	return max
}

func Min(a ...int) int {
	min := a[0]
	for i := 1; i < len(a); i++ {
		v := a[i]
		if v < min {
			min = v
		}
	}
	return min
}

func Abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func MinMax(data []int) (min, max int) {
	if len(data) == 0 {
		panic("No data")
	}
	min = math.MaxInt
	max = math.MinInt
	for _, v := range data {
		min = Min(min, v)
		max = Max(max, v)
	}
	return min, max
}

func Sqrt(v int) int {
	return int(math.Sqrt(float64(v)))
}
