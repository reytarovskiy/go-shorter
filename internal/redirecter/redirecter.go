package redirecter

import (
	"context"
	"net"
	"net/http"
	"shorter/internal/redirecter/handlers"
	"shorter/internal/storage"
	"shorter/pkg/logging"
	"strconv"
	"time"
)

type Redirecter struct {
	errors  chan error
	server  http.Server
	loggers *logging.Loggers
}

func NewRedirecter(port int, loggers *logging.Loggers, urlStorage storage.Storage) *Redirecter {
	redirectHandlers := handlers.NewRedirectHandlers(urlStorage, loggers)

	http.HandleFunc("/", redirectHandlers.Redirect)

	return &Redirecter{
		server:  http.Server{
			Addr: net.JoinHostPort("", strconv.Itoa(port)),
			Handler: nil,
		},
		errors:  make(chan error, 1),
		loggers: loggers,
	}
}

func (r *Redirecter) Start() {
	go func() {
		r.loggers.Info.Printf("Redirecter server started %s", r.server.Addr)
		r.errors <-r.server.ListenAndServe()
	}()
}

func (r *Redirecter) Notify() chan error {
	return r.errors
}

func (r *Redirecter) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := r.server.Shutdown(ctx)
	r.loggers.Info.Println("Redirecter server stopped")
	return err
}

