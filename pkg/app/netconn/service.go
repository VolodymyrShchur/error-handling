package netconn

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/lab1/errors-logs/pkg/lib/errors"
	"github.com/lab1/errors-logs/pkg/lib/validation"
	"github.com/sirupsen/logrus"
)

const retryCount = 3
const retryInterval = 10 * time.Second

type Service interface {
	CheckSet(addresses []string, timeout int) error
}

func NewService(
	logger logrus.FieldLogger,
) Service {
	return &svc{
		log: logger,
	}
}

type svc struct {
	log logrus.FieldLogger
}

func (s *svc) checkOne(address string, timeout int) error {
	t := time.Duration(timeout) * time.Second

	u, err := url.Parse(address)
	if err != nil {
		return fmt.Errorf("parse address: %v", err)
	}

	switch u.Scheme {
	case "http", "https":
		var httpClient = &http.Client{
			Timeout: t,
		}

		resp, err := httpClient.Get(u.String())
		if err != nil {
			return fmt.Errorf("http get: %v", err)
		}
		defer resp.Body.Close()
	default:
		conn, err := net.DialTimeout(u.Scheme, u.Host, t)
		if err != nil {
			return fmt.Errorf("net dial: %v", err)
		}

		err = conn.Close()
		if err != nil {
			s.log.Debugf("failed to close connection: %v", err)
		}
	}

	return nil
}

func (s *svc) CheckSet(addresses []string, timeout int) error {
	errb := errors.NewBundle()
	validation.ValidateStringSlice(errb,
		validation.ColumnStringSlice{Name: "addresses", MaxLen: 3}, addresses)
	validation.ValidateInt(errb,
		validation.ColumnInt{Name: "timeout", MaxValue: 10}, timeout)

	if !errb.IsEmpty() {
		return errors.Validation.Wrap(errb.ErrorOrNil())
	}

	err := s.checkAny(addresses, timeout)
	if err != nil {
		return fmt.Errorf("check any: %w", err)
	}

	//err := s.checkAll(addresses, timeout)
	//if err != nil {
	//	return fmt.Errorf("check all: %w", err)
	//}

	return nil
}
