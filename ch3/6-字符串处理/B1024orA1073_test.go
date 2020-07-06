package main

import "testing"

func Test_scienceCount(t *testing.T) {
	type args struct {
		input string
	}

	arg1 := args{"+1.23400E-03"}
	arg2 := args{"-1.2E+10"}
	arg3 := args{"+3.1415E+004"}
	arg4 := args{"-3.1415926E+4"}
	arg5 := args{"+3.1415926E-01"}
	arg6 := args{"-3.1415926E-0005"}


	tests := []struct {
		name    string
		args    args
		wantRes float64
	}{
		// TODO: Add test cases.
		{"t1", arg1, 0.00123400},
		{"t2", arg2, -12000000000},
		{"t3", arg3, 31415},
		{"t4", arg4, -31415.926},
		{"t5", arg5, 0.31415926},
		{"t6", arg6, -0.000031415926},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := scienceCount(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("scienceCount() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}