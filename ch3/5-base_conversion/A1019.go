// General Palindromic Number
// page 53

// 输入两个数 N 和 b， 判断N在b进制下是不是回文数

package main

func palindromicNumber(N, b int) (res bool, arr []int) {
	//arrS := utils.TtoQ(N, b)
	//for i := 0; i < len(arrS); i++ {
	//	temp, _ := strconv.Atoi(arrS[i])
	//	arr = append(arr, temp)
	//}

	arr = TtoQ(N, b)

	res = true
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			res = false
		}
	}

	return res, arr
}

func QtoT(n, q int) (res []int) {
	product := 1
	for n != 0 {
		res = append(res, n%10*product)
		product *= q
		n /= q
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func TtoQ(n, q int) (res []int) {
	for n != 0 {
		res = append(res, n%q)
		n /= q
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
