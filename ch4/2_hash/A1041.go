package hash

// Be Unique
// page 139

func beUnique(input []int) (res int) {
	res = -1 // -1 代表原题中的"None"输出
	seedCount := make(map[int]int)
	for _, n := range input {
		seedCount[n]++
	}
	for i := 0; i < len(input); i++ {
		if seedCount[input[i]] == 1 {
			res = input[i]
			break
		}
	}
	return res
}
