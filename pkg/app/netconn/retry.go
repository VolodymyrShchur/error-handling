package netconn

import (
	"github.com/lab1/errors-logs/pkg/lib/retry"
)

func (s *svc) checkOneWithRetry(address string, timeout int) error {
	return retry.Retry(retryCount, retryInterval, func() error {
		s.log.Debugf("trying connect to: %s", address)
		return s.checkOne(address, timeout)
	})
}
