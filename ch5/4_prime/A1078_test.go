package __prime

import (
	"reflect"
	"testing"
)

func Test_hashing(t *testing.T) {
	type args struct {
		input []int
		m     int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{[]int{10, 6, 4, 15}, 4}, []int{0, 1, 4, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := hashing(tt.args.input, tt.args.m); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("hashing() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
