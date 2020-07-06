// 换个格式输出整数
// page 59

package main

import (
	"fmt"
	"strconv"
)

func printAnotherInt(n int) (res string) {
	const Hundred string = "B"
	const Ten string = "S"

	for i := 0; n != 0; i++ {
		temp := n % 10
		n /= 10

		for j := 1; j <= temp; j++ {
			if i == 0 {
				res = fmt.Sprintf("%s%s", res, strconv.Itoa(j))
			} else if i == 1{
				res = fmt.Sprintf("%s%s", Ten, res)
			} else if i == 2 {
				res = fmt.Sprintf("%s%s", Hundred, res)
			}
		}
	}

	return res
}
