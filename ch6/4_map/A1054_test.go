package __map

import "testing"

func Test_dominantColor(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{"t1", args{input: [][]int{{0, 0, 255, 16777215, 24}, {24, 24, 0, 0, 24}, {24, 0, 24, 24, 24}}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := dominantColor(tt.args.input); gotMax != tt.wantMax {
				t.Errorf("dominantColor() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
