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
	http *http.Client
}

func NewAPIClient(client *http.Client) *APIClient {
	if client == nil {
		client = http.DefaultClient
	}

	return &APIClient{
		http: client,
	}
}

func (c *APIClient) GetLocation(cityName string) (*Location, error) {
	params := url.Values{}
	params.Set("name", cityName)
	params.Set("count", "1")
	params.Set("language", "en")
	params.Set("format", "json")

	requestURL := geocodeAPI + "?" + params.Encode()

	resp, err := c.http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("geocoding request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, responseError("geocoding", resp)
	}

	var geo GeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return nil, fmt.Errorf("decode geocoding response: %w", err)
	}

	if len(geo.Results) == 0 {
		return nil, fmt.Errorf("city %q not found", cityName)
	}

	return &geo.Results[0], nil
}

func (c *APIClient) GetWeather(
	lat float64,
	lon float64,
	imperial bool,
) (*Weather, error) {
	tempUnit := "celsius"
	windUnit := "kmh"

	if imperial {
		tempUnit = "fahrenheit"
		windUnit = "mph"
	}

	params := url.Values{}
	params.Set("latitude", fmt.Sprintf("%.5f", lat))
	params.Set("longitude", fmt.Sprintf("%.5f", lon))
	params.Set(
		"current",
		"temperature_2m,relative_humidity_2m,apparent_temperature,"+
			"precipitation_probability,wind_speed_10m,wind_direction_10m,"+
			"weather_code,is_day",
	)
	params.Set(
		"daily",
		"temperature_2m_max,temperature_2m_min,uv_index_max,sunrise,sunset",
	)
	params.Set("timezone", "auto")
	params.Set("forecast_days", "1")
	params.Set("temperature_unit", tempUnit)
	params.Set("wind_speed_unit", windUnit)

	requestURL := weatherAPI + "?" + params.Encode()

	resp, err := c.http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("weather request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, responseError("weather api", resp)
	}

	var weather Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, fmt.Errorf("decode weather response: %w", err)
	}

	return &weather, nil
}

func responseError(prefix string, resp *http.Response) error {
	body, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrBody))
	message := strings.TrimSpace(string(body))

	if message == "" {
		return fmt.Errorf("%s: %s", prefix, resp.Status)
	}

	return fmt.Errorf("%s: %s: %s", prefix, resp.Status, message)
}
