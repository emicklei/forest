package forest

import (
	"encoding/hex"
	"io/ioutil"
	"testing"
)

type car struct {
	Brand string
	Hp    int
}

func TestThatContentIsMarshalledToXml(t *testing.T) {
	conf := NewConfig("/")
	c := car{
		Brand: "Tesla",
		Hp:    500,
	}
	conf.Content(c, "application/xml")
	if conf.BodyReader == nil {
		t.Fatalf("expected body reader, got nil")
	}
	if conf.HeaderMap.Get("Content-Type") != "application/xml" {
		t.Fatalf("expected xml content-type")
	}
	data, _ := ioutil.ReadAll(conf.BodyReader)
	if string(data) != "<car><Brand>Tesla</Brand><Hp>500</Hp></car>" {
		t.Errorf("expected xml document, got:%s", string(data))
	}
}

func TestThatContentIsMarshalledToJson(t *testing.T) {
	conf := NewConfig("/")
	c := car{
		Brand: "Tesla",
		Hp:    500,
	}
	conf.Content(c, "application/json")
	if conf.BodyReader == nil {
		t.Fatalf("expected body reader, got nil")
	}
	if conf.HeaderMap.Get("Content-Type") != "application/json" {
		t.Fatalf("expected json content-type")
	}
	data, _ := ioutil.ReadAll(conf.BodyReader)
	if string(data) != `{"Brand":"Tesla","Hp":500}` {
		t.Errorf("expected json document, got:%s", string(data))
	}
}

func TestThatContentIsMarshalledToPlainText(t *testing.T) {
	conf := NewConfig("/")
	conf.Content("hello", "text/plain;charset=utf8")
	if conf.BodyReader == nil {
		t.Fatalf("expected body reader, got nil")
	}
	if conf.HeaderMap.Get("Content-Type") != "text/plain;charset=utf8" {
		t.Fatalf("expected plain content-type")
	}
	data, _ := ioutil.ReadAll(conf.BodyReader)
	if string(data) != "hello" {
		t.Errorf("expected plain document, got:%s", string(data))
	}
}

func TestThatContentCanBeBytes(t *testing.T) {
	conf := NewConfig("/")
	conf.Content([]byte{1, 2, 3, 4}, "application/octet-stream")
	if conf.BodyReader == nil {
		t.Fatalf("expected body reader, got nil")
	}
	if conf.HeaderMap.Get("Content-Type") != "application/octet-stream" {
		t.Fatalf("expected octet content-type")
	}
	data, _ := ioutil.ReadAll(conf.BodyReader)
	if data[0] != 1 || data[3] != 4 {
		t.Errorf("expected 1,2,3,4 bytes, got:%v", hex.Dump(data))
	}
}

func setXHeader(r *RequestConfig) {
	r.Header("X", "Y")
}

func TestThatCustomDoIsCalled(t *testing.T) {
	conf := NewConfig("/")
	conf.Do(setXHeader)
	if conf.HeaderMap.Get("X") != "Y" {
		t.Fail()
	}
}

func TestThatPathCanBeOverriden(t *testing.T) {
	conf := NewConfig("/a")
	conf.Path("/b")
	if conf.URI != "/b" {
		t.Errorf("got %v want %v", conf.URI, "/b")
	}
}

func TestThatQueryParametersCanBeAddedToTheUri(t *testing.T) {
	conf := NewConfig("/test")
	conf.Query("zoom", true)
	conf.Query("scale", 1)
	conf.Query("slash", "/")
	if conf.pathAndQuery() != "/test/scale=1&slash=%2F&zoom=true" {
		t.Errorf("got %v want %v", conf.pathAndQuery(), "/test/scale=1&slash=%2F&zoom=true")
	}
}

func TestThatPathParametersAndSubstitutedInTheUri(t *testing.T) {
	conf := NewConfig("")
	conf.Path("/{p1}/with/{p2}", "play", "/s:las:h/")
	pq := conf.pathAndQuery()
	want := "/play/with/%2Fs%3Alas%3Ah%2F"
	if pq != want {
		t.Errorf("got %s want %s", pq, want)
	}
}
