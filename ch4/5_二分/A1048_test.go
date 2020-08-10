package ch4

import "testing"

func Test_findCoins(t *testing.T) {
	type args struct {
		arr []int
		s   int
	}

	arg1 := args{
		arr: []int{1, 2, 8, 7, 2, 4, 11, 15},
		s:  15,
	}

	arg2 := args{
		arr: []int{1, 8, 7, 2, 4, 11, 15},
		s:  14,
	}

	arg3 := args{
		arr: []int{1, 8, 8, 2, 4, 11, 11, 15},
		s:  22,
	}
	
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "4 11"},
		{"t2", arg2, "No Solution"},
		{"t3", arg3, "11 11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCoins(tt.args.arr, tt.args.s); got != tt.want {
				t.Errorf("findCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
