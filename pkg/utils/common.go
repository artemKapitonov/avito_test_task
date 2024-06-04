package utils

import (
	"time"
)

// DoWithTries is doing function with tries for postgres connection.
func DoWithTries(fn func() error, attempts int, duration time.Duration) error {
	var err error
	for ; attempts > 0; attempts-- {
		if err = fn(); err != nil {
			time.Sleep(duration)

			continue
		}


		return nil
	}

	return err
}
