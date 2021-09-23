package __simple_math

import (
	"reflect"
	"testing"
)

func Test_iWantPass(t *testing.T) {

	tests := []struct {
		name    string
		inputs  []string
		wantRes []string
	}{
		{"t1",
			[]string{"PAT", "PAAT", "AAPATAA", "AAPAATAAAA", "xPATx", "PT", "Whatever", "APAAATAA"},
			[]string{"YES", "YES", "YES", "YES", "NO", "NO", "NO", "NO"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := iWantPass(tt.inputs); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("iWantPass() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
