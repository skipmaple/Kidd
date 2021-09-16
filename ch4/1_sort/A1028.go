package main

// List Sorting
// page 103

type student1028 struct {
	id    int
	name  string
	score int
}

func listSorting(stus []student1028, col int) []student1028 {
	//// 冒泡
	//for i := 0; i < len(stus)-1; i++ {
	//	for j := 0; j < len(stus)-1-i; j++ {
	//		if cmp1028(stus[j], stus[j+1], col) {
	//			stus[j], stus[j+1] = stus[j+1], stus[j]
	//		}
	//	}
	//}

	//// 插入
	//for i := 1; i < len(stus); i++ {
	//	temp, j := stus[i], i
	//	for ; j > 0 && cmp1028(stus[j-1], temp, col); j-- {
	//		stus[j] = stus[j-1]
	//	}
	//	stus[j] = temp
	//}

	// 选择
	for i := 0; i < len(stus); i++ {
		k := i
		for j := i; j < len(stus); j++ {
			if cmp1028(stus[k], stus[j], col) {
				k = j
			}
		}
		stus[i], stus[k] = stus[k], stus[i]
	}
	return stus
}

func cmp1028(a, b student1028, col int) bool {
	switch col {
	case 1:
		return a.id > b.id
	case 2:
		if a.name == b.name {
			return a.id > b.id
		} else {
			return a.name > b.name
		}
	case 3:
		if a.score == b.score {
			return a.id > b.id
		} else {
			return a.score > b.score
		}
	default:
		return false
	}
}
