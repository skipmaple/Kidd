package ch4

import "testing"

func Test_smallestNumber(t *testing.T) {
	type args struct {
		arr []string
	}

	arg1 := args{arr: []string{"32", "321", "3214", "0229", "87"}}
	arg2 := args{arr: []string{"00", "000", "00000"}}
	arg3 := args{arr: []string{"00", "020", "00000"}}

	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// TODO: Add test cases.
		{"t1", arg1, 22932132143287},
		{"t2", arg2, 0},
		{"t3", arg3, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := smallestNumber(tt.args.arr); gotRes != tt.wantRes {
				t.Errorf("smallestNumber() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
