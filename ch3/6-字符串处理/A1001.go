package main

import (
	"fmt"
	"strconv"
)

// A+B Format
// page 73

func sumFormat(a, b int) (res string) {
	sum := a + b
	sumStr := strconv.Itoa(sum)
	for i, j := len(sumStr)-1, 0; i >= 0; i, j = i-1, j+1 {
		res = fmt.Sprintf("%s%s", string(sumStr[i]), res)

		if j%3 == 2 {
			if sum > 0 {
				res = fmt.Sprintf("%s%s", ",", res)

			} else if j != len(sumStr)-2 {
				res = fmt.Sprintf("%s%s", ",", res)
			}
		}
	}

	return res
}
