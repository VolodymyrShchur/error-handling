package netconn

import (
	"github.com/lab1/errors-logs/pkg/lib/errors"
	"golang.org/x/sync/errgroup"
)

func (s *svc) checkAll(addresses []string, timeout int) error {
	eg := errgroup.Group{}

	for _, address := range addresses {
		address := address // https://golang.org/doc/faq#closures_and_goroutines
		eg.Go(func() error {
			return s.checkOneWithRetry(address, timeout)
		})
	}

	err := eg.Wait()
	if err != nil {
		return errors.NoConnection.Wrap(err)
	}

	return nil
}
