package main

import "testing"

func Test_addHogwartsCash(t *testing.T) {
	type args struct {
		p string
		a string
	}

	arg1 := args{
		p: "3.2.1",
		a: "10.16.27",
	}
	tests := []struct {
		name  string
		args  args
		wantC string
	}{
		{"t1", arg1, "14.1.28"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := addHogwartsCash(tt.args.p, tt.args.a); gotC != tt.wantC {
				t.Errorf("addHogwartsCash() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
