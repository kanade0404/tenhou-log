package operation

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/pkg/parser/xml"
	"strconv"
)

type Reach struct {
	playerIndex xml.PlayerIndex
	isComplete  bool
}

func (r *Reach) PlayerIndex() xml.PlayerIndex {
	return r.playerIndex
}

func (r *Reach) IsComplete() bool {
	return r.isComplete
}

func NewReach(who string, step string) (*Reach, error) {
	idx, err := strconv.Atoi(who)
	if err != nil {
		return nil, err
	}
	playerIndex, err := xml.NewPlayerIndex(idx)
	if err != nil {
		return nil, err
	}
	var isComplete bool
	if step == "1" {
		isComplete = false
	} else if step == "2" {
		isComplete = true
	} else {
		return nil, fmt.Errorf("step must be 1 or 2. actual: %s", step)
	}
	return &Reach{
		*playerIndex,
		isComplete,
	}, nil
}
