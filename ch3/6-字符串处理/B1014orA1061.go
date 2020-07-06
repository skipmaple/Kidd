package main

import "fmt"

// 福尔摩斯的约会
// page 66

func datingDate(secret []string) (res string) {
	week, hour := findWeekAndHour(secret[0], secret[1])
	minute := findMinute(secret[2], secret[3])
	res = fmt.Sprintf("%s %s:%s", week, hour, minute)
	return res
}

func findWeekAndHour(s1, s2 string) (string, string) {
	week := ""
	hour := ""
	weekMap := map[byte]string{
		'A': "MON", 'B': "TUE", 'C': "WED", 'D': "THU", 'E': "FRI", 'F': "SAT", 'G': "SUN",
	}

	hourMap := map[byte]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '9': 9, 'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19, 'K': 20, 'L': 21, 'M': 22, 'N': 23,
	}

	for i := 0; i < len(s1); i++ {
		if week == ""{
			if s1[i] < 'A' || s1[i] > 'G' {
				continue
			}
		} else {
			if s1[i] < 'A' || s1[i] > 'N' {
				continue
			}
		}

		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				if week == "" {
					week = weekMap[s1[i]]
				} else if hour == "" {
					hour = fmt.Sprintf("%02d", hourMap[s1[i]])
				}
				break
			}
		}

		if week != "" && hour != ""{
			break
		}
	}

	return week, hour
}

func findMinute(s1, s2 string) (res string) {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if !((s1[i] > 'a' && s1[i] < 'z') || (s1[i] > 'A' && s1[i] < 'Z')) {
			continue
		}

		if s1[i] == s2[i] {
			res = fmt.Sprintf("%02d", i)
			break
		}
	}

	return res
}
