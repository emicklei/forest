package rat

import (
	"log"
	"os"
)

// T is the interface that this package is using from standard testing.T
type T interface {
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// LoggingT provides the reporting api of testing.T
var LoggingT = logger{}

type logger struct{}

func (l logger) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l logger) Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func (l logger) Fatalf(format string, args ...interface{}) {
	log.Printf("[FATAL] "+format, args...)
	os.Exit(1)
}
