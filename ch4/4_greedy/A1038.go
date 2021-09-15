package ch4

import "strconv"

// Recover the Smallest Number
// page162

func smallestNumber(arr []string) (res int) {

	for i:= 1;i<len(arr);i++{
		temp, j := arr[i], i
		for ;j>0&&cmp1038(temp, arr[j-1]);j--{
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}

	resStr := ""
	for i:= 0;i< len(arr);i++{
		resStr += arr[i]
	}

	res,_ = strconv.Atoi(resStr)

	return res
}

func cmp1038(a,b string) bool {
	return a+b < b+a
}
