package main

import "fmt"

// 说反话
func main() {
	s := "another day another chance"
	fmt.Println(talkReverse(s))
}

func talkReverse(s string) string {
	bs := []byte(s)
	var res string
	word := make([]byte, 0)
	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] != ' ' {
			word = append(word, bs[i])
		} else {
			word = reverseWord(word)
			res += string(word) + " "
			word = make([]byte, 0)
		}

		if i == 0 {
			word = reverseWord(word)
			res += string(word)
		}
	}
	return res
}

func reverseWord(w []byte) []byte {
	for j, k := 0, len(w)-1; j < k; j, k = j+1, k-1 {
		w[j], w[k] = w[k], w[j]
	}
	return w
}
