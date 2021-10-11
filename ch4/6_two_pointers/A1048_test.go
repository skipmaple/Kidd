package __two_pointers

import "testing"

func Test_findCoins(t *testing.T) {
	type args struct {
		nums []int
		sum  int
	}
	arg1 := args{
		nums: []int{1, 2, 8, 7, 2, 4, 11, 15},
		sum:  15,
	}

	arg2 := args{
		nums: []int{1, 8, 7, 2, 4, 11, 15},
		sum:  14,
	}

	arg3 := args{
		nums: []int{1, 8, 8, 2, 4, 11, 11, 15},
		sum:  22,
	}
	tests := []struct {
		name  string
		args  args
		wantA int
		wantB int
	}{
		{"t1", arg1, 4, 11},
		{"t2", arg2, -1, -1},
		{"t3", arg3, 11, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := twoPointers(tt.args.nums, tt.args.sum)
			if gotA != tt.wantA {
				t.Errorf("findCoins() gotA = %v, want %v", gotA, tt.wantA)
			}
			if gotB != tt.wantB {
				t.Errorf("findCoins() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
