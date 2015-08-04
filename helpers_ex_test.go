package forest

import "testing"

func ExampleSkipUnless() {
	var t *testing.T

	// t implements skippeable (has method Skipf)
	SkipUnless(t, "nightly")
	// code below is executed only if the environment variable LABELS contains `nightly`

	// You run the `nightly` tests like this:
	//
	// LABELS=nightly go test -v
}
