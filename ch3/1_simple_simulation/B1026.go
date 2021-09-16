package main

import (
	"fmt"
	"math"
)

const CLK_TCK = 100.0

// 程序运行时间
// 常数CLK_TCK：机器时钟每秒所走的时钟打点数
// 计算两个打点数的间隔时分秒
func main() {
	var c1, c2 int
	var diffSec float64
	var hh, mm, ss int
	_, _ = fmt.Scanf("%d %d", &c1, &c2)
	diffSec = float64(c2-c1) / CLK_TCK
	diffSec = math.Round(diffSec)
	//fmt.Printf("%g\n", diffSec)
	ss = int(diffSec) % 60
	mm = int(diffSec) / 60 % 60
	hh = int(diffSec) / 60 / 60
	fmt.Printf("%2d:%2d:%2d\n", hh, mm, ss)
}
