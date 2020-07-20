package hash

import "strings"

// 字符统计
// page 135

func byteCount(intput string) (alpha string, max int) {
	alphaCount := make(map[string]int)
	for _, e := range intput {
		temp := strings.ToLower(string(e))
		if temp >= "a" && temp <= "z" {
			alphaCount[temp]++
		}
	}

	for k, v := range alphaCount {
		if v > max {
			max = v
			alpha = k
		} else if v == max && k < alpha{
			alpha = k
		}
	}

	return alpha, max
}
