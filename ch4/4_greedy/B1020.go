package ch4

// 月饼
// main page 118

type moonCake struct {
	store float64
	sell  float64
	price float64
}

func yueBing(need float64, haves []float64, sells []float64) (max float64) {
	moonCakes := []moonCake{}
	for i := 0; i < len(haves); i++ {
		moonCakes = append(moonCakes, moonCake{
			store: haves[i],
			sell:  sells[i],
			price: sells[i] / haves[i],
		})
	}
	moonCakes = sortPrices(moonCakes)

	for need > 0 {
		for i := 0; i < len(moonCakes); i++ {
			if moonCakes[i].store >= need {
				max += moonCakes[i].price * need
				need = 0
				break
			} else {
				max += moonCakes[i].sell
				need -= moonCakes[i].store
			}
		}
	}
	return max
}

func sortPrices(moonCakes []moonCake) []moonCake {
	for i := 0; i < len(moonCakes)-1; i++ {
		for j := 0; j < len(moonCakes)-1-i; j++ {
			if moonCakes[j].price < moonCakes[j+1].price {
				moonCakes[j+1], moonCakes[j] = moonCakes[j], moonCakes[j+1]
			}
		}
	}
	return moonCakes
}
