package __map

import (
	"sort"
)

// The Dominant Color
// page 247

func dominantColor(input [][]int) (max int) {
	res := make(map[int]int)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			res[input[i][j]]++
		}
	}

	keys := make([]int, 0, len(res))
	for key := range res {
		keys = append(keys, key)
	}

	// 这块的排序没有搞懂, (和闭包相关??)
	sort.Slice(keys, func(i, j int) bool {
		return res[keys[i]] > res[keys[j]]
	})

	//for _, key := range keys {
	//	fmt.Println(key, res[key])
	//}

	max = keys[0]

	return max
}
