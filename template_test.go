package forest

import "testing"

func TestApplyTemplate(t *testing.T) {
	who := "ratatouille"
	s := ProcessTemplate(t, `{{.}}`, who)
	if s != who {
		t.Errorf("got %s but expected %q", s, who)
	}
}

func TestApplyTemplateFail(t *testing.T) {
	m := new(mockedT)
	who := "ratatouile"
	s := ProcessTemplate(m, `{{Missing}}`, who)
	if len(s) != 0 {
		t.Errorf("expected emptiness")
	}
	expected := `failed to parse:template: temporary:1: function "Missing" not defined`
	if m.logMessage != expected {
		t.Errorf("got %q expected %q", m.fatalMessage, expected)
	}
}
