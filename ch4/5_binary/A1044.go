package ch4

// Shopping in Mars
// page 171

func shoppingInMars(arr []int, s int) (res [][]int) {
	upperRes := [][]int{}
	sumArr := []int{}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		sumArr = append(sumArr, sum)
	}

	//fmt.Println("arr   ", arr)
	//fmt.Println("sumArr", sumArr)
	//fmt.Println("0-4", LRSum(2,2,sumArr))

	for i := 0; i < len(arr); i++ {
		r := binarySearch(sumArr, i, len(arr), s)

		//if i == 2{
		//	fmt.Println("i = 2", i, r, LRSum(i,r, sumArr))
		//}

		if LRSum(i, r-1, sumArr) == s {
			res = append(res, []int{i + 1, r})
		} else if r < len(arr) {
			upperRes = append(upperRes, []int{i + 1, r + 1})
		}
	}

	if len(res) != 0 {
		return res
	} else {
		return upperRes
	}

}

func binarySearch(arr []int, left, right, s int) int {
	if left == right {
		if LRSum(0, right, arr) > s {
			return left
		} else {
			return left + 1
		}
	}
	start := left
	for left < right {
		middle := (left + right) / 2
		if LRSum(start, middle, arr) > s {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return left
}

func LRSum(l, r int, arr []int) int {
	if r <= 0 {
		return arr[0]
	} else if l == 0 {
		return arr[r]
	} else {
		return arr[r] - arr[l-1]
	}
}
