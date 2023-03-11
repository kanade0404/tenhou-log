package error

import "fmt"

func WrapError(packageName, funcName string, err error) error {
	return fmt.Errorf("error at %s.%s(): %w", packageName, funcName, err)
}
