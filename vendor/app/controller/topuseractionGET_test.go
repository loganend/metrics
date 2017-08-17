package controller

import (
	"testing"
	"app/shared/jsonconfig"
	"app/shared/database"
	"net/http/httptest"
	"net/http"
	"net/url"
)

func TestGetTopFirst(t *testing.T) {
	jsonconfig.Load("/Users/serqeycheremisin/work/src/github.com/metrics/config/"+"config.json", config)
	database.Connect(config.TestDatabase)


	req, err := http.NewRequest("GET", "/api/users",
		url.Values{"date1": {"2017-08-01"}, "date2":{"2017-08-10"}, "action":{"like"}, "limit":{"2"}})

	//req.Header.Set("Content-Type", "application/json")

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

	expected := "{\"items\":{\"2017.08.01\":[{\"id\":\"3\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"3\"},{\"id\":\"2\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"1\"}],\"2017.08.02\":[{\"id\":\"5\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"1\"}]," +
		"\"2017.08.03\":[{\"id\":\"11\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"8\"}],\"2017.08.04\":[{\"id\":\"1\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"1\"}],\"2017.08.05\":[{\"id\":\"9\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"9\"},{\"id\":\"11\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"3\"}]," +
		"\"2017.08.07\":[{\"id\":\"11\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"7\"},{\"id\":\"4\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"4\"}],\"2017.08.08\":[{\"id\":\"9\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"2\"}]," +
		"\"2017.08.09\":[{\"id\":\"5\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"1\"}],\"2017.08.10\":[{\"id\":\"5\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"2\"},{\"id\":\"6\",\"age\":\"20\",\"sex\":\"male\",\"count\":\"2\"}]}}"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
