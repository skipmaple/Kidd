package main

import (
	"reflect"
	"testing"
)

func Test_sayNumSum(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		// TODO: Add test cases.
		{"t1", args{"1234567890987654321123456789"}, []string{"yi", "san", "wu"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sayNumSum(tt.args.num); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("sayNumSum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
