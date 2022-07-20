package internalhttp

import (
	"context"
	"github.com/serge.povalyaev/calendar/internal/logger"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	logger     *logger.CalendarLogger
}

func NewServer(host, port string, logger *logger.CalendarLogger) *Server {
	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: getHandler(logger),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	return &Server{httpServer: server, logger: logger}
}

func getHandler(logger *logger.CalendarLogger) *http.ServeMux {
	handler := &ServerHandler{logger: logger}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", logging(handler.Hello, logger))

	return mux
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Close()
}
