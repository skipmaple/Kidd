package main

import (
	"reflect"
	"testing"
)

func Test_listSorting(t *testing.T) {
	type args struct {
		stus []student1028
		col  int
	}

	stus1 := []student1028{
		{7, "james", 85},
		{10, "amy", 90},
		{1, "zoe", 60},
	}

	stus2 := []student1028{
		{7, "james", 85},
		{10, "amy", 90},
		{1, "zoe", 60},
		{2, "james", 98},
	}

	stus3 := []student1028{
		{7, "james", 85},
		{10, "amy", 90},
		{1, "zoe", 60},
		{2, "james", 90},
	}

	arg1 := args{
		stus: stus1,
		col:  1,
	}

	arg2 := args{
		stus: stus2,
		col:  2,
	}

	arg3 := args{
		stus: stus3,
		col:  3,
	}

	res1 := []student1028{
		{1, "zoe", 60},
		{7, "james", 85},
		{10, "amy", 90},
	}

	res2 := []student1028{
		{10, "amy", 90},
		{2, "james", 98},
		{7, "james", 85},
		{1, "zoe", 60},
	}

	res3 := []student1028{
		{1, "zoe", 60},
		{7, "james", 85},
		{2, "james", 90},
		{10, "amy", 90},
	}
	tests := []struct {
		name    string
		args    args
		wantRes []student1028
	}{
		// TODO: Add test cases.
		{"t1", arg1, res1},
		{"t2", arg2, res2},
		{"t3", arg3, res3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := listSorting(tt.args.stus, tt.args.col); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("listSorting() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
