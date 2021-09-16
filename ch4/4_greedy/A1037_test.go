package ch4

import "testing"

func Test_magicCoupon(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	arg1 := args{
		a: []int{1, 2, 4, -1},
		b: []int{7, 6, -2, -3},
	}

	arg2 := args{
		a: []int{-2, -3, 3, 4, 5},
		b: []int{-3, -2, -4, -1, 0, 7},
	}

	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// TODO: Add test cases.
		{"t1", arg1, 43},
		{"t2", arg2, 53},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := magicCoupon(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("magicCoupon() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
