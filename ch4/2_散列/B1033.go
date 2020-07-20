package hash

import (
	"strings"
)

// 旧键盘打字
// page 130

func oldKeyboard(broken, input string) (output string) {
	shiftKeyFlag := strings.Contains(broken, "+")

	// // 性能不好
	//output = input
	//for _, s := range input {
	//	if strings.Contains(broken, strings.ToUpper(string(s))) || (shiftKeyFlag && string(s) >= "A" && string(s) <= "Z"){
	//		output = strings.ReplaceAll(output, string(s), "")
	//	}
	//}

	// 升级性能
	uniqChars := make(map[string]bool)
	for _, s := range input {
		if uniqChars[string(s)] {
			continue
		} else {
			uniqChars[string(s)] = true
		}
	}

	output = input
	for k, _ := range uniqChars {
		if strings.Contains(broken, strings.ToUpper(k)) || (shiftKeyFlag && k >= "A" && k <= "Z") {
			output = strings.Replace(output, k, "", -1)
			//output = strings.ReplaceAll(output, k, "") // 可以用这个，pat的go版本太老了，没有这个方法
		}
	}

	return output
}

//func main() {
//	var broken, input string
//	_, _ = fmt.Scanf("%s", &broken)
//	_, _ = fmt.Scanf("%s", &input)
//	fmt.Println(oldKeyboard(broken, input))
//}
