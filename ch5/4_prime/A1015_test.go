package __prime

import (
	"reflect"
	"testing"
)

func Test_reversiblePrimes(t *testing.T) {
	type args struct {
		inputs []input
	}
	tests := []struct {
		name    string
		args    args
		wantRes []bool
	}{
		// TODO: Add test cases.
		{"t1", args{[]input{{73, 10}, {23, 2}, {23, 10}}}, []bool{true, true, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := reversiblePrimes(tt.args.inputs); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("reversiblePrimes() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
