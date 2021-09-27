package main

import "fmt"
import "Kidd/utils"

// 成绩排名
// page 31

type testCase1004 struct {
	count    int
	students []utils.StudentInfo // studentInfo结构体 复用了 B041
}

func f1004() {
	t1 := testCase1004{
		count: 3,
		students: []utils.StudentInfo{
			{
				Name:  "Joe",
				Sid:   990112,
				Score: 89,
			},
			{
				Name:  "Mike",
				Sid:   991301,
				Score: 100,
			},
			{
				Name:  "Joe",
				Sid:   990830,
				Score: 95,
			},
		},
	}

	var min, max utils.StudentInfo
	min.Score, max.Score = 101, -1
	for _, s := range t1.students {
		if min.Score > s.Score {
			min = s
		}
		if max.Score < s.Score {
			max = s
		}
	}

	fmt.Println(max.Name, max.Sid)
	fmt.Println(min.Name, min.Sid)
}
