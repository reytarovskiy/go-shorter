package main

import (
	"log"
	"os"
	"shorter/internal/restapi"
	"shorter/internal/storage"
	"shorter/pkg/logging"
	"shorter/pkg/shorter"
)

func main() {
	loggers := &logging.Loggers{
		Info: log.New(os.Stderr, "INFO: ", log.Lshortfile | log.Ltime),
		Error: log.New(os.Stderr, "ERROR: ", log.Lshortfile | log.Ltime),
	}

	urlStorage := storage.NewMemoryStorage()
	urlShorter := shorter.NewRandomShorter(7)

	rapi := restapi.New(9030, loggers, urlStorage, urlShorter)
	rapi.Start()

	err := <-rapi.Notify()
	loggers.Error.Printf("Restapi server error: %v", err)

	if err = rapi.Stop(); err != nil {
		loggers.Error.Printf("Restapi server stopping error: %v", err)
	}
}
