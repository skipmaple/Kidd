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
			left = middle + 1
		}
	}
	return -1 // 表示x不存在于arr中
}

// 求第一个大于等于x的元素的下标
func lowerBound(arr []int, x int) int {
	left, right := 0, len(arr) // ！！！传入的right是len arr
	for left < right {         // ！！！只需判断 < 不再是 <=
		middle := (left + right) / 2
		if arr[middle] >= x {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

// 求第一个大于x的元素的下标
func upperBound(arr []int, x int) int {
	left, right := 0, len(arr)
	for left < right {
		middle := (left + right) / 2
		if arr[middle] > x {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

// Q: 木棒问题
// 给出N根木棒，长度均已知，切割它们来获得至少K段长度相等的木棒（长度须为整数），求这些相等木棒的最长长度能有多长
func maxLen(arr []int, k int) int {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	left, right := 0, min
	for left < right {
		middle := (left + right) / 2
		if k > calcL(arr, middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left - 1
}

func calcL(arr []int, len int) (l int) {
	//for i := 0; i < len(arr); i++ {
	//	l += arr[i]/len
	//}

	for _, v := range arr {
		l += v / len
	}

	return l
}

// 快速幂
// 求a^b%m

// 递归写法
func binaryPow(a, b, m int64) (res int64) {
	if b == 0 {
		return 1
	} else if b%2 == 1 {
		return a * binaryPow(a, b-1, m) % m
	} else {
		mul := binaryPow(a, b/2, m)
		return mul * mul % m
	}
}

// 迭代写法
func binaryPow2(a,b,m int64) (res int64) {
	res =1
	for b>0 {
		if b%2 == 1{
			res = res*a%m
		}
		a = a*a%m
		b >>= 1 // b /= 2
	}
	return res
}
