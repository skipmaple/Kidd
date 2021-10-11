package hash

import (
	"reflect"
	"testing"
)

func Test_callatzThink(t *testing.T) {
	type args struct {
		input []int
	}

	arg1 := args{[]int{3, 5, 6, 7, 8, 11}}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", arg1, []int{7, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := callatzThink(tt.args.input); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("callatzThink() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
