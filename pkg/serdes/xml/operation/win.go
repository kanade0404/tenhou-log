package operation

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"github.com/kanade0404/tenhou-log/serdes/xml/hai"
	"strconv"
	"strings"
)

func wrapError(funcName string, err error) error {
	return fmt.Errorf("error at operation.%s(): %w", funcName, err)
}

func yakuMap() map[int]string {
	return map[int]string{
		0:  "門前清自摸和",
		1:  "立直",
		2:  "一発",
		3:  "槍槓",
		4:  "嶺上開花",
		5:  "河底撈魚",
		6:  "平和",
		7:  "断幺九",
		8:  "一盃口",
		9:  "自風 東",
		10: "自風 南",
		11: "自風 西",
		12: "自風 北",
		13: "場風 東",
		14: "場風 南",
		15: "場風 西",
		16: "場風 北",
		17: "白",
		18: "發",
		19: "中",
		20: "両立直",
		21: "七対子",
		22: "混全帯幺九",
		23: "一気通貫",
		24: "三色同順",
		25: "三色同刻",
		26: "三槓子",
		27: "対々和",
		28: "三暗刻",
		29: "小三元",
		30: "混老頭",
		31: "二盃口",
		32: "純全帯幺九",
		33: "混一色",
		34: "清一色",
		35: "",
		36: "天和",
		37: "地和",
		38: "大三元",
		39: "四暗刻",
		40: "四暗刻単騎",
		41: "字一色",
		42: "緑一色",
		43: "清老頭",
		44: "九蓮宝燈",
		45: "純正九蓮宝燈",
		46: "国士無双",
		47: "国士無双１３面",
		48: "大四喜",
		49: "小四喜",
		50: "四槓子",
		51: "ドラ",
		52: "裏ドラ",
		53: "赤ドラ",
	}
}

type completeHand map[string]int
type point struct {
	hu int
	p  int
	completeHand
}
type Dora struct {
	common []*commonDora
	isRed  bool
}
type commonDora struct {
	omote hai.IHai
	ura   hai.IHai
}

func (d *Dora) String() string {
	var (
		omoteDoras, uraDoras, redDoras []string
		existUraDora                   bool
		dora                           string
	)
	for i := range d.common {
		omoteDoras = append(omoteDoras, fmt.Sprintf("{%s}", d.common[i].omote.String()))
		if d.common[i].ura != nil {
			uraDoras = append(uraDoras, fmt.Sprintf("{%s}", d.common[i].ura.String()))
			existUraDora = true
		}
	}
	dora += fmt.Sprintf("ドラ:(%s)", strings.Join(omoteDoras, ","))
	if existUraDora {
		dora += fmt.Sprintf(",裏ドラ:(%s)", strings.Join(uraDoras, ","))
	}
	if d.isRed {
		reds := hai.Reds(d.isRed)
		for i := range reds {
			redDoras = append(redDoras, fmt.Sprintf("{%s}", reds[i].String()))
		}
		dora = fmt.Sprintf(",赤ドラ:(%s)", strings.Join(redDoras, ","))
	}
	return dora
}

type player struct {
	winner xml.PlayerIndex
	loser  *xml.PlayerIndex
}
type pointSpread struct {
	before int
	after  int
}
type Win struct {
	handHais           []hai.IHai   // 面前の手牌
	callHais           [][]hai.IHai // 副露の手牌
	player             *player
	continuePoint      int
	reachPoint         int
	winHai             hai.IHai
	dora               *Dora
	winPoint           *point
	playerPointSpreads []*pointSpread
}

/*
AllHais
和了時の全ての手牌
*/
func (w *Win) AllHais() []hai.IHai {
	results := w.handHais
	for i := range w.callHais {
		results = append(results, w.callHais[i]...)
	}
	return results
}

/*
HandHais
和了時の面前の手牌
*/
func (w *Win) HandHais() []hai.IHai {
	return w.handHais
}

/*
CallHais
和了時の副露の手牌
*/
func (w *Win) CallHais() [][]hai.IHai {
	return w.callHais
}

/*
WinPlayer
和了プレイヤー番号
*/
func (w *Win) WinPlayer() xml.PlayerIndex {
	return w.player.winner
}

/*
LosePlayer
放銃プレイヤー番号
*/
func (w *Win) LosePlayer() *xml.PlayerIndex {
	return w.player.loser
}

/*
ContinuePoint
供託棒
*/
func (w *Win) ContinuePoint() int {
	return w.continuePoint
}

/*
ReachPoint
立直棒
*/
func (w *Win) ReachPoint() int {
	return w.reachPoint
}

/*
WinHai
和了牌
*/
func (w *Win) WinHai() hai.IHai {
	return w.winHai
}

/*
Dora
ドラ牌
*/
func (w *Win) Dora() *Dora {
	return w.dora
}

/*
IsHitDora
ドラ牌を持っているかどうか
*/
func (w *Win) IsHitDora() bool {
	hais := w.AllHais()
	for i := range hais {
		for j := range w.dora.common {
			if hais[i].Num() == w.dora.common[j].omote.Num() || (w.dora.common[j].ura != nil && hais[i].Num() == w.dora.common[j].ura.Num()) {
				return true
			}
		}
	}
	return false
}

/*
NewWin 和了情報
ba 積み棒,供託リーチ棒
hands カンマ区切りの副露牌以外の手牌（和了牌含む）
m 副露牌
winHai 和了牌
ten 符と和了打点と満貫情報をカンマ区切りで連結（0:満貫未満,1:満貫,2:跳満,3:倍満,4:三倍満,5:役満）
completeHand 役満以外の和了役
yakuman 役満の場合の和了役
doraHai ドラ表示牌の牌番号
doraHaiUra 裏ドラの牌番号
who 和了プレイヤーの番号
fromWho 放銃プレイヤーの番号（ツモの場合は和了プレイヤーの番号）
sc 和了による点数移動を持ち点と収支をカンマ区切り
owari 終局時の持ち点とポイント数をカンマ区切り
*/
func NewWin(ba, hands, m, machi, ten, yaku, yakuman, doraHai, doraHaiUra, who, fromWho, sc, owari string, isRedRule bool) (*Win, error) {
	// 積み棒と供託リーチ棒
	continuePoint, reachPoint, err := createContinueAndReachPoint(ba)
	if err != nil {
		return nil, wrapError("NewWin", err)
	}
	// 手牌
	handHais, err := createHandHais(hands, isRedRule)
	if err != nil {
		return nil, wrapError("NewWin", err)
	}
	// 副露牌
	callHais, err := createCallHais(m, isRedRule)
	if err != nil {
		return nil, wrapError("NewWin", err)
	}
	// 和了待ち牌
	winHai, err := createWinHai(machi, isRedRule)
	if err != nil {
		return nil, wrapError("NewWin", err)
	}
	// 和了点
	winPoint, err := createWinPoint(ten, yaku)
	if err != nil {
		return nil, wrapError("NewWin", err)
	}
	// ドラ
	dora, err := createDora(doraHai, doraHaiUra, isRedRule)
	if err != nil {
		return nil, wrapError("NewWin", err)
	}
	// 和了・放銃プレイヤー
	player, err := createPlayer(who, fromWho)
	if err != nil {
		return nil, wrapError("NewWin", err)
	}
	playerPointSpread, err := createPlayerPointSpread(sc)
	if err != nil {
		return nil, wrapError("NewWin", err)
	}
	return &Win{
		handHais:           handHais,
		callHais:           callHais,
		continuePoint:      continuePoint,
		reachPoint:         reachPoint,
		winHai:             winHai,
		winPoint:           winPoint,
		dora:               dora,
		player:             player,
		playerPointSpreads: playerPointSpread,
	}, nil
}

/*
createContinueAndReachPoint
リーチ棒と供託つみ棒の情報を取得する。
*/
func createContinueAndReachPoint(ba string) (continuePoint, reachPoint int, err error) {
	points := strings.Split(ba, ",")
	if len(points) < 2 {
		err = fmt.Errorf("argument 'ba' must be two elements. actual: %s", ba)
		return 0, 0, wrapError("createContinueAndReachPoint", err)
	}
	cp, err := strconv.Atoi(points[0])
	if err != nil {
		return 0, 0, wrapError("createContinueAndReachPoint", err)
	}
	rp, err := strconv.Atoi(points[1])
	if err != nil {
		return 0, 0, wrapError("createContinueAndReachPoint", err)
	}
	if cp != 0 {
		continuePoint = cp * 1000
	}
	if rp != 0 {
		reachPoint = rp * 1000
	}
	return continuePoint, reachPoint, nil
}

/*
createHandHais
面前の手牌情報を取得する。
*/
func createHandHais(hands string, isRedRule bool) ([]hai.IHai, error) {
	var handHais []hai.IHai
	hhs := strings.Split(hands, ",")
	for i := range hhs {
		idx, err := strconv.Atoi(hhs[i])
		if err != nil {
			return nil, wrapError("createHandHais", err)
		}
		h, err := hai.NewHai(idx, isRedRule)
		if err != nil {
			return nil, wrapError("createHandHais", err)
		}
		handHais = append(handHais, h)
	}
	return handHais, nil
}

/*
createCallHais
副露牌情報を取得する。
*/
func createCallHais(m string, isRedRule bool) ([][]hai.IHai, error) {
	if m == "" {
		return nil, nil
	}
	var callHais [][]hai.IHai
	chs := strings.Split(m, ",")
	for i := range chs {
		c, err := NewClaim(chs[i], isRedRule)
		if err != nil {
			return nil, wrapError("createCallHais", err)
		}
		callHais = append(callHais, c.hais)
	}
	return callHais, nil
}

/*
createWinPoint
和了点数情報を取得する
*/
func createWinPoint(ten string, yaku string) (*point, error) {
	tens := strings.Split(ten, ",")
	hu, err := strconv.Atoi(tens[0])
	if err != nil {
		return nil, wrapError("createWinPoint", err)
	}
	p, err := strconv.Atoi(tens[1])
	if err != nil {
		return nil, wrapError("createWinPoint", err)
	}
	completeHand, err := createCompleteHand(yaku)
	if err != nil {
		return nil, wrapError("createWinPoint", err)
	}
	return &point{
		hu:           hu,
		p:            p,
		completeHand: completeHand,
	}, nil
}

/*
createWinHai
待ち牌情報を取得する。
*/
func createWinHai(machi string, isRedRule bool) (hai.IHai, error) {
	machiHaiIndex, err := strconv.Atoi(machi)
	if err != nil {
		return nil, wrapError("createWinHai", err)
	}
	machiHai, err := hai.NewHai(machiHaiIndex, isRedRule)
	if err != nil {
		return nil, wrapError("createWinHai", err)
	}
	return machiHai, nil
}

/*
createDora
ドラ情報を取得する
*/
func createDora(d, ud string, isRedRule bool) (*Dora, error) {
	var (
		doras []*commonDora
	)
	ds := strings.Split(d, ",")
	uds := strings.Split(ud, ",")
	for i := range ds {
		haiID, err := strconv.Atoi(ds[i])
		if err != nil {
			return nil, wrapError("createDora", err)
		}
		h, err := hai.NewHai(haiID, isRedRule)
		if err != nil {
			return nil, wrapError("createDora", err)
		}
		if ud == "" {
			doras = append(doras, &commonDora{
				omote: h,
			})
		} else {
			haiID, err := strconv.Atoi(uds[i])
			if err != nil {
				return nil, wrapError("createDora", err)
			}
			udh, err := hai.NewHai(haiID, isRedRule)
			if err != nil {
				return nil, wrapError("createDora", err)
			}
			doras = append(doras, &commonDora{
				omote: h,
				ura:   udh,
			})
		}
	}
	return &Dora{
		common: doras,
		isRed:  isRedRule,
	}, nil
}

/*
createPlayer
和了プレイヤーと放銃プレイヤー情報を取得する
*/
func createPlayer(who, fromWho string) (*player, error) {
	winner, err := xml.NewPlayerIndexFromString(who)
	if err != nil {
		return nil, wrapError("createPlayer", err)
	}
	if fromWho == who {
		return &player{
			winner: *winner,
		}, nil
	}
	loser, err := xml.NewPlayerIndexFromString(fromWho)
	if err != nil {
		return nil, wrapError("createPlayer", err)
	}
	return &player{
		winner: *winner,
		loser:  loser,
	}, nil
}

/*
createPlayerPointSpread
点数移動情報を取得する。
indexは0が起家で順になる。
[起家, 南家, 西家(, 北家)]
*/
func createPlayerPointSpread(sc string) ([]*pointSpread, error) {
	var results []*pointSpread
	scs := strings.Split(sc, ",")
	l := len(scs)
	for i := 0; i < l; i += 2 {
		before, err := strconv.Atoi(scs[i])
		if err != nil {
			return nil, wrapError("createPlayerPointSpread", err)
		}
		after, err := strconv.Atoi(scs[i+1])
		if err != nil {
			return nil, wrapError("createPlayerPointSpread", err)
		}
		results = append(results, &pointSpread{
			before: before,
			after:  after,
		})
	}
	return results, nil
}

func createCompleteHand(yaku string) (completeHand, error) {
	yakus := strings.Split(yaku, ",")
	yakuLen := len(yakus)
	completeHandMap := make(completeHand)
	for i := 0; i < yakuLen; i += 2 {
		_yID := yakus[i]
		yID, err := strconv.Atoi(_yID)
		if err != nil {
			return nil, wrapError("createCompleteHand", err)
		}
		_yNum := yakus[i+1]
		yNum, err := strconv.Atoi(_yNum)
		completeHandMap[yakuMap()[yID]] = yNum
	}
	return completeHandMap, nil
}