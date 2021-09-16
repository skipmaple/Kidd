package __two_pointers

// Find Coins
// page 183

func twoPointers(arr []int, sum int) (a, b int) {
	for i := 1; i < len(arr); i++ {
		temp, j := arr[i], i
		for ; j > 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}

	//fmt.Printf("%v", arr)

	i, j := 0, len(arr)-1
	for i < j {
		if arr[i]+arr[j] == sum {
			break
		} else if arr[i]+arr[j] < sum {
			i++
		} else {
			j--
		}
	}

	//fmt.Println(i, j)

	if i < j {
		// have solution
		//println("%d %d", arr[i], arr[j])
		return arr[i], arr[j]
	} else {
		// no solution
		//println("No Solution")
		return -1, -1
	}

}
