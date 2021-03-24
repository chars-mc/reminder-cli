package client

import "fmt"

// wrapError returns a custom error and the original error
func wrapError(customMsg string, originalError error) error {
	return fmt.Errorf("%s : %v", customMsg, originalError)
}
