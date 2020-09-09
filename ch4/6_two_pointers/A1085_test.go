package __two_pointers

import "testing"

func Test_perfect_arr(t *testing.T) {
	type args struct {
		arr []int
		p   int
	}

	arg1 := args{
		arr: []int{2, 3, 20, 4, 5, 1, 6, 7, 8, 9},
		p:   8,
	}

	arg2 := args{
		arr: []int{1, 2, 3},
		p:   3,
	}

	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// TODO: Add test cases.
		{"t1", arg1, 8},
		{"t2", arg2, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := perfectArr(tt.args.arr, tt.args.p); gotRes != tt.wantRes {
				t.Errorf("perfect_arr() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
