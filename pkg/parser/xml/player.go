package xml

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/parser/str"
	"net/url"
	"strconv"
)

type PlayerName string

func DecodePlayerName(encoded string) (PlayerName, error) {
	name, err := url.QueryUnescape(encoded)
	if err != nil {
		return "", err
	}
	return PlayerName(name), nil
}

// Encode プレイヤー名をencodeする
func (n PlayerName) Encode() string {
	var result string
	// https://gist.github.com/hnaohiro/4627658
	for _, c := range n {
		if c <= 0x7f { // single byte
			result += fmt.Sprintf("%%%X", c)
		} else if c > 0x1fffff { // quaternary byte
			result += fmt.Sprintf("%%%X%%%X%%%X%%%X",
				0xf0+((c&0x1c0000)>>18),
				0x80+((c&0x3f000)>>12),
				0x80+((c&0xfc0)>>6),
				0x80+(c&0x3f),
			)
		} else if c > 0x7ff { // triple byte
			result += fmt.Sprintf("%%%X%%%X%%%X",
				0xe0+((c&0xf000)>>12),
				0x80+((c&0xfc0)>>6),
				0x80+(c&0x3f),
			)
		} else { // double byte
			result += fmt.Sprintf("%%%X%%%X",
				0xc0+((c&0x7c0)>>6),
				0x80+(c&0x3f),
			)
		}
	}
	return result
}

type Rate float64

func (r Rate) String() string {
	return fmt.Sprintf("%.2f", r)
}

type Dan uint32

func (d Dan) String() string {
	return strconv.Itoa(int(d))
}

const (
	NewBie = iota
	NineKyuu
	EightKyuu
	SevenKyuu
	SixKyuu
	FiveKyuu
	FourKyuu
	ThreeKyuu
	TwoKyuu
	OneKyuu
	OneDan
	TwoDan
	ThreeDan
	FourDan
	FiveDan
	SixDan
	SevenDan
	EightDan
	NineDan
	TenDan
	Tenhoui
)

func (d Dan) Name() string {
	switch d {
	case NewBie:
		return "新人"
	case NineKyuu:
		return "9級"
	case EightKyuu:
		return "8級"
	case SevenKyuu:
		return "7級"
	case SixKyuu:
		return "6級"
	case FiveKyuu:
		return "5級"
	case FourKyuu:
		return "4級"
	case ThreeKyuu:
		return "3級"
	case TwoKyuu:
		return "2級"
	case OneKyuu:
		return "1級"
	case OneDan:
		return "初段"
	case TwoDan:
		return "二段"
	case ThreeDan:
		return "三段"
	case FourDan:
		return "四段"
	case FiveDan:
		return "五段"
	case SixDan:
		return "六段"
	case SevenDan:
		return "七段"
	case EightDan:
		return "八段"
	case NineDan:
		return "九段"
	case TenDan:
		return "十段"
	case Tenhoui:
		return "天鳳位"
	default:
		return "UNKNOWN"
	}
}

type Sex string

const (
	F Sex = "F"
	M Sex = "M"
)

func (s Sex) Name() string {
	switch s {
	case F:
		return "男性"
	case M:
		return "女性"
	default:
		return "UNKNOWN"
	}
}

func (s Sex) String() string {
	return string(s)
}

type Player struct {
	Name PlayerName
	Rate
	Dan
	Sex
}

type PlayerIndex int

func (p *PlayerIndex) String() string {
	return strconv.Itoa(int(*p))
}

var AllPlayerIndexes = []PlayerIndex{0, 1, 2, 3}

func NewPlayerIndex(idx int) (*PlayerIndex, error) {
	if idx < 0 || idx > 3 {
		return nil, fmt.Errorf("player index must be between 0 and 3 but %d", idx)
	}
	i := PlayerIndex(idx)
	return &i, nil
}

func NewPlayerIndexFromString(pi string) (*PlayerIndex, error) {
	idx, err := strconv.Atoi(pi)
	if err != nil {
		return nil, err
	}
	i, err := NewPlayerIndex(idx)
	if err != nil {
		return nil, err
	}
	return i, nil
}

type Players []Player

func (p *Players) Unmarshal(e XmlElement) error {
	rates := str.SplitByComma(e.AttrByName("rate"))
	dans, err := str.SplitByCommaAsInt(e.AttrByName("dan"))
	if err != nil {
		return err
	}
	sexes := str.SplitByComma(e.AttrByName("sx"))

	for _, pi := range AllPlayerIndexes {
		name, err := DecodePlayerName(e.AttrByName(fmt.Sprintf("n%d", pi)))
		if err != nil {
			return err
		}
		// サンマの4人目の場合
		if name == "" {
			continue
		}
		rate, err := strconv.ParseFloat(rates[pi], 64)
		dan := dans[pi]
		sex := sexes[pi]

		*p = append(*p, Player{
			Name: name,
			Rate: Rate(rate),
			Dan:  Dan(dan),
			Sex:  Sex(sex),
		})
	}
	return nil
}

func (p Players) Marshal() (XmlElement, error) {
	var dan []string
	var rate []string
	var sx []string
	var n []string
	for i := range AllPlayerIndexes {
		// n0-n3 (サンマの場合はデータが無いのでデフォルト値を追加する)
		player := Player{}
		if i < len(p) {
			player = (p)[i]
		}
		n = append(n, player.Name.Encode())
		dan = append(dan, player.Dan.String())
		rate = append(rate, player.Rate.String())
		sx = append(sx, player.Sex.String())
	}
	return newXmlElement("UN", NewXmlAttrs(map[string]string{
		"n0":   n[0],
		"n1":   n[1],
		"n2":   n[2],
		"n3":   n[3],
		"rate": str.JoinByComma(rate),
		"dan":  str.JoinByComma(dan),
		"sx":   str.JoinByComma(sx),
	})), nil
}
