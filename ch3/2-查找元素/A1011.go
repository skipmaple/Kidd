package main

import (
	"fmt"
	"strconv"
)

// World Cup Betting
// page 36

type tCase struct {
	matches [3]matchBet
}

type matchBet struct {
	win  float64
	tie  float64
	lose float64
}

func worldCupBetting(t tCase) (string, string, string, float64) {
	res := make([]string, 3)
	income := 0.65
	for i, match := range t.matches {
		var rate float64
		res[i], rate = findBiggest(match.win, match.tie, match.lose)
		income *= rate
	}
	income,_ = strconv.ParseFloat(fmt.Sprintf("%.2f",(income - 1) * 2+0.005), 64) // ！！加0.005是为了确保 保留两位小数并且四舍五入
	return res[0], res[1], res[2], income
}

func findBiggest(win, tie, lose float64) (res string, rate float64) {
	if win > tie {
		rate = win
		res = "W"
	} else {
		rate = tie
		res = "T"
	}

	if rate < lose {
		rate = lose
		res = "L"
	}

	return res, rate
}
