package __map

import "testing"

func Test_mars2Num(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{input: "elo nov"}, 115},
		{"t2", args{input: "tam"}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := mars2Num(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("mars2Num() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_num2Mars(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{input: 29}, "hel mar"},
		{"t2", args{input: 5}, "may"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := num2Mars(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("num2Mars() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
