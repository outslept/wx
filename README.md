# wx

Minimal terminal weather client powered by Open‑Meteo. Color ASCII icons, zero configuration.

# Usage

```
wx "<city>"
```

Examples:
```
wx Moscow
wx "San Francisco"
```

If no city is provided, the app will fall back to the `WEATHER_CITY` environment variable.

# Options

- `-imperial` : Display weather in Fahrenheit and mph instead of Celsius and km/h.

Examples:
```
wx Moscow -imperial
wx -imperial "San Francisco"
```

# Example output

```

  Moscow, Russia
  Monday, September 29, 2025 • 19:44 MSK

      \  /
    _ /"".-.
      \_(   ).
      /(___(__)


   7° (feels like 4°)
   Partly cloudy

   Humidity 61%   Wind 6 km/h NE   Rain 0%   Hi / Lo 10° / 4°   UV 2

   Sunrise 06:28   Sunset 18:10

```

# Notes

Data source: Open‑Meteo geocoding and forecast. No API keys. Internet connection required. Timezone is selected automatically. Colors are automatically disabled when piping output to a file or another process.
