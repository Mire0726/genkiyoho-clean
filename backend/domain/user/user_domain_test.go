package user

import (
    "testing"
    "time"
)

func TestNewUser(t *testing.T) {
    // テスト用のタイムスタンプ
    createdAt := time.Now()
    updatedAt := createdAt

    // テストケースの定義
    tests := []struct {
        name      string
        id        string
        authToken string
        email     string
        password  string
        username  string
        createdAt time.Time
        updatedAt time.Time
        wantErr   bool
    }{
        {
            name:      "Valid user",
            id:        "1",
            authToken: "token123",
            email:     "example@example.com",
            password:  "password",
            username:  "John Doe",
            createdAt: createdAt,
            updatedAt: updatedAt,
            wantErr:   false,
        },
        {
            name:      "Empty email",
            id:        "2",
            authToken: "token1234",
            email:     "",
            password:  "password",
            username:  "Jane Doe",
            createdAt: createdAt,
            updatedAt: updatedAt,
            wantErr:   true,
        },
        {
            name:      "Empty password",
            id:        "3",
            authToken: "token12345",
            email:     "jane@example.com",
            password:  "",
            username:  "Jane Doe",
            createdAt: createdAt,
            updatedAt: updatedAt,
            wantErr:   true,
        },
    }

    // テストの実行
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := NewUser(tt.id, tt.authToken, tt.email, tt.password, tt.username, tt.createdAt, tt.updatedAt)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
