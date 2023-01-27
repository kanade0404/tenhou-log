package database

import (
	"fmt"
	"github.com/kanade0404/tenhou-log/services/ent"
)

func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
