package operation

import (
	"errors"
	"github.com/kanade0404/tenhou-log/pkg/parser/helper/custom_error"
)

var ArgumentEmptyError = errors.New("argument is empty")

func wrapError(funcName string, err error) error {
	return custom_error.WrapError("operation", funcName, err)
}
