package operation

import (
	"github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml/hai"
)

type Discard struct {
	discardUserIndex int
	hai              hai.IHai
	isRedRule        bool
	_                struct{}
}

type discardUser string

const (
	FirstDiscardUser  discardUser = "D"
	SecondDiscardUser discardUser = "E"
	ThirdDiscardUser  discardUser = "F"
	FourthDiscardUser discardUser = "G"
)

func (d discardUser) PlayerIndex() int {
	switch d {
	case FirstDiscardUser:
		return FirstUserIndex
	case SecondDiscardUser:
		return SecondUserIndex
	case ThirdDiscardUser:
		return ThirdUserIndex
	case FourthDiscardUser:
		return FourthUserIndex
	default:
		return 0
	}
}

func (d *Discard) DiscardUserIndex() int {
	return d.discardUserIndex
}
func (d *Discard) Hai() hai.IHai {
	return d.hai
}

/*
NewDiscard

打牌情報
*/
func NewDiscard(discardType string, haiID int, gameInfo *xml.GameInfo) (*Discard, error) {
	discardHai, err := hai.NewHai(haiID, gameInfo.Red)
	if err != nil {
		return nil, err
	}
	return &Discard{
		discardUserIndex: discardUser(discardType).PlayerIndex(),
		hai:              discardHai,
		isRedRule:        gameInfo.Red,
	}, nil
}
