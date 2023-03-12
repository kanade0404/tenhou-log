package custom_error

import "fmt"

func WrapError(packageName, funcName string, err error) error {
	return fmt.Errorf("custom_error at %s.%s(): %w", packageName, funcName, err)
}
