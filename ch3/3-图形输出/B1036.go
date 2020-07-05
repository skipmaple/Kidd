// 跟奥巴马一起编程
// page 43

// 输出给定字符组成的正方形

package main

import (
	"fmt"
	"io"
)

func main() {
	var columnCount, rowCount int
	var ch string
	//fmt.Println("请输入 边长 和 打印字符: ")
	for {
		_, err := fmt.Scanf("%d %s", &columnCount, &ch)
		if err != io.EOF {
			if columnCount%2 == 0 {
				rowCount = columnCount / 2
			} else {
				rowCount = (columnCount + 1) / 2
			}

			for i := 0; i < rowCount; i++ {
				if i == 0 || i == rowCount-1 {
					for j := 0; j < columnCount; j++ {
						fmt.Printf(ch)
					}
				} else {
					for j := 0; j < columnCount; j++ {
						if j == 0 || j == columnCount-1 {
							fmt.Printf(ch)
						} else {
							fmt.Printf(" ")
						}
					}
				}
				fmt.Printf("\n")
			}
		} else {
			break
		}
	}

}
