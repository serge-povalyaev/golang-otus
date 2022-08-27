package internalhttp

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/serge.povalyaev/calendar/internal/logger"
	"github.com/serge.povalyaev/calendar/internal/repository"
	"net/http"
	"strconv"
	"time"
)

type ServerHandler struct {
	logger          *logger.CalendarLogger
	eventRepository repository.EventRepository
}

func (handler *ServerHandler) Add(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var event *repository.Event
	err := json.NewDecoder(request.Body).Decode(&event)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err)

		return
	}

	event.UserID, err = strconv.Atoi(request.Header.Get("UserId"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(err)

		return
	}

	event.ID = uuid.New()
	writer.WriteHeader(http.StatusCreated)
	err = handler.eventRepository.AddEvent(event)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err)

		return
	}

	json.NewEncoder(writer).Encode(struct{ ID uuid.UUID }{event.ID})
}

func (handler *ServerHandler) Get(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	eventId, userID, err := getEventAndUser(writer, request)
	if err != nil {
		return
	}

	event, err := handler.eventRepository.GetEvent(userID, eventId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	json.NewEncoder(writer).Encode(event)
}

func (handler *ServerHandler) Update(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	eventId, userID, err := getEventAndUser(writer, request)
	if err != nil {
		return
	}

	var event *repository.Event
	err = json.NewDecoder(request.Body).Decode(&event)

	if err != nil {
		fmt.Println(request.Body)
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	event.ID = eventId
	event.UserID = userID

	err = handler.eventRepository.UpdateEvent(userID, event)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	json.NewEncoder(writer).Encode(struct{ ID uuid.UUID }{event.ID})
}

func (handler *ServerHandler) Remove(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	eventId, userID, err := getEventAndUser(writer, request)
	if err != nil {
		return
	}

	event := &repository.Event{ID: eventId}
	err = handler.eventRepository.RemoveEvent(userID, event)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}
}

func (handler *ServerHandler) Events(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	userID, err := strconv.Atoi(request.Header.Get("UserId"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	from, err := time.Parse("2006-01-02 15:04:05", request.FormValue("from"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	to, err := time.Parse("2006-01-02 15:04:05", request.FormValue("to"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	events, err := handler.eventRepository.GetUserEvents(userID, from, to)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	json.NewEncoder(writer).Encode(events)
}

func (handler *ServerHandler) EventsByDay(writer http.ResponseWriter, request *http.Request) {
	getEvents(writer, request, handler.eventRepository.GetUserEventsByDay)
}

func (handler *ServerHandler) EventsByWeek(writer http.ResponseWriter, request *http.Request) {
	getEvents(writer, request, handler.eventRepository.GetUserEventsByWeek)
}

func (handler *ServerHandler) EventsByMonth(writer http.ResponseWriter, request *http.Request) {
	getEvents(writer, request, handler.eventRepository.GetUserEventsByMonth)
}

func getEvents(
	writer http.ResponseWriter,
	request *http.Request,
	eventsFunc func(userID int, dateStart time.Time) ([]*repository.Event, error),
) {
	writer.Header().Set("Content-Type", "application/json")
	userID, err := strconv.Atoi(request.Header.Get("UserId"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	from, err := time.Parse("2006-01-02 15:04:05", request.FormValue("from"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	events, err := eventsFunc(userID, from)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return
	}

	json.NewEncoder(writer).Encode(events)
}

func getEventAndUser(writer http.ResponseWriter, request *http.Request) (uuid.UUID, int, error) {
	eventId, err := uuid.Parse(request.FormValue("eventId"))
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return uuid.Nil, 0, err
	}

	userID, err := strconv.Atoi(request.Header.Get("UserId"))
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})

		return uuid.Nil, 0, err
	}

	return eventId, userID, err
}
