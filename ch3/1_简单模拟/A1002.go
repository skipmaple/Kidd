package main

import "fmt"

// A+B for Polynomials
// 求两个多项式相加的结果
// page 25

// testCase1
var a = []float64{2, 1, 2.4, 0, 3.2}
var b = []float64{2, 2, 1.5, 1, 0.5}

// testCase2
//var a = []float64{3, 1, 2.4, 0, 3.2, 4, 4.5}
//var b = []float64{4, 2, 1.5, 1, 0.5, 4, 22, 3, 11}

func main() {
	aLen := a[0]
	bLen := b[0]

	res := make(map[int]float64)

	for i := 1; i <= int(aLen); i++ {
		res[int(a[2*i-1])] += a[2*i]
	}

	for i := 1; i <= int(bLen); i++ {
		res[int(b[2*i-1])] += b[2*i]
	}

	fmt.Printf("%d ", len(res))
	// 所有指数组成 keys数组
	keys := make([]int, 0)
	for k, _ := range res {
		keys = append(keys, k)
	}

	// 对keys从大到小排序
	for i := 0; i < len(keys)-1; i++ {
		for j := 0; j < len(keys)-1-i; j++ {
			if keys[j] < keys[j+1] {
				keys[j], keys[j+1] = keys[j+1], keys[j]
			}
		}
	}

	for i := 0; i < len(keys); i++ {
		if i != len(keys)-1 {
			fmt.Printf("%d %v ", keys[i], res[keys[i]])
		} else {
			fmt.Printf("%d %v\n", keys[i], res[keys[i]])
		}
	}
}
