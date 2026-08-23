package wx

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func Run(args []string) int {
	initColors(isTerminal(os.Stdout))

	fs := flag.NewFlagSet("wx", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	imperial := fs.Bool(
		"imperial",
		false,
		"Use imperial units (Fahrenheit, mph)",
	)

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "%sError:%s %v\n", red, reset, err)
		printUsage(os.Stderr)
		return 2
	}

	city := strings.TrimSpace(strings.Join(fs.Args(), " "))
	if city == "" {
		city = strings.TrimSpace(os.Getenv("WEATHER_CITY"))
	}

	if city == "" {
		printUsage(os.Stderr)
		return 2
	}

	client := NewAPIClient(&http.Client{
		Timeout: 10 * time.Second,
	})

	loc, err := client.GetLocation(city)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError:%s %v\n", red, reset, err)
		return 1
	}

	weather, err := client.GetWeather(
		loc.Latitude,
		loc.Longitude,
		*imperial,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%sError:%s %v\n", red, reset, err)
		return 1
	}

	showWeather(weather, loc, *imperial)
	return 0
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "%sUsage:%s wx [options] <city>\n", gray, reset)
	fmt.Fprintf(w, "%sExample:%s wx Moscow\n", gray, reset)
	fmt.Fprintf(w, "%sExample:%s wx -imperial \"San Francisco\"\n", gray, reset)
}

func isTerminal(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		return false
	}

	return info.Mode()&os.ModeCharDevice != 0
}
