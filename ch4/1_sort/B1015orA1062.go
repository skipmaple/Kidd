package main

// 德才论
// page 87

type stu struct {
	id         int
	deScore    int
	caiScore   int
	totalScore int
	//sort       int
}

func deCaiLun(lScore, hScore int, students []stu) (count int, res []stu) {
	var first, second, third, fourth []stu
	for _, stu := range students {
		stu.totalScore = stu.deScore + stu.caiScore

		// 德分 或 才分 低于最低线 不予录取
		if stu.deScore < lScore || stu.caiScore < lScore {
			continue
		}

		if stu.deScore >= hScore && stu.caiScore >= hScore {
			// 德才全尽
			first = append(first, stu)
		} else if stu.caiScore < hScore && stu.deScore >= hScore {
			// 德胜才
			second = append(second, stu)
		} else if stu.deScore < hScore && stu.caiScore < hScore && stu.deScore >= stu.caiScore {
			// 才德兼亡，但尚有德胜才
			third = append(third, stu)
		} else {
			fourth = append(fourth, stu)
		}
	}

	first = sortStudent(first)
	second = sortStudent(second)
	third = sortStudent(third)
	fourth = sortStudent(fourth)

	for _, s := range first {
		res = append(res, s)
	}

	for _, s := range second {
		res = append(res, s)
	}

	for _, s := range third {
		res = append(res, s)
	}

	for _, s := range fourth {
		res = append(res, s)
	}

	return len(res), res
}

func sortStudent(stus []stu) []stu {
	for i := 0; i < len(stus); i++ {
		k := i
		for j := k; j < len(stus); j++ {
			if cmpStudent(stus[j], stus[k]) {
				k = j
			}
		}
		stus[i], stus[k] = stus[k], stus[i]
	}

	return stus
}

func cmpStudent(a, b stu) bool {
	if a.totalScore != b.totalScore {
		return a.totalScore > b.totalScore
	} else if a.deScore != b.deScore {
		return a.deScore > b.deScore
	} else {
		return a.id < b.id
	}
}
