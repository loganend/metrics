package controller

import (
	"net/http"
	"io/ioutil"
	"app/model"
	"encoding/json"
	"log"
	"time"
)

func NewStat(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var stat model.Stat
	err = json.Unmarshal(b, &stat)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	if (stat.UID == 0 || stat.Action == "" || stat.Datetime == time.Time{}){
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	ex := model.StatCreate(stat.UID, stat.Action, stat.Datetime)

	if ex != nil {
		log.Println(ex)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
