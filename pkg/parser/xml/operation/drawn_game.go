package operation

import "errors"

type DrawnType string

const (
	nm                = "nm"
	yao9              = "yao9"
	wind4             = "kaze4"
	reach4            = "reach4"
	ron3              = "ron3"
	kan4              = "kan4"
	Nagashi DrawnType = nm
	Kyuusyu DrawnType = yao9
	Wind4   DrawnType = wind4
	Reach4  DrawnType = reach4
	Ron3    DrawnType = ron3
	Kan4    DrawnType = kan4
)

func newDrawnType(drawnType string) DrawnType {
	switch drawnType {
	case nm:
		return Nagashi
	case yao9:
		return Kyuusyu
	case wind4:
		return Wind4
	case reach4:
		return Reach4
	case ron3:
		return Ron3
	case kan4:
		return Kan4
	}
	return ""
}

type Drawn struct {
	playerPointSpreads []*pointSpread
	reachPoint         int
	continuePoint      int
	end                *End
	drawnType          DrawnType
}

/*
PlayerPointSpreads
点数移動情報
*/
func (d *Drawn) PlayerPointSpreads() []*pointSpread {
	return d.playerPointSpreads
}

/*
ReachPoint
供託立直棒
*/
func (d *Drawn) ReachPoint() int {
	return d.reachPoint
}

/*
ContinuePoint
供託積み棒
*/
func (d *Drawn) ContinuePoint() int {
	return d.continuePoint
}

/*
End
終局情報
*/
func (d *Drawn) End() *End {
	return d.end
}

/*
DrawnType
流局理由
*/
func (d *Drawn) DrawnType() DrawnType {
	return d.drawnType
}

/*
NewDrawnGame
流局情報を取得します
*/
func NewDrawnGame(ba, sc, drawnType, owari string) (*Drawn, error) {
	playerPointSpread, err := createPlayerPointSpread(sc)
	if err != nil {
		return nil, Err(err)
	}
	continuePoint, reachPoint, err := createContinueAndReachPoint(ba, len(playerPointSpread))
	if err != nil {
		return nil, Err(err)
	}
	end, err := createEnd(owari)
	if err != nil {
		if !errors.Is(err, ArgumentEmptyError) {
			return nil, Err(err)
		}
	}
	dType := newDrawnType(drawnType)
	if dType == "" {
		return nil, Err(errors.New("type（流局理由）を指定してください。"))
	}
	return &Drawn{
		playerPointSpreads: playerPointSpread,
		continuePoint:      continuePoint,
		reachPoint:         reachPoint,
		end:                end,
		drawnType:          dType,
	}, nil
}

func Err(err error) error {
	return wrapError("NewDrawnGame", err)
}
