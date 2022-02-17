package forest

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"
)

type JUnitReport struct {
	TestSuites []JUnitTestSuite `xml:"testsuite"`
}

type JUnitTestSuite struct {
	Tests      int             `xml:"tests,attr"`
	Failures   int             `xml:"failures,attr"`
	Time       float64         `xml:"time,attr"`
	Name       string          `xml:"name,attr"`
	Timestamp  time.Time       `xml:"timestamp,attr"`
	Properties []JUnitProperty `xml:"properties"`
	TestCases  []JUnitTestCase `xml:"testcase"`
}

type JUnitProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type JUnitTestCase struct {
	Classname string                `xml:"classname,attr"`
	Name      string                `xml:"name,attr"`
	Time      float64               `xml:"time,attr"`
	Failure   *JUnitTestCaseFailure `xml:"failure"`
	Skipped   *JUnitTestCaseSkipped `xml:"skipped"`
}

type JUnitTestCaseFailure struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
}

type JUnitTestCaseSkipped struct {
	Message string `xml:"message,attr"`
}

func ReadJUnitReport(filename string) (r JUnitReport, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("failed to read JUnit report:%v", err)
		return
	}
	err = xml.Unmarshal(data, &r)
	if err != nil {
		err = fmt.Errorf("failed to decode JUnit report:%v", err)
	}
	return
}
