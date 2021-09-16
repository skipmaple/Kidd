package __others

import "testing"

func Test_howManyPAT(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{"APPAPT"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := howManyPAT(tt.args.str); got != tt.want {
				t.Errorf("howManyPAT() = %v, want %v", got, tt.want)
			}
		})
	}
}
