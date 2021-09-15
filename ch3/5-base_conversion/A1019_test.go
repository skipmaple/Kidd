package main

import (
	"reflect"
	"testing"
)

func Test_palindromicNumber(t *testing.T) {
	type args struct {
		N int
		b int
	}

	arg1 := args{
		N: 27,
		b: 2,
	}

	arg2 := args{
		N: 121,
		b: 5,
	}

	tests := []struct {
		name    string
		args    args
		wantRes bool
		wantArr []int
	}{
		// TODO: Add test cases.
		{
			"t1", arg1, true, []int{1, 1, 0, 1, 1},
		},
		{
			"t2", arg2, false, []int{4, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotArr := palindromicNumber(tt.args.N, tt.args.b)
			if gotRes != tt.wantRes {
				t.Errorf("palindromicNumber() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("palindromicNumber() gotArr = %v, want %v", gotArr, tt.wantArr)
			}
		})
	}
}
