package utils

import (
	"reflect"
	"testing"
)

func TestBubbleSortMin2Max(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name       string
		args       args
		wantResArr []int
	}{
		{"t1", args{arr: []int{6, 7, 6, 7}}, []int{6, 6, 7, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResArr := BubbleSortMin2Max(tt.args.arr); !reflect.DeepEqual(gotResArr, tt.wantResArr) {
				t.Errorf("BubbleSortMin2Max() = %v, want %v", gotResArr, tt.wantResArr)
			}
		})
	}
}
