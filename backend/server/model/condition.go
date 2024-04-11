package model

import (
	"errors"
	"log"
	"time"

	"github.com/Mire0726/Genkiyoho/backend/server/db"
)

type UserCondition struct {
    UserID      string    // ユーザーID
    ConditionID int       // 条件ID
    Name        string    // 条件の名前（月経周期の場合は"PMS"、"月経期間"など、環境条件の場合は"花粉"、"PM2.5"など）
    StartDate   time.Time // 条件の開始日
    EndDate     time.Time // 条件の終了日（月経周期のみ、環境条件では使用しない）
    Duration    int       // 月経周期の期間（日数）
    CycleLength int       // 月経周期の長さ（前回の月経開始日から次の月経開始日までの日数）
    Region      string    // 環境条件の地域（環境条件のみ）
    Count       int       // 環境条件のカウント数（例えば花粉の数）
    DamagePoint int       // 条件によるダメージポイント（任意で使用）
}
type Condition struct {
    ID   int
    Name string
    Type string
}
type ConditionTypeName struct {
    Type    string
    Name string
}

// GetConditionType は指定された条件IDに基づいて条件のconditionのtypeとnameを取得します。
func GetConditionTypeName(conditionID int) (*ConditionTypeName, error) {
    var conditionType, conditionName string

    err := db.Conn.QueryRow("SELECT type, name FROM conditions WHERE id = ?", conditionID).Scan(&conditionType, &conditionName)
    if err != nil {
        log.Printf("Error retrieving condition type and name from database: %v", err)
        return nil, err
    }
    return &ConditionTypeName{Type: conditionType, Name: conditionName}, nil
}
// InsertUserCondition はユーザーの条件を登録します。
func InsertUserCondition(record *UserCondition) error {
    var conditionType,conditionName string
    err := db.Conn.QueryRow("SELECT type FROM conditions WHERE id = ?", record.ConditionID).Scan(&conditionType,&conditionName)
    if err != nil {
        log.Printf("Error retrieving condition type from database: %v", err)
        return err
    }

    switch conditionType {
    case "cycle":
        return InsertCycleCondition(record)
    case "environment":
        return InsertEnvironmentCondition(record)
    default:
        log.Printf("Invalid condition type: %s", conditionType)
        return errors.New("invalid condition type")
    }
}

// InsertCycleCondition は周期の条件を登録します。
func InsertCycleCondition(record *UserCondition) error {
    _, err := db.Conn.Exec(
        "INSERT INTO cycle_conditions (user_id,name, start_date, duration, cycle_length,damage_point) VALUES (?,?, ?, ?, ?, ?)",
        record.UserID,
        record.Name,
        record.StartDate,
        record.Duration,
        record.CycleLength,
        record.DamagePoint,
    )
    if err != nil {
        log.Printf("Error inserting into cycle_conditions database: %v", err)
        return err
    }

    log.Println("Cycle condition successfully registered.")
    return nil
}

// InsertEnvironmentCondition は環境の条件を登録します。
func InsertEnvironmentCondition(record *UserCondition) error {
    _, err := db.Conn.Exec(
        "INSERT INTO environment_conditions (user_id,date, region, count,name,damage_point) VALUES (?,?,?, ?, ?, ?)",
        record.UserID,
        record.StartDate, // このケースではStartDateをdateとして使用
        record.Region,    // UserConditionにRegionフィールドが必要
        record.Count,     // UserConditionにCountフィールドが必要
        record.Name,      // UserConditionにNameフィールドが必要
        record.DamagePoint,      // UserConditionにDamagePointフィールドが必要
    )
    if err != nil {
        log.Printf("Error inserting into environment_conditions database: %v", err)
        return err
    }

    log.Println("Environment condition successfully registered.")
    return nil
}

//　特定のユーザーのすべてのconditionを取得
// GetUserCondition は指定されたユーザーIDに関連付けられたすべての条件を取得します。
func GetUserConditions(userID string) ([]UserCondition, error) {
	var userConditions []UserCondition

	// サイクル条件を取得するクエリ
	cycleQuery := `SELECT  user_id, name, start_date, duration, cycle_length ,damage_point FROM cycle_conditions WHERE user_id = ?`
	rows, err := db.Conn.Query(cycleQuery, userID)
	if err != nil {
		log.Printf("Error retrieving cycle conditions from database: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var uc UserCondition
		err := rows.Scan(&uc.UserID, &uc.Name, &uc.StartDate, &uc.Duration, &uc.CycleLength, &uc.DamagePoint)
		if err != nil {
			log.Printf("Error scanning cycle conditions: %v", err)
			continue
		}
		// サイクル条件に特有のフィールドを設定
		uc.Region = "" // 環境条件のフィールドは空で設定
		uc.Count = 0   // 環境条件のフィールドは空で設定
		userConditions = append(userConditions, uc)
	}

	// 環境条件を取得するクエリ
	environmentQuery := `SELECT id,user_id, date AS start_date, region,name, count, damage_point FROM environment_conditions WHERE user_id = ?`
	rows, err = db.Conn.Query(environmentQuery, userID)
	if err != nil {
		log.Printf("Error retrieving environment conditions from database: %v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var uc UserCondition
		err := rows.Scan(&uc.ConditionID,&uc.UserID, &uc.StartDate, &uc.Region,&uc.Name, &uc.Count,&uc.DamagePoint)
		if err != nil {
			log.Printf("Error scanning environment conditions: %v", err)
			continue
		}
		// 環境条件に特有のフィールドを設定
		uc.Duration = 0    // サイクル条件のフィールドは空で設定
		uc.CycleLength = 0 // サイクル条件のフィールドは空で設定
		userConditions = append(userConditions, uc)
	}

	return userConditions, nil
}

// GetConditions はすべての条件を取得します。
func GetConditions() ([]Condition, error) {
    var conditions []Condition

    query := "SELECT id, name, type FROM conditions"
    rows, err := db.Conn.Query(query)
    if err != nil {
        log.Printf("Error retrieving conditions from database: %v", err)
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var c Condition
        err := rows.Scan(&c.ID, &c.Name, &c.Type)
        if err != nil {
            log.Printf("Error scanning conditions: %v", err)
            continue
        }
        conditions = append(conditions, c)
    }
    return conditions, nil
}

func GetCycleConditions() ([]Condition, error) {
    var cycle_conditions []Condition

    query := "SELECT id, name, type FROM conditions WHERE type = 'cycle'"
    rows, err := db.Conn.Query(query)
    if err != nil {
        log.Printf("Error retrieving cycle_conditions from database: %v", err)
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var c Condition
        err := rows.Scan(&c.ID, &c.Name, &c.Type)
        if err != nil {
            log.Printf("Error scanning conditions: %v", err)
            continue
        }
        cycle_conditions = append(cycle_conditions, c)
    }
    return cycle_conditions, nil
}

func GetEnvironmentConditions() ([]Condition, error) {
    var environment_conditions []Condition

    query := "SELECT id, name, type FROM conditions WHERE type = 'environment'"
    rows, err := db.Conn.Query(query)
    if err != nil {
        log.Printf("Error retrieving environment_conditions from database: %v", err)
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var c Condition
        err := rows.Scan(&c.ID, &c.Name, &c.Type)
        if err != nil {
            log.Printf("Error scanning conditions: %v", err)
            continue
        }
        environment_conditions = append(environment_conditions, c)
    }
    return environment_conditions, nil
}