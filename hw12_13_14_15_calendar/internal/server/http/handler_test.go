package internalhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/serge.povalyaev/calendar/internal/logger"
	"github.com/serge.povalyaev/calendar/internal/repository"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	eventRepository := repository.GetEventRepository("memory", nil)
	calendarLogger := logger.New("info", "")
	ts := httptest.NewServer(getHandler(calendarLogger, eventRepository))
	defer ts.Close()

	eventID := testingAdd(t, ts)
	event := testingGet(t, ts, eventID)
	event = testingUpdate(t, ts, event)
	testingEvents(t, ts)
	testingRemove(t, ts, event)
}

func testingAdd(t *testing.T, ts *httptest.Server) uuid.UUID {
	event := createNewEvent()
	jsonEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}

	requestAdd, _ := http.NewRequest(http.MethodPut, ts.URL+"/events/add", bytes.NewReader(jsonEvent))
	requestAdd.Header.Set("UserId", "1")
	responseAdd, err := http.DefaultClient.Do(requestAdd)
	require.NoError(t, err)

	responseAddJson, err := ioutil.ReadAll(responseAdd.Body)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, responseAdd.StatusCode)

	responseAddData := make(map[string]string)
	err = json.Unmarshal(responseAddJson, &responseAddData)
	require.NoError(t, err)

	id, ok := responseAddData["ID"]
	require.True(t, ok)
	require.NotNil(t, id)

	eventID, err := uuid.Parse(id)
	require.NoError(t, err)
	require.NotNil(t, eventID)

	return eventID
}

func testingGet(t *testing.T, ts *httptest.Server, eventID uuid.UUID) *repository.Event {
	requestGet, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/events/get?eventId=%s", ts.URL, eventID), nil)
	requestGet.Header.Set("UserId", "1")
	responseGet, err := http.DefaultClient.Do(requestGet)
	require.NoError(t, err)

	responseGetJson, err := ioutil.ReadAll(responseGet.Body)
	require.NoError(t, err)

	var responseGetData *repository.Event
	err = json.Unmarshal(responseGetJson, &responseGetData)
	require.NoError(t, err)
	require.Equal(t, eventID, responseGetData.ID)

	testingNotFoundByUserAndEventId(t, requestGet, ts, http.MethodGet, "/events/get", nil)

	return responseGetData
}

func testingUpdate(t *testing.T, ts *httptest.Server, event *repository.Event) *repository.Event {
	trueEventId := event.ID
	trueUserId := event.UserID
	event.ID = uuid.New()
	event.UserID = 2
	event.Title = "Go"
	event.Description = "Go"
	event.NotifyBefore = 3600
	event.DateStart, _ = time.Parse("2006-01-02 15:04:05", "2022-08-01 15:00:00")
	event.DateStart, _ = time.Parse("2006-01-02 15:04:05", "2022-08-01 18:00:00")
	jsonEvent, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}

	requestUpdate, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/events/update?eventId=%s", ts.URL, trueEventId), bytes.NewReader(jsonEvent))
	requestUpdate.Header.Set("UserId", strconv.Itoa(trueUserId))
	responseUpdate, err := http.DefaultClient.Do(requestUpdate)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, responseUpdate.StatusCode)

	requestGet, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/events/get?eventId=%s", ts.URL, trueEventId), nil)
	requestGet.Header.Set("UserId", "1")
	responseGet, _ := http.DefaultClient.Do(requestGet)
	responseGetJson, _ := ioutil.ReadAll(responseGet.Body)

	var responseGetData *repository.Event
	_ = json.Unmarshal(responseGetJson, &responseGetData)

	event.UserID = trueUserId
	event.ID = trueEventId
	require.Equal(t, event, responseGetData)

	testingNotFoundByUserAndEventId(t, requestUpdate, ts, http.MethodPost, "/events/update", bytes.NewReader(jsonEvent))

	return responseGetData
}

func testingEvents(t *testing.T, ts *httptest.Server) {
	requestEvents, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/events?from=%s&to=%s", ts.URL, url.QueryEscape("2022-08-01 00:00:00"), url.QueryEscape("2022-08-30 00:00:00")), nil)
	requestEvents.Header.Set("UserId", "1")
	responseEvents, err := http.DefaultClient.Do(requestEvents)
	require.NoError(t, err)

	responseEventsJson, err := ioutil.ReadAll(responseEvents.Body)
	require.NoError(t, err)

	var responseEventsData []*repository.Event
	err = json.Unmarshal(responseEventsJson, &responseEventsData)
	require.NoError(t, err)
	require.Equal(t, 1, len(responseEventsData))

	requestEvents, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/events?from=%s&to=%s", ts.URL, url.QueryEscape("2022-08-01 00:00:00"), url.QueryEscape("2022-08-30 00:00:00")), nil)
	requestEvents.Header.Set("UserId", "2")
	responseEvents, _ = http.DefaultClient.Do(requestEvents)
	responseEventsJson, _ = ioutil.ReadAll(responseEvents.Body)
	err = json.Unmarshal(responseEventsJson, &responseEventsData)
	require.NoError(t, err)
	require.Equal(t, 0, len(responseEventsData))

	requestEvents, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/events?from=%s&to=%s", ts.URL, url.QueryEscape("2023-08-01 00:00:00"), url.QueryEscape("2023-08-30 00:00:00")), nil)
	requestEvents.Header.Set("UserId", "1")
	responseEvents, _ = http.DefaultClient.Do(requestEvents)
	responseEventsJson, _ = ioutil.ReadAll(responseEvents.Body)
	err = json.Unmarshal(responseEventsJson, &responseEventsData)
	require.NoError(t, err)
	require.Equal(t, 0, len(responseEventsData))
}

func testingRemove(t *testing.T, ts *httptest.Server, event *repository.Event) {
	requestRemove, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/events/remove?eventId=%s", ts.URL, event.ID), nil)
	requestRemove.Header.Set("UserId", strconv.Itoa(event.UserID))
	responseRemove, err := http.DefaultClient.Do(requestRemove)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, responseRemove.StatusCode)

	requestGet, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/events/get?eventId=%s", ts.URL, event.ID), nil)
	requestGet.Header.Set("UserId", "1")
	testingNotFound(t, requestGet)

}

func testingNotFoundByUserAndEventId(t *testing.T, request *http.Request, ts *httptest.Server, method string, url string, body io.Reader) {
	request.Header.Del("UserId")
	request.Header.Set("UserId", "2")
	testingNotFound(t, request)

	request, _ = http.NewRequest(method, fmt.Sprintf("%s%s?eventId=%s", ts.URL, url, uuid.New()), body)
	request.Header.Set("UserId", "1")
	testingNotFound(t, request)
}

func testingNotFound(t *testing.T, request *http.Request) {
	responseGetInvalidUser, err := http.DefaultClient.Do(request)
	require.NoError(t, err)

	responseGetInvalidUserJson, err := ioutil.ReadAll(responseGetInvalidUser.Body)
	responseGetInvalidUserData := make(map[string]string)
	err = json.Unmarshal(responseGetInvalidUserJson, &responseGetInvalidUserData)
	require.NoError(t, err)
	_, ok := responseGetInvalidUserData["error"]
	require.True(t, ok)
	require.Equal(t, http.StatusNotFound, responseGetInvalidUser.StatusCode)
}

func createNewEvent() *repository.Event {
	event, _ := repository.NewEvent(
		uuid.New(),
		"Занятие по Golang",
		"2022-07-15 20:00:00",
		"2022-07-15 21:30:00",
		"Сервис календаря",
		1,
		time.Duration(5400),
	)

	return event
}
