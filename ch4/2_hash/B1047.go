package hash

import (
	"strconv"
	"strings"
)

// 编程团队赛
// page 138

type grade struct {
	info  string
	score int
}

func programTeamContest(grades []grade) (team, totalScore int) {
	teamGrade := make(map[int]int)
	for _, g := range grades {
		teamId, _ := strconv.Atoi(strings.Split(g.info, "-")[0])
		teamGrade[teamId] += g.score
		if teamGrade[teamId] > totalScore {
			team = teamId
			totalScore = teamGrade[teamId]
		}
	}
	return team, totalScore
}
