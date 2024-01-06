package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/* GET */
func eventsForDayHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Client %s is sending a request 'events_for_day' (%s)", req.Host, req.Method)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	response := &GETResponse{}

	currentDay := time.Now().AddDate(0, 0, -1)
	afterDay := time.Now().AddDate(0, 0, 1)

	for _, event := range events {
		if date, err := time.Parse(dateFormat, event.Date); err == nil {
			if date.After(currentDay) && date.Before(afterDay) {
				response.Result = append(response.Result, event)
			}
		}
	}

	if len(response.Result) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}
	json.NewEncoder(w).Encode(response)
}

/* GET */
func eventsForWeekHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Client %s is sending a request 'events_for_week' (%s)", req.Host, req.Method)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	response := &GETResponse{}

	previousWeek := time.Now().AddDate(0, 0, -7)
	afterWeek := time.Now().AddDate(0, 0, 7)

	for _, event := range events {
		if date, err := time.Parse(dateFormat, event.Date); err == nil {
			if date.After(previousWeek) && date.Before(afterWeek) {
				response.Result = append(response.Result, event)
			}
		}
	}

	if len(response.Result) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}
	json.NewEncoder(w).Encode(response)
}

/* GET */
func eventsForMonthHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Client %s is sending a request 'events_for_month' (%s)", req.Host, req.Method)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	response := &GETResponse{}

	previousMonth := time.Now().AddDate(0, -1, 0)
	afterMonth := time.Now().AddDate(0, 1, 0)

	for _, event := range events {
		if date, err := time.Parse(dateFormat, event.Date); err == nil {
			if date.After(previousMonth) && date.Before(afterMonth) {
				response.Result = append(response.Result, event)
			}
		}
	}

	if len(response.Result) == 0 {
		w.WriteHeader(http.StatusNoContent)
	}
	json.NewEncoder(w).Encode(response)
}

/* POST */
func createEventHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Client %s is sending a request 'create_event' (%s)", req.Host, req.Method)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	if err := req.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "unable to parse a form data",
		})
		return
	}

	if !req.Form.Has("user_id") || !req.Form.Has("date") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "invalid data",
		})
		return
	}

	name := req.Form.Get("user_id")
	date := req.Form.Get("date")
	now := EventID(time.Now().Unix())
	events[now] = Event{
		UserID: name,
		Date:   date,
	}

	log.Println("Event is created.")
}

/* POST */
func updateEventHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Client %s is sending a request 'update_event' (%s)", req.Host, req.Method)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	if err := req.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "unable to parse a form data",
		})
		return
	}

	if !req.Form.Has("user_id") || !req.Form.Has("date") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "invalid data",
		})
		return
	}

	userId := req.Form.Get("user_id")
	date := req.Form.Get("date")
	for _, event := range events {
		if event.UserID == userId {
			event.Date = date
			break
		}
	}

	log.Println("Event is updated.")
}

/* POST */
func deleteEventHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Client %s is sending a request 'delete_event' (%s)", req.Host, req.Method)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	if err := req.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "unable to parse a form data",
		})
		return
	}

	if !req.Form.Has("user_id") || !req.Form.Has("date") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "invalid data",
		})
		return
	}

	userId := req.Form.Get("user_id")
	var deletedKey EventID
	for key, event := range events {
		if event.UserID == userId {
			deletedKey = key
			break
		}
	}

	delete(events, deletedKey)

	log.Printf("Event %d is deleted.", deletedKey)
}

func SetupCalenderServiceHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/events_for_day", eventsForDayHandler)
	mux.HandleFunc("/events_for_week", eventsForWeekHandler)
	mux.HandleFunc("/events_for_month", eventsForMonthHandler)
	mux.HandleFunc("/create_event", createEventHandler)
	mux.HandleFunc("/update_event", updateEventHandler)
	mux.HandleFunc("/delete_event", deleteEventHandler)
}
