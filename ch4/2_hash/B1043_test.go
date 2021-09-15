package hash

import "testing"

func Test_printPATest(t *testing.T) {
	type args struct {
		intput string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", args{"redlesPayBestPATTopTeePHPereatitAPPT"}, "PATestPATestPTetPTePePee"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := printPATest(tt.args.intput); gotRes != tt.wantRes {
				t.Errorf("printPATest() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
