package __prime

import "testing"

func Test_countPrime(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{20}, 4},
		{"t2", args{7}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := countPrime(tt.args.num); gotRes != tt.wantRes {
				t.Errorf("countPrime() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
