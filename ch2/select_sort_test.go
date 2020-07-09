package ch2

import (
	"reflect"
	"testing"
)

func Test_selectSort(t *testing.T) {
	type args struct {
		arr []int
	}

	arg1 := args{arr: []int{12, 33, 44, 22, 1, -2, -4, 321, 22}}


	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"t1", arg1, []int{-4, -2, 1, 12, 22, 22, 33, 44, 321}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}