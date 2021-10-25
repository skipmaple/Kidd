package __prime

// Hashing
// Page 215

func hashing(input []int, m int) (res []int) {
	t := len(input)
	if !isPrime(t) {
		for {
			t += 1
			if isPrime(t) {
				break
			}
		}
	}

	for _, e := range input {
		tmp := e % t
		flag := isExist(tmp, res)
		if flag {
			for i := 1; i <= m; i++ {
				tmp = (e + i*i) % t
				if !isExist(tmp, res) {
					res = append(res, tmp)
					flag = false
					break
				}
			}

			if flag {
				res = append(res, -1)
			}
		} else {
			res = append(res, tmp)
		}
	}

	return res
}

func isExist(e int, arr []int) bool {
	for _, item := range arr {
		if e == item {
			return true
		}
	}

	return false
}
