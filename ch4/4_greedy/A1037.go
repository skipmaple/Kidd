package ch4

// Magic Coupon
// page 156

func magicCoupon(a, b []int) (res int) {
	a, b = sortList(a), sortList(b)
	for i := 0; i < len(a) && i < len(b) && a[i] < 0 && b[i] < 0; i++ {
		res += a[i] * b[i]
	}

	for i, j := len(a)-1, len(b)-1; i >= 0 && j >= 0 && a[i] > 0 && b[j] > 0; i, j = i-1, j-1 {
		res += a[i] * b[j]
	}

	return res
}

func sortList(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	return list
}
