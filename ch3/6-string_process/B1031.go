package main

import (
	"strconv"
)

// 查验身份证
// page 61

func verifyIDCardNum(ids []string) (res []string) {
	for i := 0; i < len(ids); i++ {
		if !verifyID(ids[i]) {
			res = append(res, ids[i])
		}
	}
	return res
}

func verifyID(id string) bool {
	if len(id) != 18 {
		return false
	}
	sum := 0
	weight := map[int]int{
		16: 2,
		15: 4,
		14: 8,
		13: 5,
		12: 10,
		11: 9,
		10: 7,
		9:  3,
		8:  6,
		7:  1,
		6:  2,
		5:  4,
		4:  8,
		3:  5,
		2:  1,
		1:  9,
		0:  7,
	}

	verifyMap := map[int]interface{}{
		0:  1,
		1:  0,
		2:  'X',
		3:  9,
		4:  8,
		5:  7,
		6:  6,
		7:  5,
		8:  4,
		9:  3,
		10: 2,
	}
	for i := 0; i < len(id); i++ {
		if id[i] < '0' || id[i] > '9' {
			return false
		}
		ithNum, _ := strconv.Atoi(string(id[i]))
		sum += ithNum * weight[i]
	}

	//fmt.Printf("%d %% 11 = %d, id[-1] = %s\n", sum, sum%11, string(id[len(id)-1]))

	var lastElement interface{}
	if id[len(id)-1] == 'X' {
		lastElement = 'X'
	} else {
		lastElement, _ = strconv.Atoi(string(id[len(id)-1]))
	}
	return verifyMap[sum%11] == lastElement
}

//func main() {
//	type args struct {
//		ids []string
//	}
//
//	//arg1 := args{ids: []string{"320124198808240056", "12010X198901011234", "110108196711301866", "37070419881216001X"}}
//	//arg2 := args{ids: []string{"320124198808240056", "110108196711301862"}}
//	arg3 := args{ids: []string{"320124198808240056"}}
//
//	tests := []struct {
//		name    string
//		args    args
//		wantRes []string
//	}{
//		// TODO: Add test cases.
//		//{"t1", arg1, []string{"12010X198901011234", "110108196711301866", "37070419881216001X"}},
//		//{"t2", arg2, []string{}},
//		{"t3", arg3, []string{}},
//	}
//
//	for _, tt := range tests {
//		gotRes := verifyIDCardNum(tt.args.ids)
//		if !reflect.DeepEqual(gotRes, tt.wantRes) {
//			fmt.Printf("verifyIDCardNum() gotRes = %v, want %v", gotRes, tt.wantRes)
//		}
//	}
//}
