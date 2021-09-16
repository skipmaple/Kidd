package hash

import "testing"

func Test_oldKeyboard(t *testing.T) {
	type args struct {
		broken string
		input  string
	}

	arg1 := args{
		broken: "7+IE.",
		input:  "7_This_is_a_test.",
	}

	res1 := "_hs_s_a_tst"

	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		// TODO: Add test cases.
		{"t1", arg1, res1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := oldKeyboard(tt.args.broken, tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("oldKeyboard() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
