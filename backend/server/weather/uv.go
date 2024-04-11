package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	// "github.com/joho/godotenv"
)

// UVData struct to hold the UV index response
type UVData struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	DateIso string  `json:"date_iso"`
	Date    int64   `json:"date"`
	Value   float64 `json:"value"`
}

// PrefectureCapitalCoords holds the latitude and longitude.
type PrefectureCapitalCoords struct {
	Lat float64
	Lon float64
}

// prefectureCapitalsCoordsMap maps prefecture names to their capital's coordinates.
var prefectureCapitalsCoordsMap = map[string]PrefectureCapitalCoords{
	"Hokkaido":    {Lat: 43.0642, Lon: 141.3469}, // Sapporo
	"Aomori":      {Lat: 40.8244, Lon: 140.7400}, // Aomori
	"Iwate":       {Lat: 39.7036, Lon: 141.1525}, // Morioka
	"Miyagi":      {Lat: 38.2688, Lon: 140.8719}, // Sendai
	"Akita":       {Lat: 39.7186, Lon: 140.1024}, // Akita
	"Yamagata":    {Lat: 38.2404, Lon: 140.3633}, // Yamagata
	"Fukushima":   {Lat: 37.7503, Lon: 140.4675}, // Fukushima
	"Ibaraki":     {Lat: 36.3418, Lon: 140.4468}, // Mito
	"Tochigi":     {Lat: 36.5657, Lon: 139.8836}, // Utsunomiya
	"Gunma":       {Lat: 36.3907, Lon: 139.0604}, // Maebashi
	"Saitama":     {Lat: 35.8569, Lon: 139.6489}, // Saitama
	"Chiba":       {Lat: 35.6051, Lon: 140.1233}, // Chiba
	"Tokyo":       {Lat: 35.6895, Lon: 139.6917}, // Tokyo
	"Kanagawa":    {Lat: 35.4478, Lon: 139.6425}, // Yokohama
	// Add other prefectures as needed
}

// getCoordinatesForPlace returns the coordinates of the capital city for a given prefecture.
func getCoordinatesForPlace(placeName string) (float64, float64) {
	if coords, ok := prefectureCapitalsCoordsMap[placeName]; ok {
		return coords.Lat, coords.Lon
	}
	return 0.0, 0.0
}
// UVIndexGet function modified to return only the UV index value
func UVIndexGet(placeName string) (float64, error) {
	lat, lon := getCoordinatesForPlace(placeName)
	uvApiKey := os.Getenv("WEATHER_API_KEY")
	apiURL := fmt.Sprintf("http://api.openweathermap.org/data/2.5/uvi?lat=%f&lon=%f&appid=%s", lat, lon, uvApiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading API response: %v", err)
	}

	var uvData UVData
	if err := json.Unmarshal(body, &uvData); err != nil {
		return 0, fmt.Errorf("error decoding API response: %v", err)
	}

	return uvData.Value, nil
}
func UVLevelGet(placeName string) (float64, error) {
    uvValue, err := UVIndexGet(placeName)
    if err != nil {
        return 0, err
    }

    switch {
    case uvValue <= 2:
        return 0, nil
    case uvValue <= 5:
        return 0.5, nil
    case uvValue <= 7:
        return 1, nil
    case uvValue <= 10:
        return 1.5, nil
    default:
        return 2, nil
    }
}

func TestUV() {
    placeName := "Tokyo"

    uvLevel, err := UVLevelGet(placeName)
    if err != nil {
        log.Fatalf("Failed to fetch UV level: %v", err)
    }

    fmt.Printf("UV level for %s is %f\n", placeName, uvLevel)
}