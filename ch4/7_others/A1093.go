package __others

// 有几个PAT
// page 184

func howManyPAT(str string) int {
	leftNumP := make([]int, len(str))
	leftNumP[0] = 0
	MOD := 1000000007
	for i := 0; i < len(str); i++ {
		if i > 0 {
			leftNumP[i] = leftNumP[i-1]
		}
		if str[i] == 'P' {
			leftNumP[i]++
		}
	}

	ans, rightNumT := 0, 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == 'T' {
			rightNumT++
		} else if str[i] == 'A' {
			ans = (ans + leftNumP[i]*rightNumT) % MOD
			//fmt.Printf("i = %d, ans = %d", i, ans)
		}
	}

	return ans
}
