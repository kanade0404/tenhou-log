package operation

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"github.com/kanade0404/tenhou-log/serdes/xml/hai"
	"testing"
)

func TestWin(t *testing.T) {
	fourthUser, err := xml.NewPlayerIndex(3)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	firstUser, err := xml.NewPlayerIndex(0)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	twoCharacter, err := hai.NewHai(uint(7), true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	fiveCharacter, err := hai.NewHai(uint(17), true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	oneCircle71, err := hai.NewHai(uint(71), true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	winHai, err := hai.NewHai(uint(18), true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	winHai12, err := hai.NewHai(uint(12), true)
	if err != nil {
		t.Errorf("%v", err)
	}
	test1HaiIDs := []uint{14, 18, 22, 23, 27, 31, 41, 43, 46, 49, 53, 85, 86, 87}
	var test1Hais, test2Hais, test2CallHais []hai.IHai
	for i := range test1HaiIDs {
		h, err := hai.NewHai(test1HaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		test1Hais = append(test1Hais, h)
	}
	test2HaiIDs := []uint{2, 3, 12, 19, 20, 27, 30, 34, 77, 80, 87}
	for i := range test2HaiIDs {
		h, err := hai.NewHai(test2HaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		test2Hais = append(test2Hais, h)
	}
	test2CallHaiIDs := []uint{109, 110, 108}
	for i := range test2CallHaiIDs {
		h, err := hai.NewHai(test2CallHaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		test2CallHais = append(test2CallHais, h)
	}
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
	tests := []struct {
		name    string
		args    args
		want    *Win
		wantErr bool
	}{
		{
			name: "和了系:46678m22345p444s,和了牌:5m,ドラ:6m,裏ドラ:3m,点:8000,役:立直タンヤオドラ2,和了プレイヤー:3,放銃プレイヤー:0,立直棒:1000,供託棒:0",
			args: args{
				ba:         "0,1",
				hands:      "14,18,22,23,27,31,41,43,46,49,53,85,86,87",
				m:          "",
				machi:      "18",
				ten:        "40,8000,1",
				yaku:       "1,1,8,1,52,2,53,0",
				yakuman:    "",
				doraHai:    "17",
				doraHaiUra: "7",
				who:        "3",
				fromWho:    "0",
				sc:         "170,-10,250,0,250,0,330,10",
				owari:      "",
				isRedRule:  true,
			},
			want: &Win{
				handHais: test1Hais,
				callHais: [][]hai.IHai{},
				player: &player{
					winner: *fourthUser,
					loser:  firstUser,
				},
				continuePoint: 0,
				reachPoint:    1000,
				winHai:        winHai,
				doras: []dora{
					{
						omote: fiveCharacter,
						ura:   twoCharacter,
					},
				},
				winPoint: &point{
					hu: 40,
					p:  8000,
					completeHand: map[string]int{
						"立直":  1,
						"断幺九": 1,
						"ドラ":  2,
					},
				},
				playerPointSpreads: []*pointSpread{
					{
						before: 25000,
						after:  17000,
					},
					{
						before: 25000,
						after:  25000,
					},
					{
						before: 25000,
						after:  25000,
					},
					{
						before: 25000,
						after:  34000,
					},
				},
			},
		},
		{
			name: "和了系:1156789m234sEEE,和了牌:4m,ドラ:1p,点:1000,役:場風東,和了プレイヤー:3,放銃プレイヤー:0,立直棒:0,供託棒:0",
			args: args{
				ba:         "0,0",
				hands:      "2,3,12,19,20,27,30,34,77,80,87",
				m:          "41577",
				machi:      "12",
				ten:        "30,1000,0",
				yaku:       "14,1",
				yakuman:    "",
				doraHai:    "71",
				doraHaiUra: "",
				who:        "3",
				fromWho:    "0",
				sc:         "250,-80,250,0,250,0,240,90",
				owari:      "",
				isRedRule:  true,
			},
			want: &Win{
				handHais: test2Hais,
				callHais: [][]hai.IHai{test2CallHais},
				player: &player{
					winner: *fourthUser,
					loser:  firstUser,
				},
				continuePoint: 0,
				reachPoint:    0,
				winHai:        winHai12,
				doras: []dora{
					{
						omote: oneCircle71,
						ura:   nil,
					},
				},
				winPoint: &point{
					hu: 40,
					p:  8000,
					completeHand: map[string]int{
						"場風 東": 1,
					},
				},
				playerPointSpreads: []*pointSpread{
					{
						before: 17000,
						after:  16000,
					},
					{
						before: 25000,
						after:  25000,
					},
					{
						before: 25000,
						after:  25000,
					},
					{
						before: 33000,
						after:  34000,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWin(tt.args.ba, tt.args.hands, tt.args.m, tt.args.machi, tt.args.ten, tt.args.yaku, tt.args.yakuman, tt.args.doraHai, tt.args.doraHaiUra, tt.args.who, tt.args.fromWho, tt.args.sc, tt.args.owari, tt.args.isRedRule)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := hais2Strings(got.AllHais(), tt.want.AllHais()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := callHais2Strings(got.CallHais(), tt.want.CallHais()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := hais2Strings(got.HandHais(), tt.want.HandHais()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.ContinuePoint(), tt.want.ContinuePoint()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.ReachPoint(), tt.want.ReachPoint()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := doras2Strings(got.Doras(), tt.want.Doras()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.WinPlayer(), tt.want.WinPlayer()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.LosePlayer(), tt.want.LosePlayer()); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(fmt.Sprintf("%+v", got.WinHai()), fmt.Sprintf("%+v", tt.want.WinHai())); diff != "" {
				t.Errorf("(-got+want)\n%v", diff)
			}
			if diff := cmp.Diff(got.IsHitDora(), tt.want.IsHitDora()); diff != "" {
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
	return cmp.Diff(g, w)
}
func doras2Strings(got []dora, want []dora) (diff string) {
	var g, w []string
	for i := range got {
		g = append(g, got[i].String())
	}
	for i := range want {
		w = append(w, want[i].String())
	}
	return cmp.Diff(g, w)
}
func callHais2Strings(got [][]hai.IHai, want [][]hai.IHai) (diff string) {
	var resultGot, resultWant [][]string
	for i := range got {
		var g []string
		for j := range got[i] {
			g = append(g, got[i][j].String())
		}
		resultGot = append(resultGot, g)
	}
	for i := range want {
		var w []string
		for j := range want[i] {
			w = append(w, want[i][j].String())
		}
		resultWant = append(resultWant, w)
	}
	return cmp.Diff(resultGot, resultWant)
}
