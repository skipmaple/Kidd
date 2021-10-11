package __fraction_calc

import (
	"Kidd/utils"
	"fmt"
	"math"
)

// Rational Sum
// 求分数和
// Page 203

// 求最大公约数
// utils.Gcd

type Fraction struct {
	up   int // 分子
	down int // 分母
}

// 化简
func reduction(input Fraction) Fraction {
	if input.down < 0 {
		input.up = -input.up
		input.down = -input.down
	}
	if input.up == 0 {
		input.down = 1
	} else {
		d := utils.Gcd(int(math.Abs(float64(input.up))), int(math.Abs(float64(input.down))))
		input.up /= d
		input.down /= d
	}
	return input
}

func add(a, b Fraction) (res Fraction) {
	a = reduction(a)
	b = reduction(b)
	res.up = a.up*b.down + a.down*b.up
	res.down = a.down * b.down
	return reduction(res)
}

func showResult(input Fraction) (res string) {
	input = reduction(input)
	integer := input.up / input.down
	fraction := Fraction{input.up - integer*input.down, input.down}
	if fraction.up == 0 {
		res = fmt.Sprintf("%d", integer)
	} else if integer > 0 {
		res = fmt.Sprintf("%d %d/%d", integer, fraction.up, fraction.down)
	} else if integer < 0 {
		res = fmt.Sprintf("%d %d/%d", integer, int(math.Abs(float64(fraction.up))), fraction.down)
	} else {
		res = fmt.Sprintf("%d/%d", fraction.up, fraction.down)
	}
	return res
}

func FractionSum(arr []Fraction) (res string) {
	sum := Fraction{0, 1}
	for _, fraction := range arr {
		sum = add(sum, fraction)
	}
	res = showResult(sum)
	return res
}
