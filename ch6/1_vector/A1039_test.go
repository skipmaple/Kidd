package __vector

import (
	"sort"
	"testing"
)

func Test_studentListForCourse(t *testing.T) {
	type args struct {
		stuMsgs []stuMsg
	}

	//				"ZOE1": {4, 5},
	//				"ANN0": {1, 2, 5},
	//				"BOB5": {1, 2, 3, 4, 5},
	//				"JOE4": {2},
	//				"JAY9": {1, 2, 4, 5},
	//				"FRA8": {2, 4, 5},
	//				"DON2": {4, 5},
	//				"AMY7": {5},
	//				"KAT3": {2, 4, 5},
	//				"LOR6": {1, 2, 4, 5},

	//{4, []string{"BOB5", "DON2", "FRA8", "JAY9", "KAT3", "LOR6", "ZOE1"}},
	//{1, []string{"ANN0", "BOB5", "JAY9", "LOR6"}},
	//{2, []string{"ANN0", "BOB5", "FRA8", "JAY9", "JOE4", "KAT3", "LOR6"}},
	//{3, []string{"BOB5"}},
	//{5, []string{"AMY7", "ANN0", "BOB5", "DON2", "FRA8", "JAY9", "KAT3", "LOR6", "ZOE1"}},
	tests := []struct {
		name    string
		args    args
		wantRes map[int][]string
	}{
		{"t1", args{[]stuMsg{
			{"ZOE1", []int{4, 5}},
			{"ANN0", []int{1, 2, 5}},
			{"BOB5", []int{1, 2, 3, 4, 5}},
			{"JOE4", []int{2}},
			{"JAY9", []int{1, 2, 4, 5}},
			{"FRA8", []int{2, 4, 5}},
			{"DON2", []int{4, 5}},
			{"AMY7", []int{5}},
			{"KAT3", []int{2, 4, 5}},
			{"LOR6", []int{1, 2, 4, 5}},
		}}, map[int][]string{
			4: {"BOB5", "DON2", "FRA8", "JAY9", "KAT3", "LOR6", "ZOE1"},
			1: {"ANN0", "BOB5", "JAY9", "LOR6"},
			2: {"ANN0", "BOB5", "FRA8", "JAY9", "JOE4", "KAT3", "LOR6"},
			3: {"BOB5"},
			5: {"AMY7", "ANN0", "BOB5", "DON2", "FRA8", "JAY9", "KAT3", "LOR6", "ZOE1"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := studentListForCourse(tt.args.stuMsgs); !isEqualV2(gotRes, tt.wantRes) {
				t.Errorf("studentListForCourse() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func isEqualV2(got, want map[int][]string) bool {
	if len(got) != len(want) {
		return false
	}

	for kWant, vWant := range want {
		sort.Strings(vWant)
		sortedGotV := got[kWant]
		sort.Strings(sortedGotV)

		for i := 0; i < len(sortedGotV); i++ {
			if sortedGotV[i] != vWant[i] {
				return false
			}
		}
	}

	return true
}

//1:[ANN0 BOB5 JAY9 LOR6] 2:[ANN0 BOB5 JOE4 JAY9 FRA8 KAT3 LOR6] 3:[BOB5] 4:[ZOE1 BOB5 JAY9 FRA8 DON2 KAT3 LOR6] 5:[ZOE1 ANN0 BOB5 JAY9 FRA8 DON2 AMY7 KAT3 LOR6]],
//1:[ANN0 BOB5 JAY9 LOR6] 2:[ANN0 BOB5 FRA8 JAY9 JOE4 KAT3 LOR6] 3:[BOB5] 4:[BOB5 DON2 FRA8 JAY9 KAT3 LOR6 ZOE1] 5:[AMY7 ANN0 BOB5 DON2 FRA8 JAY9 KAT3 LOR6 ZOE1]]
