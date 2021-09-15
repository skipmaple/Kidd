package hash

// 到底买不买
// page 133

func buyOrNot(have, want string) (flag bool, count int) {
	flag = true
	haveCount := make(map[string]int)
	for _, e := range have {
		haveCount[string(e)]++
	}

	wantCount := make(map[string]int)
	for _, e := range want {
		wantCount[string(e)]++
	}

	for k, v := range wantCount {
		if haveCount[k] < v {
			flag = false
			count++
		}
	}

	if flag {
		count = len(have) - len(want)
	}

	return flag, count
}
