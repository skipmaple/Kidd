package hash

import "testing"

func Test_buyOrNot(t *testing.T) {
	type args struct {
		have string
		want string
	}

	arg1 := args{
		have: "ppRYYGrrYBR2258",
		want: "YrR8RrY",
	}

	arg2 := args{
		have: "ppRYYGrrYB225",
		want: "YrR8RrY",
	}

	tests := []struct {
		name    string
		args    args
		wantRes bool
		wantExt int
	}{
		// TODO: Add test cases.
		{"t1", arg1, true, 8},
		{"t2", arg2, false, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotExt := buyOrNot(tt.args.have, tt.args.want)
			if gotRes != tt.wantRes {
				t.Errorf("buyOrNot() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotExt != tt.wantExt {
				t.Errorf("buyOrNot() gotExt = %v, want %v", gotExt, tt.wantExt)
			}
		})
	}
}
