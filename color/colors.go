package color

import (
	"runtime"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Purple = "\033[35m"
)

func init() {
	// these colours are not supported on windows, so disable it on windows
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Yellow = ""
		Purple = ""
	}
}
