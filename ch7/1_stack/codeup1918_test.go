package __stack

import (
	"Kidd/utils"
	"reflect"
	"testing"
)

func Test_strToMiddle(t *testing.T) {
	type args struct {
		input string
	}

	aNode := []node{
		node{num: 30, isNum: true},
		node{op: "/"},
		node{num: 90, isNum: true},
		node{op: "-"},
		node{num: 26, isNum: true},
		node{op: "+"},
		node{num: 97, isNum: true},
		node{op: "-"},
		node{num: 5, isNum: true},
		node{op: "-"},
		node{num: 6, isNum: true},
		node{op: "-"},
		node{num: 13, isNum: true},
		node{op: "/"},
		node{num: 88, isNum: true},
		node{op: "*"},
		node{num: 6, isNum: true},
		node{op: "+"},
		node{num: 51, isNum: true},
		node{op: "/"},
		node{num: 29, isNum: true},
		node{op: "+"},
		node{num: 79, isNum: true},
		node{op: "*"},
		node{num: 87, isNum: true},
		node{op: "+"},
		node{num: 57, isNum: true},
		node{op: "*"},
		node{num: 92, isNum: true},
	}

	var wantAStack utils.Stack
	for i := 0; i < len(aNode); i++ {
		wantAStack.Push(aNode[i])
	}

	bNode := []node{
		node{num: 1, isNum: true},
		node{op: "+"},
		node{num: 2, isNum: true},
		node{op: "*"},
		node{num: 3, isNum: true},
		node{op: "-"},
		node{num: 4, isNum: true},
	}

	var wantBStack utils.Stack
	for i := 0; i < len(bNode); i++ {
		wantBStack.Push(bNode[i])
	}

	tests := []struct {
		name string
		args args
		want utils.Stack
	}{
		// TODO: Add test cases.
		{"t1", args{
			"30 / 90 - 26 + 97 - 5 - 6 - 13 / 88 * 6 + 51 / 29 + 79 * 87 + 57 * 92",
		}, wantAStack},
		{
			"t2", args{
				"1 + 2 * 3 - 4",
			}, wantBStack,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToMiddle(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("strToMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

//  {[{30  true} {0 / false} {90  true} {0 - false} {26  true} {0 + false} {97  true} {0 - false} {5  true} {0 - false} {6  true} {0 - false} {13  true} {0 / false} {88  true} {0 * false} {6  true} {0 + false} {51  true} {0 / false} {29  true} {0 + false} {79  true} {0 * false} {87  true} {0 + false} {57  true} {0 * false} {92  true}] {{0 0} 0 0 0 0}}

//  {[{30  true} {0 / false} {90  true} {0 - false} {26  true} {0 + false} {97  true} {0 - false} {5  true} {0 - false} {6  true} {0 - false} {13  true} {0 / false} {88  true} {0 * false} {6  true} {0 + false} {51  true} {0 / false} {29  true} {0 + false} {79  true} {0 * false} {87  true} {0 + false} {57  true} {0 * false}] {{0 0} 0 0 0 0}}

func Test_middleToBack(t *testing.T) {
	type args struct {
		middle utils.Stack
	}

	aNode := []node{
		node{num: 1, isNum: true},
		node{num: 2, isNum: true},
		node{num: 3, isNum: true},
		node{op: "*"},
		node{op: "+"},
		node{num: 4, isNum: true},
		node{op: "-"},
	}

	var wantStack utils.Stack
	for i := 0; i < len(aNode); i++ {
		wantStack.Push(aNode[i])
	}

	tests := []struct {
		name     string
		args     args
		wantBack utils.Stack
	}{
		// TODO: Add test cases.
		{"t1", args{middle: strToMiddle("1 + 2 * 3 - 4")}, wantStack},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBack := middleToBack(tt.args.middle); !reflect.DeepEqual(gotBack, tt.wantBack) {
				t.Errorf("middleToBack() = %v, want %v", gotBack, tt.wantBack)
			}
		})
	}
}

func Test_calcBack(t *testing.T) {
	type args struct {
		back utils.Stack
	}

	aNode := []node{
		node{num: 1, isNum: true},
		node{num: 2, isNum: true},
		node{num: 3, isNum: true},
		node{op: "*"},
		node{op: "+"},
		node{num: 4, isNum: true},
		node{op: "-"},
	}

	var inputStack utils.Stack
	for i := 0; i < len(aNode); i++ {
		inputStack.Push(aNode[i])
	}

	tests := []struct {
		name    string
		args    args
		wantRes float64
	}{
		// TODO: Add test cases.
		{"t1", args{back: inputStack}, 3.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := calcBack(tt.args.back); gotRes != tt.wantRes {
				t.Errorf("calcBack() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_simpleCalculator(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantRes float64
	}{
		// TODO: Add test cases.
		{"t1", args{input: "1 + 2 * 3 - 4"}, 3.0},
		{"t2", args{input: "30 / 90 - 26 + 97 - 5 - 6 - 13 / 88 * 6 + 51 / 29 + 79 * 87 + 57 * 92"}, 12178.21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := simpleCalculator(tt.args.input); gotRes != tt.wantRes {
				t.Errorf("simpleCalculator() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
