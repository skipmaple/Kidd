package main

// 个位数统计
// page 60

func numCount(n int) map[int]int {
	res := make(map[int]int)
	for n!=0 {
		temp := n%10
		n /= 10
		res[temp]++
	}
	return res
}
