package utils

import (
	"fmt"
)

// E wraps a prefix error message with an actual error text to return a new error
func E(prefix string, err error) error {
	return fmt.Errorf("%s: [%v]", prefix, err)
}
