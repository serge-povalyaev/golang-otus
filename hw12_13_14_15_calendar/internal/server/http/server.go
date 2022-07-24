package internalhttp

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/serge.povalyaev/calendar/internal/logger"
	"github.com/serge.povalyaev/calendar/internal/repository"
	"log"
	"net/http"
)

type Server struct {
	httpServer      *http.Server
	logger          *logger.CalendarLogger
	eventRepository repository.EventRepository
}

func NewServer(host, port string, logger *logger.CalendarLogger, eventRepository repository.EventRepository) *Server {
	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: getHandler(logger, eventRepository),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	return &Server{httpServer: server, logger: logger, eventRepository: eventRepository}
}

func getHandler(logger *logger.CalendarLogger, eventRepository repository.EventRepository) *mux.Router {
	handler := &ServerHandler{logger: logger, eventRepository: eventRepository}
	router := mux.NewRouter()
	router.HandleFunc("/events/add", logging(handler.Add, logger)).Methods("PUT")
	router.HandleFunc("/events/get", logging(handler.Get, logger)).Methods("GET")
	router.HandleFunc("/events/update", logging(handler.Update, logger)).Methods("POST")
	router.HandleFunc("/events/remove", logging(handler.Remove, logger)).Methods("DELETE")
	router.HandleFunc("/events", logging(handler.Events, logger)).Methods("GET")
	router.HandleFunc("/events/by-day", logging(handler.EventsByDay, logger)).Methods("GET")
	router.HandleFunc("/events/by-week", logging(handler.EventsByWeek, logger)).Methods("GET")
	router.HandleFunc("/events/by-month", logging(handler.EventsByMonth, logger)).Methods("GET")

	return router
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
