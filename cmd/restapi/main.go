package main

import (
	"log"
	"os"
	"shorter/internal/restapi"
	"shorter/pkg/logging"
)

func main() {
	loggers := &logging.Loggers{
		Info: log.New(os.Stderr, "INFO: ", log.Lshortfile | log.Ltime),
		Error: log.New(os.Stderr, "ERROR: ", log.Lshortfile | log.Ltime),
	}

	rapi := restapi.New(9030, loggers)
	rapi.Start()

	err := <-rapi.Notify()
	loggers.Error.Printf("Restapi server error: %v", err)

	if err = rapi.Stop(); err != nil {
		loggers.Error.Printf("Restapi server stopping error: %v", err)
	}
}
