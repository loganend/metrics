package controller

import (
	"net/http"
	"fmt"
	"app/model"
	"log"
	"io/ioutil"
	"encoding/json"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user model.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id := fmt.Sprint(user.ID)
	age := fmt.Sprint(user.Age)
	sex := user.Sex

	fmt.Printf("id = %s, age = %s\n", id, age)

	if (id == "0" || age == "0"  || sex == ""){
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}


	ex := model.UserCreate(id, age, sex)

	if ex != nil {
		log.Println(ex)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
