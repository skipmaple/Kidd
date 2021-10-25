package __big_int_calc

import "testing"

func Test_isSameMap(t *testing.T) {
	type args struct {
		a map[int]int
		b map[int]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{map[int]int{2: 2, 1: 1}, map[int]int{1: 1, 2: 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameMap(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isSameMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_haveFunWithNumbers(t *testing.T) {
	type args struct {
		strN string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{"1234567899"}, "Yes 2469135798"},
		{"t2", args{"123"}, "No 246"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := haveFunWithNumbers(tt.args.strN); gotRes != tt.wantRes {
				t.Errorf("haveFunWithNumbers() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
