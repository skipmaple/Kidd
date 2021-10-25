package __factors

import "testing"

func Test_primeFactors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{97532468}, "97532468=2^2 * 11 * 17 * 101 * 1291"},
		{"t2", args{1}, "1=1"},
		{"t3", args{7}, "7=7"},
		{"t4", args{8}, "8=2^3"},
		{"t5", args{9}, "9=3^2"},
		{"t6", args{180}, "180=2^2 * 3^2 * 5"},
		{"t7", args{2147483647}, "2147483647=2147483647"},
		{"t8", args{2147483646}, "2147483646=2 * 3^2 * 7 * 11 * 31 * 151 * 331"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := primeFactors(tt.args.n); gotRes != tt.wantRes {
				t.Errorf("primeFactors() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
