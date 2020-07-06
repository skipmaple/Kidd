package main

import "strconv"

// 写出这个数
// page 63

func sayNumSum(num string) (res []string) {
	numMap := map[int]string{
		0: "ling", 1: "yi", 2: "er", 3: "san", 4: "si", 5: "wu", 6: "liu", 7: "qi", 8: "ba", 9: "jiu",
	}
	sum := 0
	for i := 0; i < len(num); i++ {
		n, _ := strconv.Atoi(string(num[i]))
		sum += n
	}

	// 比如sum=123, 则sumArr = [3, 2, 1]
	var sumArr []int
	for sum!=0{
		sumArr = append(sumArr, sum%10)
		sum /= 10
	}

	for i:= len(sumArr)-1;i>=0;i-- {
		res = append(res, numMap[sumArr[i]])
	}

	return res
}
