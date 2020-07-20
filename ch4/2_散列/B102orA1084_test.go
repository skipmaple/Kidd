package hash

import "testing"

func Test_brokenKeyBoard(t *testing.T) {
	type args struct {
		input  string
		output string
	}

	arg1 := args{
		input:  "7_This_is_a_test",
		output: "_hs_s_a_es",
	}

	res1 := "7TI"

	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, res1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := brokenKeyBoard(tt.args.input, tt.args.output); gotRes != tt.wantRes {
				t.Errorf("brokenKeyBoard() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}