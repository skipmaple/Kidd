package main

import (
	"reflect"
	"testing"
)

func Test_patRank(t *testing.T) {
	type args struct {
		stus     []student1075
		problems []problem
		records  []problemRecord
	}

	problems1 := []problem{
		{1, 20},
		{2, 25},
		{3, 25},
		{4, 30},
	}

	// !!! 这里是之前犯错的地方

	//var records []problemRecord
	//for i := 0; i < len(problems1); i++ {
	//	records = append(records, problemRecord{
	//		pid:   problems1[i].pid,
	//		score: -2,
	//	})
	//}
	//
	//stus1 := []student1075{
	//	{sid: 00001, records: records},
	//	{sid: 00002, records: records},
	//	{sid: 00003, records: records},
	//	{sid: 00004, records: records},
	//	{sid: 00005, records: records},
	//	{sid: 00006, records: records},
	//	{sid: 00007, records: records}, // ！！！错在所有records都是同一个内存！正确的写法应该 深拷贝
	//}

	stus1 := []student1075{
		{sid: 00001},
		{sid: 00002},
		{sid: 00003},
		{sid: 00004},
		{sid: 00005},
		{sid: 00006},
		{sid: 00007},
	}

	records1 := []problemRecord{
		{00002, 2, 12},
		{00007, 4, 17},
		{00005, 1, 19},
		{00007, 2, 25},
		{00005, 1, 20},
		{00002, 2, 2},
		{00005, 1, 15},
		{00001, 1, 18},
		{00004, 3, 25},
		{00002, 2, 25},
		{00005, 3, 22},
		{00006, 4, -1},
		{00001, 2, 18},
		{00002, 1, 20},
		{00004, 1, 15},
		{00002, 4, 18},
		{00001, 3, 4},
		{00001, 4, 2},
		{00005, 2, -1},
		{00004, 2, 0},
	}

	for i := 0; i < len(stus1); i++ {
		var records []problemRecord
		for i := 0; i < len(problems1); i++ {
			records = append(records, problemRecord{
				sid:   stus1[i].sid,
				pid:   problems1[i].pid,
				score: -2,
			})
		}
		stus1[i].records = records
	}

	arg1 := args{
		stus:     stus1,
		problems: problems1,
		records:  records1,
	}
	ans1 := []string{"1 00002 63 20 25 - 18", "2 00005 42 20 0 22 -", "2 00007 42 - 25 - 17", "2 00001 42 18 18 4 2", "5 00004 40 15 0 25 -"}

	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"t1", arg1, ans1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := patRank(tt.args.stus, tt.args.problems, tt.args.records); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("patRank() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
