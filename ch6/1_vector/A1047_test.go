package __vector

import (
	"sort"
	"testing"
)

func Test_courseListForStudent(t *testing.T) {
	type args struct {
		courseMessages []courseMsg
	}

	//tmpMap := map[string][]int{
	//	"ZOE1": {4, 5},
	//	"ANN0": {1, 2, 5},
	//	"BOB5": {1, 2, 3, 4, 5},
	//	"JOE4": {2},
	//	"JAY9": {1, 2, 4, 5},
	//	"FRA8": {2, 4, 5},
	//	"DON2": {4, 5},
	//	"AMY7": {5},
	//	"KAT3": {2, 4, 5},
	//	"LOR6": {1, 2, 4, 5},
	//	//"NON9": {},
	//}
	//
	//keys := make([]string,0, len(tmpMap))
	//for k, _ := range tmpMap {
	//	keys = append(keys, k)
	//}
	//
	//sort.Strings(keys)
	//
	//resMap := make(map[string][]int)
	//for _, k := range keys {
	//	resMap[k] = tmpMap[k]
	//}
	//
	//fmt.Println(resMap)

	tests := []struct {
		name            string
		args            args
		wantStuMessages map[string][]int
	}{
		{
			"t1",
			args{
				[]courseMsg{
					{4, []string{"BOB5", "DON2", "FRA8", "JAY9", "KAT3", "LOR6", "ZOE1"}},
					{1, []string{"ANN0", "BOB5", "JAY9", "LOR6"}},
					{2, []string{"ANN0", "BOB5", "FRA8", "JAY9", "JOE4", "KAT3", "LOR6"}},
					{3, []string{"BOB5"}},
					{5, []string{"AMY7", "ANN0", "BOB5", "DON2", "FRA8", "JAY9", "KAT3", "LOR6", "ZOE1"}},
				},
			},
			map[string][]int{
				"ZOE1": {4, 5},
				"ANN0": {1, 2, 5},
				"BOB5": {1, 2, 3, 4, 5},
				"JOE4": {2},
				"JAY9": {1, 2, 4, 5},
				"FRA8": {2, 4, 5},
				"DON2": {4, 5},
				"AMY7": {5},
				"KAT3": {2, 4, 5},
				"LOR6": {1, 2, 4, 5},
				//"NON9": {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStuMessages := courseListForStudent(tt.args.courseMessages); !isEqual(gotStuMessages, tt.wantStuMessages) {
				t.Errorf("courseListForStudent() = %v, want %v", gotStuMessages, tt.wantStuMessages)
			}
		})
	}
}

func isEqual(got, want map[string][]int) bool {
	if len(got) != len(want) {
		return false
	}

	for kWant, vWant := range want {
		sort.Ints(vWant)
		sortedGotV := got[kWant]
		sort.Ints(sortedGotV)

		for i := 0; i < len(sortedGotV); i++ {
			if sortedGotV[i] != vWant[i] {
				return false
			}
		}
	}

	return true
}

//map[AMY7:[5] ANN0:[1 2 5] BOB5:[4 1 2 3 5] DON2:[4 5] FRA8:[4 2 5] JAY9:[4 1 2 5] JOE4:[2] KAT3:[4 2 5] LOR6:[4 1 2 5] ZOE1:[4 5]],
//map[AMY7:[5] ANN0:[1 2 5] BOB5:[1 2 3 4 5] DON2:[4 5] FRA8:[2 4 5] JAY9:[1 2 4 5] JOE4:[2] KAT3:[2 4 5] LOR6:[1 2 4 5] NON9:[] ZOE1:[4 5]]
