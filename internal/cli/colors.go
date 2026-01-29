package cli

import "os"

// ANSI escape codes. Respect NO_COLOR (https://no-color.org/).
var (
	reset   = "\033[0m"
	bold    = "\033[1m"
	dim     = "\033[2m"
	cyan    = "\033[36m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	magenta = "\033[35m"
	blue    = "\033[34m"
	white   = "\033[37m"
)

func init() {
	if os.Getenv("NO_COLOR") != "" {
		reset = ""
		bold = ""
		dim = ""
		cyan = ""
		green = ""
		yellow = ""
		magenta = ""
		blue = ""
		white = ""
	}
}

func title(s string) string    { return bold + cyan + s + reset }
func section(s string) string  { return bold + yellow + s + reset }
func cmd(s string) string      { return green + s + reset }
func flagStyle(s string) string { return magenta + s + reset }
func dimmed(s string) string   { return dim + s + reset }
func accent(s string) string   { return cyan + s + reset }
