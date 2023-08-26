package utils

import (
	"time"
)

func DoWithTries(fn func() error, attemps int, duration time.Duration) (err error) {
	for attemps > 0 {
		if err = fn(); err != nil {
			time.Sleep(duration)
			attemps--

			continue
		}

		return nil
	}
	return nil
}
