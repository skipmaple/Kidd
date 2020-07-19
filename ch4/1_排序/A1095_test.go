package main

import (
	"reflect"
	"testing"
)

func Test_carsOnCampus(t *testing.T) {
	type args struct {
		carRecords []carRecord
		queries    []query
	}

	carRecords1 := []carRecord{
		{plateNum: "JH007BD",  timeStr: "18:00:01", status: true},
		{plateNum: "ZD00001",  timeStr: "11:30:08", status: false},
		{plateNum: "DB8888A",  timeStr: "13:00:00", status: false},
		{plateNum: "ZA3Q625",  timeStr: "23:59:50", status: false},
		{plateNum: "ZA133CH",  timeStr: "10:23:00", status: true},
		{plateNum: "ZD00001",  timeStr: "04:09:59", status: true},
		{plateNum: "JH007BD",  timeStr: "05:09:59", status: true},
		{plateNum: "ZA3Q625",  timeStr: "11:42:01", status: false},
		{plateNum: "JH007BD",  timeStr: "05:10:33", status: true},
		{plateNum: "ZA3Q625",  timeStr: "06:30:50", status: true},
		{plateNum: "JH007BD",  timeStr: "12:23:42", status: false},
		{plateNum: "ZA3Q625",  timeStr: "23:55:00", status: true},
		{plateNum: "JH007BD",  timeStr: "12:24:23", status: false},
		{plateNum: "ZA133CH",  timeStr: "17:11:22", status: false},
		{plateNum: "JH007BD",  timeStr: "18:07:01", status: false},
		{plateNum: "DB8888A",  timeStr: "06:30:50", status: true},
	}

	queries1 := []query{
		{time: "05:10:00"},
		{time: "06:30:50"},
		{time: "11:00:00"},
		{time: "12:23:42"},
		{time: "14:00:00"},
		{time: "18:00:00"},
		{time: "23:59:00"},
	}
	
	arg1 := args{
		carRecords: carRecords1,
		queries:    queries1,
	}
	
	res1 := []string{"1", "4", "5", "2", "1", "0", "1", "JH007BD ZD00001 07:20:09"}

	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		// TODO: Add test cases.
		{"t1", arg1, res1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := carsOnCampus(tt.args.carRecords, tt.args.queries); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("carsOnCampus() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}