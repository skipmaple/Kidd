package main

import (
	"reflect"
	"testing"
)

func Test_listGrades(t *testing.T) {
	type args struct {
		stus     []student1083
		minScore int
		maxScore int
	}

	stus1 := []student1083{
		{"Tom", "CS000001", 59},
		{"Joe", "Math990112", 89},
		{"Mike", "CS991301", 100},
		{"Mary", "EE990830", 95},
	}

	arg1 := args{
		stus:     stus1,
		minScore: 60,
		maxScore: 100,
	}

	ans1 := []string{"Mike CS991301", "Mary EE990830", "Joe Math990112"}

	stus2 := []student1083{
		{"Jean", "AA980920", 60},
		{"Ann", "CS01", 80},
	}

	arg2 := args{
		stus:     stus2,
		minScore: 90,
		maxScore: 95,
	}

	ans2 := []string{"NONE"}


	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		// TODO: Add test cases.
		{"t1", arg1, ans1},
		{"t2", arg2, ans2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := listGrades(tt.args.stus, tt.args.minScore, tt.args.maxScore); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("listGrades() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}