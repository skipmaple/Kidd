package ch4

import (
	"reflect"
	"testing"
)

func Test_minNum(t *testing.T) {
	type args struct {
		counts []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin []int
	}{
		// TODO: Add test cases.
		{"t1", args{counts: []int{2, 2, 0, 0, 0, 3, 0, 0, 1, 0}}, []int{1, 0, 0, 1, 5, 5, 5, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := minNum(tt.args.counts); !reflect.DeepEqual(gotMin, tt.wantMin) {
				t.Errorf("minNum() = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}
