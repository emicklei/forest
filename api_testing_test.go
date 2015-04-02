package rat

import "testing"

func TestPost(t *testing.T) {
	r := tsApi.POST(t, NewConfig("/echo").Body("data").Header("ECHO", "ping"))
	ExpectString(t, r, func(m string) {
		if m != "data" {
			t.Errorf("expected data but got %v", m)
		}
	})
	ExpectHeader(t, r, "ECHO", "ping")
}
func TestDelete(t *testing.T) {
	r := tsApi.DELETE(t, NewConfig("/"))
	ExpectStatus(t, r, 204)
}

func TestPut(t *testing.T) {
	r := tsApi.PUT(t, NewConfig("/"))
	ExpectStatus(t, r, 204)
}

func TestPut404(t *testing.T) {
	r := tsApi.PUT(t, NewConfig("/404"))
	ExpectStatus(t, r, 404)
}

//func TestPutWithInvalidUrlIsCaptured(t *testing.T) {
//	captureStdout(t, func() {
//		tsApi.PUT(TestingT, NewConfig("#/#"))
//	}, func(out string) {
//		if !strings.Contains(out, "no Host in request URL") {
//			t.Errorf("different error output:[%s]", out)
//		}
//	})
//}