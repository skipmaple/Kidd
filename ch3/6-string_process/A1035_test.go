package main

import (
	"reflect"
	"testing"
)

func Test_replacePassword(t *testing.T) {
	type args struct {
		n        int
		accounts []account
	}

	arg1 := args{
		n: 3,
		accounts: []account{
			{id: "Team000002", pwd: "Rlsp0dfa"},
			{id: "Team000003", pwd: "perfectpwd"},
			{id: "Team000001", pwd: "R1spOdfa"},
		},
	}

	ans1 := []account{
		{"Team000002", "RLsp%dfa"},
		{"Team000001", "R@spodfa"},
	}

	arg2 := args{
		n: 1,
		accounts: []account{
			{"team110", "abcdefg332"},
		},
	}

	ans2 := args{
		accounts: nil,
	}

	tests := []struct {
		name           string
		args           args
		wantCount      int
		wantResAccount []account
	}{
		{"t1", arg1, 2, ans1},
		{"t2", arg2, 0, ans2.accounts},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotResAccount := replacePassword(tt.args.n, tt.args.accounts)
			if gotCount != tt.wantCount {
				t.Errorf("replacePassword() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
			if !reflect.DeepEqual(gotResAccount, tt.wantResAccount) {
				t.Errorf("replacePassword() gotResAccount = %v, want %v", gotResAccount, tt.wantResAccount)
			}
		})
	}
}
