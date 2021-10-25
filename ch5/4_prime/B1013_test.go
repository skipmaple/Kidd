package __prime

import (
	"reflect"
	"testing"
)

func Test_listPrimeNum(t *testing.T) {
	type args struct {
		begin int
		end   int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{5, 27}, []int{11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := listPrimeNum(tt.args.begin, tt.args.end); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("listPrimeNum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_listPrimeNum2(t *testing.T) {
	type args struct {
		begin int
		end   int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{5, 27}, []int{11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := listPrimeNum2(tt.args.begin, tt.args.end); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("listPrimeNum2() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
