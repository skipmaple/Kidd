package __map

import "testing"

func Test_speechPatterns(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{
			input: "Can1: \"Can a can can a can? It can!\"",
		}, "can 5"},
		{"t2", args{
			input: "hello word hello word hello bi' bi bi bi di",
		}, "bi 4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := speechPatterns(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("speechPatterns() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
