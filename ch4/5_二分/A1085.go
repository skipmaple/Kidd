package ch4

// 完美数列
// page165

func perfectList(arr []int, p int) (res int) {

	//for i := 1; i < len(arr); i++ { // 进行 n-1 趟排序
	//	temp, j := arr[i], i // temp 临时存放arr[i], j从i开始往后枚举
	//	for ; j > 0 && temp < arr[j-1]; j-- {
	//		arr[j] = arr[j-1] // arr[j-1]后移一位至arr[j]
	//	}
	//	arr[j] = temp
	//}


	for i := 1; i < len(arr); i++ {
		temp, j := arr[i], i
		for ; j > 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}

	min := arr[0]

	arrPlus := append(arr, arr[len(arr)-1]+1)
	//fmt.Println(arr)

	left,right := 0, len(arr)
	for left<right {
		middle := (left+right)/2
		if arrPlus[middle] > min*p {
			right = middle
		} else {
			left = middle+1
		}
	}

	for _, v := range arr {
		if v < arrPlus[left]{
			res++
		}
	}

	return res
}
