package weather

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	List []struct {
		Weather []struct {
			Main        string `json:"main"`        // 天気の概要
			Description string `json:"description"` // 天気の詳細な説明
		} `json:"weather"`
	} `json:"list"`
}

func CheckWeather(city string) bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	weatherApiKey := os.Getenv("WEATHER_API_KEY")

	// OpenWeatherのAPIエンドポイント（都市とAPIキーを指定）
	url := "http://api.openweathermap.org/data/2.5/forecast?q=" + city + "&appid=" + weatherApiKey

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching weather data: %s", err)
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		log.Fatalf("Error decoding weather data: %s", err)
	}

	// "Rain"の天気状態をチェック
	for _, listItem := range weather.List {
		for _, w := range listItem.Weather {
			fmt.Println(w.Main)
			if w.Main == "Rain" {
				return true // 雨が検出されたらtrueを返す
			} 
	}
}


	return false // 雨が検出されなかった場合はfalseを返す
}
func Weather() {
	city := "Tokyo"
	if CheckWeather(city) {
		fmt.Printf("Rain is expected in %s.\n", city)
	} else {
		fmt.Printf("No rain is expected in %s.\n", city)
	}
}