# wx

Minimal terminal weather client powered by Open-Meteo. Color ASCII icons, zero configuration.

## Usage

```text
wx <city> [options]
```

Examples:
```text
wx Moscow
wx "San Francisco" --imperial
wx London --json
```

If no `<city>` is provided, `wx` reads the `WEATHER_CITY` environment variable.

## Options

- `-h, --help` : Show help and exit 0.
- `--imperial` : Display weather in Fahrenheit and mph instead of Celsius and km/h.
- `--json` : Output weather data and errors as JSON.
- `--no-color` : Disable ANSI color output (also respects `NO_COLOR`, `FORCE_COLOR`, and `TERM=dumb`).
- `-v, --version` : Print the CLI version.
