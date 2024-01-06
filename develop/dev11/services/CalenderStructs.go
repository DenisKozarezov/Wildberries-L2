package services

type EventID int64

type Event struct {
	UserID string
	Date   string
}

type GETResponse struct {
	Result []interface{}
}

type POSTResponse struct {
	Result []interface{}
}

type ErrorResponse struct {
	Error string
}

var events map[EventID]Event = make(map[EventID]Event, 100)

const (
	dateFormat = "2006-01-02"
)
