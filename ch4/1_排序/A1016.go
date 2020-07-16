package main

import "fmt"

// Phone Bills
// page 94

// 这个题 感觉我写的复杂了，后续复习试试能不能优化下

// 通话记录
type phoneRecord struct {
	name       string
	month      int
	timeRecord timeRecord
	status     bool
}

// 有效通话记录
type validRecord struct {
	name         string
	month        int
	childRecords []validChildRecord
	totalAmount  float64
}

type validChildRecord struct {
	start    timeRecord
	over     timeRecord
	duration int // 通话时长
	amount   float64
}

type timeRecord struct {
	day    int
	hour   int
	minute int
}

func printPhoneBills(fee []int, records []phoneRecord) (res []string) {
	var validRecords []validRecord

	// 先排序
	records = sortPhoneRecord(records)
	//fmt.Println("ordered: ", records)
	for i := 1; i < len(records); i++ {
		// 遍历排好序的记录
		// 连续两条同名 并且 同月 状态由通话中转为挂断 则生成一条有效记录
		if records[i].name == records[i-1].name && records[i].month == records[i-1].month && records[i-1].status == true && records[i].status == false {
			duration, amount := calcFee(fee, records[i-1].timeRecord, records[i].timeRecord) // 计算通话时长，和收费
			existFlag := false                                                               // 之前是否有记录flag
			for _, v := range validRecords {
				if v.name == records[i].name {
					existFlag = true
					break
				}
			}
			if existFlag { // 之前有记录
				for index, v := range validRecords {
					if v.name == records[i].name {
						tempChildRecords := v.childRecords
						tempChildRecords = append(tempChildRecords, validChildRecord{
							start:    records[i-1].timeRecord,
							over:     records[i].timeRecord,
							duration: duration,
							amount:   amount,
						})
						validRecords[index].childRecords = tempChildRecords // 这里学到了 通过v.childRecords = someValue并不能修改值，可以借助下标index来修改
						break
					}
				}
			} else { // 之前没有记录
				validRecords = append(validRecords, validRecord{
					name:  records[i].name,
					month: records[i].month,
					childRecords: []validChildRecord{{
						start:    records[i-1].timeRecord,
						over:     records[i].timeRecord,
						duration: duration,
						amount:   amount,
					}},
				})
			}
		}
	}

	//fmt.Println("valid records: ", validRecords )

	// 计算个人的总费
	for _, v := range validRecords {
		res = append(res, fmt.Sprintf("%s %02d", v.name, v.month))
		for _, childRecord := range v.childRecords {
			res = append(res, fmt.Sprintf("%02d:%02d:%02d %02d:%02d:%02d %d $%.2f", childRecord.start.day, childRecord.start.hour, childRecord.start.minute, childRecord.over.day, childRecord.over.hour, childRecord.over.minute, childRecord.duration, childRecord.amount))
			v.totalAmount += childRecord.amount
		}
		res = append(res, fmt.Sprintf("Total amount: $%.2f", v.totalAmount))
	}

	return res
}

func calcFee(fee []int, start, over timeRecord) (duration int, amount float64) {
	for start.day < over.day || start.hour < over.hour || start.minute < over.minute {
		amount += float64(fee[start.hour]) / 100.0 // 美分转换成美元
		duration++
		start.minute++
		if start.minute == 60 {
			start.minute = 0
			start.hour++
		}
		if start.hour == 24 {
			start.hour = 0
			start.day++
		}
	}
	return duration, amount
}

func sortPhoneRecord(records []phoneRecord) []phoneRecord {
	for i := 0; i < len(records)-1; i++ {
		for j := 0; j < len(records)-1-i; j++ {
			if cmpPhoneRecord(records[j], records[j+1]) {
				records[j], records[j+1] = records[j+1], records[j]
			}
		}
	}

	return records
}

func cmpPhoneRecord(a, b phoneRecord) bool {
	if a.name != b.name {
		return a.name > b.name // 名字字母序小的在前
	} else if a.month != b.month {
		return a.month > b.month // 月份小的在前
	} else if a.timeRecord.day != b.timeRecord.day {
		return a.timeRecord.day > b.timeRecord.day
	} else if a.timeRecord.hour != b.timeRecord.hour { // 小时小的
		return a.timeRecord.hour > b.timeRecord.hour
	} else if a.timeRecord.minute != b.timeRecord.minute { // 分钟小的
		return a.timeRecord.minute > b.timeRecord.minute
	}
	return false
}
