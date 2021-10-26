package __vector

// Student List for Course
// Page 235

//type courseMsg struct {
//	courseId int
//	stuIds []string
//}
//
//type stuMsg struct {
//	stuId string
//	courseIds []int
//}

func studentListForCourse(stuMsgs []stuMsg) (res map[int][]string) {
	res = make(map[int][]string)
	for _, stuMsg := range stuMsgs {
		for _, courseId := range stuMsg.courseIds {
			res[courseId] = append(res[courseId], stuMsg.stuId)
		}
	}
	return res
}
