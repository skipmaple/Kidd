package main

import "testing"

func Test_printAnotherInt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{234}, "BBSSS1234"},
		{"t2", args{23}, "SS123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := printAnotherInt(tt.args.n); gotRes != tt.wantRes {
				t.Errorf("printAnotherInt() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
