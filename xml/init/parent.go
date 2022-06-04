package init

import (
	"fmt"
	"strconv"
)

type ParentPlayer struct {
	Index uint
}

func (p ParentPlayer) Marshal() string {
	return strconv.Itoa(int(p.Index - 1))
}

func (p *ParentPlayer) Unmarshal(index string) error {
	i, err := strconv.ParseUint(index, 10, 32)
	if err != nil {
		return err
	}
	if i < 0 || i > 3 {
		return fmt.Errorf("unexpected value %d. argument must be more than equeal 0 and less than equal 3", i)
	}
	p.Index = uint(i) + 1
	return nil
}
