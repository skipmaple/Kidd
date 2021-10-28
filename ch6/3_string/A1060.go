package __string

// Are They Equal
// Page 241

func areTheyEqual(a, b string) (res string) {
	return res
}

func processNum(nStr string) (res string) {
	for i := 0; i < len(nStr); i++ {
		if nStr[i] == '0' {
			continue
		}

		res += string(nStr[i])

	}

	return res
}
