package retry

import (
	"fmt"
	"time"

	"github.com/lab1/errors-logs/pkg/lib/errors"
)

func Retry(attempts int, sleep time.Duration, f func() error) (err error) {
	errb := errors.NewBundle()
	for i := 0; i < attempts; i++ {
		if i > 0 {
			time.Sleep(sleep)
			sleep *= 2
		}
		err = f()
		if err == nil {
			return nil
		}
		errb.Add(err)
	}
	errb.Add(fmt.Errorf("execution stopped after %d attempt", attempts))

	return errb.ErrorOrNil()
}
