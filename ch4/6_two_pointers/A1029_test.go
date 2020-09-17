package __two_pointers

import "testing"

func Test_median(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}

	arg1 := args{
		arr1: []int{11, 12, 13, 14},
		arr2: []int{9, 10, 15, 16, 17},
	}

	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", arg1, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := median(tt.args.arr1, tt.args.arr2); gotRes != tt.wantRes {
				t.Errorf("median() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
