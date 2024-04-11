package model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Mire0726/Genkiyoho/backend/server/db"
)

// User 構造体の定義（重複した定義を削除）
type User struct {
    ID        string     `json:"id"`
    AuthToken string    `json:"authtoken"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Name      string    `json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// InsertUser データベースにユーザレコードを登録する
func InsertUser(record *User) error  {

    _, err := db.Conn.Exec(
        "INSERT INTO users (id,auth_token, email, password, name) VALUES (?, ?, ?, ?, ?)",
        record.ID,
        record.AuthToken,
        record.Email,
        record.Password,
        record.Name,
    )
    if err != nil {
        log.Printf("Error inserting user into database: %v", err) // ログ追加
        return err
    }
    log.Println("User successfully registered.") // 成功メッセージもログに記録
    return nil
}


// SelectUserByPrimaryKey 主キーを条件にレコードを取得する
func SelectUserByPrimaryKey(userID string) (*User, error) {
	row := db.Conn.QueryRow("SELECT * FROM users WHERE id=?", userID)
	return convertToUser(row)
}

// UpdateUserByAuthToken 認証トークンを条件にレコードを更新する
func SelectUserByAuthToken(authToken string) (*User, error) {
	row := db.Conn.QueryRow("SELECT * FROM users WHERE auth_token=?", authToken)
	// convertToUser関数を使用して、行をUserオブジェクトに変換
    user, err := convertToUser(row)
    if err != nil {
        // エラーハンドリング：ログに記録、エラーを返すなど
        log.Printf("Error fetching user by auth token: %v", err)
        return nil, err
    }
    
    return user, nil
}

// UpdateUserByPrimaryKey 主キーを条件にレコードを更新する
func UpdateUserByPrimaryKey(record *User) error {
	if _, err := db.Conn.Exec(
		"UPDATE users SET name=?,email=?,password=? WHERE id=?",
		record.Name,
        record.Email,
        record.Password,
		record.ID,
	); err != nil {
		return err
	}
	return nil
}
// GetAllUsers データベースから全ユーザを取得する
func GetAllUsers() ([]User, error) {
    rows, err := db.Conn.Query("SELECT id, auth_token, email, password, name, created_at, updated_at FROM users")
    if err != nil {
        log.Printf("Error querying users from database: %v", err)
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User

        if err := rows.Scan(&user.ID, &user.AuthToken, &user.Email, &user.Password, &user.Name, &user.CreatedAt, &user.UpdatedAt); err != nil {
            log.Printf("Error scanning user: %v", err)
            return nil, err
        }

        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error during rows iteration: %v", err)
        return nil, err
    }

    return users, nil
}

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*User, error) {
	var user User
	err := row.Scan(&user.ID, &user.AuthToken, &user.Name, &user.Email, &user.Password,&user.CreatedAt,&user.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            // レコードが見つからない場合、nilとカスタムエラーを返す
            return nil, fmt.Errorf("user not found")
        }
        // その他のエラーの場合は、エラーをそのまま返す
        return nil, fmt.Errorf("error scanning user: %v", err)
    }
    // エラーがない場合は、取得したユーザーオブジェクトとnilを返す
    return &user, nil
    
}

// AuthenticateUser は指定されたメールアドレスとパスワードに一致するユーザーをデータベースから検索します。
func AuthenticateUser(db *sql.DB, email, password string) (*User, error) {
    user := &User{}
    err := db.QueryRow("SELECT id,auth_token, email, password FROM users WHERE email = ? AND password = ?", email, password).Scan(&user.ID,&user.AuthToken, &user.Email, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // ユーザーが見つからない場合はnilを返す
        }
        log.Printf("Failed to authenticate user: %v", err)
        return nil, err
    }
    return user, nil
}