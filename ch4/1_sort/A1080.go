package main

import (
	"strconv"
	"strings"
)

// Graduate Admission
// page 117

type stu1080 struct {
	id        int
	rank      int
	geScore   int
	giScore   int
	aveScore  float64
	schoolIds []int
}

type school struct {
	id           int
	acceptCount  int
	acceptedStus []stu1080
}

func graduateAdmission(stus []stu1080, schools []school) (res []string) {
	// 算平均分
	for i, s := range stus {
		stus[i].aveScore = float64(s.geScore+s.giScore) / 2
	}

	// 排序
	for i := 1; i < len(stus); i++ {
		temp, j := stus[i], i
		for ; j > 0 && cmp1080(temp, stus[j-1]); j-- {
			stus[j] = stus[j-1]
		}
		stus[j] = temp
	}

	// 添加排名
	for i := 0; i < len(stus); i++ {
		if i == 0 {
			stus[0].rank = 1
		} else {
			if stus[i].aveScore == stus[i-1].aveScore && stus[i].geScore == stus[i-1].geScore {
				stus[i].rank = stus[i-1].rank
			} else {
				stus[i].rank = i + 1
			}
		}
	}

	// 学校录取
	for i := 0; i < len(stus); i++ {
		//fmt.Println(stus[i])
		for j := 0; j < len(stus[i].schoolIds); j++ {
			flag := false
			for k := 0; k < len(schools); k++ { // 查找学校
				if schools[k].id == stus[i].schoolIds[j] {
					if len(schools[k].acceptedStus) < schools[k].acceptCount || schools[k].acceptedStus[len(schools[k].acceptedStus)-1].rank == stus[i].rank {
						schools[k].acceptedStus = append(schools[k].acceptedStus, stus[i])
						flag = true
						break
					}
				}
			}
			if flag {
				break
			}
		}
	}

	for i := 0; i < len(schools); i++ {
		var tempIds []int
		for j := 0; j < len(schools[i].acceptedStus); j++ {
			tempIds = append(tempIds, schools[i].acceptedStus[j].id)
		}
		res = append(res, strings.Join(sortId(tempIds), " "))
	}

	return res
}

func cmp1080(a, b stu1080) bool {
	if a.aveScore != b.aveScore {
		return a.aveScore > b.aveScore
	} else {
		return a.geScore > b.geScore
	}
}

func sortId(arr []int) (res []string) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		res = append(res, strconv.Itoa(arr[i]))
	}
	return res
}
