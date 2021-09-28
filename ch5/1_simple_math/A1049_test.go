package __simple_math

import "testing"

func Test_countingOnes(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{12}, 5},
		{"t2", args{13}, 6},
		{"t2", args{20}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countingOnes(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("countingOnes() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
