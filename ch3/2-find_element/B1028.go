package main

import (
	"fmt"
	"time"
)

// 人口普查
// page 32

type person struct {
	name     string
	birthday string
}

func humanCount(list []person) (int, string, string) {
	var validNum int
	var youngest, oldest person
	nowadays := time.Now().UTC()
	longLongAgo := nowadays.AddDate(-200, 0, 0)

	y, _ := time.Parse("2006/01/02", longLongAgo.String())
	fmt.Println("young", y)
	o, _ := time.Parse("2006/01/02", nowadays.Format("2006/01/02"))

	for _, p := range list {
		birthday, _ := time.Parse("2006/01/02", p.birthday)
		if birthday.Before(longLongAgo) || birthday.After(nowadays) {
			continue
		} else {
			validNum++
			if birthday.Before(o) {
				o = birthday
				oldest = p
			}
			if birthday.After(y) {
				y = birthday
				youngest = p
			}
		}
	}

	return validNum, oldest.name, youngest.name
}
