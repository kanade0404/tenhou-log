package operation

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"github.com/kanade0404/tenhou-log/serdes/xml/hai"
)

type DrawType string

const (
	FirstUser  DrawType = "T"
	SecondUser DrawType = "U"
	ThirdUser  DrawType = "V"
	FourthUser DrawType = "W"
)
const (
	FirstUserIndex  = iota // 自身
	SecondUserIndex        // 下家
	ThirdUserIndex         // 対面
	FourthUserIndex        // 上家
)

func (t DrawType) PlayerIndex() uint {
	switch t {
	case FirstUser:
		return FirstUserIndex
	case SecondUser:
		return SecondUserIndex
	case ThirdUser:
		return ThirdUserIndex
	case FourthUser:
		return FourthUserIndex
	default:
		return 0
	}
}
func IsDrawType(drawType string) bool {
	switch drawType {
	case string(FirstUser), string(SecondUser), string(ThirdUser), string(FourthUser):
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

func NewDraw(drawType string, haiID uint, gameInfo *xml.GameInfo) (*Draw, error) {
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
