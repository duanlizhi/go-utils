package utils

import (
	"math"
)

func MinInt(a, b int) int {
	// 拉拉
	return int(math.Min(float64(a), float64(b)))
}
