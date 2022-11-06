package operation

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"github.com/kanade0404/tenhou-log/serdes/xml/hai"
	"strconv"
)

// Claim 副露情報
type Claim struct {
	userIndex xml.PlayerIndex // 誰から鳴いたか
	hais      []hai.IHai      // 鳴いた手牌
	callIndex int             // 鳴いた牌のインデックス
	*xml.GameInfo
}

func (c *Claim) CallFromUser() xml.PlayerIndex {
	return c.userIndex
}

func (c *Claim) Hais() []hai.IHai {
	return c.hais
}
func (c *Claim) CallIndex() int {
	return c.callIndex
}

type Chii struct {
	Hais      []hai.IHai
	CallIndex int
}
type Pon struct {
	FromUserIndex xml.PlayerIndex
	Hais          []hai.IHai
	CallIndex     int
}

func claimFrom(claim uint64) xml.PlayerIndex {
	switch claim {
	case 0:
		return 0
	case 1, 2, 3:
		return xml.PlayerIndex(claim + 1)
	default:
		return 0
	}
}

// HandHais http://tenhou.net/img/tehai.js
func (c *Claim) HandHais(m string) error {
	mNum, err := strconv.ParseUint(m, 10, 16)
	if err != nil {
		return err
	}
	/**
	面子コード
	- 順子
	0x0003 - 誰から鳴いたか
	0x0004 - 1
	0x0018 - 牌添字1
	0x0060 - 牌添字2
	0x0180 - 牌添字3
	0xFC00 - 順子のパターン
	- 刻子か加槓
	0x0003 - 誰から鳴いたか
	0x0004 - 0
	0x0008 - 刻子なら1
	0x0010 - 加槓なら1
	0x0060 - 牌添字
	0xFE00 - 刻子のパターン
	- 暗槓もしくは大明槓
	0x0003 - 誰から鳴いたか(暗槓は0)
	0x0004 - 0
	0x0008 - 0
	0x0010 - 0
	0x0020 - 0
	0xFF00 - 槓子のパターン
	*/
	/**
	誰から鳴いたか
	0: 鳴きなし
	1: 下家
	2: 対面
	3: 上家
	*/
	claimFromUser := mNum & 0x0003
	/**
	true: 順子
	false: 刻子
	*/
	isChow := (mNum&0x0004)>>2 == 1
	/**
	true: 刻子
	false: 槓子
	*/
	//isPung := (mNum&0x0008)>>3 == 1
	// 加槓かどうか
	//isToExtendMeldedKong := (mNum&0x0010)>>4 == 1
	if isChow {
		/**
		面子コード
		順子
		0x0003 - 誰から鳴いたか
		0x0004 - 1
		0x0018 - 牌添字1
		0x0060 - 牌添字2
		0x0180 - 牌添字3
		0xFC00 - 順子のパターン
		*/
		// 順子
		pattern := (mNum & 0xFC00) >> 10
		// 何番目の牌を鳴いたか
		c.callIndex = int(pattern % 3)
		idx, err := xml.NewPlayerIndex(int(mNum & 0x0003))
		if err != nil {
			return err
		}
		c.userIndex = idx
		haiIndexes := []uint64{
			(mNum & 0x0018) >> 3,
			(mNum & 0x0060) >> 5,
			(mNum & 0x0180) >> 7,
		}
		a := ((pattern/3/7)*9 + (pattern / 3 % 7)) * 4
		for i, index := range haiIndexes {
			h, err := hai.NewHai(uint(a+4*uint64(i)+index), c.GameInfo.Red)
			if err != nil {
				return err
			}
			c.hais = append(c.hais, h)
		}
	} else if (mNum & (1 << 3)) != 0 {
		/**
		面子コード
		刻子
		0x0003 - 誰から鳴いたか
		0x0004 - 0
		0x0008 - 1
		0x0010 - 0
		0x0060 - 牌添字
		0xFE00 - 刻子のパターン
		*/
		// 刻子
		pattern := (mNum & 0xFF00) >> 9
		// 何番目の牌を鳴いたか
		c.callIndex = int(pattern % 3)
		// 牌添字
		unUsed := (mNum & 0x0060) >> 5
		idx := int(pattern/3) * 4
		if err != nil {
			return err
		}
		i, err := xml.NewPlayerIndex(int(mNum & 0x0003))
		if err != nil {
			return err
		}
		c.userIndex = i
		var haiIndexes []int
		switch unUsed {
		case 0:
			haiIndexes = []int{idx + 1, idx + 2, idx + 3}
		case 1:
			haiIndexes = []int{idx, idx + 2, idx + 3}
		case 2:
			haiIndexes = []int{idx, idx + 1, idx + 3}
		case 3:
			haiIndexes = []int{idx, idx + 1, idx + 2}
		default:
			return fmt.Errorf("no matched 0~3. actual: %d", unUsed)
		}
		switch c.callIndex {
		case 1:
			haiIndexes = []int{haiIndexes[1], haiIndexes[0], haiIndexes[2]}
		case 2:
			haiIndexes = []int{haiIndexes[2], haiIndexes[0], haiIndexes[1]}
		default:
			return fmt.Errorf("no matched 1~2. actual: %d", c.callIndex)
		}
		if claimFromUser < 3 {
			haiIndexes = []int{haiIndexes[2], haiIndexes[0], haiIndexes[1]}
		}
		if claimFromUser < 2 {
			haiIndexes = []int{haiIndexes[2], haiIndexes[0], haiIndexes[1]}
		}
		for _, index := range haiIndexes {
			h, err := hai.NewHai(uint(index), c.GameInfo.Red)
			if err != nil {
				return err
			}
			c.hais = append(c.hais, h)
		}
	} else if (mNum & (1 << 4)) != 0 {
		/**
		面子コード
		加槓
		0x0003 - 誰から鳴いたか
		0x0004 - 0
		0x0008 - 0
		0x0010 - 1
		0x0060 - 牌添字
		0xFE00 - 刻子のパターン
		*/
		added := (mNum & 0x0060) >> 5
		a := (mNum & 0xFE00) >> 5
		r := a % 3
		c.callIndex = int(claimFromUser)
		t := int(a / 3 * 4)
		haiIndexes := []int{t, t, t}
		switch added {
		case 0:
			haiIndexes = []int{t + 1, t + 2, t + 3}
		case 1:
			haiIndexes = []int{t, t + 2, t + 2}
		case 2:
			haiIndexes = []int{t, t + 1, t + 3}
		case 3:
			haiIndexes = []int{t, t + 1, t + 2}
		default:
			return fmt.Errorf("no matched 0~3. actual: %d", added)
		}
		switch r {
		case 1:
			haiIndexes = []int{haiIndexes[1], haiIndexes[0], haiIndexes[2]}
		case 2:
			haiIndexes = []int{haiIndexes[2], haiIndexes[0], haiIndexes[1]}
		default:
			return fmt.Errorf("no matched 1~2. actual: %d", r)
		}
		for _, index := range haiIndexes {
			h, err := hai.NewHai(uint(index), c.GameInfo.Red)
			if err != nil {
				return err
			}
			c.hais = append(c.hais, h)
		}
	} else if (mNum & (1 << 5)) != 0 {
		// 北抜き
		return nil
	} else {
		/**
		面子コード
		- 暗槓もしくは大明槓
		0x0003 - 誰から鳴いたか(暗槓は0)
		0x0004 - 0
		0x0008 - 0
		0x0010 - 0
		0x0020 - 0
		0xFF00 - 槓子のパターン
		*/
		hai0 := (mNum & 0xFF00) >> 8
		if claimFromUser != 0 {
			hai0 = (hai0 &^ 3) + 3
		}
		idx, err := xml.NewPlayerIndex(int(mNum & 0x0003))
		if err != nil {
			return err
		}
		c.userIndex = idx
		t := int(hai0/4) * 4
		haiIndexes := []int{t, t, t}
		switch hai0 % 4 {
		case 0:
			haiIndexes = []int{t + 1, t + 2, t + 3}
		case 1:
			haiIndexes = []int{t, t + 2, t + 3}
		case 2:
			haiIndexes = []int{t, t + 1, t + 3}
		case 3:
			haiIndexes = []int{t, t + 1, t + 2}
		default:
			return fmt.Errorf("no matched 0~3. actual: %d", hai0%4)
		}
		if claimFromUser == 1 {
			tmp := hai0
			hai0 = uint64(haiIndexes[2])
			haiIndexes[2] = int(tmp)
		}
		if claimFromUser == 2 {
			tmp := hai0
			hai0 = uint64(haiIndexes[0])
			haiIndexes[0] = int(tmp)
		}
		for _, index := range haiIndexes {
			h, err := hai.NewHai(uint(index), c.GameInfo.Red)
			if err != nil {
				return err
			}
			c.hais = append(c.hais, h)
		}

	}
	return nil
}
