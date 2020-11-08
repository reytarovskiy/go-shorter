package main

import (
	"flag"
	"log"
	"os"
	"shorter/internal/redirecter"
	"shorter/internal/storage"
	"shorter/pkg/logging"
)

var storagePath = flag.String("xml-path", "/tmp/shorter.xml", "Storage xml path")

func main() {
	flag.Parse()

	urlStorage := storage.NewXmlStorage(*storagePath)
	loggers := &logging.Loggers{
		Info:  log.New(os.Stderr, "INFO: ", log.Lshortfile|log.Ltime),
		Error: log.New(os.Stderr, "ERROR: ", log.Lshortfile|log.Ltime),
	}

	server := redirecter.NewRedirecter(9020, loggers, urlStorage)
	server.Start()

	err := <-server.Notify()
	loggers.Error.Printf("Redirecter server error: %v", err)

	if err = server.Stop(); err != nil {
		loggers.Error.Printf("Restapi server stopping error: %v", err)
	}
}
