package ch2

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr   []int
		left  int
		right int
	}

	arg1 := args{
		arr:   []int{12, 33, 44, 22, 1, -2, -4, 321, 22},
		left:  0,
		right: 8,
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", arg1, []int{-4, -2, 1, 12, 22, 22, 33, 44, 321}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort2(t *testing.T) {
	type args struct {
		arr []int
	}
	arg2 := args{
		arr: []int{12, 33, 44, 22, 1, -2, -4, 321, 22},
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t2", arg2, []int{-4, -2, 1, 12, 22, 22, 33, 44, 321}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
