package controller

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"app/shared/database"
	"app/shared/jsonconfig"
	"app/shared/server"
	"encoding/json"
	//"github.com/DATA-DOG/go-sqlmock"
	//"github.com/jmoiron/sqlx"
	"app/model"
)

var config = &configuration{}

type configuration struct {
	Database  database.Info   `json:"Database"`
	Server    server.Server   `json:"Server"`
	TestDatabase  database.Info `json:"TestDatabase"`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}


func TestSignupFirst(t *testing.T) {

	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	var jsonStr = []byte(`{"id":"2", "age":"12", "sex":"male"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestSignupSecond(t *testing.T){

	var jsonStr = []byte(`{"id":"1", "age":"-20", "sex":"male"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}

func TestSignupThird(t *testing.T){

	var jsonStr = []byte(`{"id":"1", "age":"20", "sex":"mae"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}

func TestSignupFourth(t *testing.T){

	var jsonStr = []byte(`{"id":"-1", "age":"20", "sex":""}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}


func TestSignupFifth(t *testing.T){

	var jsonStr = []byte(`{"id":"1", "sex":"female"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}

func TestSignupSixth(t *testing.T){

	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}


func TestSignupSeventh(t *testing.T){

	var jsonStr = []byte(`{"age":"10; "sex"="male"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}

func TestSignupEighth(t *testing.T){

	var jsonStr = []byte(`{"age":"10; "sex"="male"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}


func TestSignupNinth(t *testing.T){

	var jsonStr = []byte(`{"id":"12", "age":"", "sex":""}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}

func TestSignupTenth(t *testing.T) {

	var jsonStr = []byte(`{"id":"-2", "age":"20", "sex":"male"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}

	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSignupEleventh(t *testing.T) {

	var jsonStr = []byte(`{"id":"-2", "age":"20", "sex":"male"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}

	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSignupTwelfth(t *testing.T) {

	var jsonStr = []byte(`{"id":"100000", "age":"21", "sex":"male"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}


	user, err := model.UserExist("100000")

	if err != nil{
		t.Errorf("Error, handler hasn't created a new user: %s ", err)
	}

	if user.ID != 100000 {
		t.Errorf("handler hasn't created a new user: got %v want %v", user.ID, 100000)
	}

	model.UserRemove(100000)

	exist, err := model.UserExist("100000")

	if err != nil{
		t.Errorf("%v", err)
	}

	if exist.ID == 100000 {
		t.Error("User have to be removed")
	}

}