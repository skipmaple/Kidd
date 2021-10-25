package __big_int_calc

import (
	"reflect"
	"testing"
)

func Test_isPalindromic(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{[]int{1, 2, 3, 2, 1}}, true},
		{"t2", args{[]int{1, 2, 3, 2, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromic(tt.args.arr); got != tt.want {
				t.Errorf("isPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_palindromicNumber(t *testing.T) {
	type args struct {
		nStr string
		step int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{"67", 3}, "484 2"},
		{"t2", args{"69", 3}, "1353 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := palindromicNumber(tt.args.nStr, tt.args.step); gotRes != tt.wantRes {
				t.Errorf("palindromicNumber() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_reverseArr(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{[]int{7, 6}}, []int{6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := reverseArr(tt.args.arr); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("reverseArr() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
