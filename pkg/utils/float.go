package utils

import "math"



// minFloat64 float类型精度
const minFloat64 = 0.000001

// Float64Equal 比较float64是否相等
func Float64Equal(a, b float64) bool {
	return math.Dim(a, b) < minFloat64
}
