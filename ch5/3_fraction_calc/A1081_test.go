package __fraction_calc

import "testing"

func TestFractionSum(t *testing.T) {
	type args struct {
		arr []Fraction
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{[]Fraction{{-1, 2}, {2, 4}}}, "0"},
		{"t2", args{[]Fraction{{1, 3}, {-1, 2}}}, "-1/6"},
		{"t3", args{[]Fraction{{1, 3}, {1, -6}, {1, 8}}}, "7/24"},
		{"t4", args{[]Fraction{{2, 5}, {4, 15}, {1, 30}, {-2, 60}, {8, 3}}}, "3 1/3"},
		{"t5", args{[]Fraction{{2, 5}, {4, 15}, {1, 30}, {-2, 60}, {-10, 3}}}, "-2 2/3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := FractionSum(tt.args.arr); gotRes != tt.wantRes {
				t.Errorf("FractionSum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
