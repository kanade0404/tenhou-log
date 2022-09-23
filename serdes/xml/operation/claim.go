package operation

import (
	"github.com/kanade0404/tenhou-log/serdes/xml"
	"strconv"
)

type Claim struct {
	UserIndex xml.PlayerIndex
}

func ClaimFrom(claim uint64) xml.PlayerIndex {
	switch claim {
	case 0:
		return 0
	case 1, 2, 3:
		return xml.PlayerIndex(claim + 1)
	default:
		return 0
	}
}

func (c *Claim) HandHais(m string) error {
	mNum, err := strconv.ParseUint(m, 10, 16)
	if err != nil {
		return err
	}
	ClaimFrom(mNum & 0x0003)
}
