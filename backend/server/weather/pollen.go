package weather

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	"net/http"
	"strconv"
	"time"
)


func CheckPollen(Pre string) int{
	cityCode := GetCityCodeFromPrefecture(Pre)
	// 現在の日付をYYYYMMDD形式で取得
	currentDate := time.Now().Format("20060102")
	currentDateInt,err := strconv.Atoi(currentDate)

	// APIのエンドポイントを構築
	apiURL := fmt.Sprintf("https://wxtech.weathernews.com/opendata/v1/pollen?citycode=%d&start=%d&end=%d", cityCode, currentDateInt, currentDateInt)

	// HTTP GETリクエストを送信してみましょう
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return 0
	}

	defer response.Body.Close()

    // レスポンスボディをバイトスライスとして読み込む
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return 0
	}

// バイトスライスを文字列に変換し、CSVデータとして解析
bodyString := string(bodyBytes)
r := csv.NewReader(strings.NewReader(bodyString))

// CSVのヘッダー行を読み飛ばす
_, err = r.Read()
if err != nil {
	fmt.Println("Error reading CSV header:", err)
	return 0
}

maxPollen := 0
for {
	record, err := r.Read()
	if err == io.EOF {
		break
	}
	if err != nil {
		fmt.Println("Error reading CSV record:", err)
		return 0
	}

	// 花粉飛散量の値を整数に変換
	pollen, err := strconv.Atoi(record[2])
	if err != nil || pollen == -9999 {
		// 変換エラーまたは無効なデータの場合はスキップ
		continue
	}

	// 最大花粉飛散量を更新
	if pollen > maxPollen {
		maxPollen = pollen
	}
}
	return maxPollen
}

func GetCityCodeFromPrefecture(prefecture string) int {
	// 都道府県名をキーとして、都市コードを取得
	cityCode := PrefectureToCityCode[prefecture]
	if cityCode == 0 {
		fmt.Println("Invalid prefecture name")
		return 1
	}
	return cityCode
}

var PrefectureToCityCode = map[string]int{
	"Hokkaido": 01100, // 北海道札幌市
	"Aomori": 02201,
	"Iwate": 03201,
	"Miyagi": 04100,
	"Akita": 05201,
	"Yamagata": 06201,
	"Fukushima": 07201,
	// "Ibaraki": 08201,
	// "Tochigi": 09201,
	"Gunma": 10201,
	"Saitama": 11100,
	"Chiba": 12100,
	"Tokyo": 13103, // 東京都
	"Kanagawa": 14100,
	"Niigata": 15202,
	"Toyama": 16201,
	"Ishikawa": 17201,
	"Fukui": 18201,
	"Yamanashi": 19201,
	"Nagano": 20201,
	"Gifu": 21201,
	"Shizuoka": 22100,
	"Aichi": 23100,
	"Mie": 24201,
	"Shiga": 25201,
	"Kyoto": 26100,
	"Osaka": 27100, // 大阪府大阪市
	"Hyogo": 28100,
	"Nara": 29201,
	"Wakayama": 30201,
	"Tottori": 31201,
	"Shimane": 32201,
	"Okayama": 33201,
	"Hiroshima": 34201,
	"Yamaguchi": 35203,
	"Tokushima": 36201,
	"Kagawa": 37201,
	"Ehime": 38201,
	"Kochi": 39201,
	"Fukuoka": 40130, // 福岡県福岡市
	"Saga": 41201,
	"Nagasaki": 42201,
	"Kumamoto": 43201,
	"Oita": 44201,
	"Miyazaki": 45201,
	"Kagoshima": 46201,
	"Okinawa": 47201,
}