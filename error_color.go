package rat

import (
	"fmt"

	"github.com/wsxiaoys/terminal/color"
)

// FailMessagePrefix is used for printing failure messages.
var FailMessagePrefix = "\n:-( "

// ErrorColorSyntaxCode requires the syntax defined on https://github.com/wsxiaoys/terminal/blob/master/color/color.go
// Set to an empty string to disable coloring
var ErrorColorSyntaxCode = "@{wR}"

// FatalColorSyntaxCode requires the syntax defined on https://github.com/wsxiaoys/terminal/blob/master/color/color.go
// Set to an empty string to disable coloring
var FatalColorSyntaxCode = "@{wR}"

func serrorf(format string, args ...interface{}) string {
	return Scolorf(ErrorColorSyntaxCode, format, args...)
}

func sfatalf(format string, args ...interface{}) string {
	return Scolorf(FatalColorSyntaxCode, format, args...)
}

// Scolorf returns a string colorized for terminal output using the syntaxCode (unless that's empty).
// Requires the syntax defined on https://github.com/wsxiaoys/terminal/blob/master/color/color.go
func Scolorf(syntaxCode string, format string, args ...interface{}) string {
	plainFormatted := fmt.Sprintf(format, args...)
	if len(syntaxCode) > 0 {
		// cannot pass the code as a string param
		return color.Sprintf(syntaxCode+"%s%s", FailMessagePrefix, plainFormatted)
	}
	return FailMessagePrefix + plainFormatted
}
