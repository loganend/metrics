package model

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"app/shared/database"
)

func TestUserCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	database.SQL = sqlx.NewDb(db,"sqlmock")

	mock.ExpectExec("INSERT INTO user").WithArgs("1", "20", "male").WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectCommit()

	if err = UserCreate("1", "20", "male"); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

}