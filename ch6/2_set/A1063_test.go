package __set

import "testing"

func Test_setSimilarity(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{
			a: []int{99, 87, 101},
			b: []int{87, 101, 5, 87},
		}, "50.0%"},
		{"t2", args{
			a: []int{99, 87, 101},
			b: []int{99, 101, 18, 5, 135, 18, 99},
		}, "33.3%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := setSimilarity(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("setSimilarity() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
