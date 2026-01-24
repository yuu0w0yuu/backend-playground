package middleware

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

// ステータスコードを取得するためのResponseWriterラッパー
type responseWriterInterceptor struct {
	http.ResponseWriter // データ・レスポンスコードを送るhttp.ResponseWriterを埋め込み
	statusCode int		// レスポンスコードを格納するフィールド
}

// WriteHeaderメソッドを、ステータスコードをキャプチャするようにオーバーライド
func (rw *responseWriterInterceptor) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code) // 本来のWriteHeaderを呼び出す
}

func LoggingMiddleware(next http.Handler) http.Handler {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rwi := &responseWriterInterceptor{ // ResponseWriterラッパー構造体のインスタンス生成
			ResponseWriter: w,
			statusCode: http.StatusOK,
		}

		// 次のハンドラを呼び出す
		next.ServeHTTP(rwi, r)

		duration := time.Since(start)

		// ログ出力
		logger.Info("request completed",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.Int("status", rwi.statusCode),
			slog.Duration("duration", duration),
		)
	})
}