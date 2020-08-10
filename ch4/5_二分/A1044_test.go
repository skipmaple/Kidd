package ch4

import (
	"reflect"
	"testing"
)

func Test_shoppingInMars(t *testing.T) {
	type args struct {
		arr []int
		s   int
	}

	arg1 := args{
		arr: []int{3, 2, 1, 5, 4, 6, 8, 7, 16, 10, 15, 11, 9, 12, 14, 13},
		s:   15,
	}

	arg2 := args{
		arr: []int{1, 2, 3},
		s:   3,
	}

	arg3 := args{
		arr: []int{2, 2, 2},
		s:   5,
	}

	arg4 := args{
		arr: []int{18},
		s:   10,
	}

	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		// TODO: Add test cases.
		{"t1", arg1, [][]int{{1, 5}, {4, 6}, {7, 8}, {11, 11}}},
		{"t2", arg2, [][]int{{1, 2}, {3, 3}}},
		{"t3", arg3, [][]int{{1, 3}}},
		{"t4", arg4, [][]int{{1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := shoppingInMars(tt.args.arr, tt.args.s); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("shoppingInMars() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
