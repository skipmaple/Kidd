package __map

import (
	"Kidd/utils"
	"fmt"
	"strconv"
	"strings"
)

// 火星数字
// page 244

func mars2Num(input string) (res int) {
	geMap := map[string]int{"tret": 0, "jan": 1, "feb": 2, "mar": 3, "apr": 4, "may": 5, "jun": 6, "jly": 7, "aug": 8, "sep": 9, "oct": 10, "nov": 11, "dec": 12}
	shiMap := map[string]int{"tam": 1, "hel": 2, "maa": 3, "huh": 4, "tou": 5, "kes": 6, "hei": 7, "elo": 8, "syy": 9, "lok": 10, "mer": 11, "jou": 12}
	inputRes := strings.Split(input, " ")
	if len(inputRes) == 1 {
		res = geMap[inputRes[0]]
		if res == 0 {
			res = shiMap[inputRes[0]]
		}
	}
	if len(inputRes) == 2 {
		res = shiMap[inputRes[0]]*13 + geMap[inputRes[1]]
	}

	return res
}

func num2Mars(input int) (res string) {
	geMap := map[int]string{0: "tret", 1: "jan", 2: "feb", 3: "mar", 4: "apr", 5: "may", 6: "jun", 7: "jly", 8: "aug", 9: "sep", 10: "oct", 11: "nov", 12: "dec"}
	shiMap := map[int]string{1: "tam", 2: "hel", 3: "maa", 4: "huh", 5: "tou", 6: "kes", 7: "hei", 8: "elo", 9: "syy", 10: "lok", 11: "mer", 12: "jou"}

	thirteenNum := utils.TtoQ(input, 13)
	if len(thirteenNum) >= 2 {
		ge, _ := strconv.Atoi(thirteenNum[1])
		shi, _ := strconv.Atoi(thirteenNum[0])
		res = fmt.Sprintf("%s %s", shiMap[shi], geMap[ge])
	} else {
		ge, _ := strconv.Atoi(thirteenNum[0])
		res = fmt.Sprintf("%s", geMap[ge])
	}

	return res
}
