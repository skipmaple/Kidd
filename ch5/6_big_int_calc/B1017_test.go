package __big_int_calc

import "testing"

func Test_aDivB(t *testing.T) {
	type args struct {
		strA string
		b    int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{"123456789050987654321", 7}, "17636684150141093474 3"},
		{"t2", args{"2", 7}, "0 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := aDivB(tt.args.strA, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("aDivB() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
