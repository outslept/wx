package wx

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	tempHot         = 30.0
	tempWarm        = 20.0
	tempMild        = 10.0
	tempFreezing    = 0.0
	windSectorSpan  = 45.0
	windSectorHalf  = 22.5
	windSectorCount = 8
)

func showWeather(out io.Writer, t theme, w *Weather, loc *Location, imperial bool) {
	icon := pickIcon(w.Current.WeatherCode, w.Current.IsDay == 1)
	now := time.Now()
	if w.Timezone != "" {
		if tz, err := time.LoadLocation(w.Timezone); err == nil {
			now = now.In(tz)
		}
	}
	_, _ = fmt.Fprintf(out, "\n  %s%s%s, %s%s\n  %s%s%s\n\n", t.bold, t.white, loc.Name, loc.Country, t.reset, t.faint, now.Format("Monday, January 2, 2006 - 15:04 MST"), t.reset)
	for _, line := range icon.ASCII {
		if !t.enabled {
			line = stripANSI(line)
		}
		_, _ = fmt.Fprintf(out, "   %s\n", line)
	}
	_, _ = fmt.Fprintf(out, "\n   %s%s%.0f°%s", tempColor(w.Current.Temperature, imperial, t), t.bold, w.Current.Temperature, t.reset)
	if math.Abs(w.Current.ApparentTemp-w.Current.Temperature) >= 1.0 {
		_, _ = fmt.Fprintf(out, " %s(feels like %.0f°)%s", t.faint, w.Current.ApparentTemp, t.reset)
	}
	_, _ = fmt.Fprintf(out, "\n   %s%s%s\n\n", t.faint, icon.Description, t.reset)
	windUnit := "km/h"
	if imperial {
		windUnit = "mph"
	}
	_, _ = fmt.Fprintf(out, "   %sHumidity%s %d%%   %sWind%s %.0f %s %s   %sRain%s %d%%", t.faint, t.reset, w.Current.Humidity, t.faint, t.reset, w.Current.WindSpeed, windUnit, windDirection(w.Current.WindDirection), t.faint, t.reset, w.Current.PrecipitationProbability)
	if len(w.Daily.TempMax) > 0 && len(w.Daily.TempMin) > 0 {
		_, _ = fmt.Fprintf(out, "   %sHi / Lo%s %.0f° / %.0f°", t.faint, t.reset, w.Daily.TempMax[0], w.Daily.TempMin[0])
	}
	if len(w.Daily.UVIndexMax) > 0 {
		_, _ = fmt.Fprintf(out, "   %sUV%s %.0f", t.faint, t.reset, w.Daily.UVIndexMax[0])
	}
	_, _ = fmt.Fprintln(out)
	if len(w.Daily.Sunrise) > 0 && len(w.Daily.Sunset) > 0 {
		sr, err1 := time.Parse("2006-01-02T15:04", w.Daily.Sunrise[0])
		ss, err2 := time.Parse("2006-01-02T15:04", w.Daily.Sunset[0])
		if err1 == nil && err2 == nil {
			_, _ = fmt.Fprintf(out, "\n   %sSunrise%s %s   %sSunset%s %s\n\n", t.faint, t.reset, sr.Format("15:04"), t.faint, t.reset, ss.Format("15:04"))
		}
	}
}
func tempColor(temp float64, imperial bool, t theme) string {
	if imperial {
		temp = (temp - 32) * 5 / 9
	}
	switch {
	case temp >= tempHot:
		return t.red
	case temp >= tempWarm:
		return t.yellow
	case temp >= tempMild:
		return t.green
	case temp >= tempFreezing:
		return t.brightBlue
	default:
		return t.blue
	}
}
func windDirection(degrees int) string {
	directions := [...]string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	return directions[int((float64(((degrees%360)+360)%360)+windSectorHalf)/windSectorSpan)%windSectorCount]
}
func pickIcon(code WeatherCode, isDay bool) WeatherIcon {
	if code == 0 && !isDay {
		return iconMap[100]
	}
	if icon, ok := iconMap[code]; ok {
		return icon
	}
	return unknownIcon
}
