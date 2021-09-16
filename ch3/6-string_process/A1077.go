package main

// Kuchiguse
// page 80

// 求公共后缀

func kuchiguse(strs []string) (res string) {

	for i, str := range strs {
		strs[i] = reverseStr(str)
	}

	minLen := len(strs[0])
	var minStr string
	for _, str := range strs {
		if len(str) <= minLen {
			minLen = len(str)
			minStr = str
		}
	}

	//fmt.Println(minStr, minLen)

	for i := 0; i < minLen; i++ {
		flag := true
		for _, str := range strs {
			//fmt.Println(minStr[i],str[i], str[i]==minStr[i])
			if str[i] != minStr[i] {
				flag = false
				break
			}
		}

		//fmt.Println("flag: ", flag)
		if !flag {
			if i == 0 {
				res = "nai"
			} else {
				res = reverseStr(minStr[:i])
			}
			break
		}
	}

	return res
}

func reverseStr(str string) (res string) {
	runeStr := []rune(str)
	for i, j := 0, len(runeStr)-1; i < j; i, j = i+1, j-1 {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	}

	return string(runeStr)
}

//func main() {
//	s := []string{"Itai nyan~", "Ninjin wa iyadanyan~", "uhhh nyan~"}
//	res := kuchiguse(s)
//	fmt.Println(res)
//}
