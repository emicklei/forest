package rat

import (
	"fmt"
	"os"
)

// T is the interface that this package is using from standard testing.T
type T interface {
	Parallel()
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// TestingT provides a sub-api of testing.T
var TestingT = logging{true}

type logging struct {
	// this field exists for testing this package only
	doExit bool
}

func (f logging) Logf(format string, args ...interface{}) {
	fmt.Printf("\tinfo : "+format+"\n", args...)
}

func (f logging) Errorf(format string, args ...interface{}) {
	fmt.Printf("\terror: "+format+"\n", args...)
}

func (f logging) Fatalf(format string, args ...interface{}) {
	fmt.Printf("\tfatal: "+format+"\n", args...)
	if f.doExit {
		os.Exit(1)
	}
}

// Parallel is a no-op for this package.
func (f logging) Parallel() {}
