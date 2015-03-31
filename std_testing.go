package rat

import (
	"fmt"
	"os"
)

// T is the interface that this package is using from standard testing.T
type T interface {
	Fail()
	FailNow()
}

// FailingT provides the failing api of testing.T
var FailingT = fail{}

type fail struct{}

// Fail marks the function as having failed but continues execution.
func (f fail) Fail() {}

// FailNow marks the function as having failed but continues execution.
func (f fail) FailNow() {
	os.Exit(1)
}

// Logf print a log line without file location (which would be inside this package)
func Logf(format string, args ...interface{}) {
	print(fmt.Sprintf("\tinfo : "+format+"\n", args...))
}

// Errorf print a log line without file location (which would be inside this package)
func Errorf(format string, args ...interface{}) {
	print(fmt.Sprintf("\terror: "+format+"\n", args...))
}
