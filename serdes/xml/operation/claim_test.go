package operation

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"testing"
)

func TestClaim_HandHais(t *testing.T) {
	tests := []struct {
		name       string
		UserIndex  xml.PlayerIndex
		HaiIndexes []uint
		CallIndex  int
		m          string
		GameInfo   *xml.GameInfo
		wantErr    bool
	}{
		{
			name:       "上家からm23とm4でリャンメンチー",
			UserIndex:  FourthUserIndex,
			CallIndex:  2,
			HaiIndexes: []uint{6, 9, 14},
			GameInfo: &xml.GameInfo{
				Red: true,
			},
			m: "5431",
		},
		{
			name:       "下家から南ポン",
			UserIndex:  SecondUserIndex,
			CallIndex:  2,
			HaiIndexes: []uint{112, 113, 115},
			GameInfo: &xml.GameInfo{
				Red: true,
			},
			m: "44105",
		},
		{
			name:       "8pを暗槓",
			UserIndex:  FirstUserIndex,
			CallIndex:  0,
			HaiIndexes: []uint{65, 66, 67},
			GameInfo: &xml.GameInfo{
				Red: true,
			},
			m: "16384",
		},
		{
			name:       "下家から2mを明槓",
			UserIndex:  FourthUserIndex,
			CallIndex:  0,
			HaiIndexes: []uint{4, 5, 6},
			GameInfo: &xml.GameInfo{
				Red: true,
			},
			m: "1027",
		},
		{
			name:       "抜きドラ",
			UserIndex:  FirstUserIndex,
			CallIndex:  0,
			HaiIndexes: nil,
			GameInfo: &xml.GameInfo{
				Red: true,
			},
			m: "31264",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Claim{
				GameInfo: tt.GameInfo,
			}
			if err := c.HandHais(tt.m); (err != nil) != tt.wantErr {
				t.Errorf("HandHais() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(c.CallFromUser(), tt.UserIndex); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(c.CallIndex(), tt.CallIndex); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			var haiIndexes []uint
			for _, h := range c.Hais() {
				haiIndexes = append(haiIndexes, h.ID())
			}
			if diff := cmp.Diff(haiIndexes, tt.HaiIndexes); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
