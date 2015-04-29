package forest

import "io"

// closeableReader adds a no-op Close operation to a Reader
type closeableReader struct {
	io.Reader
}

func (r *closeableReader) Close() error {
	return nil
}
