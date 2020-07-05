package main

import (
	"bytes"
	"fmt"
)

// 数字分类

func main() {
	var count int
	arr := make([]int, 0)
	_, _ = fmt.Scanf("%d", &count)
	for i := 0; i < count; i++ {
		var e int
		_, _ = fmt.Scanf("%d", &e)
		arr = append(arr, e)
	}

	classify := make([]int, 5)
	classifyCount := make([]int, 5)
	resClassify := make(map[int]interface{}, 5)

	for _, n := range arr {
		if n%5 == 0 && n%2 == 0 {
			classify[0] += n
			classifyCount[0]++
		}

		if n%5 == 1 {
			if classifyCount[1]%2 == 0 {
				classify[1] += n
			} else {
				classify[1] -= n
			}
			classifyCount[1]++
		}

		if n%5 == 2 {
			classify[2]++
			classifyCount[2]++
		}

		if n%5 == 3 {
			classify[3] += n
			classifyCount[3]++
		}

		if n%5 == 4 {
			if classify[4] < n {
				//fmt.Printf("num5 is %d\n", classify[4])
				classify[4] = n
				//fmt.Printf("num5 is %d, n is %d\n", classify[4], n)
				classifyCount[4]++
			}
		}
	}

	for i, n := range classifyCount {
		if n == 0 {
			resClassify[i] = "N"
		} else if i == 3 {
			temp := new(bytes.Buffer)
			_, _ = fmt.Fprintf(temp, "%.1f", float64(classify[3]) / float64(classifyCount[3]))
			resClassify[3] = temp
		} else {
			resClassify[i] = classify[i]
		}
	}

	fmt.Printf("%v %v %v %v %v", resClassify[0], resClassify[1], resClassify[2], resClassify[3], resClassify[4])
}
