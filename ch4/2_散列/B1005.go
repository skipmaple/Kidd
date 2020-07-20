package hash

import "sort"

// 继续(3n+1)猜想（即 卡拉兹猜想）
// page 143

func callatzThink(input []int) (res []int) {

	// // 一顿操作猛如虎，回头一看二百五。。理解错题意了，理解成各个数的卡拉兹猜想过程数组的包含关系了
	//numsCallatzArray := make(map[int][]int)
	//for _, num := range input {
	//	numsCallatzArray[num] = callatz(num)
	//}
	//
	//for k, arr := range numsCallatzArray {
	//	flag := true
	//	for k2, arr2 := range numsCallatzArray {
	//		if k != k2 && intArrContainsCheck(arr2, arr) {
	//			flag = false
	//			break
	//		}
	//	}
	//	if flag {
	//		res = append(res, k)
	//	}
	//}

	coverMap := make(map[int]bool)
	for _, n := range input {
		coverMap = callatz(n, coverMap)
	}

	for _, n := range input {
		if !coverMap[n]{
			res = append(res, n)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(res))) // 这一顿操作...

	return res
}

func callatz(num int, coverMap map[int]bool) map[int]bool {
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = 3*num + 1
		}
		coverMap[num] = true
	}
	return coverMap
}

// // 比较两个数组a,b 是否a中包含b中所有元素
//func intArrContainsCheck(arr, subArr []int) (res bool) {
//	res = true
//	for _, subNum := range subArr {
//		existFlag := false
//		for _, num := range arr {
//			if num == subNum {
//				existFlag = true
//				break
//			}
//		}
//		if !existFlag {
//			res = false
//			break
//		}
//	}
//	return res
//}
