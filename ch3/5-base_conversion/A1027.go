// Colors in Mars
// page 55

// 输入十进制rgb色彩 输出用十三进制表示

package main

import (
	"fmt"
	"github.com/UncleMaple/Kidd/utils"
	_ "github.com/UncleMaple/Kidd/utils"
)

const ToThirteen int = 13

func thirteenRGB(input []int) (res string) {
	res = "#"
	for i := 0; i < len(input); i++ {
		res = fmt.Sprintf("%s%s", res, TtoQtoString(input[i], ToThirteen))
	}

	return res
}

func TtoQtoString(n, q int) (res string) {
	arr := utils.TtoQ(n, q)

	for i := 0; i < len(arr); i++ {
		res = fmt.Sprintf("%s%s", res, arr[i])
	}

	return res
}
