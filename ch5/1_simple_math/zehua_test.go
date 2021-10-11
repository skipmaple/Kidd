package __simple_math

import (
	"reflect"
	"testing"
)

func Test_merge_meetings(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		// TODO: Add test cases.
		{"t1", args{[][]int{{0, 1}, {3, 5}, {2, 4}}}, [][]int{{0, 1}, {2, 5}}},
		{"t2", args{[][]int{{0, 5}, {3, 5}, {2, 4}}}, [][]int{{0, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := mergeMeetings(tt.args.input); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("merge_meetings() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_replace(t *testing.T) {
	type args struct {
		slice [][]int
		i     int
		e     []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"t1", args{[][]int{{0, 1}, {3, 5}, {2, 4}}, 1, []int{2, 5}}, [][]int{{0, 1}, {2, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replace(tt.args.slice, tt.args.i, tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replace() = %v, want %v", got, tt.want)
			}
		})
	}
}
