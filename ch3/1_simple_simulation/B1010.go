package main

import (
	"fmt"
)

// 一元多项式求导
// page 23

func main() {
	coefficient := make([]int, 0)
	exponent := make([]int, 0)

	for i, v := range testCase1() {
		if i%2 == 0 {
			coefficient = append(coefficient, v)
			exponent = append(exponent, testCase1()[i+1])
		}
	}

	// 对指数冒泡排序, 系数位置做到跟指数同步
	for i := 0; i < len(exponent)-1; i++ {
		for j := 0; j < len(exponent)-1-i; j++ {
			if exponent[j] < exponent[j+1] {
				exponent[j], exponent[j+1] = exponent[j+1], exponent[j]
				coefficient[j], coefficient[j+1] = coefficient[j+1], coefficient[j] // 系数同步指数位置
			}
		}
	}

	for i, v := range exponent {
		if v != 0 {
			if i != len(exponent)-1 {
				fmt.Printf("%d %d ", coefficient[i]*v, v-1)
			} else {
				fmt.Printf("%d %d\n", coefficient[i]*v, v-1)
			}
		}
	}

}

func testCase1() []int {
	return []int{3, 4, -5, 2, 6, 1, -2, 0}
}
