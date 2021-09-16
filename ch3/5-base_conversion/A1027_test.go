package main

import "testing"

func Test_thirteenRGB(t *testing.T) {
	type args struct {
		input []int
	}

	arg1 := args{input: []int{15, 43, 71}}
	arg2 := args{input: []int{13, 14, 15}}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "#123456"},
		{"t2", arg2, "#101112"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := thirteenRGB(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("thirteenRGB() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
