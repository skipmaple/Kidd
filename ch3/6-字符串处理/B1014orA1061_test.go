package main

import "testing"

func Test_datingDate(t *testing.T) {
	type args struct {
		secret []string
	}

	arg1 := args{
		secret: []string{
			"3485djDkxh4hhGE",
			"2984akDfkkkkggEdsb",
			"s&hgsfdk",
			"d&Hyscvnm",
		},
	}

	arg2 := args{
		secret: []string{
			"GM",
			"GM",
			"s&hgafdk",
			"d&Hyscvkm",
		},
	}

	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "THU 14:04"},
		{"t2", arg2, "SUN 22:07"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := datingDate(tt.args.secret); gotRes != tt.wantRes {
				t.Errorf("datingDate() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}