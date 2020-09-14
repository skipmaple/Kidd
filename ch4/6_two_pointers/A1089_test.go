package __two_pointers

import (
	"reflect"
	"testing"
)

func TestInsertOrMerge(t *testing.T) {
	type args struct {
		arrInput []int
		arrGet   []int
	}

	arg1 := args{
		arrInput: []int{3, 1, 2, 8, 7, 5, 9, 4, 6, 0},
		arrGet:   []int{1, 2, 3, 7, 8, 5, 9, 4, 6, 0},
	}

	arg2 := args{
		arrInput: []int{3, 1, 2, 8, 7, 5, 9, 4, 0, 6},
		arrGet:   []int{1, 3, 2, 8, 5, 7, 4, 9, 0, 6},
	}

	tests := []struct {
		name         string
		args         args
		wantRes      string
		wantNextStep []int
	}{
		{"t1", arg1, "Insertion Sort", []int{1, 2, 3, 5, 7, 8, 9, 4, 6, 0}},
		{"t2", arg2, "Merge Sort", []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotNextStep := InsertOrMerge(tt.args.arrInput, tt.args.arrGet)
			if gotRes != tt.wantRes {
				t.Errorf("InsertOrMerge() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotNextStep, tt.wantNextStep) {
				t.Errorf("InsertOrMerge() gotNextStep = %v, want %v", gotNextStep, tt.wantNextStep)
			}
		})
	}
}
