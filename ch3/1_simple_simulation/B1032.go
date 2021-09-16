package main

import "fmt"

// 挖掘机技术哪家强
func main() {
	var count, sid, score int
	var rSid, rScore int
	fmt.Println("请输入要比较的学校个数：")
	_, _ = fmt.Scanln(&count)
	res := make(map[int]int, count)
	for i := 0; i < count; i++ {
		fmt.Println("请输入学校id和得分：")
		_, _ = fmt.Scanln(&sid, &score)
		res[sid] += score
	}

	for sid, score := range res {
		if score > rScore {
			rSid = sid
			rScore = score
		}
	}

	fmt.Println("结果：", rSid, rScore)
}
