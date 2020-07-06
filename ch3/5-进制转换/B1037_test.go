package main

import "testing"

func Test_lookForCash(t *testing.T) {
	type args struct {
		p string
		a string
	}
	arg1 := args{
		p: "10.16.27",
		a: "14.1.28",
	}
	arg2 := args{
		p: "14.1.28",
		a: "10.16.27",
	}
	tests := []struct {
		name  string
		args  args
		wantC string
	}{
		{
			"t1",
			arg1,
			"3.2.1",
		},
		{
			"t2",
			arg2,
			"-3.2.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := lookForCash(tt.args.p, tt.args.a); gotC != tt.wantC {
				t.Errorf("lookForCash() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}