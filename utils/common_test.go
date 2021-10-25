package utils

import (
	"reflect"
	"testing"
)

func TestIntToSlice(t *testing.T) {
	type args struct {
		n        int
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "t1", args: args{12345, []int{}}, want: []int{1, 2, 3, 4, 5}},
		{name: "t2", args: args{6767, []int{}}, want: []int{6, 7, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToSlice(tt.args.n, tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToInt(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{1, 2, 3, 4, 5}}, 12345},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToInt(tt.args.arr); got != tt.want {
				t.Errorf("SliceToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"t1", args{3, 8}, 1},
		{"t2", args{3, 6}, 3},
		{"t3", args{12, 36}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Gcd(tt.args.a, tt.args.b); gotRes != tt.wantRes {
				t.Errorf("Gcd() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestString2BigNum(t *testing.T) {
	type args struct {
		nStr string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{"1234567890987"}, []int{7, 8, 9, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := String2BigNum(tt.args.nStr); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("String2BigNum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestBigNum2String(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		{"t1", args{[]int{7, 8, 9, 0, 9, 8, 7, 6, 5, 4, 3, 2, 1}}, "1234567890987"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := BigNum2String(tt.args.n); gotRes != tt.wantRes {
				t.Errorf("BigNum2String() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestBigNumAdd(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		// TODO: Add test cases.
		{"t1", args{[]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{4, 3, 2, 1}}, []int{4, 2, 1, 9, 6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := BigNumAdd(tt.args.a, tt.args.b); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("BigNumAdd() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestBigNumMin(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{[]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{4, 3, 2, 1}}, []int{6, 5, 6, 6, 6, 5, 4, 3, 2, 1}},
		{"t2", args{[]int{0, 9}, []int{2, 1}}, []int{8, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := BigNumMin(tt.args.a, tt.args.b); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("BigNumMin() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestBigNumMul(t *testing.T) {
	type args struct {
		a []int
		b int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{[]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 43}, []int{0, 7, 2, 9, 1, 4, 6, 8, 0, 3, 5}},
		{"t2", args{[]int{0, 9}, 3}, []int{0, 7, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := BigNumMul(tt.args.a, tt.args.b); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("BigNumMul() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestBigNumDiv(t *testing.T) {
	type args struct {
		a []int
		b int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"t1", args{[]int{0, 9}, 3}, []int{0, 3}},
		{"t2", args{[]int{0, 9}, 30}, []int{3}},
		{"t3", args{[]int{4, 3, 2, 1}, 7}, []int{6, 7, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := BigNumDiv(tt.args.a, tt.args.b); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("BigNumDiv() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
