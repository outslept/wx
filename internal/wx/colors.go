package wx

var (
	reset      = "\033[0m"
	bold       = "\033[1m"
	red        = "\033[31m"
	green      = "\033[32m"
	yellow     = "\033[33m"
	blue       = "\033[34m"
	white      = "\033[37m"
	gray       = "\033[90m"
	brightBlue = "\033[94m"
)

func initColors(isTTY bool) {
	if isTTY {
		return
	}

	reset = ""
	bold = ""
	red = ""
	green = ""
	yellow = ""
	blue = ""
	white = ""
	gray = ""
	brightBlue = ""
}
