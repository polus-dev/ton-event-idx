package mmath

import (
	"golang.org/x/exp/constraints"
)

func MaxInteger[T constraints.Integer](x, y T) T {
	if x < y {
		return y
	}
	return x
}

func CalcAvgInteger[T constraints.Integer](nums []T) T {
	var total T = 0
	for _, number := range nums {
		total = total + number
	}

	return total / T(len(nums))
}
