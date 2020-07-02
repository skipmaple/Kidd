package main

import "fmt"

// 查找元素
// 输入一个数n，然后输入n个数值各不相同的数，再输入一个值x，输出这个值在数组中的下标(从0开始，如果不存在，则输出-1)

// example

// input:
// 4
// 1 2 3 4
// 3

// output:
// 2

func main() {
	var len, num, foundNum int
	var arr []int
	fmt.Println("请输入数组长度：")
	_, _ = fmt.Scanln(&len)
	for i := 0; i < len; i++ {
		fmt.Printf("请输入第%d个元素:", i+1)
		_, _ = fmt.Scanln(&num)
		arr = append(arr, num)
	}

	fmt.Println("请输入要查找的元素:")
	_, _ = fmt.Scanln(&foundNum)
	for i, n := range arr {
		if n == foundNum {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(-1)
}
