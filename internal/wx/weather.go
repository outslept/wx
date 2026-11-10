package wx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	weatherAPI = "https://api.open-meteo.com/v1/forecast"
	geocodeAPI = "https://geocoding-api.open-meteo.com/v1/search"
	maxErrBody = 512
)

type APIClient struct {
	http                   *http.Client
	geocodeURL, weatherURL string
}

func NewAPIClient(client *http.Client) *APIClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &APIClient{http: client, geocodeURL: geocodeAPI, weatherURL: weatherAPI}
}
func (c *APIClient) GetLocation(cityName string) (*Location, error) {
	q := url.Values{"name": {cityName}, "count": {"1"}, "language": {"en"}, "format": {"json"}}
	var geo GeocodeResponse
	if err := c.getJSON("geocoding-api.open-meteo.com", c.geocodeURL+"?"+q.Encode(), &geo); err != nil {
		return nil, err
	}
	if len(geo.Results) == 0 {
		return nil, &CLIError{Status: "error", Reason: "not_found", Message: fmt.Sprintf("Couldn't find city %q.", sanitize(cityName)), Hint: "Try another city name", Command: "wx \"San Francisco\"", Code: 1}
	}
	loc := geo.Results[0]
	loc.Name, loc.Country = sanitize(loc.Name), sanitize(loc.Country)
	return &loc, nil
}
func (c *APIClient) GetWeather(lat, lon float64, imperial bool) (*Weather, error) {
	tempUnit, windUnit := "celsius", "kmh"
	if imperial {
		tempUnit, windUnit = "fahrenheit", "mph"
	}
	q := url.Values{"latitude": {fmt.Sprintf("%.5f", lat)}, "longitude": {fmt.Sprintf("%.5f", lon)}, "current": {"temperature_2m,relative_humidity_2m,apparent_temperature,precipitation_probability,wind_speed_10m,wind_direction_10m,weather_code,is_day"}, "daily": {"temperature_2m_max,temperature_2m_min,uv_index_max,sunrise,sunset"}, "timezone": {"auto"}, "forecast_days": {"1"}, "temperature_unit": {tempUnit}, "wind_speed_unit": {windUnit}}
	var w Weather
	if err := c.getJSON("api.open-meteo.com", c.weatherURL+"?"+q.Encode(), &w); err != nil {
		return nil, err
	}
	return &w, nil
}
func (c *APIClient) getJSON(host, reqURL string, dst any) error {
	resp, err := c.http.Get(reqURL)
	if err != nil {
		return &CLIError{Status: "error", Reason: "network_error", Message: fmt.Sprintf("Couldn't reach %s.", host), Hint: "Check your connection or see help", Command: "wx --help", Code: 1}
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrBody))
		msg := fmt.Sprintf("Failed to fetch data from %s (%s).", host, resp.Status)
		if d := sanitize(strings.TrimSpace(string(body))); d != "" {
			msg += " " + d
		}
		return &CLIError{Status: "error", Reason: "api_error", Message: msg, Code: 1}
	}
	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		return &CLIError{Status: "error", Reason: "api_error", Message: fmt.Sprintf("Failed to decode response from %s.", host), Code: 1}
	}
	return nil
}
