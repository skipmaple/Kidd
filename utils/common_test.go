package utils

import (
	"reflect"
	"testing"
)

func TestIntToSlice(t *testing.T) {
	type args struct {
		n        int
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{12345, []int{}}, want: []int{1, 2, 3, 4, 5}},
		{name: "t2", args: args{6767, []int{}}, want: []int{6, 7, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToSlice(tt.args.n, tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToInt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2, 3, 4, 5}}, 12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToInt(tt.args.arr); got != tt.want {
				t.Errorf("SliceToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
