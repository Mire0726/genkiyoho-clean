package model

import "time"

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
