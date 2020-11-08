package main

import (
	"flag"
	"log"
	"os"
	"shorter/internal/restapi"
	"shorter/internal/storage"
	"shorter/pkg/logging"
	"shorter/pkg/shorter"
)

var storagePath = flag.String("xml-path", "/tmp/shorter.xml", "Storage xml path")
var redirecterUrl = flag.String("redirecter-url", "http://localhost:9020", "Redirecter domain")

func main() {
	flag.Parse()

	loggers := &logging.Loggers{
		Info:  log.New(os.Stderr, "INFO: ", log.Lshortfile|log.Ltime),
		Error: log.New(os.Stderr, "ERROR: ", log.Lshortfile|log.Ltime),
	}

	urlStorage := storage.NewXmlStorage(*storagePath)
	urlShorter := shorter.NewRandomShorter(7)

	rapi := restapi.New(9030, *redirecterUrl, loggers, urlStorage, urlShorter)
	rapi.Start()

	err := <-rapi.Notify()
	loggers.Error.Printf("Restapi server error: %v", err)

	if err = rapi.Stop(); err != nil {
		loggers.Error.Printf("Restapi server stopping error: %v", err)
	}
}
