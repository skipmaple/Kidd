package ch4

import "fmt"

// To Fill or Not Fill
// page 151

type gasStation struct {
	price    float64
	distance float64
}

func toFillOrNot(maxTank, totalDistance, unitCanDrive int, gasStations []gasStation) (res string) {
	currentStation := gasStation{}
	nextStation := gasStation{}
	for i := 0; i < len(gasStations); i++ {
		k := i
		for j := k; j < len(gasStations); j++ {
			if cmpGasStation(gasStations[j], gasStations[k]) {
				k = j
			}
		}
		gasStations[i], gasStations[k] = gasStations[k], gasStations[i]
	}

	//for i := 0; i < len(gasStations); i++ {
	//	fmt.Println(i, gasStations[i])
	//}

	// 不存在起点加油站
	if gasStations[0].distance != 0 {
		res = "The maximum travel distance = 0.00"
		return res
	} else {
		currentStation = gasStations[0]
		nextStation = currentStation
	}

	alreadyDis := 0.0
	money := 0.0

	for i := 1; i < len(gasStations); i++ {
		currentStation = nextStation
		tempDis := alreadyDis + float64(maxTank*unitCanDrive)
		//fmt.Println("maxTank*unitCanDrive = ", tempDis)
		if tempDis < gasStations[i].distance { // 不能到下一个加油站
			res = fmt.Sprintf("The maximum travel distance = %.2f", alreadyDis)
			return res
		} else if gasStations[i].price > currentStation.price { // 下一个加油站 汽油价格大于当前加油站
			for j := i + 1; j < len(gasStations); j++ {
				if gasStations[j].price < gasStations[i].price && tempDis >= gasStations[j].distance {
					i = j
					break
				}
			}
			nextStation = gasStations[i]
		} else {
			nextStation = gasStations[i]
		}

		//fmt.Println("use ", i)
		alreadyDis += nextStation.distance - currentStation.distance
		money += (nextStation.distance - currentStation.distance) / float64(unitCanDrive) * currentStation.price
	}

	if (float64(totalDistance)-alreadyDis)/float64(unitCanDrive) > float64(maxTank) {
		res = fmt.Sprintf("The maximum travel distance = %.2f", alreadyDis+float64(maxTank*unitCanDrive))
		return res
	}
	money += (float64(totalDistance) - alreadyDis) / float64(unitCanDrive) * nextStation.price

	return fmt.Sprintf("%.2f", money)
}

// 距离 单价
func cmpGasStation(a, b gasStation) bool {
	if a.distance != b.distance {
		return a.distance < b.distance
	} else {
		return a.price < b.price
	}
}
