package main

import "testing"

func Test_numToSecret(t *testing.T) {
	type args struct {
		a string
		b string
	}

	arg1 := args{
		a: "1234567",
		b: "368782971",
	}

	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "3695Q8118"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := numToSecret(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("numToSecret() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
