package main

import "testing"

func Test_signInAndSignOut(t *testing.T) {

	type args struct {
		someDayRecord oneDayRecord
	}
	arg := args{
		someDayRecord: oneDayRecord{
			count: 3,
			signRecords: []signRecord{
				{
					"CS301111",
					"15:30:28",
					"17:00:10",
				},
				{
					"SC3021234",
					"08:00:00",
					"11:25:25",
				},
				{
					"CS301133",
					"21:45:00",
					"21:58:40",
				},
			},
		},
	}
	tests := []struct {
		name        string
		args        args
		wantFirstID string
		wantLastID  string
	}{
		{
			"t1",
			arg,
			"SC3021234",
			"CS301133",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstID, gotLastID := signInAndSignOut(tt.args.someDayRecord)
			if gotFirstID != tt.wantFirstID {
				t.Errorf("signInAndSignOut() gotFirstID = %v, want %v", gotFirstID, tt.wantFirstID)
			}
			if gotLastID != tt.wantLastID {
				t.Errorf("signInAndSignOut() gotLastID = %v, want %v", gotLastID, tt.wantLastID)
			}
		})
	}
}
