package hai

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/str"
)

type IHai interface {
	Name() string
}

type HaiType string

const (
	CharactersType HaiType = "萬子"
	BamboosType    HaiType = "索子"
	CirclesType    HaiType = "筒子"
	HonorsType     HaiType = "字牌"
)

type Hai struct {
	ID   uint
	Num  uint
	Type HaiType
	_    struct{}
}

func NewHai(id uint, isRed bool) (IHai, error) {
	if id < 0 || id > 135 {
		return nil, fmt.Errorf("unexpected argument id: %d, isRed: %t", id, isRed)
	} else if id >= 0 && id <= 107 {
		return newSuits(id, isRed)
	} else {
		return newHonours(id)
	}
}

type ISuits interface {
	IHai
	redFiveID() uint
}

type Suits struct {
	Hai
	IsRed bool
	_     struct{}
}

func newSuits(id uint, isRed bool) (ISuits, error) {
	if id < 0 || id > 107 {
		return nil, fmt.Errorf("unexpected argument id: %d, isRed: %t", id, isRed)
	} else if id >= 0 && id <= 35 {
		return newCharacters(id, isRed)
	} else if id >= 36 && id <= 71 {
		return newCircles(id, isRed)
	} else {
		return newBamboos(id, isRed)
	}
}

type ICharacters interface {
	ISuits
}

var _ ICharacters = Characters{}

// Characters 萬子
type Characters struct {
	Suits
	_ struct{}
}

func (c Characters) redFiveID() uint {
	return 16
}

func (c Characters) Name() string {
	cjkNum := str.ConvertToCjkNum(c.Num)
	if c.IsRed && c.ID == c.redFiveID() {
		return fmt.Sprintf("赤%s萬", cjkNum)
	} else {
		return fmt.Sprintf("%s萬", cjkNum)
	}
}

// newCharacters 萬子 0~35
func newCharacters(id uint, isRed bool) (ICharacters, error) {
	if id < 0 || id > 35 {
		return nil, fmt.Errorf("unexpected argument %d", id)
	} else {
		rs := (id + 1) / 4
		mod := (id + 1) % 4
		var num uint
		if mod == 0 {
			num = rs
		} else {
			num = rs + 1
		}
		return Characters{
			Suits: Suits{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: CharactersType,
				},
				IsRed: isRed,
			},
		}, nil
	}
}

type ICircles interface {
	ISuits
}

var _ ICircles = Circles{}

// Circles 筒子
type Circles struct {
	Suits
	_ struct{}
}

func (c Circles) redFiveID() uint {
	return 52
}

func (c Circles) Name() string {
	cjkNum := str.ConvertToCjkNum(c.Num)
	if c.IsRed && c.ID == c.redFiveID() {
		return fmt.Sprintf("赤%s筒", cjkNum)
	} else {
		return fmt.Sprintf("%s筒", str.ConvertToCjkNum(c.Num))
	}
}

func newCircles(id uint, isRed bool) (ICircles, error) {
	if id < 36 || id > 71 {
		return nil, fmt.Errorf("unexpected argument id: %d, isRed: %t", id, isRed)
	} else {
		rs := (id - 36 + 1) / 4
		mod := (id - 36 + 1) % 4
		var num uint
		if mod == 0 {
			num = rs
		} else {
			num = rs + 1
		}
		return Circles{
			Suits: Suits{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: CirclesType,
				},
				IsRed: isRed,
			},
		}, nil
	}
}

type IBamboos interface {
	ISuits
}

var _ IBamboos = Bamboos{}

// Bamboos 索子
type Bamboos struct {
	Suits
	_ struct{}
}

func (b Bamboos) redFiveID() uint {
	return 88
}

func (b Bamboos) Name() string {
	cjkNum := str.ConvertToCjkNum(b.Num)
	if b.IsRed && b.ID == b.redFiveID() {
		return fmt.Sprintf("赤%s索", cjkNum)
	} else {
		return fmt.Sprintf("%s索", cjkNum)
	}
}

func newBamboos(id uint, isRed bool) (IBamboos, error) {
	if id < 72 || id > 107 {
		return nil, fmt.Errorf("unexpected argument id: %d, isRed: %t", id, isRed)
	} else {
		rs := (id - 72 + 1) / 4
		mod := (id - 72 + 1) % 4
		var num uint
		if mod == 0 {
			num = rs
		} else {
			num = rs + 1
		}
		return Bamboos{
			Suits: Suits{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: BamboosType,
				},
				IsRed: isRed,
			},
		}, nil
	}
}

type IHonours interface {
	IHai
}

var _ IHonours = Honours{}

type Honours struct {
	Hai
	_ struct{}
}

func (h Honours) Name() string {
	switch (h.ID - 109) / 4 {
	case 0:
		return "東"
	case 1:
		return "南"
	case 2:
		return "西"
	case 3:
		return "北"
	case 4:
		return "白"
	case 5:
		return "發"
	case 6:
		return "中"
	default:
		return ""
	}
}

func newHonours(id uint) (IHonours, error) {
	if id < 108 || id > 136 {
		return nil, fmt.Errorf("unexpected argument id: %d", id)
	} else {
		rs := (id - 108 + 1) / 4
		mod := (id - 108 + 1) % 4
		var num uint
		if mod == 0 {
			num = rs
		} else {
			num = rs + 1
		}
		switch num {
		case 1:
			//東
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
		case 2:
			// 南
			return Honours{
				Hai: Hai{Num: num, ID: id, Type: HonorsType},
			}, nil
		case 3:
			// 西
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
		case 4:
			// 北
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
			// 白
		case 5:
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
			// 發
		case 6:
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
		case 7:
			// 中
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
		default:
			return nil, fmt.Errorf("unexpected value num: %d, id: %d", num, id)
		}
	}
}
