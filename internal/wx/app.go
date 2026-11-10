package wx

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const version = "1.0.0"

func Run(args []string) int {
	return runWith(args, os.Stdout, os.Stderr, isTerminal(os.Stdout), isTerminal(os.Stderr), NewAPIClient(&http.Client{Timeout: 10 * time.Second}))
}
func runWith(args []string, stdout, stderr io.Writer, outTTY, errTTY bool, client *APIClient) int {
	fs := flag.NewFlagSet("wx", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	imperial, jsonOut, noColor, showVer := fs.Bool("imperial", false, ""), fs.Bool("json", false, ""), fs.Bool("no-color", false, ""), fs.Bool("version", false, "")
	fs.BoolVar(showVer, "v", false, "")
	reordered, showHelp := reorderArgs(args)
	err := fs.Parse(reordered)
	outTheme, errTheme := newTheme(outTTY, *noColor), newTheme(errTTY, *noColor)
	if showHelp || errors.Is(err, flag.ErrHelp) {
		printUsage(stdout)
		return 0
	}
	if err != nil {
		writeError(stderr, errTheme, *jsonOut, &CLIError{Status: "error", Reason: "unknown_option", Message: fmt.Sprintf("Can't parse options: %v.", err), Hint: "See all options", Command: "wx --help", Code: 2})
		return 2
	}
	if *showVer {
		if *jsonOut {
			writeJSON(stdout, outTTY, map[string]string{"version": version})
		} else {
			_, _ = fmt.Fprintln(stdout, version)
		}
		return 0
	}
	pos := fs.Args()
	city := strings.TrimSpace(strings.Join(pos, " "))
	if len(pos) > 0 && city == "" {
		writeError(stderr, errTheme, *jsonOut, &CLIError{Status: "error", Reason: "invalid_arguments", Message: "Can't look up weather for an empty city name.", Hint: "Pass a city name", Command: "wx Moscow", Code: 2})
		return 2
	}
	if len(pos) == 0 {
		city = strings.TrimSpace(os.Getenv("WEATHER_CITY"))
	}
	if city == "" {
		writeError(stderr, errTheme, *jsonOut, &CLIError{Status: "action_required", Reason: "missing_city", Message: "Can't look up weather without a city.", Hint: "Pass a city name", Command: "wx <CITY>", Code: 2})
		return 2
	}
	loc, err := client.GetLocation(city)
	if err != nil {
		return handleErr(stderr, errTheme, *jsonOut, err)
	}
	weather, err := client.GetWeather(loc.Latitude, loc.Longitude, *imperial)
	if err != nil {
		return handleErr(stderr, errTheme, *jsonOut, err)
	}
	if *jsonOut {
		units := "metric"
		if *imperial {
			units = "imperial"
		}
		writeJSON(stdout, outTTY, WeatherReport{Location: *loc, Weather: *weather, Condition: pickIcon(weather.Current.WeatherCode, weather.Current.IsDay == 1).Description, Units: units})
		return 0
	}
	showWeather(stdout, outTheme, weather, loc, *imperial)
	return 0
}
func reorderArgs(args []string) ([]string, bool) {
	var flags, pos []string
	for i, a := range args {
		if a == "--" {
			pos = append(pos, args[i+1:]...)
			break
		}
		if a == "help" && i == 0 {
			return nil, true
		}
		if strings.HasPrefix(a, "-") && a != "-" {
			flags = append(flags, a)
		} else {
			pos = append(pos, a)
		}
	}
	return append(flags, pos...), false
}
func handleErr(stderr io.Writer, t theme, jsonOut bool, err error) int {
	var ce *CLIError
	if !errors.As(err, &ce) {
		ce = &CLIError{Status: "error", Reason: "internal_error", Message: err.Error(), Code: 1}
	}
	writeError(stderr, t, jsonOut, ce)
	return ce.Code
}
func writeError(w io.Writer, t theme, jsonOut bool, ce *CLIError) {
	if jsonOut {
		b, _ := json.Marshal(ce)
		_, _ = fmt.Fprintln(w, string(b))
		return
	}
	prefix := "Error:"
	if t.enabled {
		prefix = t.red + "✗" + t.reset
	}
	_, _ = fmt.Fprintf(w, "%s %s\n", prefix, ce.Message)
	if ce.Command != "" {
		_, _ = fmt.Fprintf(w, "\n  %s:\n    %s\n", ce.Hint, ce.Command)
	}
}
func writeJSON(w io.Writer, pretty bool, v any) {
	b, _ := json.Marshal(v)
	if pretty {
		b, _ = json.MarshalIndent(v, "", "  ")
	}
	_, _ = fmt.Fprintln(w, string(b))
}
func printUsage(w io.Writer) {
	_, _ = fmt.Fprint(w, "wx <city> [options]\nShow current weather for a city\n\nArguments:\n  <city>        City name, such as Moscow or \"San Francisco\"\n\nOptions:\n  -h, --help    Show help for wx\n  --imperial    Use imperial units (Fahrenheit, mph)\n  --json        Output the result as JSON\n  --no-color    Disable color output\n  -v, --version Print the CLI version\n\nEnvironment:\n  WEATHER_CITY  Default city when <city> is omitted\n\nExamples:\n  Show current weather in metric units\n    wx Moscow\n\n  Show current weather in imperial units\n    wx \"San Francisco\" --imperial\n\n  Output weather as JSON\n    wx London --json\n")
}
func isTerminal(file *os.File) bool {
	info, err := file.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice != 0
}
