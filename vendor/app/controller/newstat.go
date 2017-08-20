package controller

import (
	"net/http"
	"io/ioutil"
	"app/model"
	"encoding/json"
	"log"
	"time"
	"strconv"
	"fmt"
	"app/shared/workers"
)

func NewStat(w http.ResponseWriter, r *http.Request){
	_, err := workers.WP.AddTaskSyncTimed(func() interface{}{
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			return err
		}

		var stat model.Stat
		err = json.Unmarshal(b, &stat)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			return err
		}

		if !validateJsonNewStat(stat.UID, stat.Action, stat.Datetime) {
			w.WriteHeader(http.StatusPreconditionFailed)
			return err
		}

		uid := strconv.Itoa(int(stat.UID))
		usr, err := model.UserExist(uid)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		if usr.ID == stat.ID{
			log.Println("User exist")
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		ex := model.StatCreate(stat.UID, stat.Action, stat.Datetime)

		if ex != nil {
			log.Println(ex)
			w.WriteHeader(http.StatusBadRequest)
			return err
		} else {
			w.WriteHeader(http.StatusOK)
			return err
		}
	}, workers.RequestWaitInQueueTimeout)

	if err != nil {
		http.Error(w, fmt.Sprintf("error: %s!\n", err), 500)
	}
}

func validateJsonNewStat(user uint32, action string, ts time.Time) bool{

	if user <= 0 {
		log.Println("user id less than zero")
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
