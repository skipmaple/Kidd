package main

import (
	"reflect"
	"testing"
)

func Test_boysVsGirls(t *testing.T) {
	type args struct {
		t tCase1036
	}

	arg1 := args{
		t: tCase1036{
			count: 3,
			students: []student{
				{
					"Joe",
					"M",
					"Math990112",
					89,
				},
				{
					"Mike",
					"M",
					"CS991301",
					100,
				},
				{
					"Mary",
					"F",
					"EE990830",
					95,
				},
			},
		},
	}

	arg2 := args{
		t: tCase1036{
			count: 1,
			students: []student{
				{
					"Jean",
					"M",
					"AA980920",
					60,
				},
			},
		},
	}

	tests := []struct {
		name      string
		args      args
		wantFName string
		wantFId   string
		wantMName string
		wantMId   string
		wantDiff  interface{}
	}{
		{"t1",arg1,"Mary", "EE990830", "Joe", "Math990112", 6},
		{"t2",arg2,"Absent", "Absent", "Jean", "AA980920", "NA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFName, gotFId, gotMName, gotMId, gotDiff := boysVsGirls(tt.args.t)
			if gotFName != tt.wantFName {
				t.Errorf("boysVsGirls() gotFName = %v, want %v", gotFName, tt.wantFName)
			}
			if gotFId != tt.wantFId {
				t.Errorf("boysVsGirls() gotFId = %v, want %v", gotFId, tt.wantFId)
			}
			if gotMName != tt.wantMName {
				t.Errorf("boysVsGirls() gotMName = %v, want %v", gotMName, tt.wantMName)
			}
			if gotMId != tt.wantMId {
				t.Errorf("boysVsGirls() gotMId = %v, want %v", gotMId, tt.wantMId)
			}
			if !reflect.DeepEqual(gotDiff, tt.wantDiff) {
				t.Errorf("boysVsGirls() gotDiff = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}
