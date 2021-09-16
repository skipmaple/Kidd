package main

import (
	"fmt"
	"strconv"
)

// Product of Polynomials
// 两个多项式的乘积
// page 27

func main() {
	// testCase1
	var a = []float64{2, 1, 2.4, 0, 3.2}
	var b = []float64{2, 2, 1.5, 1, 0.5}

	aLen := a[0]
	bLen := b[0]
	res := make(map[int]float64)
	for i := 1; i <= int(aLen); i++ {
		for j := 1; j <= int(bLen); j++ {
			exp := int(a[2*i-1] + b[2*j-1])
			temp, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", a[2*i]*b[2*j]), 64)
			res[exp] += temp
			//fmt.Printf("a[%d]=%g, b[%d]=%g ", 2*i, a[2*i], 2*j, b[2*j])
			//fmt.Println("exp: ", exp, "value: ", temp)
		}
	}

	var keys []int
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

	fmt.Printf("%d ", len(res))

	for i := 0; i < len(keys); i++ {
		if i != len(keys)-1 {
			fmt.Printf("%d %.1f ", keys[i], res[keys[i]])
		} else {
			fmt.Printf("%d %.1f\n", keys[i], res[keys[i]])
		}
	}
}
