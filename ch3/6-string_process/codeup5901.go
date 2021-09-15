package main

// 判断回文字符串
//func main() {
//fmt.Printf("%v", isHuiwen("noon"))
//}

func isHuiwen(s string) bool {
	bs := []byte(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if bs[i] != bs[j] {
			return false
		}
	}
	return true

}
