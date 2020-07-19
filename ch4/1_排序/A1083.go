package main

import "fmt"

// List Grades
// page 115

type student1083 struct {
	name string
	id string
	score int
}

func listGrades(stus []student1083, minScore,maxScore int) (res []string) {
	var validStus []student1083
	for _, s := range stus {
		if s.score <= maxScore && s.score >= minScore {
			validStus = append(validStus, s)
		}
	}

	if len(validStus) == 0 {
		return []string{"NONE"}
	}

	for i:=0;i< len(validStus);i++{
		k := i;
		for j:=i;j< len(validStus);j++ {
			if cmp1083(validStus[j], validStus[k]) {
				k = j
			}
		}
		validStus[i], validStus[k] = validStus[k],validStus[i]
	}

	for i:=0;i< len(validStus);i++ {
		res = append(res, fmt.Sprintf("%s %s", validStus[i].name, validStus[i].id))
	}

	return res
}

func cmp1083(a, b student1083) bool {
	return a.score > b.score
}
