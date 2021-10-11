package __fraction_calc

import (
	"fmt"
	"math"
)

// Rational Arithmetic
// 求两个分数的和 差 积 商
// Page 205

func fmtResult(input Fraction) (res string) {
	input = reduction(input)
	integer := input.up / input.down
	fraction := Fraction{input.up - integer*input.down, input.down}
	if fraction.up == 0 {
		if integer >= 0 {
			res = fmt.Sprintf("%d", integer)
		} else {
			res = fmt.Sprintf("(%d)", integer)
		}
	} else if integer > 0 {
		res = fmt.Sprintf("%d %d/%d", integer, fraction.up, fraction.down)
	} else if integer < 0 {
		res = fmt.Sprintf("(%d %d/%d)", integer, int(math.Abs(float64(fraction.up))), fraction.down)
	} else if fraction.up < 0 {
		res = fmt.Sprintf("(%d/%d)", fraction.up, fraction.down)
	} else {
		res = fmt.Sprintf("%d/%d", fraction.up, fraction.down)
	}
	return res
}

// multiply 计算乘积
func multiply(a, b Fraction) (res Fraction) {
	a = reduction(a)
	b = reduction(b)
	res.up = a.up * b.up
	res.down = a.down * b.down
	return reduction(res)
}

func RationalArithmetic(a, b Fraction) (res string) {
	strA := fmtResult(a)
	strB := fmtResult(b)

	sum := add(a, b)
	res += fmt.Sprintf("%s + %s = %s\n", strA, strB, fmtResult(sum))

	minus := add(a, Fraction{-b.up, b.down})
	res += fmt.Sprintf("%s - %s = %s\n", strA, strB, fmtResult(minus))

	mul := multiply(a, b)
	res += fmt.Sprintf("%s * %s = %s\n", strA, strB, fmtResult(mul))

	if b.up == 0 {
		res += fmt.Sprintf("%s / %s = Inf\n", strA, strB)
	} else {
		div := multiply(a, Fraction{b.down, b.up})
		res += fmt.Sprintf("%s / %s = %s\n", strA, strB, fmtResult(div))
	}
	return res
}
