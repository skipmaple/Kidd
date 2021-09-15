package main

import (
	"fmt"
)

// 数组元素循环右移问题

// 输入两个数n m 分别代表数组长度 和 移动m个位置
// 输出移动后的数组

func main() {
	var n, m int
	arr := make([]int, 0)
	var e int

	_, _ = fmt.Scanf("%d %d", &n, &m)

	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &e)
		arr = append(arr, e)
	}

	if m >= n {
		m %= n
	}

	if m == 0 {
		fmt.Printf("%v", arr)
	} else {
		for i := n - m; i < n; i++ {
			fmt.Printf("%d ", arr[i])
		}
		for i := 0; i < n-m; i++ {
			if i != n-m {
				fmt.Printf("%d ", arr[i])
			} else {
				fmt.Printf("%d", arr[i])
			}
		}
	}

}
