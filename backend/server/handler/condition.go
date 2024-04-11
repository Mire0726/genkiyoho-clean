package handler

import (
	"errors"
	"fmt"

	"net/http"
	"time"

	"github.com/Mire0726/Genkiyoho/backend/server/context/auth"
	"github.com/Mire0726/Genkiyoho/backend/server/model"
	"github.com/labstack/echo/v4"
)

// 基本リクエスト構造体
type baseConditionRequest struct {
    ConditionID int `json:"condition_id"`
    StartDate string `json:"start_date"`
}

// サイクル条件リクエスト構造体
type cycleConditionRequest struct {
    baseConditionRequest
    Duration int `json:"duration"`
    CycleLength int `json:"cycle_length"`
    DamagePoint int `json:"damage_point"`
}

// 環境条件リクエスト構造体
type environmentConditionRequest struct {
    baseConditionRequest
    Region string `json:"region"`
    Count int `json:"count"`
    DamagePoint int `json:"damage_point"`
}

// 共通の前処理
func commonConditionPreprocess(c echo.Context, req interface{}) (string, error) {
	if err := c.Bind(req); err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "Invalid request body format")
	}
	// Contextから認証済みのユーザIDを取得
	ctx := c.Request().Context()
	userID := auth.GetUserIDFromContext(ctx)
	if userID == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
	}
	// ここで共通の前処理を行う
	return userID, nil
}


// サイクル条件の登録
func HandleCycleConditionCreate() echo.HandlerFunc {
    return func(c echo.Context) error {
        req := &cycleConditionRequest{}
        userID, err := commonConditionPreprocess(c, req)
        if err != nil {
            return err
        }
        
        userCondition, err := convertToUserCondition(req,userID)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to convert request: "+err.Error())
        }
        
        // UserIDをUserCondition構造体に設定
        userCondition.UserID = userID

        if err := model.InsertCycleCondition(userCondition); err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to insert cycle condition: "+err.Error())
        }

        return c.NoContent(http.StatusOK)
    }
}

// 環境条件の登録
func HandleEnvironmentConditionCreate() echo.HandlerFunc {
    return func(c echo.Context) error {
        req := &environmentConditionRequest{}
        fmt.Printf("Received request: %+v\n", req)
        userID, err := commonConditionPreprocess(c, req) // userID を受け取る
        if err != nil {
            return err
        }

        userCondition, err := convertToUserCondition(req, userID) // convertToUserCondition 関数に userID を渡す
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to convert request: "+err.Error())
        }

        if err := model.InsertEnvironmentCondition(userCondition); err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to insert environment condition: "+err.Error())
        }
        return c.NoContent(http.StatusOK)
    }
}

// リクエストをUserConditionに変換
func convertToUserCondition(req interface{}, userID string) (*model.UserCondition, error) {
    var uc model.UserCondition
    
    // conditionIDを取得するための一時変数を初期化します。
    var conditionID int

    // reqの型に応じて、conditionIDを取得します。
    switch v := req.(type) {
    case *cycleConditionRequest:
        conditionID = v.ConditionID
    case *environmentConditionRequest:
        conditionID = v.ConditionID
    default:
        return nil, errors.New("invalid request type")
    }

    // conditionIDを使用して、条件のタイプと名前を取得します。
    conditionTypeName, err := model.GetConditionTypeName(conditionID)
    if err != nil {
        return nil, err
    }

    // 再度、reqの型に応じた処理を行います。
    switch v := req.(type) {
    case *cycleConditionRequest:
        startDate, err := time.Parse("2006-01-02", v.StartDate)
        if err != nil {
            return nil, err
        }
        uc = model.UserCondition{
            UserID:      userID,
            ConditionID: v.ConditionID,
            Name:        conditionTypeName.Name, // Nameを設定
            StartDate:   startDate,
            Duration:    v.Duration,
            CycleLength: v.CycleLength,
            DamagePoint: v.DamagePoint,
        }
    case *environmentConditionRequest:
        startDate, err := time.Parse("2006-01-02", v.StartDate)
        if err != nil {
            return nil, err
        }
        uc = model.UserCondition{
            UserID:      userID,
            Name:        conditionTypeName.Name, // Nameを設定
            StartDate:   startDate,
            Region:      v.Region,
            Count:       v.Count,
            DamagePoint: v.DamagePoint,
        }
    }
    fmt.Println(uc)
    return &uc, nil
}


//　特定のユーザーのすべてのconditionを取得
func HandleUserConditionGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        userID := auth.GetUserIDFromContext(c.Request().Context())
        if userID == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "userID is empty")
        }
        conditions, err := model.GetUserConditions(userID)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user conditions: "+err.Error())
        }
        
        return c.JSON(http.StatusOK, conditions)
    }
}

//conditionsの取得
func HandleConditionsGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        conditions, err := model.GetConditions()
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get conditions: "+err.Error())
        }
        return c.JSON(http.StatusOK, conditions)
    }
}

func HandleCycleConditionGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        cycle_conditions, err := model.GetCycleConditions()
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get conditions: "+err.Error())
        }
        return c.JSON(http.StatusOK, cycle_conditions)
    }
}

func HandleEnvironmentConditionGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        environment_conditions, err := model.GetEnvironmentConditions()
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get conditions: "+err.Error())
        }
        return c.JSON(http.StatusOK, environment_conditions)
    }
}


