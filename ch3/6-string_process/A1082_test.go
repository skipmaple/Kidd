package main

import "testing"

func Test_readChineseNum(t *testing.T) {
	type args struct {
		n int
	}

	arg1 := args{n: -123456789}
	arg2 := args{n: 100800}
	arg3 := args{n: 1000003}

	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "负一亿二千三百四十五万六千七百八十九"},
		{"t2", arg2, "一十万零八百"},
		{"t3", arg3, "一百万零三"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := readChineseNum(tt.args.n); gotRes != tt.wantRes {
				t.Errorf("readChineseNum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
