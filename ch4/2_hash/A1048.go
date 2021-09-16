package hash

// Find Coins
// page 145

func findCoins(nums []int, sum int) (a, b int) {
	a, b = -1, -1
	numCount := make(map[int]int)
	for i := 1; i < len(nums); i++ {
		temp, j := nums[i], i
		for ; j > 0 && temp < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = temp
	}

	for _, num := range nums {
		numCount[num]++
	}

	for i := 0; i < len(nums); i++ {
		if numCount[nums[i]] > 0 && numCount[sum-nums[i]] > 0 {
			if nums[i] != sum-nums[i] || (nums[i] == sum-nums[i] && numCount[nums[i]] >= 2) {
				a = nums[i]
				b = sum - nums[i]
				break
			}
		}
	}
	return a, b
}
