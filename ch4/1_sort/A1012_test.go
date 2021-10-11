package main

import (
	"reflect"
	"testing"
)

func Test_theBestRank(t *testing.T) {
	type args struct {
		stus     []CsStudent
		queryIds []int
	}

	arg1 := args{
		stus: []CsStudent{
			{id: 310101, cScore: 98, mScore: 85, eScore: 88},
			{id: 310102, cScore: 70, mScore: 95, eScore: 88},
			{id: 310103, cScore: 82, mScore: 87, eScore: 94},
			{id: 310104, cScore: 91, mScore: 91, eScore: 91},
			{id: 310105, cScore: 85, mScore: 90, eScore: 90},
		},
		queryIds: []int{310101, 310102, 310103, 310104, 310105, 9999999},
	}

	ans1 := []string{"1 C", "1 M", "1 E", "1 A", "3 A", "N/A"}

	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"t1", arg1, ans1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := theBestRank(tt.args.stus, tt.args.queryIds); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("theBestRank() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
