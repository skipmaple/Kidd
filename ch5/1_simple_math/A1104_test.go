package __simple_math

import "testing"

func Test_sliceSum(t *testing.T) {
	type args struct {
		input []float64
	}
	tests := []struct {
		name    string
		args    args
		wantRes float64
	}{
		{"t1", args{[]float64{0.1, 0.2, 0.3, 0.4}}, 5.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sliceSum(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("sliceSum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
