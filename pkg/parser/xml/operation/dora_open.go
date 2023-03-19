package operation

import (
	"github.com/kanade0404/tenhou-log/pkg/parser/xml/hai"
	"strconv"
)

/*
NewDoraOpen
開槓情報を作成する。
*/
func NewDoraOpen(haiID string, isRed bool) (hai.IHai, error) {
	i, err := strconv.Atoi(haiID)
	if err != nil {
		return nil, wrapErr(err)
	}
	h, err := hai.NewHai(i, isRed)
	if err != nil {
		return nil, wrapErr(err)
	}
	return h, nil
}

func wrapErr(err error) error {
	return wrapError("NewDoraOpen", err)
}
