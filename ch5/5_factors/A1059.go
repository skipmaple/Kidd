package __factors

import (
	"fmt"
	"math"
)

// Prime Factors
// Page 221

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	sqr := int(math.Sqrt(float64(n)))
	for i := 2; i < sqr; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

const maxN = 100010

var pNum = 0

// 求素数表
func findPrime() (primes [maxN]int) {
	for i := 1; i < maxN; i++ {
		if isPrime(i) == true {
			primes[pNum] = i
			pNum++
		}
	}
	return primes
}

type factor struct {
	// x 为质因子
	x int
	// cnt 为其个数
	cnt int
}

func primeFactors(n int) (res string) {
	fac := [10]factor{}
	primes := findPrime()
	// num 为n的不同质因子的个数
	num := 0
	if n == 1 {
		res = "1=1"
	} else {
		res += fmt.Sprintf("%d=", n)
		sqr := int(math.Sqrt(float64(n)))
		for i := 0; i < pNum && primes[i] <= sqr; i++ {
			if n%primes[i] == 0 {
				// 记录该质因子
				fac[num].x = primes[i]
				fac[num].cnt = 0
				// 计算出质因子prime[i]的个数
				for n%primes[i] == 0 {
					fac[num].cnt++
					n /= primes[i]
				}
				// 不同质因子个数+1
				num++
			}
			if n == 1 {
				break
			}
		}

		// 如果无法被根号n以内的质因子除尽
		if n != 1 {
			// 那么一定有一个大于根号n的质因子
			fac[num].x = n
			fac[num].cnt = 1
			num++
		}

		// 按格式输出结果
		for i := 0; i < num; i++ {
			if i > 0 {
				res += " * "
			}
			res += fmt.Sprintf("%d", fac[i].x)
			if fac[i].cnt > 1 {
				res += fmt.Sprintf("^%d", fac[i].cnt)
			}
		}
	}

	return res
}
