package controller

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"app/shared/database"
)


func TestSignup(t *testing.T) {

	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database.SQL = sqlx.NewDb(db,"sqlmock")

	var jsonStr = []byte(`{"id":"14", "age":"12", "sex":"male"}`)
	req, err := http.NewRequest("POST", "/api/users",  bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Signup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//expected := `{"alive": true}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}

	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

}