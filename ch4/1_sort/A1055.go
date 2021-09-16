package main

import (
	"fmt"
)

// the world's richest
// page 106

type person struct {
	name   string
	age    int
	wealth int
}

type query1055 struct {
	amount int
	ageMin int
	ageMax int
}

func theRichest(people []person, queries []query1055) (res []string) {
	for i := 0; i < len(queries); i++ {
		res = append(res, fmt.Sprintf("Case #%d:", i+1))
		var validAgePeople []person
		for _, p := range people {
			if p.age <= queries[i].ageMax && p.age >= queries[i].ageMin {
				validAgePeople = append(validAgePeople, p)
			}
		}
		if len(validAgePeople) == 0 {
			res = append(res, "None")
		} else {
			minAmount := 0
			if len(validAgePeople) < queries[i].amount {
				minAmount = len(validAgePeople)
			} else {
				minAmount = queries[i].amount
			}
			validAgePeople = sortPeople(validAgePeople)
			for j := 0; j < minAmount; j++ {
				res = append(res, fmt.Sprintf("%s %d %d", validAgePeople[j].name, validAgePeople[j].age, validAgePeople[j].wealth))
			}
		}
	}
	return res
}

func sortPeople(people []person) []person {
	for i := 1; i < len(people); i++ {
		temp, k := people[i], i
		for ; k > 0 && cmp1055(people[k-1], temp); k-- {
			people[k] = people[k-1]
		}
		people[k] = temp
	}
	return people
}

func cmp1055(a, b person) bool {
	if a.wealth != b.wealth {
		return a.wealth < b.wealth
	} else if a.age != b.age {
		return a.age > b.age
	} else {
		return a.name > b.name
	}
}
