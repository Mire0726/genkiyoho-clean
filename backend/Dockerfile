# ビルドステージ
FROM golang:1.18 AS builder
WORKDIR /app

# backendフォルダ内のgo.modとgo.sumをコピー
COPY go.mod .
COPY go.sum .

# 依存関係のダウンロード
RUN go mod tidy

# # backendフォルダ内のソースコードをコピー
# COPY backend/ .

# .envファイルをコピー（もしルート直下にあるなら、この行は不要）
COPY . .


# アプリケーションのビルド
RUN go build -o main .

# 実行ステージ
FROM alpine:latest
WORKDIR /root/

# ビルドステージからバイナリをコピー
COPY --from=builder /app/main .

# ポート8080を公開
EXPOSE 8080

# 実行コマンド
CMD ["./main"]