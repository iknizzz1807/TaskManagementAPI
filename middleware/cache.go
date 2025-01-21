package middleware

import (
	"bytes"
	"context"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/iknizzz1807/TaskManagementAPI/utils"
)

func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		cacheKey := r.RequestURI

		cachedResponse, err := utils.RedisClient.Get(ctx, cacheKey).Result()
		if err == redis.Nil {
			rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(rec, r)

			utils.RedisClient.Set(ctx, cacheKey, rec.body.String(), 10*time.Minute)
		} else if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cachedResponse))
		}
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *responseRecorder) Write(b []byte) (int, error) {
	rec.body.Write(b)
	return rec.ResponseWriter.Write(b)
}
