package main

import "testing"

func Test_worldCupBetting(t *testing.T) {
	m1 := matchBet{
		win:  1.1,
		tie:  2.5,
		lose: 1.7,
	}

	m2 := matchBet{
		win:  1.2,
		tie:  3.0,
		lose: 1.6,
	}

	m3 := matchBet{
		win:  4.1,
		tie:  1.2,
		lose: 1.1,
	}

	type args struct {
		t tCase
	}
	tests := []struct {
		name   string
		args   args
		caseM1 string
		caseM2 string
		caseM3 string
		income float64
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				t: tCase{
					matches: [3]matchBet{
						m1,
						m2,
						m3,
					},
				},
			},
			caseM1: "T",
			caseM2: "T",
			caseM3: "W",
			income: 37.98,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := worldCupBetting(tt.args.t)
			if got != tt.caseM1 {
				t.Errorf("worldCupBetting() got = %v, caseM1 %v", got, tt.caseM1)
			}
			if got1 != tt.caseM2 {
				t.Errorf("worldCupBetting() got1 = %v, caseM1 %v", got1, tt.caseM2)
			}
			if got2 != tt.caseM3 {
				t.Errorf("worldCupBetting() got2 = %v, caseM1 %v", got2, tt.caseM3)
			}
			if got3 != tt.income {
				t.Errorf("worldCupBetting() got3 = %v, caseM1 %v", got3, tt.income)
			}
		})
	}
}