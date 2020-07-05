package main

import (
	"fmt"
	"strconv"
)

// shuffling machine
// 打乱扑克牌顺序
// page 18

func main() {
	given := make([]int, 54)
	order := make([]int, 54)

	flower := map[int]string{
		0: "S",
		1: "H",
		2: "C",
		3: "D",
		4: "J",
	}

	for i := 0; i < 54; i++ {
		given[i] = i
		//fmt.Printf("%d ", 54 - (i))
	}
	//fmt.Println()

	var count int
	_, _ = fmt.Scanf("%d", &count)
	//newOrder := make([]int, len(given))

	for i:=0;i<54;i++ {
		_, _ = fmt.Scanf("%d", &order[i])
	}

	//fmt.Println(order)

	for i := 0; i < count; i++ {
		given = shuffling(order, given)
	}


	fmt.Println("res: ", given)
	for i, v := range given {
		var res string
		res = flower[v/13] + strconv.Itoa(v%13+1)
		if i != len(given)-1 {
			fmt.Printf("%v ", res)
		} else {
			fmt.Printf("%v\n", res)
		}
	}

}

func shuffling(order []int, given []int) []int {
	newGiven := make([]int, len(given))
	//fmt.Println("order len: ", len(order))
	//fmt.Println("given is: ", given)
	//fmt.Println("order is: ", order)
	for i := 0; i < len(order); i++ {
		newGiven[order[i]-1] = given[i]
	}
	return newGiven
}
