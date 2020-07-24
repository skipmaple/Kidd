package ch4

import "fmt"

// 组最小个数
// page 148

func minNum(counts []int) (min []int) {
	var nums []int
	for i:=0;i< len(counts);i++ {
		for j:=0;j<counts[i];j++ {
			nums = append(nums, i)
		}
	}
	//// counts 由小到大排序
	//for i := 1; i < len(counts); i++ {
	//	temp, j := counts[i], i
	//	for ; j > 0 && temp < counts[j]; j-- {
	//		counts[j] = counts[j-1]
	//	}
	//	counts[j] = temp
	//}

	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			min = append(min, nums[i])
			nums = append(nums[:i], nums[i+1:]...)
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		min = append(min, nums[i])
	}

	return min
}
