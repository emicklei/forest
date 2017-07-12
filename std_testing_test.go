package forest

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestThatInfoLoggingIsPrinted(t *testing.T) {
	l := Logger{true, true, false}
	captureStdout(t, func() {
		l.Logf("%s", "logf")
	}, func(output string) {
		if output != "\tinfo : logf\n" {
			t.Errorf("different output than expected:%s", output)
		}
	})
}

func TestThatErrorLoggingIsPrinted(t *testing.T) {
	l := Logger{true, true, false}
	captureStdout(t, func() {
		l.Error("logf")
	}, func(output string) {
		if output != "\terror: [logf]\n" {
			t.Errorf("different output than expected:%q", output)
		}
	})
}

func TestThatFatalLoggingIsPrinted(t *testing.T) {
	l := Logger{true, true, false}
	captureStdout(t, func() {
		l.Fatal("logf")
	}, func(output string) {
		if output != "\tfatal: logf\n" {
			t.Errorf("different output than expected:%q", output)
		}
	})
}

func captureStdout(t *testing.T, work func(), callback func(string)) {
	old := os.Stdout // keep backup of the real stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	work()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	callback(out)
}
