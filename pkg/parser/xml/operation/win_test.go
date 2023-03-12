package operation

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml/hai"
	"sort"
	"testing"
)

type args struct {
	ba         string
	hands      string
	m          string
	machi      string
	ten        string
	yaku       string
	yakuman    string
	doraHai    string
	doraHaiUra string
	who        string
	fromWho    string
	sc         string
	owari      string
	isRedRule  bool
}
type want struct {
	allHais       []hai.IHai
	handHais      []hai.IHai
	callHais      [][]hai.IHai
	winPlayer     xml.PlayerIndex
	losePlayer    *xml.PlayerIndex
	continuePoint int
	reachPoint    int
	winHai        hai.IHai
	dora          *Dora
	isHitDora     bool
	yakuman       map[int]string
	winPoint      *point
}

func TestWin(t *testing.T) {
	fiveCharacter17, err := hai.NewHai(17, true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	oneCircle71, err := hai.NewHai(71, true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	winHai12, err := hai.NewHai(12, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	var riichWinHandHais, calledWinHandHais, calledWinCallHais []hai.IHai
	riichWinHandHaiIDs := []int{14, 18, 22, 23, 27, 31, 41, 43, 46, 49, 53, 85, 86, 87}
	for i := range riichWinHandHaiIDs {
		h, err := hai.NewHai(riichWinHandHaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		riichWinHandHais = append(riichWinHandHais, h)
	}
	calledWinHandHaiIDs := []int{2, 3, 12, 19, 20, 27, 30, 34, 77, 80, 87}
	for i := range calledWinHandHaiIDs {
		h, err := hai.NewHai(calledWinHandHaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		calledWinHandHais = append(calledWinHandHais, h)
	}
	calledWinCallHaiIDs := []int{109, 110, 108}
	for i := range calledWinCallHaiIDs {
		h, err := hai.NewHai(calledWinCallHaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		calledWinCallHais = append(calledWinCallHais, h)
	}
	firstUser, err := xml.NewPlayerIndex(0)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	var (
		riichWinHandHaiArg  = "14,18,22,23,27,31,41,43,46,49,53,85,86,87"
		calledWinHandHaiArg = "2,3,12,19,20,27,30,34,77,80,87"
		calledWinCallHaiArg = "41577"
	)
	secondUser, err := xml.NewPlayerIndex(1)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	thirdUser, err := xml.NewPlayerIndex(2)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	fourthUser, err := xml.NewPlayerIndex(3)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	winPoint20hu1100p := &point{
		hu: 20,
		p:  1100,
		completeHand: completeHand{
			"場風 南": 1,
		},
	}
	winPoint30hu1000p := &point{
		hu: 30,
		p:  1000,
		completeHand: completeHand{
			"場風 南": 1,
		},
	}
	winPoint40hu1300p := &point{
		hu: 40,
		p:  1300,
		completeHand: completeHand{
			"場風 南": 1,
		},
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "20符1100点表ドラ1プレイヤー0ツモ",
			args: args{
				ba:        "0,0",
				hands:     calledWinHandHaiArg,
				m:         calledWinCallHaiArg,
				machi:     "12",
				ten:       "20,1100,0",
				yaku:      "14,1",
				doraHai:   "71",
				who:       "0",
				fromWho:   "0",
				sc:        "250,11,250,-7,250,-4,250,-4",
				isRedRule: true,
			},
			want: want{
				allHais:   append(calledWinHandHais, calledWinCallHais...),
				handHais:  calledWinHandHais,
				callHais:  [][]hai.IHai{calledWinCallHais},
				winPlayer: *firstUser,
				winHai:    winHai12,
				dora:      &Dora{common: []*commonDora{{omote: oneCircle71}}, isRed: true},
				isHitDora: false,
				winPoint:  winPoint20hu1100p,
			},
		},
		{
			name: "積み棒300点20符1100点表ドラ1裏ドラ1プレイヤー0ロンプレイヤー1",
			args: args{
				ba:         "1,0",
				hands:      riichWinHandHaiArg,
				machi:      "12",
				ten:        "30,1000,0",
				yaku:       "14,1",
				doraHai:    "71",
				doraHaiUra: "17",
				who:        "0",
				fromWho:    "1",
				sc:         "250,11,250,-7,250,-4,250,-4",
				isRedRule:  true,
			},
			want: want{
				allHais:       riichWinHandHais,
				handHais:      riichWinHandHais,
				continuePoint: 300,
				winPlayer:     *firstUser,
				losePlayer:    secondUser,
				winHai:        winHai12,
				dora:          &Dora{common: []*commonDora{{omote: oneCircle71, ura: fiveCharacter17}}, isRed: true},
				isHitDora:     true,
				winPoint:      winPoint30hu1000p,
			},
		},
		{
			name: "リーチ棒1000点40符1300点表ドラ2裏ドラ2プレイヤー0ロンプレイヤー2",
			args: args{
				ba:         "0,1",
				hands:      riichWinHandHaiArg,
				machi:      "12",
				ten:        "40,1300,0",
				yaku:       "14,1",
				doraHai:    "71,17",
				doraHaiUra: "17,71",
				who:        "0",
				fromWho:    "2",
				sc:         "250,11,250,-7,250,-4,250,-4",
				isRedRule:  true,
			},
			want: want{
				allHais:    riichWinHandHais,
				handHais:   riichWinHandHais,
				reachPoint: 1000,
				winPlayer:  *firstUser,
				losePlayer: thirdUser,
				winHai:     winHai12,
				dora:       &Dora{common: []*commonDora{{omote: oneCircle71, ura: fiveCharacter17}, {omote: fiveCharacter17, ura: oneCircle71}}, isRed: true},
				isHitDora:  true,
				winPoint:   winPoint40hu1300p,
			},
		},
		{
			name: "積み棒300点リーチ棒1000点20符1100点表ドラ1プレイヤー0ロンプレイヤー3",
			args: args{
				ba:        "1,1",
				hands:     calledWinHandHaiArg,
				m:         calledWinCallHaiArg,
				machi:     "12",
				ten:       "20,1100,0",
				yaku:      "14,1",
				doraHai:   "71",
				who:       "0",
				fromWho:   "3",
				sc:        "250,11,250,-7,250,-4,250,-4",
				isRedRule: true,
			},
			want: want{
				allHais:       append(calledWinHandHais, calledWinCallHais...),
				handHais:      calledWinHandHais,
				callHais:      [][]hai.IHai{calledWinCallHais},
				continuePoint: 300,
				reachPoint:    1000,
				winPlayer:     *firstUser,
				losePlayer:    fourthUser,
				winHai:        winHai12,
				dora:          &Dora{common: []*commonDora{{omote: oneCircle71}}, isRed: true},
				isHitDora:     false,
				winPoint:      winPoint20hu1100p,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWin(tt.args.ba, tt.args.hands, tt.args.m, tt.args.machi, tt.args.ten, tt.args.yaku, tt.args.yakuman, tt.args.doraHai, tt.args.doraHaiUra, tt.args.who, tt.args.fromWho, tt.args.sc, tt.args.owari, tt.args.isRedRule)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWin() custom_error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := hais2Strings(got.AllHais(), tt.want.allHais); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := callHais2Strings(got.CallHais(), tt.want.callHais); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := hais2Strings(got.HandHais(), tt.want.handHais); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.ContinuePoint(), tt.want.continuePoint); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.ReachPoint(), tt.want.reachPoint); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.Dora().String(), tt.want.dora.String()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.WinPlayer(), tt.want.winPlayer); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.LosePlayer(), tt.want.losePlayer); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.winHai, tt.want.winHai, cmp.AllowUnexported(hai.Hai{})); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.IsHitDora(), tt.want.isHitDora); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.winPoint, tt.want.winPoint, cmp.AllowUnexported(point{})); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}

func hais2Strings(got []hai.IHai, want []hai.IHai) (diff string) {
	var g, w []string
	for i := range got {
		g = append(g, got[i].String())
	}
	for i := range want {
		w = append(w, want[i].String())
	}
	return cmp.Diff(g, w, cmpopts.SortSlices(func(i, j int) bool {
		return i < j
	}))
}
func callHais2Strings(got [][]hai.IHai, want [][]hai.IHai) (diff string) {
	var (
		mapGot                = make(map[uint][]string)
		mapWant               = make(map[uint][]string)
		gotIDs, wantIDs       []uint
		resultGot, resultWant [][]string
	)
	for gotIdx := range got {
		g := got[gotIdx]
		sort.Slice(g, func(i, j int) bool {
			return g[i].ID() < g[j].ID()
		})
		var gs []string
		for j := range g {
			gs = append(gs, g[j].String())
		}
		if len(gs) > 0 {
			mapGot[g[0].ID()] = gs
			gotIDs = append(gotIDs, g[0].ID())
		}
	}
	for wantIdx := range want {
		w := want[wantIdx]
		sort.Slice(w, func(i, j int) bool {
			return w[i].ID() < w[j].ID()
		})
		var ws []string
		for j := range w {
			ws = append(ws, w[j].String())
		}
		if len(ws) > 0 {
			mapWant[w[0].ID()] = ws
			wantIDs = append(wantIDs, w[0].ID())
		}
	}
	for i := range gotIDs {
		if v, ok := mapGot[gotIDs[i]]; ok {
			resultGot = append(resultGot, v)
		}
	}
	for i := range wantIDs {
		if v, ok := mapWant[wantIDs[i]]; ok {
			resultWant = append(resultWant, v)
		}
	}
	return cmp.Diff(resultGot, resultWant)
}
func TestCreateEnd(t *testing.T) {
	tests := []struct {
		owari   string
		want    *End
		wantErr bool
	}{
		{
			owari: "218,-18.2,227,2.7,148,-35.2,407,50.7",
			want: &End{
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
			},
		},
		{
			owari:   "",
			wantErr: true,
		},
		{
			owari: "430,-18.2,350,2.7,280,-35.2",
			want: &End{
				playerPoints: []*playerPoint{
					{
						gamePoint: 43000,
						winPoint:  -18.2,
					},
					{
						gamePoint: 35000,
						winPoint:  2.7,
					},
					{
						gamePoint: 28000,
						winPoint:  -35.2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		var name string
		if tt.want == nil {
			name = fmt.Sprintf("createEnd(%s) == custom_error", tt.owari)
		} else {
			name = fmt.Sprintf("createEnd(%s) == %v", tt.owari, tt.want.playerPoints)
		}
		t.Run(name, func(t *testing.T) {
			got, err := createEnd(tt.owari)
			if (err != nil) != tt.wantErr {
				t.Errorf("createEnd() custom_error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if diff := cmp.Diff(got.PlayerPoints(), tt.want.PlayerPoints(), cmp.AllowUnexported(playerPoint{})); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
		})
	}
}
