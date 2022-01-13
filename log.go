package forest

// Logf adds the actual file:line information to the log message
func Logf(t T, format string, args ...interface{}) {
	t.Helper()
	t.Logf(format, args...)
	//logf_func(t, noStackOffset, "\n"+format, args...)
}

func logfatal(t T, format string, args ...interface{}) {
	t.Helper()
	t.Logf(format, args...)
	t.FailNow()
}

func logerror(t T, format string, args ...interface{}) {
	t.Helper()
	t.Logf(format, args...)
	t.Fail()
}
