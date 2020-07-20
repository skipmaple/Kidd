package hash

import (
	"reflect"
	"testing"
)

func Test_sameGradeStudent(t *testing.T) {
	type args struct {
		grades  []int
		queries []int
	}

	arg1 := args{
		grades:  []int{60, 75, 90, 55, 75, 99, 82, 90, 75, 50},
		queries: []int{75, 90, 88},
	}

	res1 := []int{3, 2, 0}

	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		// TODO: Add test cases.
		{"t1",arg1,res1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sameGradeStudent(tt.args.grades, tt.args.queries); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("sameGradeStudent() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
