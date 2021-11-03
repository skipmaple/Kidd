package __map

import (
	"fmt"
	"sort"
	"strings"
)

// Speech Patterns
// page 249

func speechPatterns(input string) (res string) {

	// 转换成小写
	input = strings.ToLower(input)

	// 分词成数组
	wordMap := make(map[string]int)
	tmpStr := ""
	for i := 0; i < len(input); i++ {
		if check(input[i]) {
			tmpStr += string(input[i])
		} else {
			if tmpStr == "" {
				continue
			}
			wordMap[tmpStr]++
			tmpStr = ""
		}
	}

	// 遍历map，找出出现最多的单词
	keys := make([]string, 0, len(wordMap))
	for k, _ := range wordMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return wordMap[keys[i]] > wordMap[keys[j]]
	})

	res = fmt.Sprintf("%s %d", keys[0], wordMap[keys[0]])

	return res
}

func check(c byte) bool {
	if c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
		return true
	}

	return false
}
