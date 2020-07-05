package utils

// 传入两个数，传出小数和大数
func Swap(a, b int) (min, max int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return min, max
}
