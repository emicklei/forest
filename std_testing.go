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
var TestingT = fail{}

type fail struct{}

// Fail marks the function as having failed but continues execution.
func (f fail) Fail() {}

// FailNow marks the function as having failed but continues execution.
func (f fail) FailNow() {
	os.Exit(1)
}

func (f fail) Logf(format string, args ...interface{}) {
	print(fmt.Sprintf("\tinfo : "+format+"\n", args...))
}

func (f fail) Errorf(format string, args ...interface{}) {
	print(fmt.Sprintf("\terror: "+format+"\n", args...))
}

func (f fail) Fatalf(format string, args ...interface{}) {
	print(fmt.Sprintf("\tfatal: "+format+"\n", args...))
	os.Exit(1)
}

// Parallel is a no-op for rat.
func (f fail) Parallel() {}
