package netconn

import (
	"sync"

	"github.com/lab1/errors-logs/pkg/lib/errors"
)

func (s *svc) checkAny(addresses []string, timeout int) error {
	var (
		wg        sync.WaitGroup
		connected = false
	)

	wg.Add(len(addresses))

	for _, address := range addresses {
		go func(wg *sync.WaitGroup, address string) {
			err := s.checkOne(address, timeout)
			if err != nil {
				s.log.Debugf("connection to address FAILED: %s; err: %v", address, err)
			} else {
				connected = true
				s.log.Debugf("connection to address OK: %s", address)
			}

			wg.Done()
		}(&wg, address)
	}

	wg.Wait()

	if !connected {
		return errors.NoConnection
	}

	return nil
}
