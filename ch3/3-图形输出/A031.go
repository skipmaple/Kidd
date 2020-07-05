// Hello World for U
// page 46

package main

import (
	"fmt"
)

func uShape(s string) {
	var lLen, bLen int

	// lLen = rLen <= bLen
	// bLen >= 3 && bLen <= len(s)
	// 2 * lLen + bLen - 2 = len(s)

	bLen = len(s)/3 + len(s)%3
	lLen = len(s) / 3

	for i := 0; i < lLen; i++ {
		fmt.Printf("%c", s[i])
		for j := 0; j < bLen-2; j++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%c\n", s[len(s)-1-i])
	}

	fmt.Printf("%s\n", s[lLen:lLen+bLen])
}

func main() {
	//uShape("nihaoshijie")
	uShape("hello")
	//uShape("comeoncomeoncomeon!!")
}
