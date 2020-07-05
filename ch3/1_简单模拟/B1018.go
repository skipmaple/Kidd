package main

import "fmt"

// 锤子剪刀布
// page 15

func main() {
	var count int
	_, _ = fmt.Scanf("%d", &count)
	a := make(map[int]string, count)
	b := make(map[int]string, count)
	var aWin, aLose, aTie int
	var bWin, bLose, bTie int
	var aMostWinGestureCount int
	var bMostWinGestureCount int
	var aMostWinGesture string
	var bMostWinGesture string
	aWGesture := make(map[string]int)
	bWGesture := make(map[string]int)

	for i := 0; i < count; i++ {
		var tempA, tempB string
		_, _ = fmt.Scanf("%s %s", &tempA, &tempB)
		a[i], b[i] = tempA, tempB
		if a[i] == "C" && b[i] == "J" || a[i] == "J" && b[i] == "B" || a[i] == "B" && b[i] == "C" {
			aWin++
			aWGesture[a[i]]++
			bLose++
		} else if a[i] == "C" && b[i] == "B" || a[i] == "J" && b[i] == "C" || a[i] == "B" && b[i] == "J" {
			aLose++
			bWin++
			bWGesture[b[i]]++
		} else {
			aTie++
			bTie++
		}
	}

	for k, v := range aWGesture {
		if v >= aMostWinGestureCount {
			if aMostWinGesture == "" {
				aMostWinGesture = k
			} else if aMostWinGesture > k {
				aMostWinGesture = k
			}
		}
	}

	for k, v := range bWGesture {
		if v >= bMostWinGestureCount {
			if bMostWinGesture == "" {
				bMostWinGesture = k
			} else if bMostWinGesture > k {
				bMostWinGesture = k
			}
		}
	}

	fmt.Printf("%d %d %d\n", aWin, aTie, aLose)
	fmt.Printf("%d %d %d\n", bWin, bTie, bLose)
	fmt.Printf("%s %s\n", aMostWinGesture, bMostWinGesture)
}
