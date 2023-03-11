package init

import (
	"github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml/hai"
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
		i, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		newHai, err := hai.NewHai(i, h.gameInfo.Red)
		if err != nil {
			return err
		}
		h.Hands = append(h.Hands, newHai)
	}
	return nil
}
