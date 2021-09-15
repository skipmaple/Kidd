package main

import (
	"fmt"
	"strings"
)

// PAT Judge
// page 109

type student1075 struct {
	sid          int
	records      []problemRecord
	flag         bool // 是否有能通过编译的记录
	totalScore   int
	perfectCount int
	rank         int
}

type problem struct {
	pid      int
	maxScore int
}

type problemRecord struct {
	sid   int
	pid   int
	score int
}

func patRank(stus []student1075, problems []problem, records []problemRecord) (res []string) {
	// 给考生统计每道题最高分
	for i := 0; i < len(records); i++ { // 遍历答题记录
		// !!! 下边儿这一大坨，都是因为我对 go深浅赋值掌握不够扎实写的无用代码。。
		// !!! 之前一直在排查是不是这里的错误, 错怪这段恶心的箭头代码了-_-
		// !!! 错误依旧保留在测试部分，打了注释了

		//	for j := 0; j < len(stus); j++ {
		//		if stus[j].sid == records[i].sid { // 找到对应考生
		//			for k := 0; k < len(stus[j].records); k++ {
		//				if stus[j].records[k].pid == records[i].pid {
		//					if records[i].score > stus[j].records[k].score {
		//						if records[i].score != -1 {
		//							stus[j].records[k].score = records[i].score
		//							if stus[j].flag == false {
		//								stus[j].flag = true
		//							}
		//						} else {
		//							stus[j].records[k].score = 0
		//						}
		//					}
		//				}
		//			}
		//		}
		//	}

		for sIndex, s := range stus { // 查找该考生
			if s.sid == records[i].sid {
				for rIndex, r := range stus[sIndex].records { // 更新该考生对应题目最高分
					if records[i].pid == r.pid {
						if records[i].score > r.score {
							if records[i].score != -1 { // 编译通过
								stus[sIndex].records[rIndex].score = records[i].score
								if s.flag == false { // 更新有编译成功记录的flag
									stus[sIndex].flag = true
								}
							} else { // 编译不通过
								stus[sIndex].records[rIndex].score = 0
							}
						}
					}
				}
			}
		}
	}

	//fmt.Println("unsorted: ", stus)

	// 筛选出所有 有编译通过记录的考生
	var validStus []student1075
	for _, s := range stus {
		if s.flag {
			validStus = append(validStus, s)
		}
	}

	problemMaxScore := make(map[int]int)
	for _, p := range problems {
		problemMaxScore[p.pid] = p.maxScore
	}

	// 计算每个考生的总分和满分数量
	for sIndex, s := range validStus {
		for _, r := range s.records {
			if r.score == -2 {
				continue
			}
			validStus[sIndex].totalScore += r.score
			if r.score == problemMaxScore[r.pid] {
				validStus[sIndex].perfectCount++
			}
		}
	}

	// 选择排序
	for i := 0; i < len(validStus); i++ {
		k := i
		for j := i; j < len(validStus); j++ {
			if cmp1075(validStus[j], validStus[k]) {
				k = j
			}
		}
		validStus[i], validStus[k] = validStus[k], validStus[i]
	}

	fmt.Println("sorted: ", validStus)

	for i := 0; i < len(validStus); i++ {
		// 加入排名
		if i == 0 {
			validStus[0].rank = 1
		} else {
			if validStus[i].totalScore == validStus[i-1].totalScore {
				validStus[i].rank = validStus[i-1].rank
			} else {
				validStus[i].rank = i + 1
			}
		}

		var problemScores string
		for _, r := range validStus[i].records {
			if r.score == -2 {
				problemScores += fmt.Sprintf("- ")
			} else {
				problemScores += fmt.Sprintf("%d ", r.score)
			}
		}
		problemScores = strings.TrimRight(problemScores, " ")

		res = append(res, fmt.Sprintf("%d %05d %d ", validStus[i].rank, validStus[i].sid, validStus[i].totalScore)+problemScores)
	}

	return res
}

func cmp1075(a, b student1075) bool {
	if a.totalScore != b.totalScore {
		return a.totalScore > b.totalScore
	} else if a.perfectCount != b.perfectCount {
		return a.perfectCount > b.perfectCount
	} else {
		return a.sid < b.sid
	}
}
