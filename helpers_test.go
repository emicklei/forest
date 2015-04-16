package rat

import "testing"

func TestDump(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/jsonarray"))
	Dump(t, r)
	// check if we can read it again
	Dump(t, r)
}
