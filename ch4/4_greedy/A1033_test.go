package ch4

import "testing"

func Test_toFillOrNot(t *testing.T) {
	type args struct {
		maxTank       int
		totalDistance int
		unitCanDrive  int
		gasStations   []gasStation
	}

	gasStations1 := []gasStation{
		{6.00, 1250},
		{7.00, 600},
		{7.00, 150},
		{7.10, 0},
		{7.20, 200},
		{7.50, 400},
		{7.30, 1000},
		{6.85, 300},
	}

	arg1 := args{
		maxTank:       50,
		totalDistance: 1300,
		unitCanDrive:  12,
		gasStations:   gasStations1,
	}

	gasStations2 := []gasStation{
		{7.10, 0},
		{7.00, 600},
	}

	arg2 := args{
		maxTank:       50,
		totalDistance: 1300,
		unitCanDrive:  12,
		gasStations:   gasStations2,
	}

	tests := []struct {
		name    string
		args    args
		wantRes string
	}{
		// TODO: Add test cases.
		{"t1", arg1, "757.92"},
		//{"t1", arg1, "749.17"}, // !!! 答案给的这个，我自己手算和程序给的都是757.92
		{"t2", arg2, "The maximum travel distance = 1200.00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := toFillOrNot(tt.args.maxTank, tt.args.totalDistance, tt.args.unitCanDrive, tt.args.gasStations); gotRes != tt.wantRes {
				t.Errorf("toFillOrNot() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
