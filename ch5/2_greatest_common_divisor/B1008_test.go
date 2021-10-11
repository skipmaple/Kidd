package __greatest_common_divisor

import (
	"reflect"
	"testing"
)

func Test_arrLoopRight(t *testing.T) {
	type args struct {
		arr []int
		m   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"t1", args{[]int{1, 2, 3, 4, 5, 6}, 2}, []int{5, 6, 1, 2, 3, 4}},
		{"t2", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3}, []int{6, 7, 8, 1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrLoopRight(tt.args.arr, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrLoopRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
