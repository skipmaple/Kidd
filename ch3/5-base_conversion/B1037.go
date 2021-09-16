// 在霍格茨堡找零钱
// page 51

// 29 knut = 1 sickle;
//17 sickle = 1 galleon;

// or 1 g = 17 s = 17 * 29 k

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func lookForCash(p, a string) (c string) {
	const Knut int = 1
	const Sickle int = 29 * Knut
	const Galleon int = 17 * 29 * Knut

	pGalleon, pSickle, pKnut := strings.Split(p, ".")[0], strings.Split(p, ".")[1], strings.Split(p, ".")[2]
	aGalleon, aSickle, aKnut := strings.Split(a, ".")[0], strings.Split(a, ".")[1], strings.Split(a, ".")[2]

	pG, _ := strconv.Atoi(pGalleon)
	pS, _ := strconv.Atoi(pSickle)
	pK, _ := strconv.Atoi(pKnut)

	aG, _ := strconv.Atoi(aGalleon)
	aS, _ := strconv.Atoi(aSickle)
	aK, _ := strconv.Atoi(aKnut)

	pSum := pG*Galleon + pS*Sickle + pK*Knut
	aSum := aG*Galleon + aS*Sickle + aK*Knut

	diff := aSum - pSum

	dG := diff / Galleon
	diff %= Galleon

	dS := diff / Sickle
	diff %= Sickle

	dK := diff / Knut

	dS, dK = int(math.Abs(float64(dS))), int(math.Abs(float64(dK)))

	c = fmt.Sprintf("%d.%d.%d", dG, dS, dK)
	return c
}
