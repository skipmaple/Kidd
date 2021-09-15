package main

import "fmt"

// Shortest Distance
// 最短距离
// page 20

func main() {
	var count int
	_, _ = fmt.Scanf("%d", &count)
	a := make([]int, count)
	dis := make([]int, count)
	var disSum int

	// 输入各个节点的长度
	for i := 0; i < count; i++ {
		_, _ = fmt.Scanf("%d", &a[i])
		disSum += a[i]
		dis[i] = disSum
	}

	// 输入要查询的个数
	var reqCount int
	_, _ = fmt.Scanf("%d", &reqCount)

	resD := make([]int, reqCount)
	// 输入reqCount个查询节点
	for i := 0; i < reqCount; i++ {
		var point1, point2 int
		var d1, d2 int
		_, _ = fmt.Scanf("%d %d", &point1, &point2)
		point1, point2 = swap(point1-1, point2-1)
		var temp int
		if point1 == 0 {
			temp = 0
		} else {
			temp = dis[point1-1]
		}
		d1 = dis[point2-1] - temp // 这里就是需要减去1才能对应所求点的距离，也由此引入了temp[初始点的距离是0]
		d2 = disSum - d1
		resD[i], _ = swap(d1, d2)
	}

	for _, v := range resD {
		fmt.Println(v)
	}
}

func swap(a, b int) (min, max int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return min, max
}
