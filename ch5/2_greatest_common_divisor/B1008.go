package __greatest_common_divisor

import "Kidd/utils"

// 数组元素循环右移问题
// page 201
// https://pintia.cn/problem-sets/994805260223102976/problems/994805316250615808

func arrLoopRight(arr []int, m int) []int {
	n := len(arr)
	if m%n == 0 {
		return arr
	}
	d := utils.Gcd(n, m)
	// temp 临时变量, pos 存放当前处理位置, next 为下一个要处理的位置
	var temp, pos, next int

	for i := n - m; i < n-m+d; i++ {
		temp = arr[i]
		pos = i
		for {
			next = (pos - m + n) % n
			if next != i {
				arr[pos] = arr[next]
			} else {
				arr[pos] = temp
			}
			pos = next
			if pos == i {
				break
			}
		}
	}

	return arr
}
