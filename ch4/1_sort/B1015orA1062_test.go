package main

import (
	"reflect"
	"testing"
)

func Test_deCaiLun(t *testing.T) {
	type args struct {
		lScore   int
		hScore   int
		students []stu
	}

	arg1 := args{
		lScore: 60,
		hScore: 80,
		students: []stu{
			{id: 10000001, deScore: 64, caiScore: 90},
			{id: 10000002, deScore: 90, caiScore: 60},
			{id: 10000003, deScore: 85, caiScore: 80},
			{id: 10000004, deScore: 85, caiScore: 80},
			{id: 10000005, deScore: 80, caiScore: 85},
			{id: 10000006, deScore: 82, caiScore: 77},
			{id: 10000007, deScore: 83, caiScore: 76},
			{id: 10000008, deScore: 90, caiScore: 78},
			{id: 10000009, deScore: 75, caiScore: 79},
			{id: 10000010, deScore: 59, caiScore: 90},
			{id: 10000011, deScore: 88, caiScore: 45},
			{id: 10000012, deScore: 80, caiScore: 100},
			{id: 10000013, deScore: 90, caiScore: 99},
			{id: 10000014, deScore: 66, caiScore: 60},
		},
	}

	ans1 := []stu{
		{10000013, 90, 99, 189},
		{10000012, 80, 100, 180},
		{10000003, 85, 80, 165},
		{10000004, 85, 80, 165},
		{10000005, 80, 85, 165},
		{10000008, 90, 78, 168},
		{10000007, 83, 76, 159},
		{10000006, 82, 77, 159},
		{10000002, 90, 60, 150},
		{10000014, 66, 60, 126},
		{10000009, 75, 79, 154},
		{10000001, 64, 90, 154},
	}

	tests := []struct {
		name      string
		args      args
		wantCount int
		wantRes   []stu
	}{
		// TODO: Add test cases.
		{"t1", arg1, 12, ans1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotRes := deCaiLun(tt.args.lScore, tt.args.hScore, tt.args.students)
			if gotCount != tt.wantCount {
				t.Errorf("deCaiLun() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("deCaiLun() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
