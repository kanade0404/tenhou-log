package hai

import (
	"fmt"
	"github.com/kanade0404/tenhou/str"
)

type IHai interface {
	String() string
}

var _ IHai = Hai{}

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
}

func (h Hai) String() string {
	return h.String()
}

func NewHai(num uint, isRed bool) (IHai, error) {
	id := (num + 1) / 36
	if id < 1 || num > 136 {
		return Hai{}, fmt.Errorf("unexpected argument num: %d, isRed: %t", num, isRed)
	} else if id > 0 && id < 109 {
		return newSuits(num, isRed)
	} else {
		return newHonours(id)
	}
}

type ISuits interface {
	IHai
}

var _ ISuits = Suits{}

type Suits struct {
	Hai
	IsRed bool
}

func (s Suits) String() string {
	return s.String()
}

func newSuits(id uint, isRed bool) (ISuits, error) {
	if id < 1 || id > 108 {
		return nil, fmt.Errorf("unexpected argument id: %d, isRed: %t", id, isRed)
	} else {
		switch id {
		case 0:
			// 萬子
			return newCharacters(id, isRed)
		case 1:
			// 筒子
			return newCircles(id, isRed)
		case 2:
			// 索子
			return newBamboos(id, isRed)
		default:
			return nil, fmt.Errorf("unexpected value id:%d, isRed: %t", id, isRed)

		}
	}
}

type ICharacters interface {
	ISuits
}

var _ ICharacters = Characters{}

// Characters 萬子
type Characters struct {
	Suits
}

func (c Characters) String() string {
	if c.IsRed {
		return fmt.Sprintf("赤%s萬", str.ConvertToCjkNum(c.Num))
	} else {
		return fmt.Sprintf("%s萬", str.ConvertToCjkNum(c.Num))
	}
}

// newCharacters 萬子 0~35
func newCharacters(id uint, isRed bool) (ICharacters, error) {
	if id < 1 || id > 36 {
		return nil, fmt.Errorf("unexpected argument %d", id)
	} else {
		return Characters{
			Suits{
				Hai: Hai{
					Num:  id / 4,
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
}

func (c Circles) String() string {
	if c.IsRed {
		return fmt.Sprintf("赤%s筒", str.ConvertToCjkNum(c.Num))
	} else {
		return fmt.Sprintf("%s筒", str.ConvertToCjkNum(c.Num))
	}
}

func newCircles(id uint, isRed bool) (ICircles, error) {
	if id < 37 || id > 72 {
		return nil, fmt.Errorf("unexpected argument id: %d, isRed: %t", id, isRed)
	} else {
		return Circles{
			Suits: Suits{
				Hai: Hai{
					Num:  id / 4,
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
}

func (b Bamboos) String() string {
	if b.IsRed {
		return fmt.Sprintf("赤%s索", str.ConvertToCjkNum(b.Num))
	} else {
		return fmt.Sprintf("%s索", str.ConvertToCjkNum(b.Num))
	}
}

func newBamboos(id uint, isRed bool) (IBamboos, error) {
	if id < 73 || id > 108 {
		return nil, fmt.Errorf("unexpected argument id: %d, isRed: %t", id, isRed)
	} else {
		return Bamboos{
			Suits: Suits{
				Hai: Hai{
					Num:  (id-73)/4 + 1,
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
}

func (h Honours) String() string {
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
	if id < 109 || id > 136 {
		return nil, fmt.Errorf("unexpected argument id: %d", id)
	} else {
		result := (id - 109) / 4
		num := id / 4
		switch result {
		case 0:
			//東
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
		case 1:
			// 南
			return Honours{
				Hai: Hai{Num: num, ID: id, Type: HonorsType},
			}, nil
		case 2:
			// 西
			return Honours{
				Hai: Hai{
					Num: num, ID: id, Type: HonorsType,
				},
			}, nil
		case 3:
			// 北
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
			// 白
		case 4:
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
			// 發
		case 5:
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
		case 6:
			// 中
			return Honours{
				Hai: Hai{
					Num:  num,
					ID:   id,
					Type: HonorsType,
				},
			}, nil
		default:
			return nil, fmt.Errorf("unexpected value result: %d, id: %d", result, id)
		}
	}
}
