package utils

import (
	"reflect"
	"testing"
)

func TestIntToSlice(t *testing.T) {
	type args struct {
		n        int64
		sequence []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{name: "t1", args: args{12345, []int64{}}, want: []int64{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ToSlice(tt.args.n, tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToInt(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"t1", args{[]int64{1, 2, 3, 4, 5}}, 12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToInt(tt.args.arr); got != tt.want {
				t.Errorf("SliceToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
