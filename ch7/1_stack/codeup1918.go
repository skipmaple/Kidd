package __stack

import (
	"Kidd/utils"
	"math"
	"strconv"
	"strings"
)

// 简单计算器
// main page 247

// 输入格式： 整数和运算符之间用一个空格分隔

type node struct {
	num   float64 // 操作数
	op    string  // 操作符
	isNum bool    //  true 表示操作数, false 表示操作符
}

var stack *utils.Stack

func init() {
	stack = utils.NewStack()
}

func simpleCalculator(input string) (res float64) {
	middle := strToMiddle(input)
	back := middleToBack(middle)
	res = calcBack(back)

	return res
}

// 字符串转中缀表达式
func strToMiddle(input string) (res utils.Stack) {

	tmpArr := strings.Fields(input)
	for _, item := range tmpArr {
		tmpNode := node{}
		if strings.Contains("+-*/", item) {
			tmpNode.isNum = false
			tmpNode.op = item
		} else {
			tmpNode.isNum = true
			tmpNum, _ := strconv.ParseFloat(item, 64)
			tmpNode.num = tmpNum
		}
		res.Push(tmpNode)
	}

	return res
}

// 中缀表达式转后缀表达式
func middleToBack(middle utils.Stack) (back utils.Stack) {
	tmpOpStack := utils.NewStack()
	for i := 0; i < len(middle.Items); i++ {
		currentNode := middle.Items[i]
		if currentNode, ok := currentNode.(node); ok {
			if currentNode.isNum {
				//	操作数压入后缀表达式中
				back.Push(currentNode)
			} else {
				//	操作符压入临时操作符栈中
				if tmpOpStack.IsEmpty() {
					tmpOpStack.Push(currentNode)
					continue
				}

				for !tmpOpStack.IsEmpty() {
					tmpOpTop, _ := tmpOpStack.TopValue().(node)
					if opCompare(currentNode, tmpOpTop) {
						tmpOpStack.Push(currentNode)
						break
					} else {
						item := tmpOpStack.Pop()
						back.Push(item)

						// 如果临时操作符栈弹空了，就把当前操作符节点push进去
						if tmpOpStack.IsEmpty() {
							tmpOpStack.Push(currentNode)
							break
						}
					}
				}

			}
		}
	}

	for !tmpOpStack.IsEmpty() {
		back.Push(tmpOpStack.Pop())
	}

	return back
}

// 后缀表达式计算
func calcBack(back utils.Stack) (res float64) {
	tmpResStack := utils.NewStack()
	for i := 0; i < len(back.Items); i++ {
		currentNode := back.Items[i]
		if currentNode, ok := currentNode.(node); ok {
			if currentNode.isNum {
				// 操作数压入临时结果栈
				tmpResStack.Push(currentNode)
			} else {
				// 遇到操作符从临时结果栈弹出两个操作数
				// 后弹出的是第一个操作数， 先弹出的是第二个操作数
				num2Node, _ := tmpResStack.Pop().(node)
				num1Node, _ := tmpResStack.Pop().(node)
				num2 := num2Node.num
				num1 := num1Node.num
				tmpRes := calcAB(num1, num2, currentNode.op)
				tmpResStack.Push(node{
					num:   tmpRes,
					op:    "",
					isNum: true,
				})
			}
		}
	}
	resNode, _ := tmpResStack.Pop().(node)
	// float64 保留两位小数
	// 参考 https://yourbasic.org/golang/round-float-2-decimal-places/
	res = math.Round(resNode.num*100) / 100
	return res
}

// 比较操作符a优先级是否高于b
// a > b ? true : false
func opCompare(a, b node) bool {
	if a.isNum || b.isNum {
		return false
	}

	if strings.Contains("+-", b.op) && strings.Contains("*/", a.op) {
		return true
	}
	return false
}

func calcAB(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	return 0.0
}
