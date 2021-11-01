package __string

import (
	"reflect"
	"testing"
)

func Test_areTheyEqual(t *testing.T) {
	type args struct {
		a string
		b string
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{
			a: "12300",
			b: "12358.9",
			n: 3,
		}, "YES 0.123*10^5"},
		{"t2", args{
			a: "123.5678",
			b: "123.5",
			n: 4,
		}, "YES 0.1235*10^3"},
		{"t3", args{
			a: "00123.5678",
			b: "0001235",
			n: 4,
		}, "NO 0.1235*10^3 0.1235*10^4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := areTheyEqual(tt.args.a, tt.args.b, tt.args.n); gotRes != tt.wantRes {
				t.Errorf("areTheyEqual() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_processNum(t *testing.T) {
	type args struct {
		nStr string
		n    int
	}
	tests := []struct {
		name    string
		args    args
		wantRes result
	}{
		{"t1", args{
			nStr: "12358.9",
			n:    3,
		}, result{
			nStr: "123",
			e:    5,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := processNum(tt.args.nStr, tt.args.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("processNum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
