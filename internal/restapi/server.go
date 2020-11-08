package restapi

import (
	"context"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"shorter/internal/restapi/handlers"
	"shorter/internal/storage"
	"shorter/pkg/logging"
	"shorter/pkg/shorter"
	"strconv"
	"time"
)

type RestAPI struct {
	errors  chan error
	server  http.Server
	loggers *logging.Loggers
}

func New(port int, redirecterUrl string, loggers *logging.Loggers, storage storage.Storage, shorter shorter.Shorter) *RestAPI {
	shortHandlers := handlers.NewShorter(
		storage,
		loggers,
		shorter,
		redirecterUrl,
	)

	r := mux.NewRouter()
	r.HandleFunc("/short", shortHandlers.Short)

	return &RestAPI{
		server: http.Server{
			Addr: net.JoinHostPort("", strconv.Itoa(port)),
			Handler: r,
		},
		loggers: loggers,
		errors: make(chan error, 1),
	}
}

func (r *RestAPI) Start() {
	go func() {
		r.loggers.Info.Printf("Restapi server started %s", r.server.Addr)
		r.errors <-r.server.ListenAndServe()
	}()
}

func (r *RestAPI) Notify() chan error {
	return r.errors
}

func (r *RestAPI) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := r.server.Shutdown(ctx)
	r.loggers.Info.Println("Restapi server stopped")
	return err
}
