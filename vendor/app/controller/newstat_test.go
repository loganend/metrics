package controller

import (
	"bytes"
	"net/http/httptest"
	"testing"
	"net/http"
	"app/shared/jsonconfig"
	"app/shared/database"
	"app/model"
)

func TestNewStatErrorInvalidJson(t *testing.T) {

	jsonconfig.Load("../../../config/config.json", config)
	database.Connect(config.TestDatabase)


	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", "/api/users/stats",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewStat)

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


func TestNewStatStatusOk(t *testing.T) {

	var jsonStr = []byte(`{"user": "2", "action": "comment","ts": "2016-08-14T15:17:21+04:00"}`)
	req, err := http.NewRequest("POST", "/api/users/stats",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewStat)

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

	err = model.StatGet("2", "comment", "2016-08-14 11:17:21")

	if err != nil{
		t.Errorf("New statistic wasn't added to database %v", err)
	}

	model.StatRemove(2, "comment", "2016-08-14 11:17:21")
}


func TestNewStatErrorInvalidDate(t *testing.T) {

	var jsonStr = []byte(`{"user": "6", "action": "like","ts": "2017-08-14T15:17:"}`)
	req, err := http.NewRequest("POST", "/api/users/stats",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewStat)

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

func TestNewStatErrorInvalidError(t *testing.T) {

	var jsonStr = []byte(`{"user": "6", "action": "hel","ts": "2017-08-14T15:17:21+03:00"}`)
	req, err := http.NewRequest("POST", "/api/users/stats",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewStat)

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

func TestNewStatErrorInvalidJsonEmptyField(t *testing.T) {

	var jsonStr = []byte(`{"user": "6"}`)
	req, err := http.NewRequest("POST", "/api/users/stats",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewStat)

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

func TestNewStatErrorUserNotExist(t *testing.T) {

	var jsonStr = []byte(`{"user": "100000", "action": "like","ts": "2017-08-14T15:17:21+03:00"}`)
	req, err := http.NewRequest("POST", "/api/users/stats",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewStat)

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


func TestNewStatComment(t *testing.T) {

	var jsonStr = []byte(`{"user": "3", "action": "comment","ts": "2017-08-14T15:17:21+03:00"}`)
	req, err := http.NewRequest("POST", "/api/users/stats",  bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(NewStat)

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
}