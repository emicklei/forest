package rat

import "fmt"

type mockedT struct {
	errorMessage, logMessage, fatalMessage string
}

func (m *mockedT) Logf(format string, args ...interface{}) {
	m.logMessage = fmt.Sprintf(format, args...)
}
func (m *mockedT) Errorf(format string, args ...interface{}) {
	m.errorMessage = fmt.Sprintf(format, args...)
}
func (m *mockedT) Fatalf(format string, args ...interface{}) {
	m.fatalMessage = fmt.Sprintf(format, args...)
}
