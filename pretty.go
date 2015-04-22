package rat

import (
	"fmt"

	"github.com/wsxiaoys/terminal/color"
)

// FailMessagePrefix is used for printing failure messages.
var FailMessagePrefix = "\n:-( "

var ColorizeErrors = true

func prettyF(format string, args ...interface{}) string {
	plainFormatted := fmt.Sprintf(format, args...)
	if ColorizeErrors {
		return color.Sprintf("@{wR}%s%s", FailMessagePrefix, plainFormatted)
	}
	return FailMessagePrefix + plainFormatted
}
