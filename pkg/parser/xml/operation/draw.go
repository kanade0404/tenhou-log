package operation

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml/hai"
)

type DrawType string

const (
	FirstDrawUser  DrawType = "T"
	SecondDrawUser DrawType = "U"
	ThirdDrawUser  DrawType = "V"
	FourthDrawUser DrawType = "W"
)

func (t DrawType) PlayerIndex() uint {
	switch t {
	case FirstDrawUser:
		return FirstUserIndex
	case SecondDrawUser:
		return SecondUserIndex
	case ThirdDrawUser:
		return ThirdUserIndex
	case FourthDrawUser:
		return FourthUserIndex
	default:
		return 0
	}
}
func IsDrawType(drawType string) bool {
	switch drawType {
	case string(FirstDrawUser), string(SecondDrawUser), string(ThirdDrawUser), string(FourthDrawUser):
		return true
	}
	return false
}

// Draw ツモ情報
type Draw struct {
	DrawType
	Hai      hai.IHai
	GameInfo *xml.GameInfo
	_        struct{}
}

func NewDraw(drawType string, haiID int, gameInfo *xml.GameInfo) (*Draw, error) {
	drawHai, err := hai.NewHai(haiID, gameInfo.Red)
	isDrawType := IsDrawType(drawType)
	if isDrawType && err == nil {
		return &Draw{
			DrawType: DrawType(drawType),
			Hai:      drawHai,
			GameInfo: gameInfo,
		}, nil
	}
	return nil, fmt.Errorf(": error: %w, \nisDrawType: %t", err, isDrawType)
}
