package tools

import "fmt"

func ErrParamIsRequired(name, value string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, value)
}
