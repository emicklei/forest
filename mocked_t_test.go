package forest

import "fmt"

type mockedT struct {
	errorMessage, logMessage, fatalMessage string
}

func (m *mockedT) Logf(format string, args ...interface{}) {
	m.logMessage = fmt.Sprintf(format, args...)
}
func (m *mockedT) Error(args ...interface{}) {
	m.errorMessage = fmt.Sprint(args...)
}
func (m *mockedT) Fatal(args ...interface{}) {
	m.fatalMessage = fmt.Sprint(args...)
}
func (m *mockedT) FailNow() {}
func (m *mockedT) Fail()    {}
