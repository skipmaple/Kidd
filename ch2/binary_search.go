package ch2

// 二分查找
// main page 127

// 给定一个从小到大排序的数组arr，查找元素x的位置，返回x在arr中的下标
func binarySearch(arr []int, x int) int {
	left, right := 0, len(arr)
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] == x {
			return middle
		} else if arr[middle] > x {
			right = middle - 1
		} else {
			left = middle+1
		}
	}
	return -1 // 表示x不存在于arr中
}
