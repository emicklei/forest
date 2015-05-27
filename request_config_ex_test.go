package forest

import (
	"os"
	"time"
)

func ExampleRequestConfig() {
	var cfg *RequestConfig

	// set path template and header
	cfg = Path("/v1/assets/{id}", "artreyu").
		Header("Accept", "application/json")

	// set query parameters (the config will do escaping)
	cfg = NewConfig("/v1/assets").
		Query("lang", "en")

	// contents as is
	cfg = Path("/v1/assets").
		Body("some payload for POST or PUT")

	// content from file (io.Reader)
	cfg = Path("/v1/assets")
	f, _ := os.Open("payload.xml")
	cfg.BodyReader = f

	// content by marshalling (xml,json,plain text) your value
	cfg = NewConfig("/v1/assets").
		Content(time.Now(), "application/json")

}
