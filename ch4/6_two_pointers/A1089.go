package __two_pointers

import (
	"fmt"
	"github.com/UncleMaple/Kidd/utils"
)

// Insert or Merge
// page 177

func InsertOrMerge(arrInput, arrGet []int) (res string, nextStep []int) {
	originArr := make([]int, len(arrInput))
	copy(originArr, arrInput)
	flag, arr := insertSort(arrInput, arrGet)
	if flag {
		return "Insertion Sort", arr
	}

	fmt.Println(originArr)
	flag, arr = mergeSort(originArr, arrGet)
	if flag {
		return "Merge Sort", arr
	}
	return "Not Found", []int{}
}

func insertSort(arr, arrGet []int) (bool, []int) {
	flag := false
	for i := 1; i < len(arr); i++ { // 进行 n-1 趟排序
		temp, j := arr[i], i // temp 临时存放arr[i], j从i开始往后枚举
		for ; j > 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1] // arr[j-1]后移一位至arr[j]
		}
		arr[j] = temp
		if flag {
			return true, arr
		}
		if IntArrayEquals(arr, arrGet) {
			flag = true
		}
	}

	return false, arr
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func merge(arr []int, l1 int, r1 int, l2 int, r2 int) []int {
	i, j := l1, l2
	tempArr := make([]int, len(arr))
	index := 0
	for i <= r1 && j <= r2 {
		if arr[i] <= arr[j] {
			tempArr[index] = arr[i]
			index++
			i++
		} else {
			tempArr[index] = arr[j]
			index++
			j++
		}
	}

	for i <= r1 {
		tempArr[index] = arr[i]
		index++
		i++
	}

	for j <= r2 {
		tempArr[index] = arr[j]
		index++
		j++
	}

	for i := 0; i < index; i++ {
		arr[l1+i] = tempArr[i]
	}

	return arr
}

// 非递归写法
func mergeSort(arr, arrGet []int) (bool, []int) {
	flag := false
	// step 为组内元素个数，step/2为左子区间元素个数
	for step := 2; step/2 < len(arr); step *= 2 {
		// 每step个元素一组，组内前step/2 和 后step/2 个元素合并
		for i := 0; i < len(arr); i += step {
			mid := i + step/2 - 1
			if mid+1 < len(arr) {
				arr = merge(arr, i, mid, mid+1, utils.Min(i+step-1, len(arr)-1))
			}
		}
		if flag {
			return true, arr
		}
		//fmt.Println("step: ", arr)
		if IntArrayEquals(arr, arrGet) {
			flag = true
		}
	}

	//fmt.Println("归并排序：",arr)
	return false, arr
}
