package client

import "fmt"

func wrapError(customMsg string, originalError error) error {
	return fmt.Errorf("%s : %v", customMsg, originalError)
}
