package entities

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		name    string
		logText string
		want    *CompressedLogFile
		wantErr bool
	}{
		{
			name:    "カンマあり",
			logText: "{\"file\":'sca20220924.log.gz',\"size\":94787},",
			want: &CompressedLogFile{
				File: "sca20220924.log.gz",
				Size: 94787,
			},
		},
		{
			name:    "カンマなし",
			logText: "{\"file\":'sca20220924.log.gz',\"size\":94787}",
			want: &CompressedLogFile{
				File: "sca20220924.log.gz",
				Size: 94787,
			},
		},
		{
			name:    "プロパティ名がダブルクォートで囲まれていない",
			logText: "{file:'sca20220924.log.gz',size:94787}",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s Unmarshal(%s) = %+v", tt.name, tt.logText, tt.want), func(t *testing.T) {
			got, err := Unmarshal(tt.logText)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
