package __simple_math

// page 190
// 我要通过！

func iWantPass(inputs []string) (res []string) {
	for i := 0; i < len(inputs); i++ {
		numP, numT, other, locP, locT := 0, 0, 0, -1, -1
		for j := 0; j < len(inputs[i]); j++ {
			temp := inputs[i][j]
			if temp == 'P' {
				numP++
				locP = j
			} else if temp == 'T' {
				numT++
				locT = j
			} else if temp != 'A' {
				other++
			}
		}

		currentRes := "NO"
		if other != 0 || numP != 1 || numT != 1 || locT-locP <= 1 {
		} else {
			x, y, z := locP, locT-locP-1, len(inputs[i])-locT-1
			if x == z-(y-1)*x {
				currentRes = "YES"
			}
		}
		//fmt.Printf("str => %s\n", inputs[i])
		//fmt.Printf("res => %s\n", currentRes)
		res = append(res, currentRes)
	}

	return res
}
