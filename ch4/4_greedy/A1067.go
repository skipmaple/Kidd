package ch4

// Sort with Swap(0,*)
// page 159

func sortWithSwap0(list []int) (res int) {
	left := 0 // 不在本位的数量
	minLeftValue := len(list)
	index0 := len(list)
	for i := 0; i < len(list); i++ {
		if list[i] != 0 && i != list[i] {
			left++
			if list[i] < minLeftValue {
				minLeftValue = list[i]
			}
		}

		if list[i] == 0 {
			index0 = i
		}
	}

	for left > 0 {
		if index0 == 0 {
			if minLeftValue != list[minLeftValue] {
				list, index0 = swap0(list, minLeftValue)
				res++
			} else {
				minLeftValue++
			}
		}

		if index0 != 0 {
			list, index0 = swap0(list, index0)
			left--
			res++
		}
	}

	return res
}

// 交换 0 和 index0 的位置
func swap0(list []int, index0 int) ([]int, int) {
	if list[index0] == 0 {
		for i := 0; i < len(list); i++ {
			if list[i] == index0 {
				list[i], list[index0] = list[index0], list[i]
				index0 = i
			}
		}
	} else {
		for i := 0; i < len(list); i++ {
			if list[i] == index0 {
				list[i], list[0] = list[0], list[i]
				index0 = i
			}
		}
	}

	return list, index0
}
