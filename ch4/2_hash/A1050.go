package hash

import (
	"fmt"
)

// String Subtraction
// page 141

// 根据go语言原生方法 我能想到的最简单写法
//func stringSubtraction(str1, str2 string) string {
//	for _, s := range str2 {
//		str1 = strings.ReplaceAll(str1, string(s), "")
//	}
//	return str1
//}

// 用hash实现
func stringSubtraction(str1, str2 string) (res string) {
	charMap := make(map[byte]bool)
	for _, s := range str2 {
		charMap[byte(s)] = true
	}

	for _, s := range str1 {
		if !charMap[byte(s)] {
			res = fmt.Sprintf("%s%s", res, string(s))
		}
	}
	return res
}
