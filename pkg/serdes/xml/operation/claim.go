package operation

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"github.com/kanade0404/tenhou-log/serdes/xml/hai"
	"sort"
	"strconv"
)

const (
	ClaimChii      = "チー"
	ClaimPon       = "ポン"
	ClaimKan       = "暗槓"
	ClaimMinkan    = "大明槓"
	ClaimChakan    = "加槓"
	ClaimNorthDora = "北ドラ抜き"
)

// Claim 副露情報
type Claim struct {
	playerIndex xml.PlayerIndex // 誰から鳴いたか
	hais        []hai.IHai      // 鳴いた手牌
	callIndex   int             // 鳴いた牌のインデックス
	claimType   string
	isRedRule   bool
}

// CallFromPlayer 自分から見て誰から鳴いたか
func (c *Claim) CallFromPlayer() xml.PlayerIndex {
	return c.playerIndex
}

// Hais 鳴いた時の手牌
func (c *Claim) Hais() []hai.IHai {
	return c.hais
}

// CallIndex 鳴いた牌のインデックス番号
func (c *Claim) CallIndex() int {
	return c.callIndex
}

// Type 鳴きの種類
func (c *Claim) Type() string {
	return c.claimType
}

// NewClaim New Claim 副露情報を面子コードから解析する
func NewClaim(mentsuCode string, isRedRule bool) (*Claim, error) {
	c := &Claim{
		isRedRule: isRedRule,
	}
	err := c.unmarshalMentsuCode(mentsuCode)
	if err != nil {
		return nil, err
	}
	return c, nil
}

/*
chii チーの面子コード解析

0x0004 - 1

0x0018 - 牌添字1

0x0060 - 牌添字2

0x0180 - 牌添字3

0xFC00 - 順子のパターン
*/
func chii(mentsuCode uint64, isRedRule bool) (hais []hai.IHai, callIndex int, claimType string, err error) {
	claimType = ClaimChii
	isChow := (mentsuCode & 0x0004) >> 2
	if isChow != 1 {
		return hais, callIndex, claimType, fmt.Errorf("0x0004 must be 1. actual: %d", isChow)
	}
	pattern := (mentsuCode & 0xFC00) >> 10
	// 何番目の牌を鳴いたか
	callIndex = int(pattern % 3)
	// 手牌のインデックス
	haiIndexes := []uint64{
		(mentsuCode & 0x0018) >> 3,
		(mentsuCode & 0x0060) >> 5,
		(mentsuCode & 0x0180) >> 7,
	}
	a := ((pattern/3/7)*9 + (pattern / 3 % 7)) * 4
	for i, index := range haiIndexes {
		h, err := hai.NewHai(int(a+4*uint64(i)+index), isRedRule)
		if err != nil {
			return hais, callIndex, claimType, err
		}
		hais = append(hais, h)
	}
	return hais, callIndex, claimType, nil
}

/*
pon ポンの面子コード解析

0x0003 - 誰から鳴いたか

0x0004 - 0

0x0008 - 1

0x0010 - 0

0x0060 - 牌添字

0xFE00 - 刻子のパターン
*/
func pon(mentsuCode uint64, isRedRule bool) (hais []hai.IHai, callIndex int, claimType string, err error) {
	claimType = ClaimPon
	pattern := (mentsuCode & 0xFF00) >> 9
	// 何番目の牌を鳴いたか
	callIndex = int(pattern % 3)
	// 牌添字
	haiIndex := (mentsuCode & 0x0060) >> 5
	idx := int(pattern/3) * 4
	var haiIndexes []int
	switch haiIndex {
	case 0:
		haiIndexes = []int{idx + 1, idx + 2, idx + 3}
	case 1:
		haiIndexes = []int{idx, idx + 2, idx + 3}
	case 2:
		haiIndexes = []int{idx, idx + 1, idx + 3}
	case 3:
		haiIndexes = []int{idx, idx + 1, idx + 2}
	default:
		return hais, callIndex, claimType, fmt.Errorf("error at pon: no matched 0~3. actual: %d", haiIndex)
	}
	if callIndex == 1 {
		haiIndexes = []int{haiIndexes[1], haiIndexes[0], haiIndexes[2]}
	}
	if callIndex == 2 {
		haiIndexes = []int{haiIndexes[2], haiIndexes[0], haiIndexes[1]}
	}
	callPlayerIndex := callPlayerIndex(mentsuCode)
	if callPlayerIndex < 3 {
		haiIndexes = []int{haiIndexes[2], haiIndexes[0], haiIndexes[1]}
	}
	if callPlayerIndex < 2 {
		haiIndexes = []int{haiIndexes[2], haiIndexes[0], haiIndexes[1]}
	}
	for _, index := range haiIndexes {
		h, err := hai.NewHai(index, isRedRule)
		if err != nil {
			return hais, callIndex, claimType, err
		}
		hais = append(hais, h)
	}
	return hais, callIndex, claimType, nil
}

/*
chakan 加槓の面子コード解析

0x0003 - 0

0x0004 - 0

0x0008 - 0

0x0010 - 1

0x0060 - 牌添字

0xFE00 - 刻子のパターン
*/
func chakan(mentsuCode uint64, isRedRule bool) (hais []hai.IHai, callIndex int, claimType string, err error) {
	claimType = ClaimChakan
	addedHaiIndex := int((mentsuCode & 0x0060) >> 5)
	pattern := (mentsuCode & 0xFE00) >> 9
	// どの牌を鳴いたか
	haiIndex := pattern % 3
	idx := int(pattern / 3 * 4)
	haiIndexes := []int{idx, idx, idx}
	switch addedHaiIndex {
	case 0:
		haiIndexes = []int{idx + 1, idx + 2, idx + 3}
	case 1:
		haiIndexes = []int{idx, idx + 2, idx + 2}
	case 2:
		haiIndexes = []int{idx, idx + 1, idx + 3}
	case 3:
		haiIndexes = []int{idx, idx + 1, idx + 2}
	default:
		return hais, addedHaiIndex, claimType, fmt.Errorf("no matched 0~3. actual: %d", addedHaiIndex)
	}
	switch haiIndex {
	case 1:
		haiIndexes = []int{haiIndexes[1], haiIndexes[0], haiIndexes[2]}
	case 2:
		haiIndexes = []int{haiIndexes[2], haiIndexes[0], haiIndexes[1]}
	default:
		return hais, addedHaiIndex, claimType, fmt.Errorf("no matched 1~2. actual: %d", haiIndex)
	}
	haiIndexes = append(haiIndexes, idx+int(addedHaiIndex))
	sort.Slice(haiIndexes, func(i, j int) bool {
		return haiIndexes[i] < haiIndexes[j]
	})
	for _, index := range haiIndexes {
		h, err := hai.NewHai(index, isRedRule)
		if err != nil {
			return hais, 0, claimType, err
		}
		hais = append(hais, h)
	}
	return hais, addedHaiIndex, claimType, nil
}

/*
kan 暗槓あるいは大明槓の面子コード解析

0x0003 - 誰から鳴いたか(暗槓は0)

0x0004 - 0

0x0008 - 0

0x0010 - 0

0x0020 - 0

0xFF00 - 槓子のパターン
*/
func kan(mentsuCode uint64, isRedRule bool) (hais []hai.IHai, callIndex int, claimType string, err error) {
	pattern := (mentsuCode & 0xFF00) >> 8
	callPlayerIndex := callPlayerIndex(mentsuCode)
	if callPlayerIndex == 0 {
		claimType = ClaimKan
		pattern = (pattern &^ 3) + 3
	} else {
		claimType = ClaimMinkan
	}
	haiIndex := int(pattern/4) * 4
	haiIndexes := []int{haiIndex, haiIndex, haiIndex}
	switch pattern % 4 {
	case 0:
		haiIndexes = []int{haiIndex + 1, haiIndex + 2, haiIndex + 3}
	case 1:
		haiIndexes = []int{haiIndex, haiIndex + 2, haiIndex + 3}
	case 2:
		haiIndexes = []int{haiIndex, haiIndex + 1, haiIndex + 3}
	case 3:
		haiIndexes = []int{haiIndex, haiIndex + 1, haiIndex + 2}
	default:
		return hais, callIndex, claimType, fmt.Errorf("no matched 0~3. actual: %d", pattern%4)
	}
	if callPlayerIndex == 1 {
		_pattern := pattern
		pattern = uint64(haiIndexes[2])
		haiIndexes[2] = int(_pattern)
	}
	if callPlayerIndex == 2 {
		_pattern := pattern
		pattern = uint64(haiIndexes[0])
		haiIndexes[0] = int(_pattern)
	}
	haiIndexes = append(haiIndexes, int(pattern))
	sort.Slice(haiIndexes, func(i, j int) bool {
		return haiIndexes[i] < haiIndexes[j]
	})
	for _, index := range haiIndexes {
		h, err := hai.NewHai(index, isRedRule)
		if err != nil {
			return hais, callIndex, claimType, err
		}
		hais = append(hais, h)
	}
	return hais, callIndex, claimType, nil
}

func callPlayerIndex(mentsuCode uint64) int {
	return int(mentsuCode & 0x0003)
}

/*
unmarshalMentsuCode 面子コードをparseする

参考: http://tenhou.net/img/tehai.js

0x0003 - 誰から鳴いたか

0x0004 - 順子は必ず1、刻子は必ず0
*/
func (c *Claim) unmarshalMentsuCode(m string) error {
	mNum, err := strconv.ParseUint(m, 10, 16)
	if err != nil {
		return err
	}
	// 自分から見てどのプレイヤーから鳴いたか
	playerIndex, err := xml.NewPlayerIndex(callPlayerIndex(mNum))
	if err != nil {
		return err
	}
	c.playerIndex = *playerIndex
	if (mNum&0x0004)>>2 == 1 {
		hais, callIndex, claimType, err := chii(mNum, c.isRedRule)
		if err != nil {
			return err
		}
		c.hais = hais
		c.callIndex = callIndex
		c.claimType = claimType
	} else if (mNum & (1 << 3)) != 0 {
		hais, callIndex, claimType, err := pon(mNum, c.isRedRule)
		if err != nil {
			return err
		}
		c.hais = hais
		c.callIndex = callIndex
		c.claimType = claimType
	} else if (mNum & (1 << 4)) != 0 {
		hais, callIndex, claimType, err := chakan(mNum, c.isRedRule)
		if err != nil {
			return err
		}
		c.hais = hais
		c.callIndex = callIndex
		c.claimType = claimType
	} else if (mNum & (1 << 5)) != 0 {
		// 北抜き
		c.callIndex = 0
		c.claimType = ClaimNorthDora
	} else {
		hais, callIndex, claimType, err := kan(mNum, c.isRedRule)
		if err != nil {
			return err
		}
		c.hais = hais
		c.callIndex = callIndex
		c.claimType = claimType
	}
	return nil
}
