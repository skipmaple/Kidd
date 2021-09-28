package __simple_math

// page 195
// 数列片段和

func sliceSum(input []float64) (res float64) {
	for i := 1; i <= len(input); i++ {
		res += input[i-1] * float64(i) * float64(len(input)+1-i)
	}
	return res
}
