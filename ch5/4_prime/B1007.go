package __prime

import "math"

// prime
// page 209

// isPrime give a integer number judge prime
func isPrime(num int) bool {
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

func countPrime(num int) (res int) {
	for i := 3; i+2 <= num; i += 2 {
		if isPrime(i) == true && isPrime(i+2) == true {
			res += 1
		}
	}

	return res
}
