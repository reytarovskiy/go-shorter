package logging

import "log"

type Loggers struct {
	Info *log.Logger
	Error *log.Logger
}
