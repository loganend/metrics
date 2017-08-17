package controller

import (
	"testing"
	"app/shared/jsonconfig"
	"app/shared/database"
	"net/http/httptest"
	"net/http"
)

func TestGetTopStatusOkActionLike(t *testing.T) {

	jsonconfig.Load("../../../config/config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01&date2=2017-08-10&action=like&limit=2", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "{\"items\":{\"2017.08.01\":[{\"id\":\"3\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"3\"},{\"id\":\"2\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"1\"}],\"2017.08.02\":[{\"id\":\"5\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"1\"}]," +
		"\"2017.08.03\":[{\"id\":\"11\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"8\"}],\"2017.08.04\":[{\"id\":\"1\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"1\"}],\"2017.08.05\":[{\"id\":\"9\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"9\"},{\"id\":\"11\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"3\"}]," +
		"\"2017.08.07\":[{\"id\":\"11\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"7\"},{\"id\":\"4\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"4\"}],\"2017.08.08\":[{\"id\":\"9\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"2\"}]," +
		"\"2017.08.09\":[{\"id\":\"5\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"1\"}],\"2017.08.10\":[{\"id\":\"5\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"2\"},{\"id\":\"6\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"2\"}]}}"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetTopErrorDate2LessDate1(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01&date2=2017-07-10&action=like&limit=2", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

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


func TestGetTopErrorLimitEmpty(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01&date2=2017-08-10&action=like&limit=", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

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


func TestGetTopErrorLimitZero(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01&date2=2017-08-10&action=like&limit=0", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

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

func TestGetTopErrorInvalidDate(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-0ds8-01&date2=2017-08-10&action=like&limit=4", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

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

func TestGetTopErrorInvalidAction(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01&date2=2017-08-10&actsdion=like&limit=4", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

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

func TestGetTopErrorInvalidJson(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01&actsdion=like&limit=4", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

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

func TestGetTopStatusOkActionLogin(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01&date2=2017-08-10&action=login&limit=2", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "{\"items\":{}}"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


func TestGetTopStatusOkActionComment(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01&date2=2017-08-10&action=comment&limit=2", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "{\"items\":{\"2017.08.07\":[{\"id\":\"1\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"2\"}]}}"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetTopErrorDateFormat(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)

	//r := bytes.NewReader(jsonStr)

	req, err := http.NewRequest("GET", "/api/users?date1=2017-08-01T15:17:21+03:00&date2=2017-08-10&action=comment&limit=2", nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil{
		t.Errorf("expected no error got %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetTop)

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