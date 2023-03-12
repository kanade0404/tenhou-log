package operation

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewDrawnGame(t *testing.T) {
	type args struct {
		ba        string
		sc        string
		drawnType string
		owari     string
	}
	var (
		nagashiSc4     = "250,+80,250,-20,250,-20,250,-40"
		nagashiSc4Want = []*pointSpread{
			{
				before: 250,
				after:  330,
			},
			{
				before: 250,
				after:  230,
			},
			{
				before: 250,
				after:  230,
			},
			{
				before: 250,
				after:  210,
			},
		}
		sc4     = "250,0,250,0,250,0,250,0"
		sc3     = "350,0,350,0,350,0"
		sc4Want = []*pointSpread{
			{
				before: 250,
				after:  250,
			},
			{
				before: 250,
				after:  250,
			},
			{
				before: 250,
				after:  250,
			},
			{
				before: 250,
				after:  250,
			},
		}
		sc3Want = []*pointSpread{
			{
				before: 350,
				after:  350,
			},
			{
				before: 350,
				after:  350,
			},
			{
				before: 350,
				after:  350,
			},
		}
		invalidSc = "250,0,250,0,250,0,250"
		cp0rp0    = "0,0"
		cp1rp0    = "1,0"
		cp0rp1    = "0,1"
		cp1rp1    = "1,1"
		end       = "218,-18.2,227,2.7,148,-35.2,407,50.7"
		endWant   = &End{
			playerPoints: []*playerPoint{
				{
					gamePoint: 21800,
					winPoint:  -18.2,
				},
				{
					gamePoint: 22700,
					winPoint:  2.7,
				},
				{
					gamePoint: 14800,
					winPoint:  -35.2,
				},
				{
					gamePoint: 40700,
					winPoint:  50.7,
				},
			},
		}
	)
	tests := []struct {
		name    string
		args    args
		want    *Drawn
		wantErr bool
	}{
		{
			name: "四麻流し満貫",
			args: args{
				ba:        cp0rp0,
				sc:        nagashiSc4,
				drawnType: nm,
				owari:     end,
			},
			want: &Drawn{
				playerPointSpreads: nagashiSc4Want,
				drawnType:          Nagashi,
				end:                endWant,
			},
		},
		{
			name: "三麻九種九牌",
			args: args{
				ba:        cp1rp0,
				sc:        sc3,
				drawnType: yao9,
			},
			want: &Drawn{
				playerPointSpreads: sc3Want,
				continuePoint:      200,
				drawnType:          Kyuusyu,
			},
		},
		{
			name: "四麻四風連打",
			args: args{
				ba:        cp0rp1,
				sc:        sc4,
				drawnType: wind4,
			},
			want: &Drawn{
				playerPointSpreads: sc4Want,
				reachPoint:         1000,
				drawnType:          Wind4,
			},
		},
		{
			name: "四家立直",
			args: args{
				ba:        cp1rp1,
				sc:        sc4,
				drawnType: reach4,
			},
			want: &Drawn{
				playerPointSpreads: sc4Want,
				continuePoint:      300,
				reachPoint:         1000,
				drawnType:          Reach4,
			},
		},
		{
			name: "三家和了",
			args: args{
				ba:        cp0rp0,
				sc:        sc4,
				drawnType: ron3,
			},
			want: &Drawn{
				playerPointSpreads: sc4Want,
				drawnType:          Ron3,
			},
		},
		{
			name: "四槓散了",
			args: args{
				ba:        cp0rp0,
				sc:        sc4,
				drawnType: kan4,
			},
			want: &Drawn{
				playerPointSpreads: sc4Want,
				drawnType:          Kan4,
			},
		},
		{
			name: "不正なba",
			args: args{
				ba:        "",
				sc:        sc4,
				drawnType: nagashiSc4,
			},
			wantErr: true,
		},
		{
			name: "不正なsc",
			args: args{
				ba:        cp0rp0,
				sc:        invalidSc,
				drawnType: nagashiSc4,
			},
			wantErr: true,
		},
		{
			name: "不正なdrawnType",
			args: args{
				ba: cp0rp0,
				sc: sc4,
			},
			wantErr: true,
		},
		{
			name: "不正なowari",
			args: args{
				ba:        cp0rp0,
				sc:        sc4,
				drawnType: nagashiSc4,
				owari:     "218,-18.2,227,2.7",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDrawnGame(tt.args.ba, tt.args.sc, tt.args.drawnType, tt.args.owari)
			if err != nil {
				if tt.wantErr {
					return
				} else {
					t.Errorf("NewDrawnGame() custom_error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if diff := cmp.Diff(got.PlayerPointSpreads(), tt.want.playerPointSpreads, cmp.AllowUnexported(pointSpread{})); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.ContinuePoint(), tt.want.continuePoint); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.ReachPoint(), tt.want.reachPoint); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.DrawnType(), tt.want.drawnType); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.End(), tt.want.end, cmp.AllowUnexported(End{}, playerPoint{})); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
