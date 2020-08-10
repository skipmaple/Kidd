package ch4

import (
	"math"
	"strconv"
)

// Radix
// page 167

func radix(a, b string, tag, radix int) string {
	radixMap := map[string]int{}
	for i := 0; i <= 9; i++ {
		radixMap[strconv.Itoa(i)] = i
	}

	for i := 'a'; i <= 'z'; i++ {
		radixMap[string(i)] = int(i) - 87
	}

	known, unknown := "", ""
	knownInt := 0
	if tag == 1 {
		known, unknown = a, b
	} else {
		known, unknown = b, a
	}

	for i := 0; i < len(b); i++ {
		knownInt = strMapToInt(radixMap, known, radix)
	}

	left, right := 0, 35
	for left <= right {
		middle := (left + right) / 2
		if knownInt == strMapToInt(radixMap, unknown, middle) {
			return strconv.Itoa(middle)
		} else if knownInt < strMapToInt(radixMap, unknown, middle) {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return "Impossible"
}

func strMapToInt(radixMap map[string]int, str string, radix int) (res int) {
	for i := len(str)-1; i >= 0; i-- {
		res += radixMap[string(str[i])] * int(math.Pow(float64(radix), float64(len(str)-1-i)))
	}
	return res
}
