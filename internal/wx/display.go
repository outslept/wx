package wx

import (
	"fmt"
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

func showWeather(w *Weather, loc *Location, imperial bool) {
	icon := pickIcon(w.Current.WeatherCode, w.Current.IsDay == 1)

	printHeader(loc, w.Timezone)
	printIcon(icon)
	printTemperature(w)
	printCondition(icon.Description)
	printMetrics(w, imperial)
	printSunTimes(w)
}

func printHeader(loc *Location, timezone string) {
	fmt.Println()

	now := time.Now()

	if timezone != "" {
		if location, err := time.LoadLocation(timezone); err == nil {
			now = now.In(location)
		}
	}

	fmt.Printf(
		"  %s%s%s, %s%s\n",
		bold,
		white,
		loc.Name,
		loc.Country,
		reset,
	)

	fmt.Printf(
		"  %s%s%s\n\n",
		gray,
		now.Format("Monday, January 2, 2006 • 15:04 MST"),
		reset,
	)
}

func printIcon(icon WeatherIcon) {
	for _, line := range icon.ASCII {
		fmt.Printf("   %s\n", line)
	}

	fmt.Println()
}

func printTemperature(w *Weather) {
	color := tempColor(w.Current.Temperature)

	fmt.Printf(
		"   %s%s%.0f°%s",
		color,
		bold,
		w.Current.Temperature,
		reset,
	)

	if math.Abs(w.Current.ApparentTemp-w.Current.Temperature) >= 1.0 {
		fmt.Printf(
			" %s(feels like %.0f°)%s",
			gray,
			w.Current.ApparentTemp,
			reset,
		)
	}

	fmt.Println()
}

func printCondition(description string) {
	fmt.Printf(
		"   %s%s%s\n\n",
		gray,
		description,
		reset,
	)
}

func printMetrics(w *Weather, imperial bool) {
	windUnit := "km/h"
	if imperial {
		windUnit = "mph"
	}

	fmt.Printf(
		"   %sHumidity%s %d%%",
		gray,
		reset,
		w.Current.Humidity,
	)

	fmt.Printf(
		"   %sWind%s %.0f %s %s",
		gray,
		reset,
		w.Current.WindSpeed,
		windUnit,
		windDirection(w.Current.WindDirection),
	)

	fmt.Printf(
		"   %sRain%s %d%%",
		gray,
		reset,
		w.Current.PrecipitationProbability,
	)

	if len(w.Daily.TempMax) > 0 && len(w.Daily.TempMin) > 0 {
		fmt.Printf(
			"   %sHi / Lo%s %.0f° / %.0f°",
			gray,
			reset,
			w.Daily.TempMax[0],
			w.Daily.TempMin[0],
		)
	}

	if len(w.Daily.UVIndexMax) > 0 {
		fmt.Printf(
			"   %sUV%s %.0f",
			gray,
			reset,
			w.Daily.UVIndexMax[0],
		)
	}

	fmt.Println()
}

func printSunTimes(w *Weather) {
	if len(w.Daily.Sunrise) == 0 || len(w.Daily.Sunset) == 0 {
		return
	}

	location := time.Local

	if w.Timezone != "" {
		if loaded, err := time.LoadLocation(w.Timezone); err == nil {
			location = loaded
		}
	}

	const layout = "2006-01-02T15:04"

	sunrise, err := time.ParseInLocation(
		layout,
		w.Daily.Sunrise[0],
		location,
	)
	if err != nil {
		return
	}

	sunset, err := time.ParseInLocation(
		layout,
		w.Daily.Sunset[0],
		location,
	)
	if err != nil {
		return
	}

	fmt.Printf(
		"\n   %sSunrise%s %s",
		gray,
		reset,
		sunrise.Format("15:04"),
	)

	fmt.Printf(
		"   %sSunset%s %s\n\n",
		gray,
		reset,
		sunset.Format("15:04"),
	)
}

func tempColor(temp float64) string {
	switch {
	case temp >= tempHot:
		return red
	case temp >= tempWarm:
		return yellow
	case temp >= tempMild:
		return green
	case temp >= tempFreezing:
		return brightBlue
	default:
		return blue
	}
}

func windDirection(degrees int) string {
	directions := [...]string{
		"N",
		"NE",
		"E",
		"SE",
		"S",
		"SW",
		"W",
		"NW",
	}

	normalized := ((degrees % 360) + 360) % 360

	index := int(
		(float64(normalized) + windSectorHalf) / windSectorSpan,
	) % windSectorCount

	return directions[index]
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
