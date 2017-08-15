package controller

import (
	"time"
	"net/http"
	"app/model"
	"log"
	"fmt"
	"encoding/json"
)

type object struct {
	from time.Time `json:"date1,string,omitepty"`
	to time.Time   `json:"date2,string,omitepty"`
	action string  `json:"action,string,omitepty"`
	limit uint32   `json:"limit,string,omitepty"`
}



func GetTop(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("date1")
	to := r.URL.Query().Get("date2")
	action := r.URL.Query().Get("action")
	limit := r.URL.Query().Get("limit")

	mapItems := make(map[string][]model.Row)

	mapItems, ex := model.GetTopUsersByAction(action, from, to, limit)

	if ex != nil {
		log.Println(ex)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := struct {
		Items map[string][]model.Row `json:"items"`
	}{mapItems}

	fmt.Println("Db resp")
	fmt.Println(mapItems)

	resp, err := json.Marshal(data)

	fmt.Println("Json resp")
	fmt.Println(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}
