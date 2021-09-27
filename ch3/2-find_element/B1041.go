package main

import "fmt"
import "Kidd/utils"

// 考试座位号
// page 29

type testCase1041 struct {
	count         int
	studentInfos  []utils.StudentInfo
	queryCount    int
	queryTestNums []int
}

func f1041() {
	t1 := testCase1041{
		count: 4,
		studentInfos: []utils.StudentInfo{
			{
				Sid:     10120150912233,
				TestNum: 2,
				RealNum: 4,
			},
			{
				Sid:     10120150912119,
				TestNum: 4,
				RealNum: 1,
			},
			{
				Sid:     10120150912126,
				TestNum: 1,
				RealNum: 3,
			},
			{
				Sid:     10120150912002,
				TestNum: 3,
				RealNum: 2,
			},
		},
		queryCount:    2,
		queryTestNums: []int{3, 4},
	}

	for _, qNum := range t1.queryTestNums {
		for _, stu := range t1.studentInfos {
			if qNum == stu.TestNum {
				fmt.Println(stu.Sid, stu.RealNum)
			}
		}
	}

}
