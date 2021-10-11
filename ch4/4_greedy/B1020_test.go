package ch4

import "testing"

func Test_yueBing(t *testing.T) {
	type args struct {
		need  float64
		haves []float64
		sells []float64
	}

	arg1 := args{
		need:  20,
		haves: []float64{18, 15, 10},
		sells: []float64{75, 72, 45},
	}

	tests := []struct {
		name    string
		args    args
		wantMax float64
	}{
		{"t1", arg1, 94.50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := yueBing(tt.args.need, tt.args.haves, tt.args.sells); gotMax != tt.wantMax {
				t.Errorf("yueBing() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
