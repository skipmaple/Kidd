package __others

import (
	"reflect"
	"testing"
)

func Test_findOriginElement(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		wantRes []int
	}{
		{"t1", []int{1, 3, 2, 4, 5}, []int{1, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := findOriginElement(tt.arr); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("find_origin_element() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
