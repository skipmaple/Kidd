package main

import (
	"reflect"
	"testing"
)

func Test_verifyIDCardNum(t *testing.T) {
	type args struct {
		ids []string
	}

	arg1 := args{ids: []string{"320124198808240056", "12010X198901011234", "110108196711301866", "37070419881216001X"}}
	arg2 := args{ids: []string{"320124198808240056", "110108196711301862"}}
	arg3 := args{ids: []string{"320124198808240056"}}

	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		// TODO: Add test cases.
		{"t1", arg1, []string{"12010X198901011234", "110108196711301866", "37070419881216001X"}},
		{"t2", arg2, nil},
		{"t3", arg3, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := verifyIDCardNum(tt.args.ids)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("verifyIDCardNum() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
