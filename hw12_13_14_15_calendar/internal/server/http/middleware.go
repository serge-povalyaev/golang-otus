package internalhttp

import (
	"fmt"
	"github.com/serge.povalyaev/calendar/internal/logger"
	"net/http"
	"time"
)

func logging(handler http.HandlerFunc, logger *logger.CalendarLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		duration := time.Since(start)

		logger.Info(
			fmt.Sprintf(
				"%s [%s] %s %s %s (%.2fs)",
				r.RemoteAddr,
				time.Now().Format("2006-01-02 15:04:05"),
				r.Method,
				r.URL,
				r.UserAgent(),
				duration.Seconds(),
			),
		)
	}
}
