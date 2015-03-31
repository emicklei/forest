package rat

import "testing"

func TestDump(t *testing.T) {
	r := tsApi.GET(t, NewConfig("/jsonarray"))
	Dump(t, r)
}
