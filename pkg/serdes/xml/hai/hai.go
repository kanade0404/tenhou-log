package hai

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/serdes/str"
)

type IHai interface {
	Name() string
	ID() uint
	Num() uint
	Type() HaiType
	IsRed() bool
	String() string
}

type HaiType string

const (
	CharactersType  HaiType = "萬子"
	BamboosType     HaiType = "索子"
	CirclesType     HaiType = "筒子"
	HonorsType      HaiType = "字牌"
	CharactersRedID         = 16
	CirclesRedID            = 52
	BambooRedID             = 88
)

type Hai struct {
	id        uint
	num       uint
	haiType   HaiType
	isRedRule bool
	_         struct{}
}

func (h *Hai) String() string {
	return fmt.Sprintf("id:%d,num:%d,haiType:%s,isRedRule:%t", h.id, h.num, h.haiType, h.isRedRule)
}

func NewHai(id int, isRed bool) (IHai, error) {
	if id < 0 || id > 135 {
		return nil, fmt.Errorf("unexpected argument id: %d, IsRed: %t", id, isRed)
	}
	i := uint(id)
	if id >= 0 && id <= 107 {
		if id >= 0 && id <= 35 {
			return newCharacters(i, isRed)
		} else if id >= 36 && id <= 71 {
			return newCircles(i, isRed)
		} else {
			return newBamboos(i, isRed)
		}
	} else {
		return newHonours(i, isRed)
	}
}

func Reds(isRed bool) []IHai {
	if !isRed {
		return []IHai{}
	}
	return []IHai{&Characters{
		Hai: &Hai{
			id:        CharactersRedID,
			num:       5,
			haiType:   CharactersType,
			isRedRule: true,
		},
	},
		&Circles{
			Hai: &Hai{
				id:        CirclesRedID,
				num:       5,
				haiType:   CirclesType,
				isRedRule: true,
			},
		},
		&Bamboos{
			Hai: &Hai{
				id:        BambooRedID,
				num:       5,
				haiType:   BamboosType,
				isRedRule: true,
			},
		}}
}

type ICharacters interface {
	IHai
}

var _ ICharacters = &Characters{}

// Characters 萬子
type Characters struct {
	*Hai
	_ struct{}
}

func (c *Characters) ID() uint {
	return c.id
}

func (c *Characters) Num() uint {
	return c.num
}

func (c *Characters) Type() HaiType {
	return CharactersType
}

func (c *Characters) IsRed() bool {
	return c.id == CharactersRedID && c.isRedRule
}

func (c *Characters) Name() string {
	cjkNum := str.ConvertToCjkNum(c.Num())
	if c.IsRed() && c.isRedRule {
		return fmt.Sprintf("赤%s萬", cjkNum)
	} else {
		return fmt.Sprintf("%s萬", cjkNum)
	}
}

// newCharacters 萬子 0~35
func newCharacters(id uint, isRed bool) (IHai, error) {
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
		return &Characters{
			Hai: &Hai{
				num:       num,
				id:        id,
				haiType:   CharactersType,
				isRedRule: isRed,
			},
		}, nil
	}
}

type ICircles interface {
	IHai
}

var _ ICircles = &Circles{}

// Circles 筒子
type Circles struct {
	*Hai
	_ struct{}
}

func (c *Circles) ID() uint {
	return c.id
}

func (c *Circles) Num() uint {
	return c.num
}

func (c *Circles) Type() HaiType {
	return CirclesType
}

func (c *Circles) IsRed() bool {
	return c.id == CirclesRedID && c.isRedRule
}

func (c *Circles) Name() string {
	cjkNum := str.ConvertToCjkNum(c.Num())
	if c.IsRed() && c.isRedRule {
		return fmt.Sprintf("赤%s筒", cjkNum)
	} else {
		return fmt.Sprintf("%s筒", str.ConvertToCjkNum(c.Num()))
	}
}

func newCircles(id uint, isRed bool) (IHai, error) {
	if id < 36 || id > 71 {
		return nil, fmt.Errorf("unexpected argument id: %d, IsRed: %t", id, isRed)
	} else {
		rs := (id - 36 + 1) / 4
		mod := (id - 36 + 1) % 4
		var num uint
		if mod == 0 {
			num = rs
		} else {
			num = rs + 1
		}
		return &Circles{
			Hai: &Hai{
				num:       num,
				id:        id,
				haiType:   CirclesType,
				isRedRule: isRed,
			},
		}, nil
	}
}

type IBamboos interface {
	IHai
}

var _ IBamboos = &Bamboos{}

// Bamboos 索子
type Bamboos struct {
	*Hai
	_ struct{}
}

func (b *Bamboos) ID() uint {
	return b.id
}

func (b *Bamboos) Num() uint {
	return b.num
}

func (b *Bamboos) Type() HaiType {
	return BamboosType
}

func (b *Bamboos) IsRed() bool {
	return b.id == BambooRedID && b.isRedRule
}

func (b *Bamboos) Name() string {
	cjkNum := str.ConvertToCjkNum(b.Num())
	if b.IsRed() && b.isRedRule {
		return fmt.Sprintf("赤%s索", cjkNum)
	} else {
		return fmt.Sprintf("%s索", cjkNum)
	}
}

func newBamboos(id uint, isRed bool) (IHai, error) {
	if id < 72 || id > 107 {
		return nil, fmt.Errorf("unexpected argument id: %d, IsRed: %t", id, isRed)
	} else {
		rs := (id - 72 + 1) / 4
		mod := (id - 72 + 1) % 4
		var num uint
		if mod == 0 {
			num = rs
		} else {
			num = rs + 1
		}
		return &Bamboos{
			Hai: &Hai{
				num:       num,
				id:        id,
				haiType:   BamboosType,
				isRedRule: isRed,
			},
		}, nil
	}
}

type IHonours interface {
	IHai
}

var _ IHonours = &Honours{}

type Honours struct {
	Hai
	_ struct{}
}

func (h *Honours) String() string {
	return fmt.Sprintf("{id:%d,num:%d,haiType:%s}", h.id, h.num, h.haiType)
}

func (h *Honours) ID() uint {
	return h.id
}

func (h *Honours) Num() uint {
	return h.num
}

func (h *Honours) Type() HaiType {
	return HonorsType
}

func (h *Honours) Name() string {
	switch (h.ID() - 109) / 4 {
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

func (h *Honours) IsRed() bool {
	return false
}

func newHonours(id uint, isRed bool) (IHai, error) {
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
			return &Honours{
				Hai: Hai{
					num:       num,
					id:        id,
					haiType:   HonorsType,
					isRedRule: isRed,
				},
			}, nil
		case 2:
			// 南
			return &Honours{
				Hai: Hai{num: num, id: id, haiType: HonorsType, isRedRule: isRed},
			}, nil
		case 3:
			// 西
			return &Honours{
				Hai: Hai{
					num:       num,
					id:        id,
					haiType:   HonorsType,
					isRedRule: isRed,
				},
			}, nil
		case 4:
			// 北
			return &Honours{
				Hai: Hai{
					num:       num,
					id:        id,
					haiType:   HonorsType,
					isRedRule: isRed,
				},
			}, nil
			// 白
		case 5:
			return &Honours{
				Hai: Hai{
					num:       num,
					id:        id,
					haiType:   HonorsType,
					isRedRule: isRed,
				},
			}, nil
			// 發
		case 6:
			return &Honours{
				Hai: Hai{
					num:       num,
					id:        id,
					haiType:   HonorsType,
					isRedRule: isRed,
				},
			}, nil
		case 7:
			// 中
			return &Honours{
				Hai: Hai{
					num:       num,
					id:        id,
					haiType:   HonorsType,
					isRedRule: isRed,
				},
			}, nil
		default:
			return nil, fmt.Errorf("unexpected value num: %d, id: %d", num, id)
		}
	}
}
