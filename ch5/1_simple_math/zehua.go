package __simple_math

import "fmt"

// merge 给定二维数组中重合的元素
// input: [[0,1], [2,4], [3,5]]
// want:  [[0,1], [2,5]]

func mergeMeetings(input [][]int) (res [][]int) {
	// 第一个数组元素排序
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1-i; j++ {
			if input[j][0] > input[j+1][0] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}

	// merge slice中重合的元素
	res = mergeArray(input)

	return res
}

func mergeArray(input [][]int) [][]int {
	flag := true
	fmt.Println(input)
	var res [][]int
	for flag {
		flag = false

		for i := 0; i < len(input)-1; i++ {
			fmt.Printf("i: %d, input[%d][1]: %v, input[%d][0]: %v\n", i, i, input[i][1], i, input[i+1][0])
			if input[i][1] > input[i+1][0] {
				res = replace(input, i, []int{input[i][0], input[i+1][1]})
				flag = true
			}
		}

		//fmt.Printf("\n")
		//fmt.Printf("input: %v\n", input)
		//fmt.Printf("res  : %v\n", res)

		input = res
	}
	return res
}

func replace(slice [][]int, i int, e []int) (res [][]int) {
	res = make([][]int, len(slice))
	copy(res, slice)
	res[i] = e
	copy(res[i+1:][:], res[i+2:][:])

	return res[:len(res)-1][:]
}

//func judgeSameSliceValue(slice1, slice2 [][]int) bool {
//	if len(slice1) != len(slice2) {
//		return false
//	}
//
//	for i := 0; i < len(slice1); i++ {
//		for j := 0; j < len(slice2); j++ {
//			if slice1[i][j] != slice2[i][j] {
//				return false
//			}
//		}
//	}
//
//	return true
//}
