package ch2

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}

	arg1 := args{
		arr: []int{1, 3, 4, 6, 7, 8, 10, 11, 12, 15},
		x:   6,
	}

	arg2 := args{
		arr: []int{1, 3, 4, 6, 7, 8, 10, 11, 12, 15},
		x:   9,
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"t1", arg1, 3},
		{"t2", arg2, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lowerBound(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}

	arg1 := args{
		arr: []int{1, 3, 3, 3, 6},
		x:   3,
	}

	arg2 := args{
		arr: []int{1, 3, 3, 3, 6},
		x:   7,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"t1", arg1, 1},
		{"t2", arg2, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowerBound(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("lowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_upperBound(t *testing.T) {
	type args struct {
		arr []int
		x   int
	}

	arg1 := args{
		arr: []int{1, 3, 3, 3, 6},
		x:   3,
	}

	arg2 := args{
		arr: []int{1, 3, 3, 3, 6},
		x:   7,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"t1", arg1, 4},
		{"t2", arg2, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upperBound(tt.args.arr, tt.args.x); got != tt.want {
				t.Errorf("upperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxLen(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}

	arg1 := args{
		arr: []int{10, 24, 15},
		k:   7,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"t1", arg1, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLen(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("maxLen() = %v, want %v", got, tt.want)
			}
		})
	}
}



