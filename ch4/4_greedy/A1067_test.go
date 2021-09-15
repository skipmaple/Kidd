package ch4

import "testing"

func Test_sortWithSwap0(t *testing.T) {
	type args struct {
		list []int
	}

	arg1 := args{list: []int{0, 1, 2, 3, 4}}
	arg2 := args{list: []int{1, 2, 3, 4, 0}}
	arg3 := args{list: []int{0, 2, 1, 4, 3, 6, 5, 8, 7}}

	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		// TODO: Add test cases.
		{"t1", arg1, 0},
		{"t2", arg2, 4},
		{"t3", arg3, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sortWithSwap0(tt.args.list); gotRes != tt.wantRes {
				t.Errorf("sortWithSwap0() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
