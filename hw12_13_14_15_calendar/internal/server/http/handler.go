package internalhttp

import (
	"github.com/serge.povalyaev/calendar/internal/logger"
	"io"
	"net/http"
	"time"
)

type ServerHandler struct {
	logger *logger.CalendarLogger
}

func (handler *ServerHandler) Hello(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(2 * time.Second)
	_, _ = io.WriteString(writer, "<html><body>Hello World!</body></html>")
}
