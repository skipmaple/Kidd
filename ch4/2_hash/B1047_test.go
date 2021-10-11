package hash

import "testing"

func Test_programTeamContest(t *testing.T) {
	type args struct {
		grades []grade
	}

	arg1 := args{
		grades: []grade{
			{"3-10", 99},
			{"11-5", 87},
			{"102-1", 0},
			{"102-3", 100},
			{"11-9", 89},
			{"3-2", 61},
		},
	}

	tests := []struct {
		name           string
		args           args
		wantTeam       int
		wantTotalScore int
	}{
		{"t1", arg1, 11, 176},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTeam, gotTotalScore := programTeamContest(tt.args.grades)
			if gotTeam != tt.wantTeam {
				t.Errorf("programTeamContest() gotTeam = %v, want %v", gotTeam, tt.wantTeam)
			}
			if gotTotalScore != tt.wantTotalScore {
				t.Errorf("programTeamContest() gotTotalScore = %v, want %v", gotTotalScore, tt.wantTotalScore)
			}
		})
	}
}
