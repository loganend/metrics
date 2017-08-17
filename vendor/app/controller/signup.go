package controller

import (
	"net/http"
	"fmt"
	"app/model"
	"log"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		//http.Error(w, err.Error(), 500)
		return
	}

	var user model.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		//http.Error(w, err.Error(), 500)
		return
	}

	id := fmt.Sprint(user.ID)
	age := fmt.Sprint(user.Age)
	sex := user.Sex

	if !validateJsonSignup(id, age, sex){
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}

	usr, err := model.UserExist(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user_id, err := strconv.Atoi(id)

	if int(usr.ID) == user_id{
		log.Println("User exist")
		w.WriteHeader(http.StatusBadRequest)
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

func validateJsonSignup(id, age, sex string) bool{
	var err error

	 _, err = strconv.Atoi(id)

	if err != nil{
		fmt.Println("id is not int")
		return false
	}

	agei , err := strconv.Atoi(age)

	fmt.Println(agei)

	if err != nil{
		fmt.Println("id is not int")
		return false
	}

	if agei <= 0 {
		fmt.Println("Age less than zero")
		return false
	}

	switch sex{
	case "male":
	case "female":
	default:
		return false
	}

	return true
}
