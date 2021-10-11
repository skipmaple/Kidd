package hash

import "testing"

func Test_beUnique(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{input: []int{5, 31, 5, 88, 67, 88, 17}}, 31},
		{"t1", args{input: []int{888, 666, 666, 888, 888}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := beUnique(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("beUnique() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
