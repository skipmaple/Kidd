package __simple_math

import (
	"reflect"
	"testing"
)

func Test_mathBlackHole(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"t1", args{6767}, []string{"7766-6677=1089\n", "9810-0189=9621\n", "9621-1269=8352\n", "8532-2358=6174\n"}},
		{"t2", args{2222}, []string{"2222-2222=0000\n"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := mathBlackHole(tt.args.input); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("mathBlackHole() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
