package init

import (
	"strconv"
	"strings"
)

type Point struct {
	PlayerIndex uint
	Point       uint
	_           struct{}
}
type Points struct {
	Points []Point
	_      struct{}
}

func (p *Points) Marshal() string {
	var results []string
	for _, point := range p.Points {
		results = append(results, strconv.Itoa(int(point.Point/100)))
	}
	return strings.Join(results, ",")
}

func (p *Points) Unmarshal(point string) error {
	var points []Point
	splitPoints := strings.Split(point, ",")

	for i, point := range splitPoints {
		pnt, err := strconv.ParseUint(point, 10, 32)
		if err != nil {
			return err
		}
		points = append(points, Point{
			PlayerIndex: uint(i + 1),
			Point:       uint(pnt * 100),
		})
	}
	p.Points = points
	return nil
}
