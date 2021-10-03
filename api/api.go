package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/lab1/errors-logs/pkg/app/netconn"
	"github.com/lab1/errors-logs/pkg/lib/httpserver"
	"github.com/sirupsen/logrus"
)

type config struct {
	port int
}

type API struct {
	conf config

	server *httpserver.Server
	logger logrus.FieldLogger

	connService netconn.Service
}

type Listener interface {
	Listen()
}

// nolint:funlen // allow many handlers on init.
func Init(logger logrus.FieldLogger, port int) *API {
	connService := netconn.NewService(logger.WithField("service", "netconn"))

	a := &API{
		conf: config{
			port: port,
		},
		server:      httpserver.New(),
		logger:      logger,
		connService: connService,
	}
	// Middleware
	a.server.AddMiddleware(a.LogMiddleware)

	a.server.HandlePOST("/check-internet", a.checkInternetHandler)

	return a
}

func (a *API) Listen(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		<-ctx.Done()

		if err := a.server.Shutdown(context.Background()); err != nil {
			a.logger.Error(err, "server shutdown")
		}

		close(done)
	}()

	addr := fmt.Sprintf(":%d", a.conf.port)

	a.logger.WithField("addr", addr).Info("server listen")

	err := a.server.Listen(addr)
	if err != nil {
		a.logger.Error(err, "server start")
		close(done)
	}

	<-done
}

// LogMiddleware provides access log for multi requests.
func (a *API) LogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)
		a.logger.
			WithField("method", r.Method).
			WithField("url", r.URL.String()).
			WithField("duration", time.Since(start)).
			Debug("access")
	}
}
