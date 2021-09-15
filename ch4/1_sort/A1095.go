package main

import (
	"fmt"
	"strconv"
	"time"
)

// Cars on Campus
// page 123

type carRecord struct {
	plateNum string
	timeStr  string
	time     time.Time
	status   bool // true equals in, false equals out
}

type validCarRecord struct {
	plateNum  string
	in        time.Time
	out       time.Time
	inFormat  string
	outFormat string
	duration  time.Duration
}

type query struct {
	time  string
	count int
}

func carsOnCampus(carRecords []carRecord, queries []query) (res []string) {
	timeFormat := "15:04:05"
	var validCarRecords []validCarRecord

	for i, record := range carRecords {
		parseTime, _ := time.Parse(timeFormat, record.timeStr)
		//fmt.Println(parseTime)
		carRecords[i].time = parseTime
		//fmt.Println(record, time.Unix(parseTime.Unix(),0).Truncate(time.Second).Format(timeFormat))
	}

	carRecords = sortCarRecords(carRecords)

	for i := 0; i < len(carRecords); i++ {
		fmt.Println(carRecords[i].plateNum, carRecords[i].status, carRecords[i].time.Truncate(time.Second).Format(timeFormat))
	}

	for i := 0; i < len(carRecords)-1; i++ {
		if carRecords[i].plateNum == carRecords[i+1].plateNum && carRecords[i].status && !carRecords[i+1].status {
			validCarRecords = append(validCarRecords, validCarRecord{
				plateNum: carRecords[i].plateNum,
				in:       carRecords[i].time,
				out:      carRecords[i+1].time,
				duration: carRecords[i+1].time.Sub(carRecords[i].time),
			})
		}
	}

	for i := 0; i < len(validCarRecords); i++ {
		validCarRecords[i].inFormat = validCarRecords[i].in.Truncate(time.Second).Format(timeFormat)
		validCarRecords[i].outFormat = validCarRecords[i].out.Truncate(time.Second).Format(timeFormat)
	}

	for i := 0; i < len(queries); i++ {
		parseTime, _ := time.Parse(timeFormat, queries[i].time)
		for _, r := range validCarRecords {
			if (r.in.Before(parseTime) || r.in.Equal(parseTime)) && r.out.After(parseTime) {
				queries[i].count++
			}
		}
		res = append(res, strconv.Itoa(queries[i].count))
	}

	var maxLengthPlatNum string
	var maxLengthTime time.Duration
	for _, r := range validCarRecords {
		fmt.Println(r.plateNum, r.inFormat, r.outFormat, formatToHHMMSS(int64(r.duration.Seconds())))
		if r.duration > maxLengthTime {
			maxLengthTime = r.duration
			maxLengthPlatNum = r.plateNum
		} else if r.duration == maxLengthTime {
			maxLengthPlatNum = fmt.Sprintf("%s %s", maxLengthPlatNum, r.plateNum)
		}
	}

	res = append(res, fmt.Sprintf("%s %s", maxLengthPlatNum, formatToHHMMSS(int64(maxLengthTime.Seconds()))))

	return res
}

func sortCarRecords(carRecords []carRecord) []carRecord {
	for i := 0; i < len(carRecords); i++ {
		k := i
		for j := i; j < len(carRecords); j++ {
			if cmp1095(carRecords[k], carRecords[j]) {
				k = j
			}
		}
		carRecords[i], carRecords[k] = carRecords[k], carRecords[i]
	}
	return carRecords
}

func cmp1095(a, b carRecord) bool {
	if a.plateNum != b.plateNum {
		return a.plateNum > b.plateNum
	} else {
		return a.time.After(b.time)
	}
}

func formatToHHMMSS(t int64) string {
	return fmt.Sprintf("%02d:%02d:%02d", t/60/60, t%(60*60)/60, t%60)
}

// 我瞪着大眼比对了好久。。 没有找见 车牌号JH007BD 停车时长为07:20:09的记录

// my code output: [1 4 5 2 1 0 1 ZD00001 07:20:09],
// want:           [1 4 5 2 1 0 1 JH007BD ZD00001 07:20:09]

// 题目网址 https://pintia.cn/problem-sets/994805342720868352/problems/994805371602845696

// 这是测试用例中 所有 JH007BD 车牌号的记录

// JH007BD 05:09:59 in
// JH007BD 05:10:33 in
// JH007BD 12:23:42 out
// JH007BD 12:24:23 out
// JH007BD 18:00:01 in
// JH007BD 18:07:01 out

// 这是排序后的记录

/*  车牌     进/出  时间点
DB8888A true 06:30:50
DB8888A false 13:00:00
JH007BD true 05:09:59
JH007BD true 05:10:33
JH007BD false 12:23:42
JH007BD false 12:24:23
JH007BD true 18:00:01
JH007BD false 18:07:01
ZA133CH true 10:23:00
ZA133CH false 17:11:22
ZA3Q625 true 06:30:50
ZA3Q625 false 11:42:01
ZA3Q625 true 23:55:00
ZA3Q625 false 23:59:50
ZD00001 true 04:09:59
ZD00001 false 11:30:08
*/

// 这是所有 有效停车记录

/*  车牌     进       出        时长
DB8888A 06:30:50 13:00:00 06:29:10
JH007BD 05:10:33 12:23:42 07:13:09
JH007BD 18:00:01 18:07:01 00:07:00
ZA133CH 10:23:00 17:11:22 06:48:22
ZA3Q625 06:30:50 11:42:01 05:11:11
ZA3Q625 23:55:00 23:59:50 00:04:50
ZD00001 04:09:59 11:30:08 07:20:09
*/
