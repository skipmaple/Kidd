package ch4

import "testing"

func Test_radix(t *testing.T) {
	type args struct {
		a     string
		b     string
		tag   int
		radix int
	}

	arg1 := args{
		a:     "6",
		b:     "110",
		tag:   1,
		radix: 10,
	}

	arg2 := args{
		a:     "1",
		b:     "ab",
		tag:   1,
		radix: 2,
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", arg1, "2"},
		{"t2", arg2, "Impossible"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := radix(tt.args.a, tt.args.b, tt.args.tag, tt.args.radix); got != tt.want {
				t.Errorf("radix() = %v, want %v", got, tt.want)
			}
		})
	}
}
