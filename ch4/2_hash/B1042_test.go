package hash

import "testing"

func Test_byteCount(t *testing.T) {
	type args struct {
		intput string
	}

	tests := []struct {
		name      string
		args      args
		wantAlpha string
		wantMax   int
	}{
		// TODO: Add test cases.
		{"t1", args{"This is a simple TEST.  There ARE numbers and other symbols 1&2&3..."}, "e", 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAlpha, gotMax := byteCount(tt.args.intput)
			if gotAlpha != tt.wantAlpha {
				t.Errorf("byteCount() gotAlpha = %v, want %v", gotAlpha, tt.wantAlpha)
			}
			if gotMax != tt.wantMax {
				t.Errorf("byteCount() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
