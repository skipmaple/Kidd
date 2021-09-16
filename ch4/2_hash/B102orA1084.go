package hash

import "strings"

// 旧键盘
// page 128

func brokenKeyBoard(input, output string) (res string) {
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	for _, s := range input {
		map1[strings.ToUpper(string(s))]++
	}

	for _, s := range output {
		map2[strings.ToUpper(string(s))]++
	}

	for k1, v1 := range map1 {
		if map2[k1] != v1 {
			res += k1
		}
	}

	return res
}
