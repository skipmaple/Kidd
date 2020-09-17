package __two_pointers

// Median
// page 181
// 给定两个递增序列， 求合并成一个递增序列后的中位数

func median(arr1, arr2 []int) (res int) {
	l1, r1, l2, r2 := 0, len(arr1)-1, 0, len(arr2)-1
	combineArr := make([]int, len(arr1)+len(arr2))
	index := 0
	for l1 <= r1 && l2 <= r2 {
		if arr1[l1] <= arr2[l2] {
			combineArr[index] = arr1[l1]
			index++
			l1++
		} else {
			combineArr[index] = arr2[l2]
			index++
			l2++
		}
	}
	for l1 <= r1 {
		combineArr[index] = arr1[l1]
		index++
		l1++
	}
	for l2 <= r2 {
		combineArr[index] = arr2[l2]
		index++
		l2++
	}
	//fmt.Println(combineArr)
	res = combineArr[(len(arr1)+len(arr2))/2]

	return res
}
