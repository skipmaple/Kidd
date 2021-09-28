package __simple_math

// page197
// elevator

func elevator(input []int) (res int) {
	now, to := 0, 0

	for i := 0; i < len(input); i++ {
		to = input[i]
		if to > now {
			res += (to-now)*6 + 5
		} else if to < now {
			res += (now-to)*4 + 5
		} else {
			res += 5
		}
		now = to
	}
	return res
}
