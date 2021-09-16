package ch4

import "fmt"

// Find Coins
// page 175
// question see page 146

func findCoins(arr []int, s int) string {
	for i := 0; i < len(arr); i++ {
		k := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		r := binarySearch1048(i, len(arr)-1, arr, s)
		if r != -1 && r != i {
			fmt.Println(i, r)
			return fmt.Sprintf("%d %d", arr[i], arr[r])
		}
	}

	return "No Solution"
}

func binarySearch1048(left, right int, arr []int, s int) int {
	start := left
	for left <= right {
		middle := (left + right) / 2
		if arr[start]+arr[middle] == s {
			return middle
		} else if arr[start]+arr[middle] > s {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
