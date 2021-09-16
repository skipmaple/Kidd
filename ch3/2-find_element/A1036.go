package main

// Boys vs Girls
// page 40

type tCase1036 struct {
	count    int
	students []student
}

type student struct {
	name   string
	gender string
	id     string
	grade  int
}

func boysVsGirls(t tCase1036) (fName, fId, mName, mId string, diff interface{}) {
	minMale := student{
		name:   "Absent",
		id:     "Absent",
		gender: "M",
		grade:  100,
	}
	maxFemale := student{
		name:   "Absent",
		id:     "Absent",
		gender: "F",
		grade:  0,
	}
	for _, stu := range t.students {
		if stu.gender == "F" {
			if stu.grade > maxFemale.grade {
				maxFemale = stu
			}
		}

		if stu.gender == "M" {
			if stu.grade < minMale.grade {
				minMale = stu
			}
		}
	}

	if maxFemale.name == "Absent" || minMale.name == "Absent" {
		diff = "NA"
	} else {
		diff = maxFemale.grade - minMale.grade
	}
	fName, fId, mName, mId = maxFemale.name, maxFemale.id, minMale.name, minMale.id
	return fName, fId, mName, mId, diff
}
