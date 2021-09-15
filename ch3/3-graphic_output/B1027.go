// 打印沙漏
// page 44

package main

import (
	"fmt"
	"math"
)

func main() {
	//printShaLou(19, "*")
	//printShaLou(55, "*")
	printShaLou(100, "*")
}

func printShaLou(count int, ch string) {
	res, times, rest := shaLou(count)
	for i := 0; i < len(res); i++ {
		// 打印空格
		if i < times {
			for k := 0; k < i; k++ {
				fmt.Printf(" ")
			}
		} else {
			// i > times
			for k := times - (i - times); k > 0; k-- {
				fmt.Printf(" ")
			}
		}
		// 打印字符
		for j := 0; j < res[i]; j++ {
			fmt.Printf("%s", ch)
		}
		// 换行
		fmt.Printf("\n")
	}
	fmt.Println(rest)
}

func shaLou(count int) (res []int, times, rest int) {
	for i := 0; ; i++ {
		if composeShaLou(i) < count && composeShaLou(i+1) > count {
			times = i
			rest = count - composeShaLou(times)
			break
		}
	}

	for i := -times; i <= times; i++ {
		temp := int(math.Abs(float64(i)))
		lineCount := 2*temp + 1
		res = append(res, lineCount)
	}

	return res, times, rest

}

// 给定i次，给出组成沙漏的字符总数
func composeShaLou(i int) (count int) {
	if i == 0 {
		return 1
	}
	return composeShaLou(i-1) + 4*i + 2
}
