package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/serge.povalyaev/calendar/internal/repository"
	"time"
)

type Server struct {
	eventRepository repository.EventRepository
}

const (
	dateFormat  = "2006-02-01 15:04:05"
	ErrNewEvent = iota
	ErrAddEvent
	ErrParseUUID
	ErrGetEvent
	ErrParseDate
	ErrUpdateEvent
	ErrRemoveEvent
	ErrGetUserEvents
)

func NewServer(eventRepository repository.EventRepository) *Server {
	return &Server{eventRepository: eventRepository}
}

func (s Server) AddEvent(ctx context.Context, request *AddEventRequest) (*AddEventResponse, error) {
	response := &AddEventResponse{}
	event, err := repository.NewEvent(
		uuid.New(),
		request.Event.Title,
		request.Event.DateStart,
		request.Event.DateFinish,
		request.Event.Description,
		int(request.UserId),
		time.Duration(request.Event.NotifyBefore)*time.Second,
	)

	if err != nil {
		response.Error = &Error{Code: ErrNewEvent, Message: err.Error()}

		return response, err
	}

	err = s.eventRepository.AddEvent(event)
	if err != nil {
		response.Error = &Error{Code: ErrAddEvent, Message: err.Error()}

		return response, err
	}

	response.Id = event.ID.String()

	return response, nil
}

func (s Server) GetEvent(ctx context.Context, request *GetEventRequest) (*GetEventResponse, error) {
	response := &GetEventResponse{}

	eventID, err := uuid.Parse(request.Id)
	if err != nil {
		response.Error = &Error{Code: ErrParseUUID, Message: err.Error()}

		return response, err
	}

	event, err := s.eventRepository.GetEvent(int(request.UserId), eventID)
	if err != nil {
		response.Error = &Error{Code: ErrGetEvent, Message: err.Error()}

		return response, err
	}

	response.Event = &Event{
		Id:        request.Id,
		UserId:    request.UserId,
		EventData: transformEventToEventData(event),
	}

	return response, nil
}

func (s Server) UpdateEvent(ctx context.Context, request *UpdateEventRequest) (*Error, error) {
	var event *repository.Event

	eventID, err := uuid.Parse(request.Id)
	if err != nil {
		e := &Error{Code: ErrParseUUID, Message: err.Error()}

		return e, err
	}

	dateStart, err := time.Parse(dateFormat, request.EventData.DateStart)
	if err != nil {
		e := &Error{Code: ErrParseDate, Message: err.Error()}

		return e, err
	}

	dateFinish, err := time.Parse(dateFormat, request.EventData.DateFinish)
	if err != nil {
		e := &Error{Code: ErrParseDate, Message: err.Error()}

		return e, err
	}

	event = &repository.Event{
		ID:           eventID,
		UserID:       int(request.UserId),
		Title:        request.EventData.Title,
		Description:  request.EventData.Description,
		DateStart:    dateStart,
		DateFinish:   dateFinish,
		NotifyBefore: int(request.EventData.NotifyBefore),
	}

	err = s.eventRepository.UpdateEvent(event.UserID, event)
	if err != nil {
		e := &Error{Code: ErrUpdateEvent, Message: err.Error()}

		return e, err
	}

	return &Error{}, nil
}

func (s Server) RemoveEvent(ctx context.Context, request *RemoveEventRequest) (*Error, error) {
	eventID, err := uuid.Parse(request.Id)
	if err != nil {
		e := &Error{Code: ErrParseUUID, Message: err.Error()}

		return e, err
	}

	event := &repository.Event{ID: eventID}
	err = s.eventRepository.RemoveEvent(int(request.UserId), event)
	if err != nil {
		e := &Error{Code: ErrRemoveEvent, Message: err.Error()}

		return e, err
	}

	return &Error{}, nil
}

func (s Server) GetUserEvents(ctx context.Context, request *GetUserEventsRequest) (*GetUserEventsResponse, error) {
	response := &GetUserEventsResponse{}

	from, err := time.Parse("2006-01-02 15:04:05", request.From)
	if err != nil {
		response.Error = &Error{Code: ErrParseDate, Message: err.Error()}

		return response, err
	}

	to, err := time.Parse("2006-01-02 15:04:05", request.To)
	if err != nil {
		response.Error = &Error{Code: ErrParseDate, Message: err.Error()}

		return response, err
	}

	events, err := s.eventRepository.GetUserEvents(int(request.UserId), from, to)
	if err != nil {
		response.Error = &Error{Code: ErrGetUserEvents, Message: err.Error()}

		return response, err
	}

	response.Events = transformEventsList(events)

	return response, nil
}

func (s Server) GetUserEventsByDay(ctx context.Context, request *GetUserEventsShortRequest) (*GetUserEventsResponse, error) {
	return getEvents(request, s.eventRepository.GetUserEventsByDay)
}

func (s Server) GetUserEventsByWeek(ctx context.Context, request *GetUserEventsShortRequest) (*GetUserEventsResponse, error) {
	return getEvents(request, s.eventRepository.GetUserEventsByWeek)
}

func (s Server) GetUserEventsByMonth(ctx context.Context, request *GetUserEventsShortRequest) (*GetUserEventsResponse, error) {
	return getEvents(request, s.eventRepository.GetUserEventsByMonth)
}

func (s Server) mustEmbedUnimplementedEventServiceServer() {}

func getEvents(request *GetUserEventsShortRequest, eventsFunc func(UserID int, from time.Time) ([]*repository.Event, error)) (*GetUserEventsResponse, error) {
	response := &GetUserEventsResponse{}

	from, err := time.Parse("2006-01-02 15:04:05", request.From)
	if err != nil {
		response.Error = &Error{Code: ErrParseDate, Message: err.Error()}

		return response, err
	}

	events, err := eventsFunc(int(request.UserId), from)
	if err != nil {
		response.Error = &Error{Code: ErrGetUserEvents, Message: err.Error()}

		return response, err
	}

	response.Events = transformEventsList(events)

	return response, nil
}

func transformEventToEventData(event *repository.Event) *EventData {
	return &EventData{
		Title:        event.Title,
		DateStart:    event.DateStart.Format(dateFormat),
		DateFinish:   event.DateFinish.Format(dateFormat),
		Description:  event.Description,
		NotifyBefore: int64(event.NotifyBefore),
	}
}

func transformEventsList(events []*repository.Event) []*Event {
	eventsList := make([]*Event, 0, len(events))
	for _, event := range events {
		eventsList = append(eventsList, &Event{
			Id:        event.ID.String(),
			UserId:    int64(event.UserID),
			EventData: transformEventToEventData(event),
		})
	}

	return eventsList
}
