package main

import (
	"reflect"
	"testing"
)

func Test_printPhoneBills(t *testing.T) {
	type args struct {
		fee     []int
		records []phoneRecord
	}

	arg1 := args{
		fee: []int{10, 10, 10, 10, 10, 10, 20, 20, 20, 15, 15, 15, 15, 15, 15, 15, 20, 30, 20, 15, 15, 10, 10, 10},
		records: []phoneRecord{
			{
				name:  "CYLL",
				month: 01,
				timeRecord: timeRecord{
					day:    01,
					hour:   06,
					minute: 01,
				},
				status: true,
			},
			{
				name:  "CYLL",
				month: 01,
				timeRecord: timeRecord{
					day:    28,
					hour:   16,
					minute: 05,
				},
				status: false,
			},
			{
				name:  "CYJJ",
				month: 01,
				timeRecord: timeRecord{
					day:    01,
					hour:   07,
					minute: 00,
				},
				status: false,
			},
			{
				name:  "CYLL",
				month: 01,
				timeRecord: timeRecord{
					day:    01,
					hour:   8,
					minute: 03,
				},
				status: false,
			},
			{
				name:  "CYJJ",
				month: 01,
				timeRecord: timeRecord{
					day:    01,
					hour:   05,
					minute: 59,
				},
				status: true,
			},
			{
				name:  "aaa",
				month: 01,
				timeRecord: timeRecord{
					day:    01,
					hour:   01,
					minute: 03,
				},
				status: true,
			},
			{
				name:  "aaa",
				month: 01,
				timeRecord: timeRecord{
					day:    02,
					hour:   00,
					minute: 01,
				},
				status: true,
			},
			{
				name:  "CYLL",
				month: 01,
				timeRecord: timeRecord{
					day:    28,
					hour:   15,
					minute: 41,
				},
				status: true,
			},
			{
				name:  "aaa",
				month: 01,
				timeRecord: timeRecord{
					day:    05,
					hour:   02,
					minute: 24,
				},
				status: true,
			},
			{
				name:  "aaa",
				month: 01,
				timeRecord: timeRecord{
					day:    04,
					hour:   23,
					minute: 59,
				},
				status: false,
			},
		},
	}

	res1 := []string{"CYJJ 01", "01:05:59 01:07:00 61 $12.10", "Total amount: $12.10", "CYLL 01", "01:06:01 01:08:03 122 $24.40", "28:15:41 28:16:05 24 $3.85", "Total amount: $28.25", "aaa 01", "02:00:01 04:23:59 4318 $638.80", "Total amount: $638.80"}

	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"t1", arg1, res1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := printPhoneBills(tt.args.fee, tt.args.records); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("printPhoneBills() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
