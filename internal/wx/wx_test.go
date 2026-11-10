package wx

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTempColorAndWind(t *testing.T) {
	th := newTheme(true, false)
	if got := tempColor(32, true, th); got != th.brightBlue {
		t.Fatalf("expected 32F brightBlue, got %q", got)
	}
	if got := tempColor(32, false, th); got != th.red {
		t.Fatalf("expected 32C red, got %q", got)
	}
	if windDirection(0) != "N" || windDirection(45) != "NE" || windDirection(-45) != "NW" {
		t.Fatalf("unexpected wind direction")
	}
}
func TestCLIContracts(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "search") {
			if r.URL.Query().Get("name") == "Missing" {
				_, _ = w.Write([]byte(`{"results":[]}`))
				return
			}
			_, _ = w.Write([]byte(`{"results":[{"name":"Moscow","country":"Russia","latitude":55.75,"longitude":37.61}]}`))
			return
		}
		_, _ = w.Write([]byte(`{"current":{"temperature_2m":32,"relative_humidity_2m":60,"apparent_temperature":28,"precipitation_probability":10,"wind_speed_10m":5,"wind_direction_10m":45,"weather_code":0,"is_day":1},"daily":{"temperature_2m_max":[35],"temperature_2m_min":[25],"uv_index_max":[2],"sunrise":["2025-09-29T06:28"],"sunset":["2025-09-29T18:10"]},"timezone":"UTC"}`))
	}))
	defer srv.Close()
	client := NewAPIClient(srv.Client())
	client.geocodeURL, client.weatherURL = srv.URL+"/search", srv.URL+"/forecast"
	var out, errBuf bytes.Buffer
	if code := runWith([]string{"--help"}, &out, &errBuf, false, false, client); code != 0 || errBuf.Len() != 0 || !strings.Contains(out.String(), "--imperial") {
		t.Fatalf("help failed")
	}
	t.Setenv("WEATHER_CITY", "Moscow")
	out.Reset()
	errBuf.Reset()
	if code := runWith([]string{"   "}, &out, &errBuf, false, false, client); code != 2 {
		t.Fatalf("expected exit 2, got %d", code)
	}
	out.Reset()
	errBuf.Reset()
	if code := runWith([]string{"Moscow", "--imperial"}, &out, &errBuf, false, false, client); code != 0 || strings.Contains(out.String(), "\033[") || !strings.Contains(out.String(), "mph") {
		t.Fatalf("piped imperial failed")
	}
	out.Reset()
	errBuf.Reset()
	if code := runWith([]string{"Moscow", "--json"}, &out, &errBuf, false, false, client); code != 0 || errBuf.Len() != 0 {
		t.Fatalf("json failed")
	}
	var report WeatherReport
	if err := json.Unmarshal(out.Bytes(), &report); err != nil || report.Location.Name != "Moscow" {
		t.Fatalf("invalid json: %v", err)
	}
	out.Reset()
	errBuf.Reset()
	if code := runWith([]string{"Missing", "--json"}, &out, &errBuf, false, false, client); code != 1 || out.Len() != 0 || !strings.Contains(errBuf.String(), `"reason":"not_found"`) {
		t.Fatalf("json error failed")
	}
}
