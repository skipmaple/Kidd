package main

import "testing"

func Test_kuchiguse(t *testing.T) {
	type args struct {
		strs []string
	}

	arg1 := args{strs: []string{"Itai nyan~", "Ninjin wa iyadanyan~", "uhhh nyan~"}}
	arg2 := args{strs: []string{"Itai!", "hhhh T_T", "uhhh"}}

	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "nyan~"},
		{"t2", arg2, "nai"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := kuchiguse(tt.args.strs); gotRes != tt.wantRes {
				t.Errorf("kuchiguse() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}