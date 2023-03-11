package operation

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"testing"
)

func TestClaim_NewClaim(t *testing.T) {
	tests := []struct {
		name       string
		userIndex  xml.PlayerIndex
		haiIndexes []uint
		callIndex  int
		m          string
		claimType  string
		isRedRule  bool
		wantErr    bool
	}{
		{
			name:       "上家からm23とm4でリャンメンチー",
			userIndex:  FourthUserIndex,
			callIndex:  2,
			haiIndexes: []uint{6, 9, 14},
			isRedRule:  true,
			claimType:  ClaimChii,
			m:          "5431",
		},
		{
			name:       "下家から南ポン",
			userIndex:  SecondUserIndex,
			callIndex:  2,
			haiIndexes: []uint{112, 113, 115},
			isRedRule:  true,
			claimType:  ClaimPon,
			m:          "44105",
		},
		{
			name:       "8pを暗槓",
			userIndex:  FirstUserIndex,
			callIndex:  0,
			haiIndexes: []uint{64, 65, 66, 67},
			isRedRule:  true,
			claimType:  ClaimKan,
			m:          "16384",
		},
		{
			name:       "上家から2mを明槓",
			userIndex:  FourthUserIndex,
			callIndex:  0,
			haiIndexes: []uint{4, 5, 6, 7},
			isRedRule:  true,
			claimType:  ClaimMinkan,
			m:          "1027",
		},
		{
			name:       "抜きドラ",
			userIndex:  FirstUserIndex,
			callIndex:  0,
			haiIndexes: nil,
			isRedRule:  true,
			claimType:  ClaimNorthDora,
			m:          "31264",
		},
		{
			name:       "下家からポンした東を加カン",
			userIndex:  SecondUserIndex,
			callIndex:  2,
			haiIndexes: []uint{108, 109, 110, 111},
			isRedRule:  true,
			claimType:  ClaimChakan,
			m:          "42065",
		},
		{
			name:       "から東をポン",
			userIndex:  SecondUserIndex,
			callIndex:  0,
			haiIndexes: []uint{109, 110, 108},
			isRedRule:  true,
			m:          "41577",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClaim(tt.m, tt.isRedRule)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalMentsuCode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(c.CallFromPlayer(), tt.userIndex); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(c.CallIndex(), tt.callIndex); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			var haiIndexes []uint
			for _, h := range c.Hais() {
				haiIndexes = append(haiIndexes, h.ID())
			}
			if diff := cmp.Diff(haiIndexes, tt.haiIndexes); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
