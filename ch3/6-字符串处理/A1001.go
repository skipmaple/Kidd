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
	for i := 0; i < len(sumStr); i++ {
		res = fmt.Sprintf("%s%s", res, string(sumStr[i]))

		if (i+1)%3 == 0 && i != len(sumStr)-1 {
			if sum > 0 {
				res = fmt.Sprintf("%s%s", ",", res)
			} else {
				if i != len(sumStr)-2 {
					res = fmt.Sprintf("%s%s", ",", res)
				}
			}
		}
	}

	return res
}
