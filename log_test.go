package forest

import "fmt"

func init() {
	fmt.Println("disable scanStackForFile for testing only")
	scanStackForFile = false
	logf_func = logf_test
}

func logf_test(t T, stackOffset int, format string, args ...interface{}) {
	t.Logf(format, args...)
}
