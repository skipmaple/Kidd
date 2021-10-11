package __fraction_calc

import "testing"

func TestRationalArithmetic(t *testing.T) {
	type args struct {
		a Fraction
		b Fraction
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{Fraction{2, 3}, Fraction{-4, 2}}, "2/3 + (-2) = (-1 1/3)\n2/3 - (-2) = 2 2/3\n2/3 * (-2) = (-1 1/3)\n2/3 / (-2) = (-1/3)\n"},
		{"t2", args{Fraction{5, 3}, Fraction{0, 6}}, "1 2/3 + 0 = 1 2/3\n1 2/3 - 0 = 1 2/3\n1 2/3 * 0 = 0\n1 2/3 / 0 = Inf\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := RationalArithmetic(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("RationalArithmetic() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
