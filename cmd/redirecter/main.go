package main

import (
	"log"
	"os"
	"shorter/internal/redirecter"
	"shorter/internal/storage"
	"shorter/pkg/logging"
)

func main() {
	urlStorage := storage.NewXmlStorage("/tmp/shorter.xml")
	loggers := &logging.Loggers{
		Info: log.New(os.Stderr, "INFO: ", log.Lshortfile | log.Ltime),
		Error: log.New(os.Stderr, "ERROR: ", log.Lshortfile | log.Ltime),
	}

	server := redirecter.NewRedirecter(9020, loggers, urlStorage)
	server.Start()

	err := <-server.Notify()
	loggers.Error.Printf("Redirecter server error: %v", err)

	if err = server.Stop(); err != nil {
		loggers.Error.Printf("Restapi server stopping error: %v", err)
	}
}
