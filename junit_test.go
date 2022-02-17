package forest

import (
	"testing"
)

func TestReadJunitReport(t *testing.T) {
	r, _ := ReadJUnitReport("junit-example.xml")
	t.Log(len(r.TestSuites))
	tc := 0
	for _, each := range r.TestSuites {
		//t.Log("props:", len(each.Properties))
		//t.Log("cases:", len(each.TestCases))
		for range each.TestCases {
			//t.Log(other.Name)
			tc++
		}
	}
	t.Log(tc)
}
