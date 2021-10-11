package hash

import "testing"

func Test_stringSubtraction(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}

	arg1 := args{
		str1: "They are students.",
		str2: "aeiou",
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", arg1, "Thy r stdnts."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringSubtraction(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("stringSubtraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
