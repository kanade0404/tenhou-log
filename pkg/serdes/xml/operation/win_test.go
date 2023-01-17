package operation

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"github.com/kanade0404/tenhou-log/serdes/xml/hai"
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

/*
		待ちテストケース
		- リャンメン
		  - 14
		  - 25
		  - 36
		  - 47
		  - 58
		  - 69
		  - 147
		  - 258
		  - 369
		- シャンポン
		  - 東南
		  - 西北
		  - 白發
		- 単騎
		  - 中
	    - ノベタン
	      - 14
	      - 25
	      - 36
	      - 47
	      - 58
	      - 69
		- カンチャン
		  - 間2
		  - 間3
		  - 間4
		  - 間5
		  - 間6
		  - 間7
		  - 間8
		- ペンチャン
		  - 辺3
		  - 辺7
		- 複雑系
		  - 123(2333)
		  - 789(8889)
	      - 13(1222)
		  - 23(2444)
		  - 78(8666)
		  - 14W(23444WW)
*/
func haiTestCases() []struct {
	name       string
	handHaiIDs []int
	callHaiIDs []int
	machi      int
} {
	return []struct {
		name       string
		handHaiIDs []int
		callHaiIDs []int
		machi      int
	}{
		{
			"リャンメン(14m待ち1和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 4, 8},
			[]int{},
			0,
		},
		{
			"リャンメン(25m待ち2和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 8, 12},
			[]int{},
			4,
		},
		{
			"リャンメン(36m待ち3和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 12, 16},
			[]int{},
			8,
		},
		{
			"リャンメン(47m待ち4和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 16, 20},
			[]int{},
			12,
		},
		{
			"リャンメン(58m待ち5和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 20, 24},
			[]int{},
			16,
		},
		{
			"リャンメン(69m待ち6和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 24, 28},
			[]int{},
			20,
		},
		{
			"シャンポン(東南待ち東和了)",
			[]int{109, 110, 113, 114, 116, 117, 118, 120, 121, 122, 0, 4, 8},
			[]int{},
			108,
		},
		{
			"シャンポン(西北待ち西和了)",
			[]int{108, 109, 110, 112, 113, 114, 117, 118, 121, 122, 0, 4, 8},
			[]int{},
			116,
		},
		{
			"シャンポン(白發待ち白和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 124, 125, 128, 129},
			[]int{},
			126,
		},
		{
			"単騎(中待ち中和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 122, 132},
			[]int{},
			133,
		},
		{
			"ノベタン(14待ち1和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 0, 4, 8, 12},
			[]int{},
			3,
		},
		{
			"ノベタン(25待ち2和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 4, 8, 12, 16},
			[]int{},
			5,
		},
		{
			"ノベタン(36待ち3和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 8, 12, 16, 20},
			[]int{},
			9,
		},
		{
			"ノベタン(47待ち4和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 12, 16, 20, 24},
			[]int{},
			13,
		},
		{
			"ノベタン(58待ち5和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 16, 20, 24, 28},
			[]int{},
			17,
		},
		{
			"ノベタン(69待ち6和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 20, 24, 28, 31},
			[]int{},
			21,
		},
		{
			"カンチャン(2待ち2和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 0, 8},
			[]int{},
			4,
		},
		{
			"カンチャン(3待ち3和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 4, 12},
			[]int{},
			8,
		},
		{
			"カンチャン(4待ち4和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 8, 16},
			[]int{},
			12,
		},
		{
			"カンチャン(5待ち5和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 12, 20},
			[]int{},
			16,
		},
		{
			"カンチャン(6待ち6和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 16, 24},
			[]int{},
			20,
		},
		{
			"カンチャン(7待ち7和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 20, 28},
			[]int{},
			24,
		},
		{
			"カンチャン(8待ち8和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 24, 32},
			[]int{},
			28,
		},
		{
			"ペンチャン(3待ち3和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 0, 4},
			[]int{},
			8,
		},
		{
			"ペンチャン(7待ち7和了)",
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 120, 121, 28, 32},
			[]int{},
			24,
		},
		{
			"複雑系(123待ち1和了)",
			// 2333
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 4, 8, 9, 10},
			[]int{},
			0,
		},
		{
			"複雑系(13待ち1和了)",
			// 1222
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 1, 4, 5, 6},
			[]int{},
			0,
		},
		{
			"複雑系(79待ち9和了)",
			// 8889
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 28, 29, 30, 32},
			[]int{},
			33,
		},
		{
			"複雑系(23待ち3和了)",
			// 2444
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 4, 12, 13, 14},
			[]int{},
			8,
		},
		{
			"複雑系(78待ち7和了)",
			// 6668
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 118, 20, 21, 22, 28},
			[]int{},
			24,
		},
		{
			"複雑系(14西待ち1和了)",
			// 23444西西
			[]int{108, 109, 110, 112, 113, 114, 116, 117, 4, 8, 12, 13, 14},
			[]int{},
			0,
		},
	}

}

func testWin_AllHais_Args(hai, m, machi string) args {
	return args{
		ba:         "0,0",
		hands:      hai,
		m:          m,
		machi:      machi,
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
	}
}

func TestWin_AllHais(t *testing.T) {
	tests := []struct {
		name string
		args
		wantIDs []int
		wantErr bool
	}{
		{
			name:    "47m待ち副露なし",
			args:    testWin_AllHais_Args("2,3,12,19,20,27,30,34,77,80,87", "41577", "12"),
			wantIDs: []int{2, 3, 12, 19, 20, 27, 30, 34, 77, 80, 87, 109, 110, 108},
		},
	}
	for _, tt := range tests {
		win, err := NewWin(tt.args.ba, tt.args.hands, tt.args.m, tt.args.machi, tt.args.ten, tt.args.yaku, tt.args.yakuman, tt.args.doraHai, tt.args.doraHaiUra, tt.args.who, tt.args.fromWho, tt.args.sc, tt.args.owari, tt.args.isRedRule)
		if (err != nil) != tt.wantErr {
			t.Errorf("NewWin() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		var want []hai.IHai
		for i := range tt.wantIDs {
			h, err := hai.NewHai(tt.wantIDs[i], tt.isRedRule)
			if err != nil {
				t.Errorf("%v", err)
				return
			}
			want = append(want, h)
		}
		if diff := hais2Strings(win.AllHais(), want); diff != "" {
			t.Errorf("(-got+want)\n%v", diff)
			return
		}
	}
}

func TestWin_HandHais(t *testing.T) {
	//TODO
}

func TestWin_CallHais(t *testing.T) {
	//TODO
}

func TestWin_Doras(t *testing.T) {
	//TODO
}

func TestWin_ContinuePoint(t *testing.T) {
	//TODO
}

func TestWin_ReachPoint(t *testing.T) {
	//TODO
}

func TestWin_WinHai(t *testing.T) {
	//TODO
}

func TestWin_WinPlayer(t *testing.T) {
	//TODO
}

func TestWin_LosePlayer(t *testing.T) {
	//TODO
}

func TestWin_IsHitDora(t *testing.T) {
	//TODO
}

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
	twoCharacter, err := hai.NewHai(7, true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	fiveCharacter, err := hai.NewHai(17, true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	oneCircle71, err := hai.NewHai(71, true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	winHai, err := hai.NewHai(18, true)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	winHai12, err := hai.NewHai(12, true)
	if err != nil {
		t.Errorf("%v", err)
	}
	test1HaiIDs := []int{14, 18, 22, 23, 27, 31, 41, 43, 46, 49, 53, 85, 86, 87}
	var test1Hais, test2Hais, test2CallHais []hai.IHai
	for i := range test1HaiIDs {
		h, err := hai.NewHai(test1HaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		test1Hais = append(test1Hais, h)
	}
	test2HaiIDs := []int{2, 3, 12, 19, 20, 27, 30, 34, 77, 80, 87}
	for i := range test2HaiIDs {
		h, err := hai.NewHai(test2HaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		test2Hais = append(test2Hais, h)
	}
	test2CallHaiIDs := []int{109, 110, 108}
	for i := range test2CallHaiIDs {
		h, err := hai.NewHai(test2CallHaiIDs[i], true)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		test2CallHais = append(test2CallHais, h)
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
				dora: &Dora{
					common: []*commonDora{
						{
							omote: fiveCharacter,
							ura:   twoCharacter,
						},
					},
					isRed: true,
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
				dora: &Dora{
					common: []*commonDora{
						{
							omote: oneCircle71,
							ura:   nil,
						},
					},
					isRed: true,
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
			if diff := cmp.Diff(got.Dora().String(), tt.want.Dora().String()); diff != "" {
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
