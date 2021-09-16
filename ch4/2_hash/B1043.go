package hash

import (
	"fmt"
	"strings"
)

// 输出PATest
// page 136

func printPATest(intput string) (res string) {
	outputFmt := "PATest"
	alphaCount := make(map[string]int)
	for _, s := range outputFmt {
		alphaCount[string(s)] = strings.Count(intput, string(s))
	}

	maxCount := 0
	for _, count := range alphaCount {
		if count > maxCount {
			maxCount = count
		}
	}

	for i := 0; i < maxCount; i++ {
		for _, s := range outputFmt {
			if alphaCount[string(s)] > 0 {
				res = fmt.Sprintf("%s%s", res, string(s))
				alphaCount[string(s)]--
			}
		}
	}

	return res
}
