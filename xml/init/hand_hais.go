package init

import (
	"github.com/kanade0404/tenhou-log/xml"
	"github.com/kanade0404/tenhou-log/xml/hai"
	"strconv"
	"strings"
)

type HandHais struct {
	Hands    []hai.IHai
	gameInfo *xml.GameInfo
	_        struct{}
}

func NewHandHais(hands []hai.IHai, gameInfo *xml.GameInfo) HandHais {
	return HandHais{
		Hands:    hands,
		gameInfo: gameInfo,
	}
}

func (h HandHais) Marshal() string {
	var results []string
	for _, hand := range h.Hands {
		results = append(results, strconv.Itoa(int(hand.ID())))
	}
	return strings.Join(results, ",")
}

func (h *HandHais) Unmarshal(hais string) error {
	haiIds := strings.Split(hais, ",")
	for _, id := range haiIds {
		i, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return err
		}
		hai, err := hai.NewHai(uint(i), h.gameInfo.Red)
		if err != nil {
			return err
		}
		h.Hands = append(h.Hands, hai)
	}
	return nil
}
