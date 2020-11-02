package handlers

import (
	"net/http"
	"shorter/internal/storage"
	"shorter/pkg/logging"
	"strings"
)

type RedirectHandlers struct {
	urlStorage storage.Storage
	loggers    *logging.Loggers
}

func NewRedirectHandlers(urlStorage storage.Storage, loggers *logging.Loggers) *RedirectHandlers {
	return &RedirectHandlers{
		urlStorage: urlStorage,
		loggers: loggers,
	}
}

func (r *RedirectHandlers) Redirect(w http.ResponseWriter, req *http.Request) {
	shorts := strings.Split(req.URL.Path, "/")
	if len(shorts) > 2 {
		w.WriteHeader(http.StatusBadRequest)
		r.loggers.Error.Printf("Bad short len: %d", len(shorts))
		return
	}

	url := r.urlStorage.Get(shorts[1])
	if url == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, req, *url, http.StatusTemporaryRedirect)
}
