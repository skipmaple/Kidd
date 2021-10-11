package main

import (
	"reflect"
	"testing"
)

func Test_graduateAdmission(t *testing.T) {
	type args struct {
		stus    []stu1080
		schools []school
	}

	stus1 := []stu1080{
		{id: 0, geScore: 100, giScore: 100, schoolIds: []int{0, 1, 2}},
		{id: 1, geScore: 60, giScore: 60, schoolIds: []int{2, 3, 5}},
		{id: 2, geScore: 100, giScore: 90, schoolIds: []int{0, 3, 4}},
		{id: 3, geScore: 90, giScore: 100, schoolIds: []int{1, 2, 0}},
		{id: 4, geScore: 90, giScore: 90, schoolIds: []int{5, 1, 3}},
		{id: 5, geScore: 80, giScore: 90, schoolIds: []int{1, 0, 2}},
		{id: 6, geScore: 80, giScore: 80, schoolIds: []int{0, 1, 2}},
		{id: 7, geScore: 80, giScore: 80, schoolIds: []int{0, 1, 2}},
		{id: 8, geScore: 80, giScore: 70, schoolIds: []int{1, 3, 2}},
		{id: 9, geScore: 70, giScore: 80, schoolIds: []int{1, 2, 3}},
		{id: 10, geScore: 100, giScore: 100, schoolIds: []int{0, 2, 4}},
	}

	schools1 := []school{
		{id: 0, acceptCount: 2},
		{id: 1, acceptCount: 1},
		{id: 2, acceptCount: 2},
		{id: 3, acceptCount: 2},
		{id: 4, acceptCount: 2},
		{id: 5, acceptCount: 3},
	}

	arg1 := args{
		stus:    stus1,
		schools: schools1,
	}

	res1 := []string{"0 10", "3", "5 6 7", "2 8", "", "1 4"}

	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"t1", arg1, res1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := graduateAdmission(tt.args.stus, tt.args.schools); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("graduateAdmission() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
