package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shorter/internal/storage"
	"shorter/pkg/logging"
	"shorter/pkg/shorter"
)

type shortRequest struct {
	Url string `json:"url"`
}

type shortResponse struct {
	Short string `json:"short"`
	Url string `json:"url"`
}

type Shorter struct {
	storage storage.Storage
	loggers *logging.Loggers
	shorter shorter.Shorter
}

func NewShorter(storage storage.Storage, loggers *logging.Loggers, shorter shorter.Shorter) *Shorter {
	return &Shorter{
		storage: storage,
		loggers: loggers,
		shorter: shorter,
	}
}

func (s *Shorter) Short(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := shortRequest{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		s.loggers.Error.Printf("invalid request decode: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	short := s.shorter.Short(request.Url)
	s.storage.Store(short, request.Url)

	json.NewEncoder(w).Encode(&shortResponse{
		Short: short,
		Url: fmt.Sprintf("http://localhost/%s", short),
	})
}