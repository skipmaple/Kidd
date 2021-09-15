package main

import "testing"

func Test_sumFormat(t *testing.T) {
	type args struct {
		a int
		b int
	}

	arg1 := args{
		a: -1000000,
		b: 9,
	}

	arg2 := args{
		a: 999999,
		b: 888888,
	}

	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "-999,991"},
		{"t2", arg2, "1,888,887"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sumFormat(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("sumFormat() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
