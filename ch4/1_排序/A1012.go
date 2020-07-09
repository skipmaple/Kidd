package main

import "fmt"

// The Best Rank
// page 90

type CsStudent struct {
	id     int
	cScore int
	cSort  int
	mScore int
	mSort  int
	eScore int
	eSort  int
	aScore float64
	aSort  int
}

//func main() {
//	stus := []CsStudent{
//		{id: 310101, cScore: 98, mScore: 85, eScore: 88},
//		{id: 310102, cScore: 70, mScore: 95, eScore: 88},
//		{id: 310103, cScore: 82, mScore: 87, eScore: 94},
//		{id: 310104, cScore: 91, mScore: 91, eScore: 91},
//		{id: 310105, cScore: 85, mScore: 90, eScore: 90},
//	}
//	queryId := []int{310101, 310102, 310103, 310104, 310105, 9999999}
//
//	fmt.Println(theBestRank(stus,queryId))
//	fmt.Println(stus)
//}

func theBestRank(stus []CsStudent, queryIds []int) (res []string) {
	var stuIds []int
	for i:=0;i< len(stus);i++ {
		stus[i].aScore = float64(stus[i].cScore+stus[i].mScore+stus[i].eScore) / 3
		stuIds = append(stuIds, stus[i].id)
	}

	stus = sortCsStudent(stus, "cScore")
	//fmt.Println("按c成绩排序：", stus)
	stus = rankCsStudent(stus, "cScore", len(stus))

	stus = sortCsStudent(stus, "mScore")
	//fmt.Println("按m成绩排序：", stus)
	stus = rankCsStudent(stus, "mScore", len(stus))

	stus = sortCsStudent(stus, "eScore")
	//fmt.Println("按e成绩排序：", stus)
	stus = rankCsStudent(stus, "eScore", len(stus))

	stus = sortCsStudent(stus, "aScore")
	//fmt.Println("按a成绩排序：", stus)
	stus = rankCsStudent(stus, "aScore", len(stus))

	//fmt.Println("按自定义成绩排序：", sortCsStudent(stus, "cScore"))


	for i := 0; i < len(queryIds); i++ {
		if contains(stuIds, queryIds[i]) {
			res = append(res, findBestRank(queryIds[i], stus))
		} else {
			res = append(res, "N/A")
		}
	}

	return res
}

func findBestRank(id int, stus []CsStudent) string {
	var stu CsStudent
	for _, s := range stus {
		if s.id == id {
			stu = s
		}
	}

	sortArr := []int{stu.aSort, stu.cSort, stu.mSort, stu.eSort}

	return findBestSort(sortArr)
}

func findBestSort(sortArr []int) string {
	maxSort := 99
	var subjectIndex int
	subjectMap := map[int]string{0: "A", 1: "C", 2: "M", 3: "E"}
	for i, v := range sortArr {
		if v < maxSort {
			maxSort = v
			subjectIndex = i
		}
	}

	return fmt.Sprintf("%d %s", maxSort, subjectMap[subjectIndex])
}

func contains(arr []int, n int) bool {
	for _, e := range arr {
		if e == n {
			return true
		}
	}
	return false
}

func rankCsStudent(stus []CsStudent, category string, len int) []CsStudent{
	switch category {
	case "cScore":
		stus[0].cSort = 1
		for i := 1; i < len; i++ {
			if stus[i].cScore == stus[i-1].cScore {
				stus[i].cSort = stus[i-1].cSort
			} else {
				stus[i].cSort = i + 1
			}
		}
		break
	case "mScore":
		stus[0].mSort = 1
		for i := 1; i < len; i++ {
			if stus[i].mScore == stus[i-1].mScore {
				stus[i].mSort = stus[i-1].mSort
			} else {
				stus[i].mSort = i + 1
			}
		}
		break
	case "eScore":
		stus[0].eSort = 1
		for i := 1; i < len; i++ {
			if stus[i].eScore == stus[i-1].eScore {
				stus[i].eSort = stus[i-1].eSort
			} else {
				stus[i].eSort = i + 1
			}
		}
		break
	case "aScore":
		stus[0].aSort = 1
		for i := 1; i < len; i++ {
			if stus[i].aScore == stus[i-1].aScore {
				stus[i].aSort = stus[i-1].aSort
			} else {
				stus[i].aSort = i + 1
			}
		}
		break
	}

	return stus
}

//func sortCsStudent(stus []CsStudent, category string) []CsStudent {
//	// 插入排序
//	for i := 1; i < len(stus); i++ {
//		temp, j := stus[i], i
//		for ; j > 0 && cmpCsStudent(temp, stus[j], category); j-- {
//			stus[j] = stus[j-1]
//		}
//		stus[j] = temp
//	}
//
//	return stus
//}

func sortCsStudent(stus []CsStudent, cat string) []CsStudent {
	for i := 0; i < len(stus); i++ {
		k := i
		for j := k; j < len(stus); j++ {
			if cmpCsStudent(stus[j], stus[k],cat) {
				k = j
			}
		}
		stus[i], stus[k] = stus[k], stus[i]
	}

	return stus
}

func cmpCsStudent(a, b CsStudent, category string) bool {
	switch category {
	case "cScore":
		if a.cScore != b.cScore {
			return a.cScore > b.cScore
		} else {
			return a.id < b.id
		}
	case "mScore":
		if a.mScore != b.mScore {
			return a.mScore > b.mScore
		} else {
			return a.id < b.id
		}
	case "eScore":
		if a.eScore != b.eScore {
			return a.eScore > b.eScore
		} else {
			return a.id < b.id
		}
	case "aScore":
		if a.aScore != b.aScore {
			return a.aScore > b.aScore
		} else {
			return a.id < b.id
		}
	default:
		return false
	}
}


