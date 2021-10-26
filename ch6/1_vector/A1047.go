package __vector

// Course List for Student
// Page 232

type courseMsg struct {
	courseId int
	stuIds   []string
}

type stuMsg struct {
	stuId     string
	courseIds []int
}

func courseListForStudent(courseMessages []courseMsg) (stuMessages map[string][]int) {
	stuMessages = make(map[string][]int)

	for _, courseMessage := range courseMessages {
		for _, stuId := range courseMessage.stuIds {
			stuMessages[stuId] = append(stuMessages[stuId], courseMessage.courseId)
		}
	}

	return stuMessages
}
