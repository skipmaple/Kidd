package main

import (
	"reflect"
	"testing"
)

func Test_stuRank(t *testing.T) {
	type args struct {
		locationCount int
		locations     []testLocation
	}

	arg1 := args{
		locationCount: 2,
		locations: []testLocation{
			{
				count: 5,
				students: []student{
					{id: "1234567890001", score: 95, locationNumber: 1},
					{id: "1234567890002", score: 100, locationNumber: 1},
					{id: "1234567890003", score: 95, locationNumber: 1},
					{id: "1234567890004", score: 77, locationNumber: 1},
					{id: "1234567890005", score: 85, locationNumber: 1},
				},
			},
			{
				count: 4,
				students: []student{
					{id: "1234567890013", score: 65, locationNumber: 2},
					{id: "1234567890014", score: 25, locationNumber: 2},
					{id: "1234567890015", score: 100, locationNumber: 2},
					{id: "1234567890016", score: 85, locationNumber: 2},
				},
			},
		},
	}

	ans1 := []student{
		{"1234567890002", 100, 1, 1, 1},
		{"1234567890015", 100, 2, 1, 1},
		{"1234567890001", 95, 1, 2, 3},
		{"1234567890003", 95, 1, 2, 3},
		{"1234567890005", 85, 1, 4, 5},
		{"1234567890016", 85, 2, 2, 5},
		{"1234567890004", 77, 1, 5, 7},
		{"1234567890013", 65, 2, 3, 8},
		{"1234567890014", 25, 2, 4, 9},
	}

	tests := []struct {
		name         string
		args         args
		wantTotal    int
		wantTotalStu []student
	}{
		{"t1", arg1, 9, ans1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, gotTotalStu := stuRank(tt.args.locationCount, tt.args.locations)
			if gotTotal != tt.wantTotal {
				t.Errorf("stuRank() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if !reflect.DeepEqual(gotTotalStu, tt.wantTotalStu) {
				t.Errorf("stuRank() gotTotalStu = %v, want %v", gotTotalStu, tt.wantTotalStu)
			}
		})
	}
}
