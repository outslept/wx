package wx

type WeatherCode int

type Location struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeocodeResponse struct {
	Results []Location `json:"results"`
}

type Weather struct {
	Current struct {
		Temperature              float64     `json:"temperature_2m"`
		Humidity                 int         `json:"relative_humidity_2m"`
		ApparentTemp             float64     `json:"apparent_temperature"`
		PrecipitationProbability int         `json:"precipitation_probability"`
		WindSpeed                float64     `json:"wind_speed_10m"`
		WindDirection            int         `json:"wind_direction_10m"`
		WeatherCode              WeatherCode `json:"weather_code"`
		IsDay                    int         `json:"is_day"`
	} `json:"current"`

	Daily struct {
		TempMax    []float64 `json:"temperature_2m_max"`
		TempMin    []float64 `json:"temperature_2m_min"`
		UVIndexMax []float64 `json:"uv_index_max"`
		Sunrise    []string  `json:"sunrise"`
		Sunset     []string  `json:"sunset"`
	} `json:"daily"`

	Timezone             string `json:"timezone"`
	TimezoneAbbreviation string `json:"timezone_abbreviation"`
}

type WeatherIcon struct {
	Description string
	ASCII       []string
}
