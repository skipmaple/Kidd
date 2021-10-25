package __prime

import "math"

// list prime numbers
// page 211

func isPrimeNum(num int) bool {
	if num <= 1 {
		return false
	}

	sqr := int(math.Sqrt(float64(1.0 * num)))
	for i := 2; i <= sqr; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// 暴力破解法
func allPrime() (res []int) {
	maxN := 10000
	for i := 1; i <= maxN; i++ {
		if isPrimeNum(i) {
			res = append(res, i)
		}
	}

	return res
}

func listPrimeNum(begin, end int) (res []int) {
	allPrime := allPrime()
	for i := begin; i <= end; i++ {
		res = append(res, allPrime[i-1])
	}
	return res
}

// 筛法 (???)
func findPrime(num int) []int {
	maxN := 1000001
	p := make([]bool, maxN)
	prime := make([]int, maxN)
	index := 0
	for i := 2; i < maxN; i++ {
		if p[i] == false {
			prime[index] = i
			index += 1

			if index > num {
				break
			}

			for j := i + i; j < maxN; j += i {
				p[j] = true
			}
		}
	}

	return prime
}

func listPrimeNum2(begin, end int) (res []int) {
	primes := findPrime(end)
	for i := begin; i <= end; i++ {
		res = append(res, primes[i-1])
	}
	return res
}
