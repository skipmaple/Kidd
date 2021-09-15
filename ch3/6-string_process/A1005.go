package main

import (
	"fmt"
	"strconv"
)

// Spell it Right
// page 75

func spellNum(n string) (res string) {
	numMap := map[int]string{0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}
	sum := 0
	for i := 0; i < len(n); i++ {
		temp, _ := strconv.Atoi(string(n[i]))
		sum += temp
	}

	sumStr := strconv.Itoa(sum)
	for i := 0; i < len(sumStr); i++ {
		temp, _ := strconv.Atoi(string(sumStr[i]))
		res += fmt.Sprintf("%s", numMap[temp])
		if i != len(sumStr)-1 {
			res += " "
		}
		//res = fmt.Sprintf("%s %s", res, numMap[temp])
	}

	return res
}
