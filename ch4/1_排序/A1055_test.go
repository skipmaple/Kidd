package main

import (
	"reflect"
	"testing"
)

func Test_theRichest(t *testing.T) {
	type args struct {
		people  []person
		queries []query1055
	}

	people1 := []person{
		{"Zoe_Bill", 35, 2333},
		{"Bob_Volk", 24, 5888},
		{"Anny_Cin", 95, 999999},
		{"Williams", 30, -22},
		{"Cindy", 76, 76000},
		{"Alice", 18, 88888},
		{"Joe_Mike", 32, 3222},
		{"Michael", 5, 300000},
		{"Rosemary", 40, 5888},
		{"Dobby", 24, 5888},
		{"Billy", 24, 5888},
		{"Nobody", 5, 0},
	}

	queries1 := []query1055{
		{4, 15, 45},
		{4, 30, 35},
		{4, 5, 95},
		{1, 45, 50},
	}

	arg1 := args{
		people:  people1,
		queries: queries1,
	}

	ans1 := []string{"Case #1:", "Alice 18 88888", "Billy 24 5888", "Bob_Volk 24 5888", "Dobby 24 5888", "Case #2:", "Joe_Mike 32 3222", "Zoe_Bill 35 2333", "Williams 30 -22", "Case #3:", "Anny_Cin 95 999999", "Michael 5 300000", "Alice 18 88888", "Cindy 76 76000", "Case #4:", "None"}

	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		// TODO: Add test cases.
		{"t1", arg1,ans1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := theRichest(tt.args.people, tt.args.queries); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("theRichest() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
