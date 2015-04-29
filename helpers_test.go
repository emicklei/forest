package forest

import "testing"

func TestDump(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/jsonarray"))
	Dump(t, r)
	// check if we can read it again
	Dump(t, r)
}
