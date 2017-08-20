package controller

import (
	"time"
	"net/http"
	"app/model"
	"log"
	"fmt"
	"encoding/json"
	"app/shared/workers"
)

type object struct {
	from time.Time `json:"date1,string,omitepty"`
	to time.Time   `json:"date2,string,omitepty"`
	action string  `json:"action,string,omitepty"`
	limit uint32   `json:"limit,string,omitepty"`
}

func GetTop(w http.ResponseWriter, r *http.Request){
	_, err := workers.WP.AddTaskSyncTimed(func() interface{}{
		from := r.URL.Query().Get("date1")
		to := r.URL.Query().Get("date2")
		action := r.URL.Query().Get("action")
		limit := r.URL.Query().Get("limit")

		if!validateJsonTop(from, to, action, limit){
			w.WriteHeader(http.StatusPreconditionFailed)
			return new(error)
		}

		mapItems := make(map[string][]model.Row)

		mapItems, err := model.GetTopUsersByAction(action, from, to, limit)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		data := struct {
			Items map[string][]model.Row `json:"items"`
		}{mapItems}


		resp, err := json.Marshal(data)


		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
		w.Write(resp)
		return err
	}, workers.RequestWaitInQueueTimeout)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}
}

func validateJsonTop(from, to, action, limit string ) bool{
	var err error

	if limit == "" {
		fmt.Println("limit is empty")
		return false
	}
	if limit == "0" {
		fmt.Println("limit is nil")
		return false
	}

	_, err = time.Parse("2006-01-02", from)
	if err != nil {
		return false
	}
	_, err = time.Parse("2006-01-02", to)
	if err != nil {
		return false
	}
	if from > to {
		fmt.Println("Dates incorrect")
		return false
	}

	switch action {
	case "like":
	case "comment":
	case "login":
	default: return false
	}
	return true
}

