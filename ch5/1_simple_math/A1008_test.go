package __simple_math

import "testing"

func Test_elevator(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{[]int{2, 3, 1}}, 41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := elevator(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("elevator() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
