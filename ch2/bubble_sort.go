package main

import "fmt"

func main()  {
	a := []int{12, 33, 44, 22, 1, -2, -4, 321, 22}
	fmt.Println(bubbleSort(a))
}

// 冒泡排序
func bubbleSort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	return a
}
