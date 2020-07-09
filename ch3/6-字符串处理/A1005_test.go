package main

import "testing"

func Test_spellNum(t *testing.T) {
	type args struct {
		n string
	}

	arg1 := args{n: "12345"}

	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "one five"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := spellNum(tt.args.n); gotRes != tt.wantRes {
				t.Errorf("spellNum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}