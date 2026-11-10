package wx

import (
	"os"
	"regexp"
)

var ansiRe = regexp.MustCompile(`\x1b\[[0-9;]*m`)

type theme struct {
	enabled                                                         bool
	reset, bold, faint, red, green, yellow, blue, white, brightBlue string
}

func newTheme(isTTY, noColor bool) theme {
	enabled := isTTY && os.Getenv("TERM") != "dumb"
	if fc := os.Getenv("FORCE_COLOR"); fc != "" {
		enabled = fc != "0"
	}
	if os.Getenv("NO_COLOR") != "" || noColor || !enabled {
		return theme{}
	}
	return theme{true, "\033[0m", "\033[1m", "\033[2m", "\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[37m", "\033[94m"}
}
func stripANSI(s string) string { return ansiRe.ReplaceAllString(s, "") }
func sanitize(s string) string {
	var out []rune
	for _, r := range stripANSI(s) {
		if r >= 0x20 && r != 0x7f && (r < 0x80 || r > 0x9f) {
			out = append(out, r)
		}
	}
	return string(out)
}
