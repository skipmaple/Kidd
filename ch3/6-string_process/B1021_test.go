package main

import (
	"reflect"
	"testing"
)

func Test_numCount(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		// TODO: Add test cases.
		{"t1", args{100311}, map[int]int{0: 2, 1: 3, 3: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numCount(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
