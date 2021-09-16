package hash

import (
	"fmt"
)

// 统计同成绩学生
// page 131

func sameGradeStudent(grades []int, queries []int) (res []int) {
	gradeCount := make(map[int]int)
	for _, grade := range grades {
		gradeCount[grade]++
	}

	for _, query := range queries {
		res = append(res, gradeCount[query])
	}
	return res
}

func main() {
	stuCount := 0
	_, _ = fmt.Scanf("%d", &stuCount)
	grades := make([]int, stuCount)
	for i := 0; i < stuCount; i++ {
		_, _ = fmt.Scanf("%d", grades[i])
	}

	queryCount := 0
	_, _ = fmt.Scanf("%d", &queryCount)
	queries := make([]int, queryCount)
	for i := 0; i < queryCount; i++ {
		_, _ = fmt.Scanf("%d", queries[i])
	}

	res := sameGradeStudent(grades, queries)
	for i, e := range res {
		if i != len(res)-1 {
			fmt.Printf("%d ", e)
		} else {
			fmt.Printf("%d\n", e)
		}
	}
}
