package __factors

import (
	"reflect"
	"testing"
)

func Test_consecutiveFactors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes result
	}{
		{"t1", args{630}, result{3, "5 * 6 * 7"}},
		{"t2", args{2}, result{1, "2"}},
		{"t3", args{4}, result{1, "2"}},
		{"t4", args{6}, result{2, "2 * 3"}},
		{"t5", args{8}, result{1, "2"}},
		{"t6", args{14}, result{1, "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := consecutiveFactors(tt.args.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("consecutiveFactors() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
