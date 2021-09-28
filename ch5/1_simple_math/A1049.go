package __simple_math

import (
	"strconv"
)

// page 198
// Counting Ones

func countingOnes(input int) (res int) {
	for i := 1; i <= input; i++ {
		temp := []byte(strconv.Itoa(i))
		for _, e := range temp {
			if e == '1' {
				res += 1
			}
		}
		//fmt.Printf("%s\n", temp)
	}
	return res
}
