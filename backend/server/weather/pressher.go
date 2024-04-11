package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/joho/godotenv"
)

// OpenWeather APIからのレスポンスを格納するための構造体
type PressureResponse struct {
	Main struct {
		Pressure float64 `json:"pressure"` // 気圧
	} `json:"main"`
}

// 指定した都市の現在の気圧を基にブール値を返す関数
func CheckPressure(city string) bool {

	// err := godotenv.Load() 
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// OpenWeather APIキー（自分のAPIキーに置き換えてください）
	const pressureApiKey = "WEATHER_API_KEY"

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, pressureApiKey)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching weather data: %s", err)
	}
	defer resp.Body.Close()


	var pressure PressureResponse
	if err := json.NewDecoder(resp.Body).Decode(&pressure); err != nil {
		log.Fatalf("Error decoding pressure data: %s", err)
	}
	

	currentPressure := pressure.Main.Pressure
	averagePressure := 1013.0 // 平均気圧

	// 気圧が平均から6〜10ヘクトパスカル下がっているか判定
	if currentPressure >= averagePressure-15 && currentPressure <= averagePressure-5 {
		return true
	}

	return false
}

func Pressure() {
	city := "Tokyo"
	if CheckPressure(city) {
		fmt.Printf("The pressure in %s has dropped below the average threshold.\n", city)
	} else {
		fmt.Printf("The pressure in %s is within the normal range.\n", city)
	}

}