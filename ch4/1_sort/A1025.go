package main

// PAT Ranking
// main page 103

type student struct {
	id             string
	score          int
	locationNumber int // 考场号
	localRank      int // 考场排名
	totalRank      int // 总排名
}

// 考场
type testLocation struct {
	count    int       // 考场人数
	students []student // 考场学生
}

func stuRank(locationCount int, locations []testLocation) (total int, totalStu []student) {
	for i := 0; i < locationCount; i++ {
		// 分考场对考生成绩排序
		locations[i].students = rankLocation(locations[i].students, "location")
		// 加入总列表
		for _, stu := range locations[i].students {
			totalStu = append(totalStu, stu)
		}
	}

	totalStu = rankLocation(totalStu, "total")
	return len(totalStu), totalStu
}

// 学生排名
func rankLocation(students []student, sort string) []student {
	students = sortStu(students) // 分由高到低排序，分数相同则id小的在前面

	if sort == "location" {
		students[0].localRank = 1
		for i := 1; i < len(students); i++ {
			if students[i].score == students[i-1].score {
				students[i].localRank = students[i-1].localRank
			} else {
				students[i].localRank = i + 1
			}
		}
	} else if sort == "total" {
		students[0].totalRank = 1
		for i := 1; i < len(students); i++ {
			if students[i].score == students[i-1].score {
				students[i].totalRank = students[i-1].totalRank
			} else {
				students[i].totalRank = i + 1
			}
		}
	}

	return students
}

func cmp(a, b student) bool {
	if a.score != b.score {
		return a.score > b.score
	} else {
		return a.id < b.id
	}
}

// 排序
func sortStu(students []student) []student {
	for j := 1; j < len(students); j++ {
		temp, k := students[j], j
		for ; k > 0 && cmp(temp, students[k-1]); k-- {
			students[k] = students[k-1]
		}
		students[k] = temp
	}

	return students
}
