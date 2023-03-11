package operation

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml/hai"
	"testing"
)

func TestDiscard(t *testing.T) {
	west, err := hai.NewHai(119, true)
	if err != nil {
		panic(err)
	}
	oneBamboo, err := hai.NewHai(75, true)
	if err != nil {
		panic(err)
	}
	type args struct {
		discardType string
		haiID       int
		gameInfo    *xml.GameInfo
	}
	tests := []struct {
		name    string
		args    args
		want    *Discard
		wantErr bool
	}{
		{
			name: "親が西を捨てる",
			args: args{
				discardType: "D",
				haiID:       119,
				gameInfo: &xml.GameInfo{
					Red: true,
				},
			},
			want: &Discard{
				discardUserIndex: 0,
				hai:              west,
				isRedRule:        true,
			},
		},
		{
			name: "南家が1索を捨てる",
			args: args{
				discardType: "U",
				haiID:       75,
				gameInfo: &xml.GameInfo{
					Red: true,
				},
			},
			want: &Discard{
				discardUserIndex: 0,
				hai:              oneBamboo,
				isRedRule:        true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDiscard(tt.args.discardType, tt.args.haiID, tt.args.gameInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDiscard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got.DiscardUserIndex(), tt.want.DiscardUserIndex()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.Hai().String(), tt.want.Hai().String()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
