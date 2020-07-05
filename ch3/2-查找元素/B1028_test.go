package main

import "testing"

func Test_humanCount(t *testing.T) {
	type args struct {
		list []person
	}

	p1 := args{list: []person{
		{
			"John",
			"2001/05/12",
		},
		{
			"Tom",
			"1814/09/06",
		},
		{
			"Ann",
			"2121/01/20",
		},
		{
			"James",
			"1814/0905",
		},
		{
			"Steve",
			"1967/11/20",
		},
	},
	}
	tests := []struct {
		name         string
		args         args
		validNmu     int
		oldestName   string
		youngestName string
	}{
		{
			name: "t1",
			args: p1,
			validNmu: 2,
			oldestName: "Steve",
			youngestName: "John",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := humanCount(tt.args.list)
			if got != tt.validNmu {
				t.Errorf("humanCount() got = %v, validNmu %v", got, tt.validNmu)
			}
			if got1 != tt.oldestName {
				t.Errorf("humanCount() got1 = %v, validNmu %v", got1, tt.oldestName)
			}
			if got2 != tt.youngestName {
				t.Errorf("humanCount() got2 = %v, validNmu %v", got2, tt.youngestName)
			}
		})
	}
}
