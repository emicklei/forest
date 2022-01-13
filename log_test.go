package forest

func logf_test(t T, stackOffset int, format string, args ...interface{}) {
	t.Logf(format, args...)
}
