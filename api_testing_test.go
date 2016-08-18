package forest

import (
	"net/http"
	"testing"
)

func TestPost(t *testing.T) {
	r := tsAPI.POST(t, Path("/echo").Body("data").Header("ECHO", "ping"))
	ExpectString(t, r, func(m string) {
		if m != "data" {
			t.Errorf("expected data but got %v", m)
		}
	})
	ExpectHeader(t, r, "ECHO", "ping")
}
func TestDelete(t *testing.T) {
	r := tsAPI.DELETE(t, Path("/"))
	ExpectStatus(t, r, 204)
}

func TestPut(t *testing.T) {
	r := tsAPI.PUT(t, Path("/"))
	ExpectStatus(t, r, 204)
}

func TestPatch(t *testing.T) {
	r := tsAPI.PATCH(t, Path("/"))
	ExpectStatus(t, r, 204)
}

func TestPut404(t *testing.T) {
	r := tsAPI.PUT(t, Path("/{code}", 404))
	ExpectStatus(t, r, 404)
}

func TestPut404UsingDo(t *testing.T) {
	r, err := tsAPI.Do("PUT", Path("/{code}", 404))
	if err != nil {
		t.Fail()
	}
	ExpectStatus(t, r, 404)
}

func TestDoWithInvalidUrl(t *testing.T) {
	_, err := tsAPI.Do("HEAD", Path("::"))
	if err == nil {
		t.Fail()
	}
}

func TestBasicAuth(t *testing.T) {
	cfg := NewConfig("/").BasicAuth("a", "b")
	req, _ := http.NewRequest("GET", "/", nil)
	setBasicAuth(cfg, req)
	if len(req.Header.Get("Authorization")) == 0 {
		t.Errorf("expected Authorization header")
	}
}
