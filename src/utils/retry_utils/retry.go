package retry_utils

import (
	"time"
)

func RetryBackOff(maxRetries int, do func() error) error {
	var retryError error
	for i := 1; i <= maxRetries; i++ {
		retryError := do()
		if retryError == nil {
			return nil
		}

		if i < maxRetries {
			waitTime := time.Microsecond * 500
			time.Sleep(waitTime)
		}
	}

	return retryError
}
