package forest

import (
	"fmt"
	"os"
	"testing"
)

func TestDump(t *testing.T) {
	r := tsAPI.GET(t, NewConfig("/jsonarray"))
	Dump(t, r)
	// check if we can read it again
	Dump(t, r)
}

type skipper struct {
	*testing.T
	skipped bool
	reason  string
}

func (s *skipper) Skipf(f string, args ...interface{}) {
	s.skipped = true
	s.reason = fmt.Sprintf(f, args...)
}

func TestSkipUnless(t *testing.T) {
	os.Setenv("LABELS", "test,check")
	for _, each := range []struct {
		s      *skipper
		labels []string
		skip   bool
	}{
		{
			&skipper{t, false, ""},
			[]string{"any", "pass"},
			true,
		},
		{
			&skipper{t, false, ""},
			[]string{"test", "pass"},
			false,
		},
		{
			&skipper{t, false, ""},
			[]string{"any", "check"},
			false,
		},
		{
			&skipper{t, false, ""},
			[]string{},
			true,
		},
		{
			&skipper{t, false, ""},
			[]string{"check"},
			false,
		},
	} {
		SkipUnless(each.s, each.labels...)
		if each.s.skipped != each.skip {
			t.Errorf("unexpected skip")
		}
	}
}
