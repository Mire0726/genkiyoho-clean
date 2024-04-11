package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Mire0726/Genkiyoho/backend/server/context/auth"
	"github.com/Mire0726/Genkiyoho/backend/server/model"
	"github.com/Mire0726/Genkiyoho/backend/server/weather"
	"github.com/labstack/echo/v4"
)
func isInCurrentCycle(startDate time.Time, duration, cycleLength int) bool {
    // cycleLengthが0の場合は、即座にfalseを返すか、適切なデフォルト値を設定します。
    if cycleLength <= 0 {
        return false // または、適切なデフォルト値に基づいたロジックをここに追加
    }

    today := time.Now()

    // 開始日から今日までの総日数を計算
    totalDays := today.Sub(startDate).Hours() / 24

    // 総日数を周期の長さで割った余りを計算
    currentCycleDay := int(totalDays) % cycleLength

    // 現在の日付が活動期間内かどうかを確認
    return currentCycleDay < duration
}



func HandleUserTodayCycleConditionGet(c echo.Context) error {
	userID := auth.GetUserIDFromContext(c.Request().Context())
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
	}
	conditions, err := model.GetUserConditions(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
	}

	var todayConditions_c []model.UserCondition
	for _, condition := range conditions {
		startDate, err := time.Parse(time.RFC3339, condition.StartDate.Format(time.RFC3339)) // Convert condition.StartDate to string
		if err != nil {
			continue // 日付のパースに失敗した場合は、このコンディションをスキップします。
		}
		if isInCurrentCycle(startDate, condition.Duration, condition.CycleLength) {
			todayConditions_c = append(todayConditions_c, condition)
		}
	}

	return c.JSON(http.StatusOK, todayConditions_c)
}

func HandleUserTodayEnvironmentConditionGet(c echo.Context) error {
	userID := auth.GetUserIDFromContext(c.Request().Context())
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
	}
	conditions, err := model.GetUserConditions(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
	}
	fmt.Println(conditions)
	var todayConditions_e []model.UserCondition
	for _, condition := range conditions { 
		c:= condition.Region
		if condition.Name == "花粉" {
			if weather.CheckPollen(c) > 0 {
				todayConditions_e = append(todayConditions_e, condition)
			}
		}
		if condition.Name == "気圧の不調" {
			if weather.CheckPressure(c) {
				todayConditions_e = append(todayConditions_e, condition)
			}
		}
		if condition.Name == "紫外線" {
			uvLevel, err := weather.UVLevelGet(c)
			if err == nil && uvLevel >= 2.0 {
				todayConditions_e = append(todayConditions_e, condition)
			}
		}
	}

	return c.JSON(http.StatusOK, todayConditions_e)
}

func HandleUserTodayConditionGet(c echo.Context) error {
	userID := auth.GetUserIDFromContext(c.Request().Context())
	if userID == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
	}
	conditions, err := model.GetUserConditions(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
	}

	// 各条件のconditionnameとdamagepointを格納するためのスライス
	var conditionDetails []map[string]interface{}

	for _, condition := range conditions {
		if condition.Name == "月経中" ||  condition.Name == "PMS" || condition.Name == "PMDD" {
		// 現在のサイクルに一致するかチェック
		startDate,err:= time.Parse(time.RFC3339, condition.StartDate.Format(time.RFC3339))
		if err == nil && isInCurrentCycle(startDate, condition.Duration, condition.CycleLength) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
		}
		} else {
		// 環境条件に一致するかチェック
		t := weather.CheckPollen(condition.Region)
		if condition.Name == "花粉"  && t>0 {
			if t>0 && t<30{
				conditionDetail := map[string]interface{}{
					"condition_name": condition.Name,
					"damage_point":   condition.DamagePoint,
				}
				conditionDetails = append(conditionDetails, conditionDetail)
			} else if t>=30 && t<60{
				conditionDetail := map[string]interface{}{
					"condition_name": condition.Name,
					"damage_point":   condition.DamagePoint*2,
				}
				conditionDetails = append(conditionDetails, conditionDetail)
			} else if t>=60 && t<100{
				conditionDetail := map[string]interface{}{
					"condition_name": condition.Name,
					"damage_point":   condition.DamagePoint*3,
				}
				conditionDetails = append(conditionDetails, conditionDetail)
			}
		}

		if condition.Name=="気圧の不調" && weather.CheckPressure(condition.Region) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
		}

		if condition.Name=="雨による不調" && weather.CheckWeather(condition.Region) {
			conditionDetail := map[string]interface{}{
				"condition_name": condition.Name,
				"damage_point":   condition.DamagePoint,
			}
			conditionDetails = append(conditionDetails, conditionDetail)
		}

		if condition.Name=="紫外線" {
			uvLevel, err := weather.UVLevelGet(condition.Region)
			if err == nil && uvLevel >= 2.0 {
				conditionDetail := map[string]interface{}{
					"condition_name": condition.Name,
					"damage_point":   condition.DamagePoint*int(uvLevel),
				}
				conditionDetails = append(conditionDetails, conditionDetail)
			}
		}
	}

	}
	return c.JSON(http.StatusOK, conditionDetails)
}

func HandleUserTodayPointGet(c echo.Context) error {
    userID := auth.GetUserIDFromContext(c.Request().Context())
    if userID == "" {
        return echo.NewHTTPError(http.StatusUnauthorized, "User ID is empty")
    }
    
    conditions, err := model.GetUserConditions(userID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
    }

    var totalDamagePoints int
    for _, condition := range conditions {
        startDate, err := time.Parse(time.RFC3339, condition.StartDate.Format(time.RFC3339))
        if err != nil {
            continue // If start date parsing fails, skip this condition.
        }

        // Check if the condition is within the current cycle
        if isInCurrentCycle(startDate, condition.Duration, condition.CycleLength) {
            totalDamagePoints += condition.DamagePoint
        }


		
        // Additional checks for environmental conditions
        switch condition.Name {
        case "花粉":
            pollenCount := weather.CheckPollen(condition.Region)
            switch {
            case pollenCount > 0 && pollenCount < 30:
                totalDamagePoints += condition.DamagePoint
            case pollenCount >= 30 && pollenCount < 60:
                totalDamagePoints += condition.DamagePoint * 2
            case pollenCount >= 60:
                totalDamagePoints += condition.DamagePoint * 3
            }

        case "気圧の不調":
            if weather.CheckPressure(condition.Region) {
                totalDamagePoints += condition.DamagePoint
            }

        case "雨による不調":
            if weather.CheckWeather(condition.Region) {
                totalDamagePoints += condition.DamagePoint
            }
        
		case "紫外線":
			if uvLevel, err := weather.UVLevelGet(condition.Region); err == nil && uvLevel >= 2.0 {
			uvLevel, err := weather.UVLevelGet(condition.Region)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get UV level: "+err.Error())
			}
			totalDamagePoints += condition.DamagePoint * int(uvLevel)
		}
		}
	}
    
	genkiHP := 100 - totalDamagePoints
    if genkiHP < 0 {
        genkiHP = 0 // Ensure that Genki HP does not go below 0
    }

return c.JSON(http.StatusOK, genkiHP)

}
